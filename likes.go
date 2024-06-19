package handlers

import (
    "net/http"

    "github.com/prince12ella/forum-go/internal/models"
)

// LikePost gère la requête de like sur un post
func LikePost(w http.ResponseWriter, r *http.Request) {
    // Parse user session from cookie or authorization header
    user := getUserFromSession(r)

    // Parse post ID from request
    postID := r.FormValue("post_id") // Assuming post_id is passed as a form value

    // Create or update like in database
    like := models.Like{
        UserID: user.ID,
        PostID: postID,
        Type:   models.LikeType,
    }

    // Save the like in database (assuming db package handles database operations)
    err := db.SaveLike(&like)
    if err != nil {
        http.Error(w, "Failed to save like", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusCreated)
}

// DislikePost gère la requête de dislike sur un post
func DislikePost(w http.ResponseWriter, r *http.Request) {
    // Parse user session from cookie or authorization header
    user := getUserFromSession(r)

    // Parse post ID from request
    postID := r.FormValue("post_id") // Assuming post_id is passed as a form value

    // Create or update dislike in database
    dislike := models.Like{
        UserID: user.ID,
        PostID: postID,
        Type:   models.DislikeType,
    }

    // Save the dislike in database
    err := db.SaveLike(&dislike)
    if err != nil {
        http.Error(w, "Failed to save dislike", http.StatusInternalServerError)
        return
    }

    // Return success response
    w.WriteHeader(http.StatusCreated)
}
