package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIClient(t *testing.T) {
	// Создаем тестовый сервер
	handler := NewAPIHandler()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/user":
			if r.Method == http.MethodGet {
				handler.GetUserHandler(w, r)
			} else if r.Method == http.MethodPost {
				handler.CreateUserHandler(w, r)
			}
		case "/users":
			handler.GetAllUsersHandler(w, r)
		}
	}))
	defer server.Close()

	client := NewAPIClient(server.URL)

	t.Run("клиент получает пользователя", func(t *testing.T) {
		user, err := client.GetUser(1)
		if err != nil {
			t.Fatalf("Ошибка получения пользователя: %v", err)
		}

		if user.Name != "Гоша" {
			t.Errorf("Ожидалось имя 'Гоша', получено '%s'", user.Name)
		}
	})

	t.Run("клиент создает пользователя", func(t *testing.T) {
		user, err := client.CreateUser("Тестовый", "test@example.com", 30)
		if err != nil {
			t.Fatalf("Ошибка создания пользователя: %v", err)
		}

		if user.Name != "Тестовый" {
			t.Errorf("Ожидалось имя 'Тестовый', получено '%s'", user.Name)
		}
	})

	t.Run("клиент получает всех пользователей", func(t *testing.T) {
		users, err := client.GetAllUsers()
		if err != nil {
			t.Fatalf("Ошибка получения пользователей: %v", err)
		}

		if len(users) < 2 {
			t.Errorf("Ожидалось минимум 2 пользователя, получено %d", len(users))
		}
	})
}

func TestAPIClient_ErrorCases(t *testing.T) {
	// Сервер возвращающий ошибки
	errorServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Внутренняя ошибка сервера"})
	}))
	defer errorServer.Close()

	client := NewAPIClient(errorServer.URL)

	t.Run("клиент обрабатывает ошибки сервера", func(t *testing.T) {
		_, err := client.GetUser(1)
		if err == nil {
			t.Error("Ожидалась ошибка, но получен nil")
		}
	})
}

// Benchmark тесты
func BenchmarkGetUserHandler(b *testing.B) {
	handler := NewAPIHandler()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "/user?id=1", nil)
		rr := httptest.NewRecorder()
		handler.GetUserHandler(rr, req)
	}
}
