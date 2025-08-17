package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Все статьи (GET /articles)
func GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// Одна статья (GET /articles/{id})
func GetArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range articles {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	http.Error(w, "Статья не найдена", http.StatusNotFound)
}

// Создание статьи (POST /articles)
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var article Article

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

// Обновление статьи (PUT /articles/{id})
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)

			var article Article
			if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			article.ID = params["id"]
			articles = append(articles, article)
			json.NewEncoder(w).Encode(article)
			return
		}
	}

	http.Error(w, "Статья не найдена", http.StatusNotFound)
}

// Удаление статьи (DELETE /articles/{id})
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range articles {
		if item.ID == params["id"] {
			articles = append(articles[:index], articles[index+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Статья не найдена", http.StatusNotFound)
}
