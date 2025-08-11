package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Автоматическая регистрация pprof
	"runtime"
	"sync"
	"time"
)

var (
	// Счетчик утекших горутин
	leakCounter int
	// Мьютекс для безопасного доступа к счетчику
	leakCounterMux sync.Mutex
	// Канал для остановки утечки
	stopLeak = make(chan struct{})
)

// leakingGoroutine имитирует утечку горутин
func leakingGoroutine(id int) {
	for {
		select {
		case <-stopLeak:
			// Выход при получении сигнала остановки
			return
		default:
			// Имитация работы (10 секунд сна)
			time.Sleep(10 * time.Second)

			// Безопасное увеличение счетчика
			leakCounterMux.Lock()
			leakCounter++
			leakCounterMux.Unlock()
		}
	}
}

// startLeakHandler запускает утечку горутин
func startLeakHandler(w http.ResponseWriter, r *http.Request) {
	// Запускаем 50 утекающих горутин
	for i := 0; i < 50; i++ {
		go leakingGoroutine(i)
	}
	w.Write([]byte("🟢 Запущено 50 утекающих горутин!"))
}

// stopLeakHandler останавливает утечку горутин
func stopLeakHandler(w http.ResponseWriter, r *http.Request) {
	close(stopLeak) // Закрытие канала останавливает все горутины
	w.Write([]byte("🔴 Утечка горутин остановлена!"))
}

// statusHandler показывает текущий статус
func statusHandler(w http.ResponseWriter, r *http.Request) {
	// Формируем HTML-страницу с информацией
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Утечка горутин</title>
    <style>
        body { 
            font-family: Arial, sans-serif; 
            background: #1a1a2e;
            color: #e6e6e6;
            padding: 20px;
            text-align: center;
        }
        .container {
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
            background: #16213e;
            border-radius: 15px;
            box-shadow: 0 0 20px rgba(0,0,0,0.5);
        }
        h1 { color: #4cc9f0; }
        .stats {
            display: flex;
            justify-content: space-around;
            margin: 30px 0;
        }
        .stat-box {
            background: #0f3460;
            padding: 20px;
            border-radius: 10px;
            width: 45%%;
        }
        .stat-value {
            font-size: 2.5rem;
            font-weight: bold;
            margin: 10px 0;
        }
        .leaking { color: #f05454; animation: pulse 1.5s infinite; }
        .normal { color: #16c79a; }
        .controls { margin: 30px 0; }
        button {
            padding: 15px 30px;
            margin: 0 10px;
            font-size: 1.1rem;
            border: none;
            border-radius: 50px;
            cursor: pointer;
            transition: all 0.3s;
            font-weight: bold;
        }
        .btn-start {
            background: linear-gradient(45deg, #ff416c, #ff4b2b);
            color: white;
        }
        .btn-stop {
            background: linear-gradient(45deg, #11998e, #38ef7d);
            color: white;
        }
        button:hover {
            transform: scale(1.05);
            box-shadow: 0 5px 15px rgba(0,0,0,0.3);
        }
        .instructions {
            text-align: left;
            background: rgba(255,255,255,0.05);
            padding: 20px;
            border-radius: 10px;
            margin: 20px 0;
        }
        @keyframes pulse {
            0%% { opacity: 1; }
            50%% { opacity: 0.7; }
            100%% { opacity: 1; }
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>День 16: Поиск утечек горутин</h1>
        
        <div class="stats">
            <div class="stat-box">
                <h2>Всего горутин</h2>
                <div class="stat-value leaking">%d</div>
                <p>runtime.NumGoroutine()</p>
            </div>
            <div class="stat-box">
                <h2>Утекших горутин</h2>
                <div class="stat-value leaking">%d</div>
                <p>Счетчик утечек</p>
            </div>
        </div>
        
        <div class="controls">
            <button class="btn-start" onclick="startLeak()">Запустить утечку</button>
            <button class="btn-stop" onclick="stopLeak()">Остановить утечку</button>
        </div>
        
        <div class="instructions">
            <h3>Как обнаружить утечку:</h3>
            <ol>
                <li>Нажмите "Запустить утечку"</li>
                <li>Наблюдайте рост числа горутин</li>
                <li>Соберите профиль: 
                    <code>go tool pprof http://localhost:8080/debug/pprof/goroutine</code>
                </li>
                <li>В pprof выполните:
                    <ul>
                        <li><code>top</code> - топ по количеству горутин</li>
                        <li><code>list leakingGoroutine</code> - найти проблемный код</li>
                        <li><code>web</code> - визуализация (требуется Graphviz)</li>
                    </ul>
                </li>
            </ol>
            <p><strong>Версия Go: %s</strong></p>
        </div>
        
        <p>Автообновление через 5 секунд...</p>
    </div>
    
    <script>
        function startLeak() {
            fetch('/start-leak')
                .then(response => response.text())
                .then(data => alert(data))
                .catch(err => alert('Ошибка: ' + err));
        }
        
        function stopLeak() {
            fetch('/stop-leak')
                .then(response => response.text())
                .then(data => alert(data))
                .catch(err => alert('Ошибка: ' + err));
        }
        
        // Автообновление страницы каждые 5 секунд
        setTimeout(() => {
            location.reload();
        }, 5000);
    </script>
</body>
</html>
`, runtime.NumGoroutine(), leakCounter, runtime.Version())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/", statusHandler)
	http.HandleFunc("/start-leak", startLeakHandler)
	http.HandleFunc("/stop-leak", stopLeakHandler)

	port := ":8080"
	log.Printf("Сервер запущен на http://localhost%s", port)
	log.Println("Для анализа утечек выполните:")
	log.Println("  go tool pprof http://localhost:8080/debug/pprof/goroutine")

	// Запуск HTTP сервера
	log.Fatal(http.ListenAndServe(port, nil))
}
