package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Автоматически регистрирует обработчики pprof
)

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/cpu", cpuLoadHandler)
	http.HandleFunc("/mem", memLoadHandler)
	http.HandleFunc("/goroutines", goroutineLoadHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	fmt.Println("Доступные эндпоинты:")
	fmt.Println("  /              - Главная страница")
	fmt.Println("  /cpu           - Нагрузка CPU")
	fmt.Println("  /mem           - Нагрузка памяти")
	fmt.Println("  /goroutines    - Создание горутин")
	fmt.Println("\nПрофилирование:")
	fmt.Println("  /debug/pprof   - Веб-интерфейс pprof")
	fmt.Println("\nПримеры команд для профилирования:")
	fmt.Println("  go tool pprof http://localhost:8080/debug/pprof/profile")
	fmt.Println("  go tool pprof http://localhost:8080/debug/pprof/heap")
	fmt.Println("  go tool pprof http://localhost:8080/debug/pprof/goroutine")

	log.Fatal(http.ListenAndServe(":8080", nil))
}