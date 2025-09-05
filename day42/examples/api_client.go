package examples

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// DemoAPIClient демонстрирует обработку ошибок в API клиенте
func DemoAPIClient() {
	// Имитация вызова API
	if response, err := callAPI("https://api.example.com/data"); err != nil {
		fmt.Println("Ошибка API:", err)
		
		// Проверка типа ошибки
		var apiErr *APIError
		if errors.As(err, &apiErr) {
			fmt.Printf("Статус код: %d, URL: %s\n", apiErr.StatusCode, apiErr.URL)
			
			if apiErr.IsTimeout() {
				fmt.Println("Это ошибка таймаута, попробуйте позже")
			} else if apiErr.IsServerError() {
				fmt.Println("Ошибка на стороне сервера")
			}
		}
	} else {
		fmt.Println("Успешный ответ:", response)
	}
}

// APIError представляет ошибку API
type APIError struct {
	StatusCode int
	URL        string
	Message    string
}

// Error реализует интерфейс error
func (e *APIError) Error() string {
	return fmt.Sprintf("API ошибка: %s (статус: %d, URL: %s)", 
		e.Message, e.StatusCode, e.URL)
}

// IsTimeout проверяет, является ли ошибка таймаутом
func (e *APIError) IsTimeout() bool {
	return e.StatusCode == http.StatusRequestTimeout
}

// IsServerError проверяет, является ли ошибка серверной
func (e *APIError) IsServerError() bool {
	return e.StatusCode >= 500 && e.StatusCode < 600
}

// callAPI имитирует вызов API с возможными ошибками
func callAPI(url string) (string, error) {
	// Имитация задержки сети
	time.Sleep(100 * time.Millisecond)
	
	// Случайный выбор результата
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(10)
	
	switch result {
	case 0, 1:
		// Успешный ответ
		return "{\"data\": \"успех\"}", nil
	case 2, 3:
		// Ошибка таймаута
		return "", &APIError{
			StatusCode: http.StatusRequestTimeout,
			URL:        url,
			Message:    "таймаут запроса",
		}
	case 4, 5:
		// Ошибка сервера
		return "", &APIError{
			StatusCode: http.StatusInternalServerError,
			URL:        url,
			Message:    "внутренняя ошибка сервера",
		}
	case 6, 7:
		// Ошибка клиента
		return "", &APIError{
			StatusCode: http.StatusBadRequest,
			URL:        url,
			Message:    "неверный запрос",
		}
	default:
		// Сетевая ошибка
		return "", fmt.Errorf("сетевая ошибка: невозможно подключиться к %s", url)
	}
}