package handlers

import (
    "testing"
    "net/http"
    "net/http/httptest"
    "strings"
)

func TestCreateCategory(t *testing.T) {
    req, err := http.NewRequest("POST", "/createCategory", strings.NewReader(`{"name":"testCategory"}`))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateCategory)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}

func TestCreatePostWithCategory(t *testing.T) {
    // This assumes that you have a session and a user created already
    req, err := http.NewRequest("POST", "/createPost", strings.NewReader(`{"content":"testPost", "categories":[1]}`))
    if err != nil {
        t.Fatal(err)
    }
    req.AddCookie(&http.Cookie{Name: "session_token", Value: "valid-session-id"})
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreatePost)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
    }
}
