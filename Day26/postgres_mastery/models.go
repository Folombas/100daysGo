package main

import (
	"time"
)

// User модель пользователя
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Product модель товара
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

// UserWithProducts пользователь с товарами
type UserWithProducts struct {
	User     User      `json:"user"`
	Products []Product `json:"products"`
}