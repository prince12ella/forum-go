package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/prince12ella/forum-go/internal/db"
    "github.com/prince12ella/forum-go/internal/models"
)

func CreateCategory(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var category models.Category
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err = db.GetDB().Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
}

func GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    categoryID := r.URL.Query().Get("category_id")
    rows, err := db.GetDB().Query(`
        SELECT p.id, p.user_id, p.content 
        FROM posts p 
        JOIN post_categories pc ON p.id = pc.post_id 
        WHERE pc.category_id = ?`, categoryID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var posts []models.Post
    for rows.Next() {
        var post models.Post
        err = rows.Scan(&post.ID, &post.UserID, &post.Content)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        posts = append(posts, post)
    }

    if err = rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(posts)
}
