package main

import (
	//"fmt"
	//"log"
)

// UserRepository операции с пользователями
type UserRepository struct{}

// CreateUser создает нового пользователя
func (r *UserRepository) CreateUser(user *User) error {
	query := `
		INSERT INTO users (name, email) 
		VALUES ($1, $2) 
		RETURNING id, created_at, updated_at
	`
	
	err := db.QueryRow(query, user.Name, user.Email).
		Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
		
	return err
}

// GetUserByID возвращает пользователя по ID
func (r *UserRepository) GetUserByID(id int) (*User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		WHERE id = $1
	`
	
	user := &User{}
	err := db.QueryRow(query, id).
		Scan(&user.ID, &user.Name, &user.Email, 
			&user.CreatedAt, &user.UpdatedAt)
			
	return user, err
}

// GetAllUsers возвращает всех пользователей
func (r *UserRepository) GetAllUsers() ([]User, error) {
	query := `
		SELECT id, name, email, created_at, updated_at 
		FROM users 
		ORDER BY id
	`
	
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email,
			&user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// UpdateUser обновляет данные пользователя
func (r *UserRepository) UpdateUser(user *User) error {
	query := `
		UPDATE users 
		SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP 
		WHERE id = $3
		RETURNING updated_at
	`
	
	return db.QueryRow(query, user.Name, user.Email, user.ID).
		Scan(&user.UpdatedAt)
}

// DeleteUser удаляет пользователя
func (r *UserRepository) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	return err
}

// ProductRepository операции с товарами
type ProductRepository struct{}

// CreateProduct создает новый товар
func (r *ProductRepository) CreateProduct(product *Product) error {
	query := `
		INSERT INTO products (name, price, user_id) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at
	`
	
	return db.QueryRow(query, product.Name, product.Price, product.UserID).
		Scan(&product.ID, &product.CreatedAt)
}

// GetUserWithProducts возвращает пользователя с товарами
func (r *UserRepository) GetUserWithProducts(userID int) (*UserWithProducts, error) {
	// Начинаем транзакцию
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Получаем пользователя
	user := &User{}
	err = tx.QueryRow(`
		SELECT id, name, email, created_at, updated_at 
		FROM users WHERE id = $1
	`, userID).Scan(&user.ID, &user.Name, &user.Email, 
		&user.CreatedAt, &user.UpdatedAt)
		
	if err != nil {
		return nil, err
	}

	// Получаем товары пользователя
	rows, err := tx.Query(`
		SELECT id, name, price, created_at 
		FROM products 
		WHERE user_id = $1
	`, userID)
		
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, 
			&product.Price, &product.CreatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	// Коммитим транзакцию
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &UserWithProducts{
		User:     *user,
		Products: products,
	}, nil
}