package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Task struct {
	Name     string
	Duration time.Duration
	Deadline time.Time
}

func main() {
	fmt.Println("🚀 HomeContext Challenge: Управление конкурирующими контекстами")
	fmt.Println("================================================================")

	// Создаём родительский контекст с таймаутом на весь день (8 часов)
	parentCtx, cancel := context.WithTimeout(context.Background(), 8*time.Hour)
	defer cancel()

	// Задачи дня
	tasks := []Task{
		{"🧹 Пропылесосить комнаты", 30 * time.Minute, time.Now().Add(45 * time.Minute)},
		{"🧼 Помыть полы", 45 * time.Minute, time.Now().Add(90 * time.Minute)},
		{"📚 Урок Go: Контексты", 120 * time.Minute, time.Now().Add(150 * time.Minute)},
		{"🍳 Приготовить ужин", 40 * time.Minute, time.Now().Add(180 * time.Minute)},
	}

	var wg sync.WaitGroup
	taskCh := make(chan string, len(tasks))
	score := 0
	var mu sync.Mutex

	for _, task := range tasks {
		wg.Add(1)
		go func(t Task) {
			defer wg.Done()

			// Создаём контекст с дедлайном задачи
			ctx, cancelTask := context.WithDeadline(parentCtx, t.Deadline)
			defer cancelTask()

			select {
			case <-time.After(t.Duration):
				// Задача выполнена в срок
				mu.Lock()
				score += 10
				mu.Unlock()
				taskCh <- fmt.Sprintf("✅ %s выполнено за %v (+10 очков)", t.Name, t.Duration)

			case <-ctx.Done():
				// Задача отменена (дедлайн или родительский контекст)
				reason := "дедлайн"
				if ctx.Err() == context.Canceled {
					reason = "отмена родительского контекста"
					mu.Lock()
					score -= 20
					mu.Unlock()
				} else {
					mu.Lock()
					score -= 5
					mu.Unlock()
				}
				taskCh <- fmt.Sprintf("❌ %s отменено (%s) (-5 очков)", t.Name, reason)
			}
		}(task)
	}

	// Симуляция внешнего события (Гоша отвлёкся на видео через 90 минут)
	go func() {
		time.Sleep(90 * time.Minute)
		fmt.Println("\n⚠️  ВНИМАНИЕ: Гоша начал смотреть видео про отпуск!")
		fmt.Println("   Родительский контекст отменяется...")
		cancel()
	}()

	wg.Wait()
	close(taskCh)

	fmt.Println("\n📊 Результаты выполнения задач:")
	fmt.Println("-------------------------------")
	for result := range taskCh {
		fmt.Println(result)
	}

	mu.Lock()
	finalScore := score
	mu.Unlock()

	fmt.Printf("\n🎯 Итоговый счёт: %d очков\n", finalScore)

	switch {
	case finalScore >= 100:
		fmt.Println("🔥 ФЕНОМЕНАЛЬНО! Гоша становится мастером контекстов!")
	case finalScore >= 50:
		fmt.Println("👍 ХОРОШО! Но есть куда стремиться.")
	default:
		fmt.Println("💪 ЗАВТРА НОВЫЙ ДЕНЬ! Учимся на ошибках.")
	}
}
