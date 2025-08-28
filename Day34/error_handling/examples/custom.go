package examples

import (
	"fmt"
	"time"
)

// CustomErrors демонстрирует создание и использование кастомных ошибок
func CustomErrors() error {
	fmt.Println("Работа с кастомными ошибками...")

	// Пример 1: Проверка пользователя
	user := User{Name: "Гоша", Age: 37}
	if err := validateUser(user); err != nil {
		if err, ok := err.(*ValidationError); ok {
			fmt.Printf("Ошибка валидации: %s (код: %d)\n", err.Message, err.Code)
		} else {
			fmt.Printf("Неизвестная ошибка: %v\n", err)
		}
	}

	// Пример 2: Проверка взрослого пользователя
	user = User{Name: "Петр", Age: 25}
	if err := validateUser(user); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Printf("Пользователь %s прошел валидацию\n", user.Name)
	}

	// Пример 3: Сетевая ошибка
	if err := checkNetwork(); err != nil {
		if err, ok := err.(*NetworkError); ok {
			fmt.Printf("Сетевая ошибка: %s (статус: %d, время: %v)\n", 
				err.Message, err.StatusCode, err.Timestamp)
		}
	}

	return nil
}

// User представляет структуру пользователя
type User struct {
	Name string
	Age  int
}

// ValidationError представляет кастомную ошибку валидации
type ValidationError struct {
	Code    int
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации %d: %s", e.Code, e.Message)
}

// NetworkError представляет кастомную сетевую ошибку
type NetworkError struct {
	StatusCode int
	Message    string
	Timestamp  time.Time
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("сетевая ошибка %d: %s", e.StatusCode, e.Message)
}

// validateUser проверяет корректность данных пользователя
func validateUser(user User) error {
	if user.Name == "" {
		return &ValidationError{Code: 1001, Message: "имя не может быть пустым"}
	}
	if user.Age < 18 {
		return &ValidationError{Code: 1002, Message: "возраст должен быть не менее 18 лет"}
	}
	return nil
}

// checkNetwork имитирует проверку сетевого соединения
func checkNetwork() error {
	return &NetworkError{
		StatusCode: 500,
		Message:    "внутренняя ошибка сервера",
		Timestamp:  time.Now(),
	}
}