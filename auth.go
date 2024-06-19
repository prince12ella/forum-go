package handlers

import (
    "encoding/json"
    "errors"
    "net/http"
    "time"
    "github.com/google/uuid"
    "github.com/prince12ella/forum-go/internal/db"
    "github.com/prince12ella/forum-go/internal/models"
    "golang.org/x/crypto/bcrypt"
)

var sessions = map[string]int{} // sessionID -> userID

func Register(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    user.Password = string(hashedPassword)

    _, err = db.GetDB().Exec("INSERT INTO users (email, username, password) VALUES (?, ?, ?)", user.Email, user.Username, user.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    row := db.GetDB().QueryRow("SELECT id, password FROM users WHERE username = ?", user.Username)
    var storedUser models.User
    err = row.Scan(&storedUser.ID, &storedUser.Password)
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        http.Error(w, "Invalid username or password", http.StatusUnauthorized)
        return
    }

    sessionID := uuid.New().String()
    sessions[sessionID] = storedUser.ID

    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   sessionID,
        Expires: time.Now().Add(24 * time.Hour),
    })

    w.WriteHeader(http.StatusOK)
}

func CheckSession(r *http.Request) (int, error) {
    cookie, err := r.Cookie("session_token")
    if err != nil {
        return 0, err
    }

    userID, exists := sessions[cookie.Value]
    if !exists {
        return 0, errors.New("invalid session")
    }

    return userID, nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    cookie, err := r.Cookie("session_token")
    if err != nil {
        http.Error(w, "Session not found", http.StatusUnauthorized)
        return
    }

    delete(sessions, cookie.Value)

    http.SetCookie(w, &http.Cookie{
        Name:    "session_token",
        Value:   "",
        Expires: time.Now().AddDate(0, 0, -1),
    })

    w.WriteHeader(http.StatusOK)
}
