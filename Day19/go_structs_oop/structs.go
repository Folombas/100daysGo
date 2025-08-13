package main

import "fmt"

// Person - базовая структура (аналог класса)
type Person struct {
    name string
    age  int
}

// NewPerson - конструктор
func NewPerson(name string, age int) *Person {
    return &Person{name: name, age: age}
}

// Метод с получателем по значению
func (p Person) Introduce() string {
    return fmt.Sprintf("Меня зовут %s, мне %d лет", p.name, p.age)
}

// Метод с получателем по указателю (может изменять структуру)
func (p *Person) Birthday() {
    p.age++
}