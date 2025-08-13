package main

import (
    "encoding/json"
    "net/http"
    "text/template"
)

// WebController - структура для обработчиков
type WebController struct{}

// HomeHandler - главная страница
func (wc *WebController) HomeHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

// AnimalHandler - обработчик API для животных
func (wc *WebController) AnimalHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    
    species := r.URL.Query().Get("species")
    if species == "" {
        http.Error(w, `{"error": "Не указано животное"}`, http.StatusBadRequest)
        return
    }
    
    animal := AnimalFactory(species)
    response := map[string]string{
        "species": species,
        "sound":   animal.Speak(),
        "oopInfo": "Демонстрация ООП в Go: интерфейсы и структуры",
    }
    
    json.NewEncoder(w).Encode(response)
}