package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetUserHandler(t *testing.T) {
	handler := NewAPIHandler()

	t.Run("успешное получение пользователя", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user?id=1", nil)
		rr := httptest.NewRecorder()

		handler.GetUserHandler(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Ожидался статус 200, получен %d", rr.Code)
		}

		var user User
		if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
			t.Fatalf("Невозможно распарсить ответ: %v", err)
		}

		if user.Name != "Гоша" {
			t.Errorf("Ожидалось имя 'Гоша', получено '%s'", user.Name)
		}
	})

	t.Run("пользователь не найден", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user?id=999", nil)
		rr := httptest.NewRecorder()

		handler.GetUserHandler(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Ожидался статус 404, получен %d", rr.Code)
		}
	})

	t.Run("отсутствует параметр id", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user", nil)
		rr := httptest.NewRecorder()

		handler.GetUserHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Ожидался статус 400, получен %d", rr.Code)
		}
	})

	t.Run("некорректный метод", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/user?id=1", nil)
		rr := httptest.NewRecorder()

		handler.GetUserHandler(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("Ожидался статус 405, получен %d", rr.Code)
		}
	})
}

func TestCreateUserHandler(t *testing.T) {
	handler := NewAPIHandler()

	t.Run("успешное создание пользователя", func(t *testing.T) {
		userJSON := `{"name": "Новый пользователь", "email": "new@example.com", "age": 25}`
		req := httptest.NewRequest("POST", "/user", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler.CreateUserHandler(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("Ожидался статус 201, получен %d", rr.Code)
		}

		var user User
		if err := json.Unmarshal(rr.Body.Bytes(), &user); err != nil {
			t.Fatalf("Невозможно распарсить ответ: %v", err)
		}

		if user.Name != "Новый пользователь" {
			t.Errorf("Ожидалось имя 'Новый пользователь', получено '%s'", user.Name)
		}
	})

	t.Run("некорректный JSON", func(t *testing.T) {
		invalidJSON := `{"name": "Тест", "email":}`
		req := httptest.NewRequest("POST", "/user", strings.NewReader(invalidJSON))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler.CreateUserHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Ожидался статус 400, получен %d", rr.Code)
		}
	})

	t.Run("отсутствуют обязательные поля", func(t *testing.T) {
		userJSON := `{"email": "test@example.com"}`
		req := httptest.NewRequest("POST", "/user", strings.NewReader(userJSON))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler.CreateUserHandler(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Ожидался статус 400, получен %d", rr.Code)
		}
	})
}

func TestGetAllUsersHandler(t *testing.T) {
	handler := NewAPIHandler()

	req := httptest.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()

	handler.GetAllUsersHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Ожидался статус 200, получен %d", rr.Code)
	}

	var users []User
	if err := json.Unmarshal(rr.Body.Bytes(), &users); err != nil {
		t.Fatalf("Невозможно распарсить ответ: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Ожидалось 2 пользователя, получено %d", len(users))
	}
}
