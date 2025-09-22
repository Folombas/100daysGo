package main

import "fmt"

// Speaker интерфейс для объектов, которые могут "говорить"
type Speaker interface {
	Speak() string
}

// Introduce представляет объекты, которые могут представиться
type Introduce interface {
	Speaker
	IntroduceYourself() string
}

// Реализация интерфейса Speaker для Employee
func (e Employee) Speak() string {
	return fmt.Sprintf("Привет, меня зовут %s и я %s", e.FirstName, e.Position)
}

// Реализация интерфейса Introduce для Employee
func (e Employee) IntroduceYourself() string {
	return fmt.Sprintf("Я %s %s, работаю на позиции %s", e.FirstName, e.LastName, e.Position)
}

// Реализация для Manager
func (m Manager) IntroduceYourself() string {
	return fmt.Sprintf("Я %s %s, руковожу департаментом %s", m.FirstName, m.LastName, m.Department)
}

// ProcessSpeaker обрабатывает объект, реализующий интерфейс Speaker
func ProcessSpeaker(s Speaker) {
	fmt.Println(s.Speak())
}

// ProcessIntroduce обрабатывает объект, реализующий интерфейс Introduce
func ProcessIntroduce(i Introduce) {
	fmt.Println(i.IntroduceYourself())
}

func demoInterfaces() {
	emp := Employee{
		ID:        5,
		FirstName: "Дмитрий",
		LastName:  "Смирнов",
		Position:  "Аналитик",
		Salary:    120000,
	}

	manager := Manager{
		Employee: Employee{
			ID:        6,
			FirstName: "Ольга",
			LastName:  "Васильева",
			Position:  "Руководитель отдела",
			Salary:    280000,
		},
		Department: "Аналитика",
		TeamSize:   8,
	}

	fmt.Println("Демонстрация интерфейсов:")
	ProcessSpeaker(emp)
	ProcessIntroduce(emp)
	ProcessIntroduce(manager)
}
