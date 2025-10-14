package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Episode представляет эпизод из мультфильма "Ну, Погоди!"
type Episode struct {
	ID          int
	Description string
	Complexity  int // Сложность поимки Зайца (1-10)
}

// Worker представляет Волка-работягу
type Worker struct {
	ID     int
	Name   string
	Skills int // Навыки поимки (1-10)
}

// Result представляет результат попытки поймать Зайца
type Result struct {
	WorkerID  int
	EpisodeID int
	Success   bool
	Message   string
}

func main() {
	fmt.Println("🎬 День 81: Worker Pools - Параллельные ловцы Зайцев!")
	fmt.Println("🚪 Мы в 5D-вселенной 'Ну, Погоди!' с Пахомычем...")
	fmt.Println("🐺 Волк просит помочь обработать все эпизоды параллельно!")

	// Создаем эпизоды (задачи)
	episodes := generateEpisodes(20)
	fmt.Printf("📺 Сгенерировано %d эпизодов для обработки\n", len(episodes))

	// Создаем Волков-работяг (worker pool)
	workers := createWolfWorkers(4)
	fmt.Printf("🐺 Создан пул из %d Волков-обработчиков\n\n", len(workers))

	// Каналы для работы
	jobs := make(chan Episode, len(episodes))
	results := make(chan Result, len(episodes))

	// Запускаем worker pool
	var wg sync.WaitGroup

	// Запускаем воркеров
	for _, worker := range workers {
		wg.Add(1)
		go processEpisodes(worker, jobs, results, &wg)
	}

	// Отправляем эпизоды в работу
	go func() {
		for _, episode := range episodes {
			jobs <- episode
		}
		close(jobs)
	}()

	// Ждем завершения и собираем результаты
	go func() {
		wg.Wait()
		close(results)
	}()

	// Анализируем результаты
	analyzeResults(results, workers, episodes)
}

// generateEpisodes создает случайные эпизоды мультфильма
func generateEpisodes(count int) []Episode {
	descriptions := []string{
		"Заяц на стройке с кирпичами",
		"Заяц в лифте с морковкой",
		"Заяц на пляже с мячиком",
		"Заяц в универмаге с тележкой",
		"Заяц на катке с коньками",
		"Заяц в кинотеатре с попкорном",
		"Заяц в бассейне с надувным кругом",
		"Заяц на детской площадке с качелями",
		"Заяц в цирке с клоунами",
		"Заяц в парке с голубями",
	}

	episodes := make([]Episode, count)
	for i := 0; i < count; i++ {
		episodes[i] = Episode{
			ID:          i + 1,
			Description: descriptions[rand.Intn(len(descriptions))],
			Complexity:  rand.Intn(10) + 1,
		}
	}
	return episodes
}

// createWolfWorkers создает пул Волков-обработчиков
func createWolfWorkers(count int) []Worker {
	names := []string{"Волк-Альфа", "Волк-Бета", "Волк-Гамма", "Волк-Дельта", "Волк-Эпсилон"}
	workers := make([]Worker, count)
	for i := 0; i < count; i++ {
		workers[i] = Worker{
			ID:     i + 1,
			Name:   names[i],
			Skills: rand.Intn(8) + 3, // Навыки от 3 до 10
		}
	}
	return workers
}

// processEpisodes обрабатывает эпизоды (основная логика worker pool)
func processEpisodes(worker Worker, jobs <-chan Episode, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for episode := range jobs {
		// Имитируем обработку
		processingTime := time.Duration(rand.Intn(500)+100) * time.Millisecond

		fmt.Printf("🐺 %s начал обрабатывать эпизод %d: %s\n",
			worker.Name, episode.ID, episode.Description)

		time.Sleep(processingTime)

		// Определяем успех поимки
		success := worker.Skills >= episode.Complexity
		var message string

		if success {
			message = fmt.Sprintf("✅ УСПЕХ! %s поймал Зайца в '%s' за %v",
				worker.Name, episode.Description, processingTime)
		} else {
			message = fmt.Sprintf("❌ ПРОВАЛ! %s не смог догнать Зайца в '%s' (навыки: %d, сложность: %d)",
				worker.Name, episode.Description, worker.Skills, episode.Complexity)
		}

		results <- Result{
			WorkerID:  worker.ID,
			EpisodeID: episode.ID,
			Success:   success,
			Message:   message,
		}
	}
}

// analyzeResults анализирует и выводит итоги работы
func analyzeResults(results <-chan Result, workers []Worker, episodes []Episode) {
	fmt.Println("\n📊 АНАЛИЗ РЕЗУЛЬТАТОВ РАБОТЫ WOLF POOL:")
	fmt.Println("========================================")

	stats := struct {
		totalProcessed int
		successCount   int
		failureCount   int
		workerStats    map[int]int
	}{
		workerStats: make(map[int]int),
	}

	for result := range results {
		stats.totalProcessed++
		if result.Success {
			stats.successCount++
			stats.workerStats[result.WorkerID]++
		} else {
			stats.failureCount++
		}
		fmt.Println(result.Message)
	}

	// Выводим статистику
	fmt.Println("\n📈 ИТОГОВАЯ СТАТИСТИКА:")
	fmt.Printf("Обработано эпизодов: %d\n", stats.totalProcessed)
	fmt.Printf("Успешных поимок: %d\n", stats.successCount)
	fmt.Printf("Неудачных попыток: %d\n", stats.failureCount)
	fmt.Printf("Эффективность: %.1f%%\n", float64(stats.successCount)/float64(stats.totalProcessed)*100)

	fmt.Println("\n🏆 СТАТИСТИКА ВОЛКОВ:")
	for _, worker := range workers {
		count := stats.workerStats[worker.ID]
		fmt.Printf("%s (навыки: %d) - поймал Зайцев: %d\n",
			worker.Name, worker.Skills, count)
	}

	fmt.Println("\n🎉 Пахомыч доволен: 'Вот это параллельная обработка! Теперь и я понимаю worker pools!'")
	fmt.Println("🚪 Возвращаемся через портал в нашу бытовку...")
}
