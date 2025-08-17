package main

// Структура статьи
type Article struct {
    ID      string `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

// Хранилище в памяти
var articles = []Article{
    {"1", "Привет, Go!", "Go - отличный язык для API"},
    {"2", "REST за 5 минут", "Создаём API с помощью gorilla/mux"},
}