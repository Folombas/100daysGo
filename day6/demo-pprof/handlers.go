package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// homeHandler - главная страница с инструкцией
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>pprof Demo - 100 Days of Go</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 800px; margin: 40px auto; }
		.card { padding: 20px; margin: 20px 0; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); }
		.cpu { background-color: #ffebee; }
		.mem { background-color: #e3f2fd; }
		.goroutines { background-color: #e8f5e9; }
		.profiling { background-color: #fff8e1; }
		a { color: #1a73e8; text-decoration: none; }
		button { padding: 10px 15px; background: #1a73e8; color: white; border: none; border-radius: 4px; cursor: pointer; }
	</style>
</head>
<body>
	<h1>🔥 Профилирование Go-приложений с pprof</h1>
	
	<div class="card cpu">
		<h2>Тест CPU</h2>
		<p>Создает нагрузку на процессор вычислениями</p>
		<button onclick="load('cpu')">Запустить тест CPU</button>
	</div>
	
	<div class="card mem">
		<h2>Тест памяти</h2>
		<p>Аллоцирует 500MB памяти в куче</p>
		<button onclick="load('mem')">Запустить тест памяти</button>
	</div>
	
	<div class="card goroutines">
		<h2>Тест горутин</h2>
		<p>Создает 1000 спящих горутин</p>
		<button onclick="load('goroutines')">Запустить тест горутин</button>
	</div>
	
	<div class="card profiling">
		<h2>Профилирование</h2>
		<p><a href="/debug/pprof" target="_blank">Веб-интерфейс pprof</a></p>
		<p>Примеры команд:</p>
		<ul>
			<li>CPU: <code>go tool pprof http://localhost:8080/debug/pprof/profile</code></li>
			<li>Память: <code>go tool pprof http://localhost:8080/debug/pprof/heap</code></li>
			<li>Горутины: <code>go tool pprof http://localhost:8080/debug/pprof/goroutine</code></li>
		</ul>
	</div>
	
	<script>
		function load(endpoint) {
			fetch('/' + endpoint)
				.then(res => res.text())
				.then(data => alert(data))
				.catch(err => console.error(err));
		}
	</script>
</body>
</html>
`
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, html)
}

// memLoadHandler - создает нагрузку на память
func memLoadHandler(w http.ResponseWriter, r *http.Request) {
	// Аллоцируем 500MB
	data := make([]byte, 500*1024*1024)
	for i := range data {
		data[i] = byte(i % 256)
	}
	fmt.Fprintf(w, "Выделено 500MB памяти!")
}

// goroutineLoadHandler - создает множество горутин
func goroutineLoadHandler(w http.ResponseWriter, r *http.Request) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(5 * time.Minute) // Горутина живет 5 минут
		}(i)
	}
	wg.Wait()
	fmt.Fprintf(w, "Создано 1000 горутин!")
}