package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

type ArrayDemo struct {
	Title       string
	Description string
	Code        string
	Result      string
}

type PageData struct {
	Title string
	Demos []ArrayDemo
}

func setupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Обслуживание статических файлов
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("./static/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/",
		http.FileServer(http.Dir("./static/js"))))

	// Маршруты
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/api/arrays", arraysHandler).Methods("GET", "POST")

	return router
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Указываем правильный путь к шаблону
	path := filepath.Join("static", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Ошибка загрузки шаблона: "+err.Error(), http.StatusInternalServerError)
		return
	}

	demos := []ArrayDemo{
		{
			Title:       "Создание массива",
			Description: "Создаем массив из 5 целых чисел",
			Code:        "var numbers [5]int\nnumbers = [5]int{1, 2, 3, 4, 5}",
			Result:      "[1 2 3 4 5]",
		},
		{
			Title:       "Короткое объявление",
			Description: "Создаем массив с кратким синтаксисом",
			Code:        "fruits := [3]string{\"Яблоко\", \"Груша\", \"Апельсин\"}",
			Result:      "[Яблоко Груша Апельсин]",
		},
		{
			Title:       "Многоточие в размере",
			Description: "Компилятор сам посчитает размер массива",
			Code:        "autoSize := [...]int{10, 20, 30, 40}",
			Result:      "[10 20 30 40] (размер: 4)",
		},
		{
			Title:       "Доступ к элементам",
			Description: "Получаем и изменяем элементы массива",
			Code:        "numbers := [3]int{1, 2, 3}\nfirst := numbers[0]\nnumbers[2] = 99",
			Result:      "first = 1, numbers = [1 2 99]",
		},
		{
			Title:       "Длина массива",
			Description: "Узнаем длину массива с помощью len()",
			Code:        "arr := [4]string{\"a\", \"b\", \"c\", \"d\"}\nlength := len(arr)",
			Result:      "Длина массива: 4",
		},
		{
			Title:       "Итерация по массиву",
			Description: "Перебираем элементы массива с помощью for",
			Code:        "arr := [3]int{10, 20, 30}\nfor i := 0; i < len(arr); i++ {\n    fmt.Println(arr[i])\n}",
			Result:      "10\n20\n30",
		},
		{
			Title:       "Итерация range",
			Description: "Используем range для итерации",
			Code:        "arr := [3]string{\"a\", \"b\", \"c\"}\nfor index, value := range arr {\n    fmt.Printf(\"Индекс: %d, Значение: %s\\n\", index, value)\n}",
			Result:      "Индекс: 0, Значение: a\nИндекс: 1, Значение: b\nИндекс: 2, Значение: c",
		},
	}

	data := PageData{
		Title: "Массивы в Go: От основ к мастерству",
		Demos: demos,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Ошибка выполнения шаблона: "+err.Error(), http.StatusInternalServerError)
	}
}

func arraysHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		json.NewEncoder(w).Encode(map[string]string{"status": "success"})
		return
	}

	response := map[string]interface{}{
		"array":   [5]int{1, 2, 3, 4, 5},
		"message": "Пример массива в Go",
	}

	json.NewEncoder(w).Encode(response)
}
