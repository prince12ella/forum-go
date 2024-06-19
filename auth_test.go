package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
)

func TestRegister(t *testing.T) {
    req, err := http.NewRequest("POST", "/register", strings.NewReader(`{"email":"test@example.com","username":"testuser","password":"testpassword"}`))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Register)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}

func TestLogin(t *testing.T) {
    req, err := http.NewRequest("POST", "/login", strings.NewReader(`{"username":"testuser","password":"testpassword"}`))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Login)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

func TestLogout(t *testing.T) {
    req, err := http.NewRequest("POST", "/logout", nil)
    if err != nil {
        t.Fatal(err)
    }
    req.AddCookie(&http.Cookie{Name: "session_token", Value: "valid-session-id"})
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(Logout)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}
