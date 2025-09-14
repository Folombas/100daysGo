package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Demo data structures
type UseCase struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Example     string `json:"example"`
	Code        string `json:"code"`
}

type DemoResult struct {
	Operation string      `json:"operation"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

// Global demo data
var useCases = []UseCase{
	{
		Title:       "Web Services & APIs",
		Description: "Go excels at building high-performance web services and REST APIs with built-in HTTP server",
		Category:    "Web Development",
		Example:     "REST API with JSON responses",
		Code: `package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    ` + "`json:\"id\"`" + `
    Name string ` + "`json:\"name\"`" + `
}

func main() {
    http.HandleFunc("/api/users", func(w http.ResponseWriter, r *http.Request) {
        users := []User{{ID: 1, Name: "Alice"}, {ID: 2, Name: "Bob"}}
        json.NewEncoder(w).Encode(users)
    })
    http.ListenAndServe(":8080", nil)
}`,
	},
	{
		Title:       "Microservices Architecture",
		Description: "Go's lightweight nature makes it perfect for microservices with fast startup times",
		Category:    "Architecture",
		Example:     "Service discovery and communication",
		Code: `package main

import (
    "context"
    "fmt"
    "net/http"
    "time"
)

type Service struct {
    Name    string
    Port    int
    Healthy bool
}

func healthCheck(ctx context.Context, service *Service) {
    ticker := time.NewTicker(30 * time.Second)
    defer ticker.Stop()
    
    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            resp, err := http.Get(fmt.Sprintf("http://localhost:%d/health", service.Port))
            service.Healthy = err == nil && resp.StatusCode == 200
        }
    }
}`,
	},
	{
		Title:       "Concurrent Programming",
		Description: "Go's goroutines and channels make concurrent programming simple and efficient",
		Category:    "Concurrency",
		Example:     "Worker pool pattern",
		Code: `package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // Start workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Collect results
    for a := 1; a <= 5; a++ {
        <-results
    }
}`,
	},
	{
		Title:       "DevOps & Tooling",
		Description: "Go is widely used for building DevOps tools like Docker, Kubernetes, and CI/CD pipelines",
		Category:    "DevOps",
		Example:     "Docker-like container runtime",
		Code: `package main

import (
    "fmt"
    "os/exec"
    "syscall"
)

type Container struct {
    ID     string
    Image  string
    Status string
}

func (c *Container) Start() error {
    cmd := exec.Command("docker", "run", "-d", c.Image)
    output, err := cmd.Output()
    if err != nil {
        return err
    }
    c.ID = string(output[:12])
    c.Status = "running"
    return nil
}

func (c *Container) Stop() error {
    cmd := exec.Command("docker", "stop", c.ID)
    return cmd.Run()
}`,
	},
	{
		Title:       "System Programming",
		Description: "Go provides low-level system access while maintaining high-level abstractions",
		Category:    "Systems",
		Example:     "File system monitoring",
		Code: `package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

type FileWatcher struct {
    Path     string
    Callback func(string)
}

func (fw *FileWatcher) Watch() {
    for {
        filepath.Walk(fw.Path, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if !info.IsDir() {
                fw.Callback(path)
            }
            return nil
        })
        time.Sleep(1 * time.Second)
    }
}

func main() {
    watcher := &FileWatcher{
        Path: "/tmp",
        Callback: func(path string) {
            fmt.Printf("File changed: %s\n", path)
        },
    }
    watcher.Watch()
}`,
	},
	{
		Title:       "Cloud Native Applications",
		Description: "Go is the language of choice for cloud-native applications and serverless functions",
		Category:    "Cloud",
		Example:     "AWS Lambda function",
		Code: `package main

import (
    "context"
    "encoding/json"
    "github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
    Name string ` + "`json:\"name\"`" + `
}

type Response struct {
    Message string ` + "`json:\"message\"`" + `
}

func HandleRequest(ctx context.Context, request Request) (Response, error) {
    return Response{
        Message: fmt.Sprintf("Hello %s from Go Lambda!", request.Name),
    }, nil
}

func main() {
    lambda.Start(HandleRequest)
}`,
	},
}

func main() {
	r := mux.NewRouter()

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// API routes
	r.HandleFunc("/api/usecases", getUseCases).Methods("GET")
	r.HandleFunc("/api/demo/{type}", handleDemo).Methods("POST")
	r.HandleFunc("/ws", handleWebSocket)

	// Main page
	r.HandleFunc("/", homeHandler).Methods("GET")

	fmt.Println("🚀 Go Use Cases Educational Server starting on :3000")
	fmt.Println("📚 Lesson: 'Go Beyond Basics - Exploring Real-World Applications'")
	fmt.Println("🌐 Open http://localhost:3000 in your browser")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func getUseCases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(useCases)
}

func handleDemo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	demoType := vars["type"]

	var result DemoResult
	result.Timestamp = time.Now()

	switch demoType {
	case "concurrency":
		result.Operation = "Goroutine Demo"
		result.Data = map[string]interface{}{
			"goroutines": 1000,
			"duration":   "1ms",
			"memory":     "2MB",
		}
	case "performance":
		result.Operation = "Performance Test"
		result.Data = map[string]interface{}{
			"requests_per_second": 50000,
			"latency":             "0.1ms",
			"cpu_usage":           "15%",
		}
	case "memory":
		result.Operation = "Memory Management"
		result.Data = map[string]interface{}{
			"heap_size":   "10MB",
			"gc_pauses":   "0.5ms",
			"allocations": "1000/s",
		}
	default:
		result.Operation = "Unknown Demo"
		result.Data = map[string]string{"error": "Invalid demo type"}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	// Send real-time updates
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		update := map[string]interface{}{
			"type":      "update",
			"timestamp": time.Now(),
			"data": map[string]interface{}{
				"active_goroutines": 50 + int(time.Now().Unix()%100),
				"memory_usage":      "15.2MB",
				"cpu_usage":         "12.5%",
			},
		}

		if err := conn.WriteJSON(update); err != nil {
			log.Println("WebSocket write failed:", err)
			return
		}
	}
}
