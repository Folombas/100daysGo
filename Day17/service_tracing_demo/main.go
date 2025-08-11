package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

// Структура для конфигурации трассировки
type TracingConfig struct {
	Duration    time.Duration
	Workers     int
	MaxTasks    int
	TraceFile   string
	Description string
}

// Результат трассировки
type TraceResult struct {
	Status      string
	TraceFile   string
	GoVersion   string
	Duration    time.Duration
	Workers     int
	Tasks       int
	Description string
}

// Глобальные переменные
var (
	tracingActive bool
	tracingMutex  sync.Mutex
	traceFile     string
)

// worker имитирует выполнение задач
func worker(ctx context.Context, id int, wg *sync.WaitGroup, taskCh <-chan int) {
	defer wg.Done()
	defer trace.StartRegion(ctx, "worker_lifecycle").End()

	// Используем пустую структуру вместо taskID
	for range taskCh {
		// Регистрируем выполнение задачи
		region := trace.StartRegion(ctx, "process_task")

		// Имитация различных типов задач
		switch rand.Intn(4) {
		case 0: // CPU-bound задача
			trace.Log(ctx, "task_type", "cpu_bound")
			doCPUWork(10000000)
		case 1: // I/O-bound задача
			trace.Log(ctx, "task_type", "io_bound")
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		case 2: // Сетевая задача
			trace.Log(ctx, "task_type", "network")
			time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
		case 3: // Ожидание ресурсов
			trace.Log(ctx, "task_type", "waiting")
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}

		region.End()
	}
}

// CPU-intensive задача
func doCPUWork(n int) {
	total := 0
	for i := 0; i < n; i++ {
		total += i * i
	}
	_ = total
}

// Запуск трассировки
func startTracing(config TracingConfig) error {
	tracingMutex.Lock()
	defer tracingMutex.Unlock()

	if tracingActive {
		return fmt.Errorf("трассировка уже запущена")
	}

	// Создаем файл для трассировки
	f, err := os.Create(config.TraceFile)
	if err != nil {
		return err
	}

	// Начинаем трассировку
	if err := trace.Start(f); err != nil {
		f.Close()
		return err
	}

	tracingActive = true
	traceFile = config.TraceFile

	// Запускаем фоновую задачу для остановки трассировки
	go func() {
		time.Sleep(config.Duration)
		stopTracing()
	}()

	return nil
}

// Остановка трассировки
func stopTracing() {
	tracingMutex.Lock()
	defer tracingMutex.Unlock()

	if tracingActive {
		trace.Stop()
		tracingActive = false
	}
}

// Обработчик API для запуска трассировки
func startTraceHandler(w http.ResponseWriter, r *http.Request) {
	var config TracingConfig
	if err := json.NewDecoder(r.Body).Decode(&config); err != nil {
		http.Error(w, "Ошибка декодирования JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if config.Duration == 0 {
		config.Duration = 5 * time.Second
	}
	if config.Workers == 0 {
		config.Workers = 10
	}
	if config.MaxTasks == 0 {
		config.MaxTasks = 50
	}
	if config.TraceFile == "" {
		config.TraceFile = "trace.out"
	}

	if err := startTracing(config); err != nil {
		http.Error(w, "Ошибка запуска трассировки: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Создаем контекст для трассировки
	ctx := context.Background()

	// Создаем канал для задач
	taskCh := make(chan int, config.MaxTasks)

	// Запускаем воркеров
	var wg sync.WaitGroup
	wg.Add(config.Workers)

	for i := 0; i < config.Workers; i++ {
		go worker(ctx, i, &wg, taskCh)
	}

	// Отправляем задачи
	go func() {
		for i := 0; i < config.MaxTasks; i++ {
			taskCh <- 1 // Просто отправляем сигнал вместо ID
		}
		close(taskCh)
	}()

	// Ждем завершения воркеров в фоне
	go func() {
		wg.Wait()
	}()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":   "Трассировка запущена",
		"file":     config.TraceFile,
		"duration": config.Duration.String(),
		"workers":  fmt.Sprintf("%d", config.Workers),
		"tasks":    fmt.Sprintf("%d", config.MaxTasks),
	})
}

// Обработчик для остановки трассировки
func stopTraceHandler(w http.ResponseWriter, r *http.Request) {
	stopTracing()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "Трассировка остановлена",
		"file":   traceFile,
	})
}

// Статус трассировки
func traceStatusHandler(w http.ResponseWriter, r *http.Request) {
	status := "не активна"
	if tracingActive {
		status = "активна"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":    status,
		"file":      traceFile,
		"goVersion": runtime.Version(),
	})
}

// Главная страница
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := struct {
		Title     string
		GoVersion string
	}{
		Title:     "День 17: Трассировка поведения сервиса",
		GoVersion: runtime.Version(),
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Ошибка шаблона: "+err.Error(), http.StatusInternalServerError)
	}
}

// Скачивание файла трассировки
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	if traceFile == "" {
		http.Error(w, "Файл трассировки не найден", http.StatusNotFound)
		return
	}

	file, err := os.Open(traceFile)
	if err != nil {
		http.Error(w, "Ошибка открытия файла: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+traceFile)
	w.Header().Set("Content-Type", "application/octet-stream")

	if _, err := io.Copy(w, file); err != nil {
		http.Error(w, "Ошибка отправки файла: "+err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Регистрация обработчиков
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/api/start-trace", startTraceHandler)
	http.HandleFunc("/api/stop-trace", stopTraceHandler)
	http.HandleFunc("/api/trace-status", traceStatusHandler)
	http.HandleFunc("/download", downloadHandler)

	port := ":8080"
	log.Printf("🚀 Сервер запущен на http://localhost%s", port)
	log.Println("🔍 Для трассировки откройте веб-интерфейс:")
	log.Printf("   http://localhost%s", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
