package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"error_handling/utils"
)

// StartWebServer запускает веб-сервер с обработкой ошибок
func StartWebServer() error {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/divide", divideHandler)
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/panic", panicHandler)

	fmt.Println("Веб-сервер запущен на http://localhost:8080")
	return http.ListenAndServe(":8080", nil)
}

// homeHandler обрабатывает главную страницу
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>Обработка ошибок в Go</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 800px; margin: 0 auto; }
		.example { background: #f5f5f5; padding: 20px; margin: 20px 0; border-radius: 8px; }
		.error { color: #d32f2f; background: #ffebee; padding: 10px; border-radius: 4px; }
		.success { color: #388e3c; background: #e8f5e9; padding: 10px; border-radius: 4px; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Демонстрация обработки ошибок в Go</h1>
		
		<div class="example">
			<h2>Деление чисел</h2>
			<form action="/divide" method="get">
				<label>Числитель: <input type="number" name="a" value="10"></label><br>
				<label>Знаменатель: <input type="number" name="b" value="2"></label><br>
				<button type="submit">Разделить</button>
			</form>
		</div>
		
		<div class="example">
			<h2>Информация о пользователе</h2>
			<form action="/user" method="get">
				<label>Имя: <input type="text" name="name" value="Иван"></label><br>
				<label>Возраст: <input type="number" name="age" value="25"></label><br>
				<button type="submit">Проверить</button>
			</form>
		</div>
		
		<div class="example">
			<h2>Тест паники</h2>
			<p><a href="/panic">Вызвать панику</a> (сервер должен восстановиться)</p>
		</div>
	</div>
</body>
</html>`
	
	fmt.Fprint(w, html)
}

// divideHandler обрабатывает деление чисел
func divideHandler(w http.ResponseWriter, r *http.Request) {
	defer utils.RecoverFromPanic(w, r)
	
	aStr := r.URL.Query().Get("a")
	bStr := r.URL.Query().Get("b")
	
	a, err := strconv.ParseFloat(aStr, 64)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "неверный формат числа a")
		return
	}
	
	b, err := strconv.ParseFloat(bStr, 64)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "неверный формат числа b")
		return
	}
	
	result, err := utils.Divide(a, b)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	fmt.Fprintf(w, `<div class="success">Результат: %.2f / %.2f = %.2f</div><p><a href="/">Назад</a></p>`, a, b, result)
}

// userHandler обрабатывает проверку пользователя
func userHandler(w http.ResponseWriter, r *http.Request) {
	defer utils.RecoverFromPanic(w, r)
	
	name := r.URL.Query().Get("name")
	ageStr := r.URL.Query().Get("age")
	
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		utils.SendError(w, http.StatusBadRequest, "неверный формат возраста")
		return
	}
	
	if err := utils.ValidateUser(name, age); err != nil {
		utils.SendError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	fmt.Fprintf(w, `<div class="success">Пользователь %s прошел проверку</div><p><a href="/">Назад</a></p>`, name)
}

// panicHandler вызывает панику для демонстрации восстановления
func panicHandler(w http.ResponseWriter, r *http.Request) {
	defer utils.RecoverFromPanic(w, r)
	
	// Имитация паники
	panic("демонстрация паники в веб-обработчике")
}