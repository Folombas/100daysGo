package main

import (
	"fmt"
	"strings"
	"time"
)

// Task представляет задачу в расписании
type Task struct {
	Name        string
	Scheduled   time.Time
	Completed   bool
	CompletedAt time.Time
}

// DailyRoutine представляет расписание дня
type DailyRoutine struct {
	Date     time.Time
	Tasks    []Task
	Mood     string
	Progress float64
}

// NewRoutine создает новое расписание
func NewRoutine(date time.Time, mood string) DailyRoutine {
	return DailyRoutine{
		Date:  date,
		Mood:  mood,
		Tasks: []Task{},
	}
}

// AddTask добавляет задачу в расписание
func (dr *DailyRoutine) AddTask(name string, scheduled time.Time) {
	dr.Tasks = append(dr.Tasks, Task{
		Name:      name,
		Scheduled: scheduled,
	})
}

// CompleteTask отмечает задачу выполненной
func (dr *DailyRoutine) CompleteTask(taskName string) {
	for i := range dr.Tasks {
		if dr.Tasks[i].Name == taskName && !dr.Tasks[i].Completed {
			dr.Tasks[i].Completed = true
			dr.Tasks[i].CompletedAt = time.Now()
			dr.updateProgress()
			fmt.Printf("✅ Задача '%s' выполнена в %s!\n",
				taskName, dr.Tasks[i].CompletedAt.Format("15:04"))
			dr.generateDopamine()
			return
		}
	}
	fmt.Printf("⚠️ Задача '%s' не найдена или уже выполнена\n", taskName)
}

// updateProgress пересчитывает прогресс дня
func (dr *DailyRoutine) updateProgress() {
	completed := 0
	for _, task := range dr.Tasks {
		if task.Completed {
			completed++
		}
	}
	if len(dr.Tasks) > 0 {
		dr.Progress = float64(completed) / float64(len(dr.Tasks)) * 100
	}
}

// generateDopamine имитирует выброс дофамина
func (dr *DailyRoutine) generateDopamine() {
	motivations := []string{
		"💡 Новый нейронный путь активирован!",
		"🚀 Прогресс чувствуется в каждой клетке!",
		"🎯 Еще один шаг к Go-мастерству!",
		"⚡ Компиляция успеха завершена!",
		"🔥 Очередной баг лени пофикшен!",
	}
	index := time.Now().Unix() % int64(len(motivations))
	fmt.Println(motivations[index])
}

// Display показывает текущее состояние расписания
func (dr *DailyRoutine) Display() {
	fmt.Printf("\n📅 День: %s\n", dr.Date.Format("2 January 2006"))
	fmt.Printf("😶 Настроение: %s\n", dr.Mood)
	fmt.Printf("📊 Прогресс: %.1f%%\n\n", dr.Progress)

	fmt.Println("📋 Расписание:")
	for i, task := range dr.Tasks {
		status := "⏳"
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("%d. %s %s [%s]\n",
			i+1, status, task.Name, task.Scheduled.Format("15:04"))
	}
}

// createGoshaSchedule создает расписание Гоши
func createGoshaSchedule() DailyRoutine {
	today := time.Date(2026, 2, 2, 0, 0, 0, 0, time.UTC)
	routine := NewRoutine(today, "Борется с хаосом")

	// Добавляем задачи
	routine.AddTask("Подъем и уборка постели", time.Date(2026, 2, 2, 10, 30, 0, 0, time.UTC))
	routine.AddTask("Бритье и душ", time.Date(2026, 2, 2, 11, 0, 0, 0, time.UTC))
	routine.AddTask("Работа курьером", time.Date(2026, 2, 2, 12, 0, 0, 0, time.UTC))
	routine.AddTask("Оплата коммуналки", time.Date(2026, 2, 2, 15, 0, 0, 0, time.UTC))
	routine.AddTask("Урок Go: Common Usecases", time.Date(2026, 2, 2, 19, 0, 0, 0, time.UTC))

	return routine
}

// runDailyTasks выполняет задачи дня
func runDailyTasks(routine *DailyRoutine) {
	tasks := []string{
		"Подъем и уборка постели",
		"Бритье и душ",
		"Работа курьером",
		"Оплата коммуналки",
	}

	fmt.Println("\n🎯 Начинаем день:")
	for _, task := range tasks {
		routine.CompleteTask(task)
		time.Sleep(500 * time.Millisecond) // Уменьшил задержку для быстрого выполнения
	}

	fmt.Println("\n⏰ 19:00 - время для Go!")
	routine.Mood = "Сфокусирован на коде"
	routine.CompleteTask("Урок Go: Common Usecases")
}

// showInsights показывает инсайты дня
func showInsights() {
	insights := []string{
		"1. Рутина — это не враг, а runtime environment",
		"2. Каждая выполненная задача — это успешный тест",
		"3. Дисциплина компилируется в мастерство",
		"4. Go учит: маленькие packages > монолитный хаос",
		"5. Завтра: новый день, новый module!",
	}

	fmt.Println("\n✨ Сегодняшние инсайты:")
	for _, insight := range insights {
		fmt.Println(insight)
		time.Sleep(300 * time.Millisecond) // Уменьшил задержку
	}
}

func main() {
	fmt.Println("🎮 Daily Routine Manager v2.0")
	fmt.Println("=", strings.Repeat("=", 40))

	// Создаем расписание Гоши
	routine := createGoshaSchedule()
	routine.Display()

	// Выполняем задачи дня
	runDailyTasks(&routine)

	// Показываем итоги
	fmt.Println("\n", strings.Repeat("=", 40))
	routine.Display()

	// Показываем инсайты
	showInsights()

	fmt.Println("\n🚀 Завтрашняя цель: Day 91 — Concurrency Patterns!")
}
