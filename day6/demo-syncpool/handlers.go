package main

import (
	"encoding/json"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// Тяжелый объект для демонстрации (увеличенный размер)
type HeavyObject struct {
	ID      int
	Content [40960]byte // 40KB данных
}

// Результаты тестирования
type TestResult struct {
	Duration time.Duration `json:"duration"`
	Allocs   uint64        `json:"allocs"`
	Memory   uint64        `json:"memory"`
	WithPool bool          `json:"with_pool"`
}

var pool = sync.Pool{
	New: func() interface{} {
		return new(HeavyObject)
	},
}

const maxWorkers = 1000

// Главная страница (с встроенным HTML)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса главной страницы")

	html := `<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <title>Sync.Pool Демо</title>
    <style>
        body {
            font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
            background-color: #f5f7fa;
            margin: 0;
            padding: 20px;
            color: #333;
        }
        
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            padding: 30px;
            border-radius: 15px;
            box-shadow: 0 5px 25px rgba(0,0,0,0.1);
        }
        
        h1 {
            color: #2c3e50;
            text-align: center;
            margin-bottom: 30px;
        }
        
        .controls {
            display: flex;
            justify-content: center;
            gap: 20px;
            margin-bottom: 40px;
        }
        
        button {
            background: #3498db;
            color: white;
            border: none;
            padding: 15px 30px;
            font-size: 18px;
            border-radius: 8px;
            cursor: pointer;
            transition: all 0.3s;
        }
        
        button:hover {
            background: #2980b9;
            transform: translateY(-3px);
            box-shadow: 0 5px 15px rgba(0,0,0,0.1);
        }
        
        .results {
            display: flex;
            gap: 30px;
            margin-bottom: 40px;
        }
        
        .result-card {
            flex: 1;
            background: #f8f9fa;
            border-radius: 12px;
            padding: 20px;
            box-shadow: 0 3px 10px rgba(0,0,0,0.05);
            transition: all 0.3s;
        }
        
        .result-card.active {
            transform: scale(1.03);
            box-shadow: 0 10px 25px rgba(0,0,0,0.1);
            border: 2px solid #3498db;
        }
        
        .result-card h2 {
            text-align: center;
            color: #2c3e50;
            margin-top: 0;
        }
        
        .result-content {
            font-size: 16px;
            line-height: 1.6;
        }
        
        .comparison {
            background: #f8f9fa;
            padding: 20px;
            border-radius: 12px;
        }
        
        .bars {
            display: flex;
            height: 60px;
            margin-top: 20px;
            border-radius: 8px;
            overflow: hidden;
        }
        
        .bar {
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-weight: bold;
            transition: width 1s ease-in-out;
            flex-direction: column;
            text-align: center;
            font-size: 14px;
            padding: 5px;
            box-sizing: border-box;
        }
        
        .without-pool-bar {
            background: linear-gradient(to right, #e74c3c, #c0392b);
        }
        
        .with-pool-bar {
            background: linear-gradient(to right, #2ecc71, #27ae60);
        }
        
        .info {
            background-color: #e8f4fc;
            border-left: 4px solid #3498db;
            padding: 15px;
            margin-top: 30px;
            border-radius: 0 8px 8px 0;
        }
        
        .info h3 {
            margin-top: 0;
            color: #2c3e50;
        }
        
        .info ul {
            padding-left: 20px;
        }
        
        .memory-saving {
            font-size: 12px;
            display: block;
            margin-top: 3px;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Демонстрация sync.Pool в Go</h1>
        <div class="controls">
            <button id="without-pool-btn">Тест БЕЗ Pool</button>
            <button id="with-pool-btn">Тест С Pool</button>
        </div>
        
        <div class="results">
            <div id="without-pool" class="result-card">
                <h2>Без Pool</h2>
                <div class="result-content">Запустите тест...</div>
            </div>
            <div id="with-pool" class="result-card">
                <h2>С Pool</h2>
                <div class="result-content">Запустите тест...</div>
            </div>
        </div>
        
        <div class="comparison">
            <h2>Сравнение производительности</h2>
            <div class="bars">
                <div class="bar without-pool-bar" style="width: 50%">Без Pool: 0ms</div>
                <div class="bar with-pool-bar" style="width: 50%">С Pool: 0ms</div>
            </div>
        </div>
        
        <div class="info">
            <h3>Что демонстрирует этот тест?</h3>
            <p>Sync.Pool позволяет повторно использовать объекты, что дает два основных преимущества:</p>
            <ul>
                <li><strong>Сокращение аллокаций:</strong> Меньше созданий/удалений объектов</li>
                <li><strong>Экономия памяти:</strong> Меньше нагрузка на сборщик мусора</li>
            </ul>
            <p>Эти преимущества особенно заметны при высокой нагрузке и частом создании временных объектов.</p>
        </div>
    </div>
    
    <script>
        console.log("Sync.Pool demo script loaded!");
        
        // Получаем элементы кнопок
        const withoutPoolBtn = document.getElementById('without-pool-btn');
        const withPoolBtn = document.getElementById('with-pool-btn');
        
        // Функция для запуска тестов
        async function runTest(withPool) {
            console.log("Starting test with pool=" + withPool);
            
            const button = withPool ? withPoolBtn : withoutPoolBtn;
            const originalText = button.textContent;
            button.textContent = 'Выполняется...';
            button.disabled = true;
            
            try {
                // Увеличиваем количество итераций для более заметной разницы
                const iterations = 500000;
                console.log("Sending request to /test?pool=" + withPool + "&iterations=" + iterations);
                
                const startTime = Date.now();
                const response = await fetch("/test?pool=" + withPool + "&iterations=" + iterations);
                console.log("Response received in " + (Date.now() - startTime) + "ms");
                
                if (!response.ok) {
                    throw new Error("HTTP error! status: " + response.status);
                }
                
                const result = await response.json();
                console.log("Received result:", result);
                
                updateResults(result);
                updateComparison();
            } catch (error) {
                console.error('Test error:', error);
                alert("Error: " + error.message);
            } finally {
                button.textContent = originalText;
                button.disabled = false;
            }
        }

        function updateResults(result) {
            console.log("Updating results for " + (result.with_pool ? "with pool" : "without pool"));
            
            const containerId = result.with_pool ? 'with-pool' : 'without-pool';
            const container = document.getElementById(containerId);
            
            // Конвертируем наносекунды в миллисекунды
            const timeMs = result.duration / 1000000;
            const memoryMB = result.memory / (1024 * 1024);
            
            // Корректное отображение экономии памяти
            const economyStatus = result.with_pool ? '✅ Да' : '❌ Нет';
            
            container.querySelector('.result-content').innerHTML = 
                '<strong>Время выполнения:</strong> ' + timeMs.toFixed(2) + ' ms<br>' +
                '<strong>Аллокации:</strong> ' + result.allocs.toLocaleString() + '<br>' +
                '<strong>Память:</strong> ' + memoryMB.toFixed(2) + ' MB<br>' +
                '<strong>Экономия памяти:</strong> ' + economyStatus;
            
            container.classList.add('active');
            setTimeout(() => container.classList.remove('active'), 2000);
        }

        function updateComparison() {
            console.log("Updating comparison chart");
            
            const withoutPoolElement = document.querySelector('#without-pool .result-content');
            const withPoolElement = document.querySelector('#with-pool .result-content');
            
            if (!withoutPoolElement || !withPoolElement) {
                console.log("Result elements not found");
                return;
            }
            
            // Если нет данных для сравнения
            if (!withoutPoolElement.textContent.includes('Время') || 
                !withPoolElement.textContent.includes('Время')) {
                console.log("Not enough data for comparison");
                return;
            }
            
            // Извлекаем время из текста
            const withoutTimeText = withoutPoolElement.textContent.match(/Время выполнения: (\d+\.\d+)/);
            const withTimeText = withPoolElement.textContent.match(/Время выполнения: (\d+\.\d+)/);
            
            if (!withoutTimeText || !withTimeText || withoutTimeText.length < 2 || withTimeText.length < 2) {
                console.log("Could not parse times");
                return;
            }
            
            const withoutTime = parseFloat(withoutTimeText[1]);
            const withTime = parseFloat(withTimeText[1]);
            
            // Извлекаем данные о памяти
            const withoutMemoryText = withoutPoolElement.textContent.match(/Память: (\d+\.\d+)/);
            const withMemoryText = withPoolElement.textContent.match(/Память: (\d+\.\d+)/);
            
            let memoryImprovement = 0;
            if (withoutMemoryText && withMemoryText && withoutMemoryText.length > 1 && withMemoryText.length > 1) {
                const withoutMemory = parseFloat(withoutMemoryText[1]);
                const withMemory = parseFloat(withMemoryText[1]);
                memoryImprovement = ((withoutMemory - withMemory) / withoutMemory) * 100;
            }
            
            const timeImprovement = ((withoutTime - withTime) / withoutTime) * 100;
            let timeImprovementText = '';
            let memoryImprovementText = '';
            
            if (timeImprovement > 0) {
                timeImprovementText = '<span style="color:#2ecc71">(улучшение на ' + timeImprovement.toFixed(1) + '%)</span>';
            } else if (timeImprovement < 0) {
                timeImprovementText = '<span style="color:#e74c3c">(ухудшение на ' + Math.abs(timeImprovement).toFixed(1) + '%)</span>';
            }
            
            if (memoryImprovement > 0) {
                memoryImprovementText = '<span class="memory-saving">Экономия памяти: ' + memoryImprovement.toFixed(1) + '%</span>';
            } else if (memoryImprovement < 0) {
                memoryImprovementText = '<span class="memory-saving">Перерасход памяти: ' + Math.abs(memoryImprovement).toFixed(1) + '%</span>';
            }
            
            console.log("Times: without=" + withoutTime + "ms, with=" + withTime + "ms, improvement=" + timeImprovement.toFixed(1) + "%, memory=" + memoryImprovement.toFixed(1) + "%");
            
            const maxTime = Math.max(withoutTime, withTime, 1);
            const withoutWidth = (withoutTime / maxTime) * 100;
            const withWidth = (withTime / maxTime) * 100;
            
            const withoutBar = document.querySelector('.without-pool-bar');
            const withBar = document.querySelector('.with-pool-bar');
            
            withoutBar.style.width = withoutWidth + '%';
            withoutBar.innerHTML = 'Без Pool: ' + withoutTime.toFixed(2) + 'ms';
            
            withBar.style.width = withWidth + '%';
            withBar.innerHTML = 'С Pool: ' + withTime.toFixed(2) + 'ms ' + timeImprovementText + memoryImprovementText;
        }

        // Назначаем обработчики кнопок
        withoutPoolBtn.addEventListener('click', () => runTest(false));
        withPoolBtn.addEventListener('click', () => runTest(true));
        
        console.log("Event listeners registered");
    </script>
</body>
</html>`

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

// Обработчик теста
func testHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Обработка запроса /test")

	withPool := r.URL.Query().Get("pool") == "true"

	// Получаем количество итераций
	iterationsParam := r.URL.Query().Get("iterations")
	iterations := 500000 // Значение по умолчанию
	if iters, err := strconv.Atoi(iterationsParam); err == nil && iters > 0 {
		iterations = iters
	}

	// Ограничение для безопасности
	if iterations > 2000000 {
		iterations = 2000000
	}

	log.Printf("Запуск теста: pool=%v, iterations=%d", withPool, iterations)

	start := time.Now()
	var memStatsStart, memStatsEnd runtime.MemStats
	var allocs uint64

	runtime.GC()
	runtime.ReadMemStats(&memStatsStart)

	if withPool {
		runTestWithPool(iterations)
	} else {
		runTestWithoutPool(iterations)
	}

	runtime.ReadMemStats(&memStatsEnd)
	duration := time.Since(start)

	// Рассчитываем аллокации
	allocs = memStatsEnd.Mallocs - memStatsStart.Mallocs
	memoryUsed := memStatsEnd.TotalAlloc - memStatsStart.TotalAlloc

	result := TestResult{
		Duration: duration,
		Allocs:   allocs,
		Memory:   memoryUsed,
		WithPool: withPool,
	}

	log.Printf("Результат теста: duration=%v, allocs=%d, memory=%d bytes",
		duration, allocs, memoryUsed)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Printf("Ошибка кодирования JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Сложные операции с объектом
func processObject(obj *HeavyObject, id int) {
	obj.ID = id
	// Более сложные операции для увеличения нагрузки
	for j := 0; j < len(obj.Content); j++ {
		// Явное преобразование типов для избежания предупреждений
		value := int(obj.Content[j])

		if j%2 == 0 {
			value = (value + id) % 256
		} else {
			value = (value - id) % 256
			// Обеспечиваем положительное значение
			if value < 0 {
				value += 256
			}
		}

		obj.Content[j] = byte(value)
	}
}

func runTestWithPool(iterations int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }()

			// Получаем объект из пула
			obj := pool.Get().(*HeavyObject)

			// Выполняем сложные операции
			processObject(obj, id)

			// Возвращаем объект в пул
			pool.Put(obj)
		}(i)
	}
	wg.Wait()
}

func runTestWithoutPool(iterations int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	for i := 0; i < iterations; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }()

			// Создаем новый объект
			obj := new(HeavyObject)

			// Выполняем сложные операции
			processObject(obj, id)
		}(i)
	}
	wg.Wait()
}
