package main

import (
	"html/template"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync"
	"time"
)

var (
	leakCounter    int
	leakCounterMux sync.Mutex
	stopLeak       = make(chan struct{})
)

// Утекающая горутина
func leakingGoroutine(id int) {
	for {
		select {
		case <-stopLeak:
			return
		default:
			// Имитация работы
			time.Sleep(10 * time.Second)
			
			// Увеличиваем счетчик утечек
			leakCounterMux.Lock()
			leakCounter++
			leakCounterMux.Unlock()
		}
	}
}

// Эндпоинт для запуска утечки
func startLeakHandler(w http.ResponseWriter, r *http.Request) {
	count := 50
	for i := 0; i < count; i++ {
		go leakingGoroutine(i)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Запущено 50 утекающих горутин!"))
}

// Эндпоинт для остановки утечки
func stopLeakHandler(w http.ResponseWriter, r *http.Request) {
	close(stopLeak)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("🛑 Утечка горутин остановлена!"))
}

// Данные для шаблона
type PageData struct {
	Title        string
	GoVersion    string
	NumGoroutine int
	LeakCount    int
}

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/start-leak", startLeakHandler)
	http.HandleFunc("/stop-leak", stopLeakHandler)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		
		// Собираем текущую статистику
		data := PageData{
			Title:        "День 16: Поиск утечек горутин",
			GoVersion:    runtime.Version(),
			NumGoroutine: runtime.NumGoroutine(),
			LeakCount:    leakCounter,
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
		}
	})

	port := ":8080"
	log.Printf("🚀 Сервер запущен на http://localhost%s", port)
	log.Println("🔍 Для анализа утечек выполните:")
	log.Println("   go tool pprof http://localhost:8080/debug/pprof/goroutine")
	log.Fatal(http.ListenAndServe(port, nil))
}