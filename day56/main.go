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

type Example struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Code        string `json:"code"`
	Result      string `json:"result,omitempty"`
}

type PointerDemo struct {
	Value    int    `json:"value"`
	Pointer  *int   `json:"pointer"`
	Action   string `json:"action"`
	NewValue int    `json:"newValue,omitempty"`
}

func main() {
	r := chi.NewRouter()

	// CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	})

	r.Use(corsMiddleware.Handler)

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Routes
	r.Get("/", homeHandler)
	r.Get("/examples", examplesHandler)
	r.Post("/pointer-demo", pointerDemoHandler)

	// Start server
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

func examplesHandler(w http.ResponseWriter, r *http.Request) {
	examples := []Example{
		{
			Title:       "Базовый пример указателя",
			Description: "Создание указателя на переменную и получение значения через указатель",
			Code: `x := 42
p := &x
fmt.Println("Значение x:", x)    // 42
fmt.Println("Адрес x:", &x)      // 0x...
fmt.Println("Значение p:", p)    // 0x...
fmt.Println("Значение *p:", *p)  // 42`,
		},
		{
			Title:       "Изменение значения через указатель",
			Description: "Модификация значения переменной через указатель",
			Code: `x := 42
p := &x
*p = 100
fmt.Println("Новое значение x:", x)  // 100`,
		},
		{
			Title:       "Указатели в функциях",
			Description: "Использование указателей для модификации переменных в функциях",
			Code: `func increment(p *int) {
    *p = *p + 1
}

x := 10
increment(&x)
fmt.Println("После инкремента:", x)  // 11`,
		},
		{
			Title:       "Сравнение с Python",
			Description: "В Python нет прямых указателей, но есть mutable/immutable объекты",
			Code: `# В Python мы работаем с ссылками на объекты
def modify_list(lst):
    lst.append(4)

my_list = [1, 2, 3]
modify_list(my_list)
print(my_list)  # [1, 2, 3, 4]

# Но с числами это не работает
def modify_number(n):
    n += 1

x = 10
modify_number(x)
print(x)  # 10 (не изменилось)`,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(examples)
}

func pointerDemoHandler(w http.ResponseWriter, r *http.Request) {
	var demo PointerDemo
	if err := json.NewDecoder(r.Body).Decode(&demo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := ""

	switch demo.Action {
	case "create":
		demo.Pointer = &demo.Value
		result = fmt.Sprintf("Создан указатель на значение %d. Адрес: %p", demo.Value, demo.Pointer)
	case "dereference":
		if demo.Pointer != nil {
			result = fmt.Sprintf("Значение по указателю %p: %d", demo.Pointer, *demo.Pointer)
		} else {
			result = "Указатель не инициализирован (nil)"
		}
	case "modify":
		if demo.Pointer != nil {
			oldValue := *demo.Pointer
			*demo.Pointer = demo.NewValue
			result = fmt.Sprintf("Изменено значение по адресу %p: %d → %d", demo.Pointer, oldValue, demo.NewValue)
		} else {
			result = "Нельзя изменить значение: указатель nil"
		}
	case "compare":
		p1 := &demo.Value
		p2 := &demo.Value
		p3 := &demo.NewValue

		result = fmt.Sprintf("p1 (%p) == p2 (%p): %t\n", p1, p2, p1 == p2)
		result += fmt.Sprintf("p1 (%p) == p3 (%p): %t\n", p1, p3, p1 == p3)
		result += fmt.Sprintf("*p1 == *p3: %t", *p1 == *p3)
	}

	response := map[string]interface{}{
		"result":   result,
		"typeInfo": getTypeInfo(demo.Pointer),
		"pointer":  demo.Pointer,
		"value":    demo.Value,
		"newValue": demo.NewValue,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(response)
}

func getTypeInfo(p *int) string {
	if p == nil {
		return "Тип: *int, значение: nil"
	}
	return fmt.Sprintf("Тип: %T, значение: %d, адрес: %p", p, *p, p)
}
