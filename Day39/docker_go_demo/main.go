package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		startServer()
		return
	}

	printDockerInfo()
}

func startServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<head>
	<title>Docker + Go Demo</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.info { background: #f5f5f5; padding: 20px; border-radius: 8px; }
	</style>
</head>
<body>
	<h1>🚀 Docker + Go Demo Application</h1>
	<div class="info">
		<h2>System Information:</h2>
		<p>Go Version: %s</p>
		<p>OS/Arch: %s/%s</p>
		<p>Hostname: %s</p>
		<p>Running in Docker: %t</p>
		<p>Server Uptime: %s</p>
		<p>Current Time: %s</p>
	</div>
	<div class="info">
		<h2>Endpoints:</h2>
		<ul>
			<li><a href="/health">/health</a> - Health check</li>
			<li><a href="/stats">/stats</a> - Container statistics</li>
			<li><a href="/api/version">/api/version</a> - Version info</li>
		</ul>
	</div>
</body>
</html>
		`, runtime.Version(), runtime.GOOS, runtime.GOARCH, getHostname(), isRunningInDocker(), getUptime(), time.Now().Format("2006-01-02 15:04:05"))
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "healthy", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"go_routines": %d,
			"memory_alloc": "%v MB",
			"running_in_docker": %t,
			"hostname": "%s"
		}`, runtime.NumGoroutine(), getMemoryUsage(), isRunningInDocker(), getHostname())
	})

	http.HandleFunc("/api/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"go_version": "%s",
			"platform": "%s/%s",
			"dockerized": %t
		}`, runtime.Version(), runtime.GOOS, runtime.GOARCH, isRunningInDocker())
	})

	port := getEnv("PORT", "8080")
	fmt.Printf("🌐 Server starting on port %s...\n", port)
	fmt.Printf("📊 Endpoints:\n")
	fmt.Printf("   http://localhost:%s/\n", port)
	fmt.Printf("   http://localhost:%s/health\n", port)
	fmt.Printf("   http://localhost:%s/stats\n", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("❌ Server error: %v\n", err)
	}
}

func printDockerInfo() {
	fmt.Println("🐳 Docker + Go Demonstration")
	fmt.Println("============================")
	fmt.Println()
	fmt.Println("Available Docker configurations:")
	fmt.Println("1. Single-stage build: docker build -f single-stage/Dockerfile -t go-single .")
	fmt.Println("2. Multi-stage build: docker build -f multi-stage/Dockerfile -t go-multi .")
	fmt.Println("3. With healthcheck: docker build -f healthcheck/Dockerfile -t go-health .")
	fmt.Println("4. Docker Compose: docker-compose -f compose-app/docker-compose.yml up")
	fmt.Println()
	fmt.Println("To run the server directly: go run main.go server")
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func isRunningInDocker() bool {
	_, err := os.Stat("/.dockerenv")
	return err == nil
}

func getUptime() string {
	return time.Since(startTime).Round(time.Second).String()
}

func getMemoryUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%.2f", float64(m.Alloc)/1024/1024)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

var startTime = time.Now()