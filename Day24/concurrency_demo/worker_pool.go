package main

import (
	"fmt"
	"sync"
)

// Task - задача для воркера
type Task struct {
	ID  int
	Data string
}

// Worker - обработчик задач
func Worker(id int, tasks <-chan Task, wg *sync.WaitGroup) {
	defer wg.Done()
	
	for task := range tasks {
		fmt.Printf("Воркер %d начал задачу %d: %s\n", id, task.ID, task.Data)
		processData(task.Data) // Имитация обработки
		fmt.Printf("Воркер %d завершил задачу %d\n", id, task.ID)
	}
}

// RunWorkerPool - запуск пула воркеров
func RunWorkerPool() {
	const numWorkers = 3
	const numTasks = 10
	
	var wg sync.WaitGroup
	tasks := make(chan Task, numTasks)
	
	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, tasks, &wg)
	}
	
	// Отправка задач
	for i := 1; i <= numTasks; i++ {
		tasks <- Task{
			ID:   i,
			Data: fmt.Sprintf("Данные задачи %d", i),
		}
	}
	
	close(tasks) // Закрываем канал
	wg.Wait()    // Ожидаем завершения
}