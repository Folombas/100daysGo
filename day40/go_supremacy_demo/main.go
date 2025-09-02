package main

import (
	"fmt"
	"runtime"
	"time"
)

// Горутина для демонстрации конкурентности
func calculateFibonacci(n int, ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	ch <- a
}

// Структура для демонстрации статической типизации
type Developer struct {
	Name     string
	Language string
	Score    int
}

// Метод структуры
func (d *Developer) Celebrate() string {
	return fmt.Sprintf("Разработчик %s празднует превосходство %s!", d.Name, d.Language)
}

func main() {
	fmt.Println("🚀 День 40: Go vs Python - Превосходство в компиляции, типизации и кроссплатформенности")
	fmt.Println("================================================================================")

	// Демонстрация статической типизации
	fmt.Println("\n1. 🔥 СТАТИЧЕСКАЯ ТИПИЗАЦИЯ (ошибки обнаруживаются при компиляции)")
	
	// Попробуйте раскомментировать следующую строку - программа не скомпилируется!
	// var pythonDev Developer = "Несоответствие типов"
	
	goDev := Developer{
		Name:     "Гоша",
		Language: "Go",
		Score:    100,
	}
	fmt.Println(goDev.Celebrate())

	// Демонстрация производительности и конкурентности
	fmt.Println("\n2. ⚡ ПРОИЗВОДИТЕЛЬНОСТЬ И КОНКУРЕНТНОСТЬ")
	start := time.Now()
	
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go calculateFibonacci(1000, ch)
	}
	
	// Получаем результаты
	for i := 0; i < 1000; i++ {
		<-ch
	}
	
	elapsed := time.Since(start)
	fmt.Printf("1000 горутин обработаны за: %v\n", elapsed)

	// Демонстрация кроссплатформенности
	fmt.Println("\n3. 🌍 КРОССПЛАТФОРМЕННОСТЬ")
	fmt.Printf("ОС: %s\nАрхитектура: %s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println("Один исходный код → любой бинарный файл!")

	// Демонстрация встроенного тестирования и бенчмарков
	fmt.Println("\n4. 🧪 ВСТРОЕННОЕ ТЕСТИРОВАНИЕ И БЕНЧМАРКИ")
	fmt.Println("Запустите: go test -bench=. -benchmem")

	fmt.Println("\n5. 📦 ЕДИНЫЙ БИНАРНЫЙ ФАЙЛ")
	fmt.Println("Никаких зависимостей - один исполняемый файл содержит всё!")
	fmt.Printf("Размер бинарного файла ~5-10MB vs Python + виртуальное окружение + зависимости ~200-500MB\n")

	fmt.Println("\n🎉 Вывод: Go предлагает скорость компиляции, безопасность типов,")
	fmt.Println("   невероятную производительность и истинную кроссплатформенность!")
}