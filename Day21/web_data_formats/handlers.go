package main

import (
	"html/template"
	"net/http"
)

// Главная страница
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

// Обработка конвертации
func convertHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	inputData := r.FormValue("inputData")
	format := r.FormValue("format")
	
	var result string
	var err error
	
	switch format {
	case "json":
		result, err = convertToJSON(inputData)
	case "xml":
		result, err = convertToXML(inputData)
	default:
		result = "Неизвестный формат"
	}
	
	if err != nil {
		result = "Ошибка: " + err.Error()
	}
	
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, struct {
		Input  string
		Format string
		Result string
	}{
		Input:  inputData,
		Format: format,
		Result: result,
	})
}