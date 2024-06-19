package handlers

import (
    "net/http"
    "strconv"

    "github.com/prince12ella/forum-go/internal/models"
)

// GetPostsByCategory récupère les posts par catégorie
func GetPostsByCategory(w http.ResponseWriter, r *http.Request) {
    // Récupérer l'ID de la catégorie depuis la requête
    categoryID, err := strconv.Atoi(r.FormValue("category_id"))
    if err != nil {
        http.Error(w, "Invalid category ID", http.StatusBadRequest)
        return
    }

    // Récupérer les posts par catégorie depuis la base de données
    posts, err := db.GetPostsByCategory(categoryID)
    if err != nil {
        http.Error(w, "Failed to fetch posts by category", http.StatusInternalServerError)
        return
    }

    // Convertir les posts en JSON et envoyer la réponse
    // (Utilisez un package d'encodage JSON comme "encoding/json")
    // Par exemple :
    // json.NewEncoder(w).Encode(posts)
    // Assurez-vous d'écrire correctement la réponse et de gérer les erreurs si nécessaire
}
