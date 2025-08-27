package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// User структура для демонстрации
type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

// GetInfo возвращает информацию о пользователе
func (u User) GetInfo() (string, bool) {
	fullName := u.FirstName + " " + u.LastName
	isAdult := u.Age >= 18
	return fullName, isAdult
}

// GetEmail возвращает email пользователя
func (u User) GetEmail() (string, error) {
	if u.Email == "" {
		return "", fmt.Errorf("email не установлен")
	}
	return u.Email, nil
}

// swap меняет местами два числа
func swap(a, b int) (int, int) {
	return b, a
}

// calculateStats вычисляет статистику по числам
func calculateStats(numbers ...int) (min, max, sum, avg int) {
	if len(numbers) == 0 {
		return 0, 0, 0, 0
	}
	
	min = numbers[0]
	max = numbers[0]
	sum = 0
	
	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}
	
	if len(numbers) > 0 {
		avg = sum / len(numbers)
	}
	return
}

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/calculate", calculateHandler)
	http.HandleFunc("/user", userHandler)
	
	// Запускаем сервер
	fmt.Println("Веб-демонстрация запущена на http://localhost:8080")
	fmt.Println("Доступные endpoints:")
	fmt.Println("  / - главная страница")
	fmt.Println("  /calculate?numbers=1,2,3 - расчет статистики")
	fmt.Println("  /user - информация о пользователе")
	
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Вычисляем результат swap заранее
	swapA, swapB := swap(5, 10)
	
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>Множественные возвращаемые значения в Go</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 800px; margin: 0 auto; }
		.example { background: #f5f5f5; padding: 20px; margin: 20px 0; border-radius: 8px; }
		.result { background: #e0f7fa; padding: 15px; border-radius: 5px; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Демонстрация множественных возвращаемых значений в Go</h1>
		
		<div class="example">
			<h2>Пример 1: Обмен значений</h2>
			<p>swap(5, 10) = ` + fmt.Sprintf("%d, %d", swapA, swapB) + `</p>
		</div>
		
		<div class="example">
			<h2>Пример 2: Статистика чисел</h2>
			<form action="/calculate" method="get">
				<label>Введите числа через запятую:</label>
				<input type="text" name="numbers" value="1,2,3,4,5">
				<button type="submit">Рассчитать</button>
			</form>
		</div>
		
		<div class="example">
			<h2>Пример 3: Информация о пользователе</h2>
			<p><a href="/user">Посмотреть информацию</a></p>
		</div>
	</div>
</body>
</html>`
	
	fmt.Fprint(w, html)
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	numbersStr := r.URL.Query().Get("numbers")
	
	// Парсим числа из строки
	var numbers []int
	current := ""
	for _, char := range numbersStr {
		if char == ',' {
			if num, err := strconv.Atoi(current); err == nil {
				numbers = append(numbers, num)
			}
			current = ""
		} else if char != ' ' {
			current += string(char)
		}
	}
	if current != "" {
		if num, err := strconv.Atoi(current); err == nil {
			numbers = append(numbers, num)
		}
	}
	
	// Вычисляем статистику
	min, max, sum, avg := calculateStats(numbers...)
	
	// Выводим результат
	fmt.Fprintf(w, `<div style="font-family: Arial, sans-serif; padding: 20px;">
		<h2>Результаты расчета</h2>
		<p>Числа: %v</p>
		<div style="background: #e0f7fa; padding: 15px; border-radius: 5px;">
			<p>Минимум: %d</p>
			<p>Максимум: %d</p>
			<p>Сумма: %d</p>
			<p>Среднее: %d</p>
		</div>
		<p><a href="/">Назад</a></p>
	</div>`, numbers, min, max, sum, avg)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	// Создаем пользователя
	user := User{
		FirstName: "Мария",
		LastName:  "Сидорова",
		Age:       22,
		Email:     "maria.sidorova@example.com",
	}
	
	// Получаем информацию
	fullName, isAdult := user.GetInfo()
	email, err := user.GetEmail()
	
	// Выводим результат
	fmt.Fprintf(w, `<div style="font-family: Arial, sans-serif; padding: 20px;">
		<h2>Информация о пользователе</h2>
		<div style="background: #e0f7fa; padding: 15px; border-radius: 5px;">
			<p>Полное имя: %s</p>
			<p>Возраст: %d</p>
			<p>Совершеннолетний: %t</p>
			<p>Email: %s</p>
			<p>Ошибка email: %v</p>
		</div>
		<p><a href="/">Назад</a></p>
	</div>`, fullName, user.Age, isAdult, email, err)
}