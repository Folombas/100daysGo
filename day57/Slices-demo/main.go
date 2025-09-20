package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/cors"
)

type SliceOperation struct {
	Operation  string `json:"operation"`
	Input      []int  `json:"input"`
	Input2     []int  `json:"input2,omitempty"`
	Index      int    `json:"index,omitempty"`
	Value      int    `json:"value,omitempty"`
	Start      int    `json:"start,omitempty"`
	End        int    `json:"end,omitempty"`
	Result     []int  `json:"result"`
	Explanation string `json:"explanation"`
}

func main() {
	r := chi.NewRouter()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	r.Use(corsMiddleware.Handler)

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	r.Get("/", homeHandler)
	r.Post("/slice-operation", sliceOperationHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, nil)
}

func sliceOperationHandler(w http.ResponseWriter, r *http.Request) {
	var op SliceOperation
	if err := json.NewDecoder(r.Body).Decode(&op); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	switch op.Operation {
	case "create":
		op.Result = op.Input
		op.Explanation = "Создан новый срез с указанными значениями"
	case "append":
		op.Result = append(op.Input, op.Value)
		op.Explanation = fmt.Sprintf("Добавлено значение %d в конец среза", op.Value)
	case "access":
		if op.Index >= 0 && op.Index < len(op.Input) {
			op.Result = []int{op.Input[op.Index]}
			op.Explanation = fmt.Sprintf("Получен элемент с индексом %d: %d", op.Index, op.Input[op.Index])
		} else {
			op.Explanation = "Ошибка: индекс вне диапазона"
		}
	case "update":
		if op.Index >= 0 && op.Index < len(op.Input) {
			op.Result = make([]int, len(op.Input))
			copy(op.Result, op.Input)
			op.Result[op.Index] = op.Value
			op.Explanation = fmt.Sprintf("Обновлен элемент с индексом %d на значение %d", op.Index, op.Value)
		} else {
			op.Explanation = "Ошибка: индекс вне диапазона"
		}
	case "slice":
		if op.Start >= 0 && op.End <= len(op.Input) && op.Start <= op.End {
			op.Result = op.Input[op.Start:op.End]
			op.Explanation = fmt.Sprintf("Получен подсрез с индекса %d по %d", op.Start, op.End)
		} else {
			op.Explanation = "Ошибка: неверные границы среза"
		}
	case "length":
		op.Result = []int{len(op.Input)}
		op.Explanation = fmt.Sprintf("Длина среза: %d элементов", len(op.Input))
	case "capacity":
		op.Result = []int{cap(op.Input)}
		op.Explanation = fmt.Sprintf("Емкость среза: %d элементов", cap(op.Input))
	case "iterate":
		op.Result = op.Input
		op.Explanation = "Итерация по всем элементам среза"
	case "copy":
		op.Result = make([]int, len(op.Input))
		copy(op.Result, op.Input)
		op.Explanation = "Создана копия среза"
	case "concat":
		op.Result = append(op.Input, op.Input2...)
		op.Explanation = "Объединение двух срезов"
	default:
		op.Explanation = "Неизвестная операция"
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(op)
}
