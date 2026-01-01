package main

import (
	"fmt"
	"strings"
)

// Shape - интерфейс для геометрических фигур
type Shape interface {
	Area() float64    // Площадь фигуры
	Describe() string // Описание фигуры
}

// Rectangle - структура прямоугольника
type Rectangle struct {
	Width, Height float64
}

// Circle - структура круга
type Circle struct {
	Radius float64
}

// Value receiver для Rectangle: вычисляет площадь
// (не изменяет исходный объект, работает с копией)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Value receiver для Rectangle: генерирует описание
func (r Rectangle) Describe() string {
	return fmt.Sprintf("прямоугольник %dx%d (площадь: %.1f)",
		int(r.Width), int(r.Height), r.Area())
}

// Value receiver для Circle: вычисляет площадь
func (c Circle) Area() float64 {
	return 3.141592 * c.Radius * c.Radius
}

// Value receiver для Circle: генерирует описание
func (c Circle) Describe() string {
	return fmt.Sprintf("круг радиуса %.1f (площадь: %.1f)",
		c.Radius, c.Area())
}

// Ключевое отличие Value vs Pointer Receivers:
// Метод с value receiver НЕ может изменить исходную структуру
func (r Rectangle) Scale(factor float64) {
	r.Width *= factor // Изменится только копия!
	r.Height *= factor
}

func main() {
	// Инициализация фигур
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7.5}

	// Демонстрация работы с интерфейсом
	shapes := []Shape{rect, circle}

	fmt.Println("🔥 ДЕНЬ 59: VALUE RECEIVERS В ДЕЙСТВИИ 🔥")
	fmt.Println(strings.Repeat("=", 45))

	fmt.Println("\n🎯 ПОЧЕМУ ЭТО ВАЖНО:")
	fmt.Println("- Value receivers НЕ изменяют исходный объект")
	fmt.Println("- Методы с value receivers работают с КОПИЕЙ структуры")
	fmt.Println("- Такие методы могут вызываться как на значении, так и на указателе")

	fmt.Println("\n📐 РАБОТА С ИНТЕРФЕЙСОМ Shape:")
	for _, s := range shapes {
		fmt.Printf("• %s\n", s.Describe())
	}

	fmt.Println("\n⚠️ ОСОБЕННОСТЬ VALUE RECEIVERS:")
	originalRect := rect
	rect.Scale(2) // Попытка изменения через value receiver

	fmt.Printf("Исходный прямоугольник: %s\n", originalRect.Describe())
	fmt.Printf("После Scale(2): %s\n", rect.Describe())
	fmt.Println("→ Изменений НЕТ! Value receiver работает с копией")

	// Практическое применение: безопасное вычисление статистики
	fmt.Println("\n💡 РЕАЛЬНЫЙ КЕЙС В 2026:")
	courierStats := Rectangle{Width: 150, Height: 200} // ширина=заказы/день, высота=км
	fmt.Printf("Статистика курьера (безопасно): %s\n",
		courierStats.Describe())

	fmt.Println("\n🚀 СЛЕДУЮЩИЙ УРОВЕНЬ (Day60):")
	fmt.Println("→ Изучим POINTER RECEIVERS: когда НУЖНО изменять исходный объект")
	fmt.Println("→ Как правильно комбинировать value и pointer receivers")

	fmt.Println("\n✅ УРОК ДНЯ:")
	fmt.Println("«Value receivers — твой инструмент для ИДЕНТИФИКАЦИИ и БЕЗОПАСНЫХ вычислений.")
	fmt.Println("Когда нужно ИЗМЕНЯТЬ состояние — используй pointer receivers»")
}
