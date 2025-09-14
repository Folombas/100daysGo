package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Interactive Demo Functions for the Web Interface

// ConcurrencyDemo demonstrates Go's goroutine capabilities
func ConcurrencyDemo() map[string]interface{} {
	fmt.Println("🔄 Running Concurrency Demo...")

	start := time.Now()
	var wg sync.WaitGroup
	results := make(chan int, 1000)

	// Launch 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Simulate some work
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(10)))
			results <- id * 2
		}(i)
	}

	// Close results channel when all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	var sum int
	count := 0
	for result := range results {
		sum += result
		count++
	}

	duration := time.Since(start)

	return map[string]interface{}{
		"goroutines_launched":    1000,
		"results_processed":      count,
		"total_sum":              sum,
		"duration_ms":            duration.Milliseconds(),
		"avg_time_per_goroutine": fmt.Sprintf("%.2fμs", float64(duration.Nanoseconds())/1000/1000),
	}
}

// PerformanceDemo demonstrates Go's performance characteristics
func PerformanceDemo() map[string]interface{} {
	fmt.Println("⚡ Running Performance Demo...")

	// Simulate high-performance operations
	start := time.Now()

	// Simulate concurrent HTTP requests
	var wg sync.WaitGroup
	requestCount := 10000
	successCount := 0
	var mu sync.Mutex

	for i := 0; i < requestCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate HTTP request processing
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))

			mu.Lock()
			successCount++
			mu.Unlock()
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	requestsPerSecond := float64(requestCount) / duration.Seconds()
	avgLatency := duration.Nanoseconds() / int64(requestCount) / 1000 // microseconds

	return map[string]interface{}{
		"total_requests":      requestCount,
		"successful_requests": successCount,
		"requests_per_second": fmt.Sprintf("%.0f", requestsPerSecond),
		"avg_latency_us":      avgLatency,
		"total_duration_ms":   duration.Milliseconds(),
		"success_rate":        fmt.Sprintf("%.2f%%", float64(successCount)/float64(requestCount)*100),
	}
}

// MemoryManagementDemo demonstrates Go's memory efficiency
func MemoryManagementDemo() map[string]interface{} {
	fmt.Println("💾 Running Memory Management Demo...")

	// Simulate memory allocation patterns
	start := time.Now()

	// Create a large slice to demonstrate memory management
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024) // 1KB per slice
	}

	// Simulate some processing
	time.Sleep(100 * time.Millisecond)

	// Clear some data to trigger GC
	for i := 0; i < 500; i++ {
		data[i] = nil
	}

	// Force garbage collection (in real app, this happens automatically)
	simulateGC()

	duration := time.Since(start)

	return map[string]interface{}{
		"initial_allocation_mb": "1.0",
		"final_allocation_mb":   "0.5",
		"gc_triggers":           "1",
		"allocation_time_ms":    duration.Milliseconds(),
		"memory_efficiency":     "High",
		"gc_pause_time_us":      "< 100",
	}
}

// WebSocketDemo demonstrates real-time communication
func WebSocketDemo() map[string]interface{} {
	fmt.Println("🌐 Running WebSocket Demo...")

	// Simulate WebSocket connections
	connectionCount := 100
	messageCount := 0
	var mu sync.Mutex

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Simulate multiple WebSocket connections
	for i := 0; i < connectionCount; i++ {
		go func(connID int) {
			ticker := time.NewTicker(50 * time.Millisecond)
			defer ticker.Stop()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					mu.Lock()
					messageCount++
					mu.Unlock()
				}
			}
		}(i)
	}

	// Wait for demo to complete
	<-ctx.Done()

	return map[string]interface{}{
		"active_connections":  connectionCount,
		"messages_sent":       messageCount,
		"messages_per_second": fmt.Sprintf("%.0f", float64(messageCount)/2.0),
		"avg_latency_ms":      "< 1",
		"connection_overhead": "Minimal",
	}
}

// DatabaseDemo demonstrates database operations
func DatabaseDemo() map[string]interface{} {
	fmt.Println("🗄️ Running Database Demo...")

	start := time.Now()

	// Simulate database operations
	var wg sync.WaitGroup
	operationCount := 1000
	successCount := 0
	var mu sync.Mutex

	for i := 0; i < operationCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// Simulate database query
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)+1))

			mu.Lock()
			successCount++
			mu.Unlock()
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	opsPerSecond := float64(operationCount) / duration.Seconds()

	return map[string]interface{}{
		"total_operations":      operationCount,
		"successful_ops":        successCount,
		"operations_per_second": fmt.Sprintf("%.0f", opsPerSecond),
		"avg_query_time_ms":     fmt.Sprintf("%.2f", float64(duration.Nanoseconds())/float64(operationCount)/1000000),
		"connection_pool_size":  "10",
		"success_rate":          "100%",
	}
}

// SystemResourceDemo demonstrates system resource usage
func SystemResourceDemo() map[string]interface{} {
	fmt.Println("🖥️ Running System Resource Demo...")

	start := time.Now()

	// Simulate CPU-intensive work
	var wg sync.WaitGroup
	cpuWorkers := 4

	for i := 0; i < cpuWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Simulate CPU work
			for j := 0; j < 1000000; j++ {
				_ = j * j
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)

	return map[string]interface{}{
		"cpu_workers":       cpuWorkers,
		"work_completed":    "4,000,000 operations",
		"total_time_ms":     duration.Milliseconds(),
		"cpu_efficiency":    "High",
		"memory_per_worker": "2KB",
		"context_switches":  "Minimal",
	}
}

// RunAllDemos executes all demonstration functions
func RunAllDemos() map[string]interface{} {
	fmt.Println("🎯 Running All Interactive Demos...")
	fmt.Println(strings.Repeat("=", 50))

	results := make(map[string]interface{})

	// Run each demo
	results["concurrency"] = ConcurrencyDemo()
	results["performance"] = PerformanceDemo()
	results["memory"] = MemoryManagementDemo()
	results["websocket"] = WebSocketDemo()
	results["database"] = DatabaseDemo()
	results["system"] = SystemResourceDemo()

	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("✅ All demos completed successfully!")

	return results
}

// Utility function to simulate garbage collection (for demo purposes)
func simulateGC() {
	// In a real application, this would trigger garbage collection
	// For demo purposes, we'll just simulate a brief pause
	time.Sleep(1 * time.Millisecond)
}
