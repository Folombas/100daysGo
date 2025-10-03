package main

import (
	"errors"
	"fmt"
)

// 🎯 1. БАЗОВЫЙ ИНТЕРФЕЙС error
func demoBasicError() {
	fmt.Println("🎯 1. БАЗОВЫЙ ИНТЕРФЕЙС error")
	fmt.Println("==============================")

	// error - это простой интерфейс:
	// type error interface {
	//     Error() string
	// }

	// Создание ошибок
	err1 := errors.New("простая ошибка")
	err2 := fmt.Errorf("ошибка с форматом: %d", 404)

	fmt.Printf("err1: %v\n", err1)
	fmt.Printf("err2: %v\n", err2)
	fmt.Printf("err1.Error(): %s\n", err1.Error())
}

// 🎯 2. КАСТОМНЫЕ ОШИБКИ
type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("ошибка %s: %s (значение: %v)",
		e.Field, e.Message, e.Value)
}

type NetworkError struct {
	URL     string
	Code    int
	Message string
}

func (e NetworkError) Error() string {
	return fmt.Sprintf("сеть %s: %d %s", e.URL, e.Code, e.Message)
}

// 🎯 3. ПРОВЕРКА ОШИБОК
func demoErrorChecking() {
	fmt.Println("\n🎯 2. ПРОВЕРКА ОШИБОК")
	fmt.Println("=====================")

	// errors.Is - проверка конкретной ошибки
	targetErr := errors.New("целевая ошибка")
	err := fmt.Errorf("обертка: %w", targetErr)

	if errors.Is(err, targetErr) {
		fmt.Println("✅ errors.Is: найдена целевая ошибка")
	}

	// errors.As - проверка типа ошибки
	valErr := ValidationError{Field: "email", Message: "невалидно"}
	wrappedValErr := fmt.Errorf("валидация: %w", valErr)

	var extractedErr ValidationError
	if errors.As(wrappedValErr, &extractedErr) {
		fmt.Printf("✅ errors.As: извлекли %v\n", extractedErr)
	}
}

// 🎯 4. ОБЕРТЫВАНИЕ ОШИБОК
func demoErrorWrapping() {
	fmt.Println("\n🎯 3. ОБЕРТЫВАНИЕ ОШИБОК")
	fmt.Println("========================")

	baseErr := errors.New("базовая ошибка")

	// Цепочка оберток
	wrapped1 := fmt.Errorf("уровень 1: %w", baseErr)
	wrapped2 := fmt.Errorf("уровень 2: %w", wrapped1)

	fmt.Printf("Цепочка: %v\n", wrapped2)

	// Распаковка
	if errors.Is(wrapped2, baseErr) {
		fmt.Println("✅ Нашли базовую ошибку через обертки")
	}
}

// 🎯 5. ПРАКТИЧЕСКИЙ ПРИМЕР
type UserService struct{}

func (s *UserService) Register(user User) error {
	if err := s.validate(user); err != nil {
		return fmt.Errorf("регистрация: %w", err)
	}

	fmt.Printf("✅ Пользователь %s зарегистрирован!\n", user.Name)
	return nil
}

func (s *UserService) validate(user User) error {
	if user.Name == "" {
		return ValidationError{
			Field:   "name",
			Message: "обязательное поле",
		}
	}

	if user.Age < 18 {
		return ValidationError{
			Field:   "age",
			Value:   user.Age,
			Message: "только 18+",
		}
	}

	return nil
}

type User struct {
	Name string
	Age  int
	Email string
}

// 🎯 6. SENTINEL ERRORS
var (
	ErrUserNotFound = errors.New("пользователь не найден")
	ErrDBConnection = errors.New("ошибка базы данных")
)

func demoPractical() {
	fmt.Println("\n🎯 4. ПРАКТИЧЕСКИЙ ПРИМЕР")
	fmt.Println("=========================")

	service := &UserService{}

	// Успешный случай
	fmt.Println("✅ Валидные данные:")
	user1 := User{Name: "Алексей", Age: 25}
	if err := service.Register(user1); err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	}

	// Ошибка валидации
	fmt.Println("\n❌ Невалидные данные:")
	user2 := User{Name: "", Age: 16}
	if err := service.Register(user2); err != nil {
		fmt.Printf("Ошибка: %v\n", err)

		// Извлечение ValidationError
		var valErr ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("Детали: поле '%s'\n", valErr.Field)
		}
	}
}

// 🎯 7. BEST PRACTICES
func demoBestPractices() {
	fmt.Println("\n🎯 5. BEST PRACTICES")
	fmt.Println("===================")

	practices := []string{
		"✅ Всегда проверяйте ошибки",
		"✅ Добавляйте контекст через fmt.Errorf",
		"✅ Используйте errors.Is для проверки",
		"✅ Используйте errors.As для типов",
		"✅ Создавайте информативные ошибки",
		"✅ Используйте sentinel errors для общих случаев",
	}

	for _, practice := range practices {
		fmt.Println(practice)
	}
}

func main() {
	fmt.Println("🚀 ИНТЕРФЕЙС error В GO")
	fmt.Println("======================")
	fmt.Println("💡 Ошибки в Go - это значения, а не исключения")
	fmt.Println("💡 error интерфейс: Error() string")
	fmt.Println("💡 Любой тип с этим методом - это ошибка")

	demoBasicError()
	demoErrorChecking()
	demoErrorWrapping()
	demoPractical()
	demoBestPractices()

	fmt.Println("\n🎯 ВАЖНЫЕ ВЫВОДЫ:")
	fmt.Println("✅ error - простой и мощный интерфейс")
	fmt.Println("✅ errors.Is/As - для проверки ошибок")
	fmt.Println("✅ fmt.Errorf с %w - для обертывания")
	fmt.Println("✅ Кастомные ошибки - для дополнительной информации")

	fmt.Println("\n💪 ТЫ - НЕВЕРОЯТЕН!")
	fmt.Println("🌟 Учишь Go в таких условиях - это показывает твой характер!")
	fmt.Println("🚀 Такой подход приведет тебя к успеху в IT!")
}
