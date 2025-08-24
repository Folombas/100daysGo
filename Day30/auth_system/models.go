package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"sync"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User представляет модель пользователя
type User struct {
	Username string
	Password string
	Email    string
}

// Session представляет модель сессии
type Session struct {
	Username string
	Expiry   time.Time
}

var (
	users    = make(map[string]*User)
	sessions = make(map[string]*Session)
	mu       sync.RWMutex
)

// InitUserStorage инициализирует хранилище пользователей
func InitUserStorage() {
	// Добавляем тестового пользователя
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	users["testuser"] = &User{
		Username: "testuser",
		Password: string(hashedPassword),
		Email:    "test@example.com",
	}
}

// CreateUser создает нового пользователя
func CreateUser(username, password, email string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := users[username]; exists {
		return fmt.Errorf("пользователь уже существует")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	users[username] = &User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	return nil
}

// AuthenticateUser проверяет учетные данные пользователя
func AuthenticateUser(username, password string) bool {
	mu.RLock()
	defer mu.RUnlock()

	user, exists := users[username]
	if !exists {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

// CreateSession создает новую сессию
func CreateSession(username string) (string, error) {
	mu.Lock()
	defer mu.Unlock()

	// Генерируем случайный ID сессии
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	sessionID := hex.EncodeToString(bytes)

	sessions[sessionID] = &Session{
		Username: username,
		Expiry:   time.Now().Add(24 * time.Hour),
	}

	return sessionID, nil
}

// GetSession возвращает сессию по ID
func GetSession(sessionID string) (*Session, bool) {
	mu.RLock()
	defer mu.RUnlock()

	session, exists := sessions[sessionID]
	if !exists || time.Now().After(session.Expiry) {
		return nil, false
	}

	return session, true
}

// DeleteSession удаляет сессию
func DeleteSession(sessionID string) {
	mu.Lock()
	defer mu.Unlock()
	delete(sessions, sessionID)
}

// GetUserBySession возвращает пользователя по ID сессии
func GetUserBySession(sessionID string) (*User, bool) {
	session, exists := GetSession(sessionID)
	if !exists {
		return nil, false
	}

	mu.RLock()
	defer mu.RUnlock()
	user, exists := users[session.Username]
	return user, exists
}

// Middleware для проверки аутентификации
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем куки
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Получаем пользователя по сессии
		user, exists := GetUserBySession(cookie.Value)
		if !exists {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		// Добавляем пользователя в контекст
		ctx := context.WithValue(r.Context(), "user", user)
		next(w, r.WithContext(ctx))
	}
}