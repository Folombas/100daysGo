package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// User представляет модель пользователя
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// UserStorage имитирует базу данных пользователей
type UserStorage struct {
	users map[int]User
	nextID int
}

func NewUserStorage() *UserStorage {
	return &UserStorage{
		users: map[int]User{
			1: {ID: 1, Name: "Гоша", Email: "gosha@example.com", Age: 37},
			2: {ID: 2, Name: "Мария", Email: "maria@example.com", Age: 28},
		},
		nextID: 3,
	}
}

// GetUser возвращает пользователя по ID
func (s *UserStorage) GetUser(id int) (*User, error) {
	user, exists := s.users[id]
	if !exists {
		return nil, fmt.Errorf("пользователь с ID %d не найден", id)
	}
	return &user, nil
}

// CreateUser создает нового пользователя
func (s *UserStorage) CreateUser(name, email string, age int) *User {
	user := User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
		Age:   age,
	}
	s.users[s.nextID] = user
	s.nextID++
	return &user
}

// GetAllUsers возвращает всех пользователей
func (s *UserStorage) GetAllUsers() []User {
	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// APIHandler обработчик HTTP запросов
type APIHandler struct {
	storage *UserStorage
}

func NewAPIHandler() *APIHandler {
	return &APIHandler{
		storage: NewUserStorage(),
	}
}

// GetUserHandler обработчик для получения пользователя
func (h *APIHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "Параметр id обязателен", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID", http.StatusBadRequest)
		return
	}

	user, err := h.storage.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// CreateUserHandler обработчик для создания пользователя
func (h *APIHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var user struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Некорректный JSON", http.StatusBadRequest)
		return
	}

	if user.Name == "" || user.Email == "" {
		http.Error(w, "Имя и email обязательны", http.StatusBadRequest)
		return
	}

	newUser := h.storage.CreateUser(user.Name, user.Email, user.Age)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// GetAllUsersHandler обработчик для получения всех пользователей
func (h *APIHandler) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	users := h.storage.GetAllUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
