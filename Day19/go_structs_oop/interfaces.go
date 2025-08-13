package main

import "fmt"

// Speaker - интерфейс (реализуется неявно)
type Speaker interface {
    Speak() string
}

// Animal - независимая структура
type Animal struct {
    species string
}

func (a Animal) Speak() string {
    return fmt.Sprintf("Я %s! Гав!", a.species)
}

// Полиморфная функция
func MakeSound(s Speaker) string {
    return s.Speak()
}

// Реализация интерфейса для Person
func (p Person) Speak() string {
    return "Привет, я человек!"
}