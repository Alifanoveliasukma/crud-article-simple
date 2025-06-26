package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Alifanoveliasukma/crud-article-simple/models"
	"github.com/gorilla/mux"
)

// Create
func CreateArticleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	var article models.Article 
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	err = models.CreateArticle(db, article.Judul, article.GambarURL, article.Kategori, article.Content, article.CreatedAt)
	if err != nil {
		http.Error(w, "Failed to create article", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Article created successfully")
	}
}

func GetArticleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		articleID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest)
			return
		}

		article, err := models.GetArticle(db, articleID)
		if err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(article)

	}
}

func UpdateArticleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		idStr := vars["id"]

		articleID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
			return
		}

		var article models.Article
		err = json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		err = models.UpdateArticle(db, articleID, article.Judul, article.GambarURL, article.Kategori, article.Content, article.CreatedAt)
		if err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}
		fmt.Fprintln(w, "Article updated successfully")
	}
}

func DeleteArticleHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		articleID, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
			return
		}

		err = models.DeleteArticle(db, articleID)
		if err != nil {
			http.Error(w, "Article not found", http.StatusNotFound)
			return
		}

		fmt.Fprintln(w, "User deleted successfully")

	}
}