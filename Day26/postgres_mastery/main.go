package main

import (
	"fmt"
	"log"
)

func main() {
	// Инициализация БД
	err := InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}
	defer CloseDB()

	// Создание таблиц
	err = CreateTables()
	if err != nil {
		log.Fatal("Ошибка создания таблиц:", err)
	}

	// Работа с пользователями
	userRepo := UserRepository{}
	productRepo := ProductRepository{}

	// Создание пользователя
	newUser := User{Name: "Гоша Гошник", Email: "gosha@gofer.com"}
	err = userRepo.CreateUser(&newUser)
	if err != nil {
		log.Fatal("Ошибка создания пользователя:", err)
	}
	fmt.Printf("Создан пользователь: %+v\n", newUser)

	// Создание товара
	newProduct := Product{
		Name:   "Ноутбук",
		Price:  999.99,
		UserID: newUser.ID,
	}
	err = productRepo.CreateProduct(&newProduct)
	if err != nil {
		log.Fatal("Ошибка создания товара:", err)
	}
	fmt.Printf("Создан товар: %+v\n", newProduct)

	// Получение пользователя с товарами
	userWithProducts, err := userRepo.GetUserWithProducts(newUser.ID)
	if err != nil {
		log.Fatal("Ошибка получения данных:", err)
	}

	fmt.Printf("\nПользователь с товарами:\n")
	fmt.Printf("User: %s (%s)\n", userWithProducts.User.Name, userWithProducts.User.Email)
	fmt.Printf("Товары:\n")
	for _, product := range userWithProducts.Products {
		fmt.Printf("  - %s: $%.2f\n", product.Name, product.Price)
	}

	// Получение всех пользователей
	users, err := userRepo.GetAllUsers()
	if err != nil {
		log.Fatal("Ошибка получения пользователей:", err)
	}

	fmt.Printf("\nВсе пользователи в базе:\n")
	for _, user := range users {
		fmt.Printf("- %s (%s)\n", user.Name, user.Email)
	}
}