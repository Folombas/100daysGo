package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof" // Автоматическая регистрация pprof
	//"runtime"
	"time"
)

// Генерация случайной матрицы
func generateMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)
	for i := range matrix {
		matrix[i] = make([]float64, n)
		for j := range matrix[i] {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}

// Перемножение матриц (CPU-intensive)
func multiplyMatrices(a, b [][]float64) [][]float64 {
	n := len(a)
	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

// Горутина для создания нагрузки
func createLoad() {
	for {
		// CPU нагрузка
		a := generateMatrix(100)
		b := generateMatrix(100)
		_ = multiplyMatrices(a, b)

		// Память (аллокации)
		_ = make([]byte, 10<<20) // 10 MB

		// Задержка
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	// Запускаем 4 горутины, создающие нагрузку
	for i := 0; i < 4; i++ {
		go createLoad()
	}

	// Настройка HTTP сервера
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
		<h1>День 15: Профилирование с pprof</h1>
		<ul>
			<li><a href="/debug/pprof/">Главная pprof</a></li>
			<li><a href="/debug/pprof/heap">Heap (память)</a></li>
			<li><a href="/debug/pprof/profile?seconds=10">CPU профиль (10 сек)</a></li>
			<li><a href="/debug/pprof/goroutine">Горутины</a></li>
			<li><a href="/debug/pprof/block">Блокировки</a></li>
		</ul>
		`)
	})

	port := ":8080"
	fmt.Printf("Сервер запущен на http://localhost%s\n", port)
	fmt.Println("Доступные профили:")
	fmt.Printf("  CPU:  go tool pprof http://localhost%s/debug/pprof/profile?seconds=30\n", port)
	fmt.Printf("  Память: go tool pprof http://localhost%s/debug/pprof/heap\n", port)
	fmt.Printf("  Горутины: go tool pprof http://localhost%s/debug/pprof/goroutine\n", port)

	log.Fatal(http.ListenAndServe(port, nil))
}