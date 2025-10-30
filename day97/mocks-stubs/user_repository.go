package main

import "errors"

// User модель пользователя
type User struct {
    ID    string
    Name  string
    Email string
    Balance float64
}

// UserRepository интерфейс для работы с пользователями
type UserRepository interface {
    FindByID(id string) (*User, error)
    Save(user *User) error
    UpdateBalance(userID string, newBalance float64) error
}

// RealUserRepository реальная реализация репозитория пользователей
type RealUserRepository struct {
    // В реальной реализации здесь были бы подключения к БД
}

func NewRealUserRepository() *RealUserRepository {
    return &RealUserRepository{}
}

func (r *RealUserRepository) FindByID(id string) (*User, error) {
    // Симуляция запроса к базе данных
    if id == "user_123" {
        return &User{
            ID:      "user_123",
            Name:    "Гоша",
            Email:   "gosha@example.com",
            Balance: 1000.0,
        }, nil
    }
    return nil, errors.New("пользователь не найден")
}

func (r *RealUserRepository) Save(user *User) error {
    // Симуляция сохранения в базу данных
    return nil
}

func (r *RealUserRepository) UpdateBalance(userID string, newBalance float64) error {
    // Симуляция обновления баланса
    return nil
}
