package main

import (
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"
)

// setUTF8Header устанавливает заголовок Content-Type с указанным типом и кодировкой UTF-8
func setUTF8Header(w http.ResponseWriter, contentType string) {
	w.Header().Set("Content-Type", contentType+"; charset=utf-8")
}

// homeHandler - главная страница с инструкцией
func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<meta charset="UTF-8">
	<title>pprof Demo - 100 Days of Go</title>
	<style>
		body { font-family: Arial, sans-serif; max-width: 800px; margin: 40px auto; line-height: 1.6; }
		.card { padding: 20px; margin: 20px 0; border-radius: 8px; box-shadow: 0 4px 6px rgba(0,0,0,0.1); }
		.cpu { background-color: #ffebee; }
		.mem { background-color: #e3f2fd; }
		.goroutines { background-color: #e8f5e9; }
		.profiling { background-color: #fff8e1; }
		a { color: #1a73e8; text-decoration: none; font-weight: bold; }
		button { padding: 12px 18px; background: #1a73e8; color: white; border: none; 
                border-radius: 4px; cursor: pointer; font-size: 16px; margin: 5px 0; 
                transition: background 0.3s; }
		button:hover { background: #0d62c9; }
		code { background: #f5f5f5; padding: 3px 6px; border-radius: 3px; font-family: monospace; }
		.container { display: flex; flex-wrap: wrap; justify-content: space-between; }
		.col { flex: 1; min-width: 300px; padding: 10px; }
		h1 { color: #1a237e; }
		h2 { color: #283593; border-bottom: 2px solid #5c6bc0; padding-bottom: 5px; }
		ul { padding-left: 20px; }
		li { margin: 8px 0; }
		.footer { text-align: center; margin-top: 30px; color: #666; }
	</style>
</head>
<body>
	<h1>🔥 Профилирование Go-приложений с pprof</h1>
	
	<div class="container">
		<div class="col">
			<div class="card cpu">
				<h2>Тест CPU</h2>
				<p>Создает интенсивную нагрузку на процессор сложными вычислениями</p>
				<button onclick="load('cpu')">Запустить тест CPU</button>
				<p><small>Выполняет 50 миллионов итераций вычислений</small></p>
			</div>
			
			<div class="card mem">
				<h2>Тест памяти</h2>
				<p>Аллоцирует 500MB памяти в куче</p>
				<button onclick="load('mem')">Запустить тест памяти</button>
				<p><small>Создает большой массив байтов</small></p>
			</div>
		</div>
		
		<div class="col">
			<div class="card goroutines">
				<h2>Тест горутин</h2>
				<p>Создает 5000 спящих горутин</p>
				<button onclick="load('goroutines')">Запустить тест горутин</button>
				<p><small>Горутины будут работать 5 минут</small></p>
			</div>
			
			<div class="card profiling">
				<h2>Профилирование</h2>
				<p><a href="/debug/pprof" target="_blank">Веб-интерфейс pprof</a></p>
				<p>Примеры команд для анализа:</p>
				<ul>
					<li>CPU: <code>go tool pprof -http=:8081 http://localhost:8080/debug/pprof/profile</code></li>
					<li>Память: <code>go tool pprof -http=:8081 http://localhost:8080/debug/pprof/heap</code></li>
					<li>Горутины: <code>go tool pprof -http=:8081 http://localhost:8080/debug/pprof/goroutine</code></li>
					<li>Блокировки: <code>go tool pprof -http=:8081 http://localhost:8080/debug/pprof/block</code></li>
				</ul>
				<p>Для просмотра в браузере добавьте флаг <code>-http=:8081</code></p>
			</div>
		</div>
	</div>
	
	<div class="footer">
		<p>Демо профилирования для айти-марафона "100 Days of Go" | Сегодня: ` + time.Now().Format("2006-01-02") + `</p>
		<p><a href="https://github.com/Folombas/100daysGo" target="_blank">100 Days of Go</a></p>
	</div>
	
	<script>
		function load(endpoint) {
			const button = event.target;
			const originalText = button.textContent;
			
			button.textContent = "Выполняется...";
			button.disabled = true;
			
			fetch('/' + endpoint)
				.then(res => res.text())
				.then(data => {
					alert(data);
					button.textContent = originalText;
					button.disabled = false;
				})
				.catch(err => {
					alert("Ошибка: " + err);
					button.textContent = originalText;
					button.disabled = false;
				});
		}
	</script>
</body>
</html>
`
	setUTF8Header(w, "text/html")
	fmt.Fprint(w, html)
}

// cpuLoadHandler - создает интенсивную нагрузку на CPU
func cpuLoadHandler(w http.ResponseWriter, r *http.Request) {
	total := 0.0
	const iterations = 50000000 // 50 миллионов итераций

	start := time.Now()
	for i := 0; i < iterations; i++ {
		// Сложные вычисления для создания нагрузки
		val := math.Pow(math.Sin(float64(i)), math.Cos(float64(i)))
		total += val
	}
	duration := time.Since(start)

	setUTF8Header(w, "text/plain")
	fmt.Fprintf(w, "Интенсивная нагрузка CPU завершена!\nИтераций: %d\nВремя выполнения: %s", iterations, duration)
}

// memLoadHandler - создает нагрузку на память
func memLoadHandler(w http.ResponseWriter, r *http.Request) {
	const megaBytes = 500
	const bytes = megaBytes * 1024 * 1024

	// Создаем большой массив
	data := make([]byte, bytes)
	for i := range data {
		data[i] = byte(i % 256)
	}

	// Предотвращаем оптимизацию
	result := data[len(data)-1]

	setUTF8Header(w, "text/plain")
	fmt.Fprintf(w, "Выделено %dMB памяти!\nПоследний байт: %d", megaBytes, result)
}

// goroutineLoadHandler - создает множество горутин
func goroutineLoadHandler(w http.ResponseWriter, r *http.Request) {
	const numGoroutines = 5000
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(5 * time.Minute)
		}(i)
	}

	setUTF8Header(w, "text/plain")
	fmt.Fprintf(w, "Создано %d горутин! Они завершатся через 5 минут.", numGoroutines)
}
