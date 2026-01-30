package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Жизненные роли Гоши
type LifeRole string

const (
	CourierRole    LifeRole = "🚚 Курьер"
	StudentRole    LifeRole = "👨‍💻 Студент Go"
	ProgrammerRole LifeRole = "💼 Программист"
	EditorRole     LifeRole = "🎬 Монтажер"
)

// Событие в жизни
type LifeEvent struct {
	ID          int
	Role        LifeRole
	Description string
	Duration    time.Duration
	Reward      int // XP или деньги
	Priority    int // 1-10
	Completed   bool
}

// Каналы жизни Гоши
type LifeChannels struct {
	WorkChan      chan LifeEvent // Работа курьером
	StudyChan     chan LifeEvent // Учеба Go
	TempationChan chan string    // Искушения отвлечься
	DecisionChan  chan string    // Решения
	ProgressChan  chan int       // Прогресс обучения
}

func main() {
	fmt.Println("🎯 День 87: SELECT ALGORITHM LIFE 🎯")
	fmt.Println("Философия выбора в конкурентной жизни")
	fmt.Println(strings.Repeat("=", 70) + "\n")

	rand.Seed(time.Now().UnixNano())

	// Инициализация каналов жизни
	channels := LifeChannels{
		WorkChan:      make(chan LifeEvent, 5), // Работа может накапливаться
		StudyChan:     make(chan LifeEvent, 3), // Учеба в приоритете
		TempationChan: make(chan string, 10),   // Искушения везде
		DecisionChan:  make(chan string, 5),    // Решения принимаем быстро
		ProgressChan:  make(chan int, 100),     // Прогресс обучения
	}

	// Статистика жизни
	stats := struct {
		Money               int
		XP                  int
		DaysPassed          int
		CourierJobs         int
		StudySessions       int
		TemptationsResisted int
		DecisionsMade       int
	}{
		Money:      500,
		XP:         2900, // 29 дней уже учит
		DaysPassed: 87,
	}

	var wg sync.WaitGroup
	output := make(chan string, 50)

	// Горутина: генерация работы курьером
	wg.Add(1)
	go func() {
		defer wg.Done()
		jobs := []LifeEvent{
			{1, CourierRole, "Доставить заказ из центра на окраину", 2 * time.Hour, 300, 3, false},
			{2, CourierRole, "Забрать три посылки со склада", 90 * time.Minute, 450, 4, false},
			{3, CourierRole, "Срочная доставка документов", 45 * time.Minute, 200, 7, false},
			{4, CourierRole, "Развезти еду по офисам", 3 * time.Hour, 600, 5, false},
			{5, CourierRole, "Межгородская посылка", 4 * time.Hour, 800, 6, false},
		}

		for i := 0; i < len(jobs); i++ {
			time.Sleep(time.Duration(rand.Intn(1500)+800) * time.Millisecond)
			channels.WorkChan <- jobs[i]
			output <- fmt.Sprintf("📦 ПОСТУПИЛА РАБОТА: %s (+%d руб.)",
				jobs[i].Description, jobs[i].Reward)
		}
		close(channels.WorkChan)
	}()

	// Горутина: учебные задания по Go
	wg.Add(1)
	go func() {
		defer wg.Done()
		studies := []LifeEvent{
			{6, StudentRole, "Изучить select statement", 60 * time.Minute, 100, 9, false},
			{7, StudentRole, "Практика с каналами", 45 * time.Minute, 80, 8, false},
			{8, StudentRole, "Чтение документации по контекстам", 30 * time.Minute, 50, 7, false},
			{9, StudentRole, "Решение задач на LeetCode", 90 * time.Minute, 120, 8, false},
			{10, StudentRole, "Просмотр видео-курса", 120 * time.Minute, 150, 6, false},
		}

		for i := 0; i < len(studies); i++ {
			time.Sleep(time.Duration(rand.Intn(2000)+1000) * time.Millisecond)
			channels.StudyChan <- studies[i]
			output <- fmt.Sprintf("📚 УЧЕБНОЕ ЗАДАНИЕ: %s (+%d XP)",
				studies[i].Description, studies[i].Reward)
		}
		close(channels.StudyChan)
	}()

	// Горутина: искушения отвлечься
	wg.Add(1)
	go func() {
		defer wg.Done()
		temptations := []string{
			"🎮 'А может, поиграть в новую игру?'",
			"🎬 'Посмотреть новый сериал на Netflix?'",
			"🍺 'Сходить в бар с друзьями?'",
			"✈️  'Смонтировать видео из отпуска 2019 года?'",
			"🛌 'Просто поспать подольше?'",
			"📱 'Поскроллить соцсети часок?'",
			"🛒 'Пойти по магазинам без повода?'",
			"🎵 'Послушать музыку вместо учебы?'",
		}

		for i := 0; i < 8; i++ {
			time.Sleep(time.Duration(rand.Intn(2500)+1500) * time.Millisecond)
			temptation := temptations[rand.Intn(len(temptations))]
			channels.TempationChan <- temptation
			output <- fmt.Sprintf("😈 ИСКУШЕНИЕ: %s", temptation)
		}
		close(channels.TempationChan)
	}()

	// Горутина: мотивационные решения
	wg.Add(1)
	go func() {
		defer wg.Done()
		decisions := []string{
			"💪 'Нет! Сначала учеба, потом развлечения!'",
			"🎯 'Алгоритм жизни: работа → учеба → карьера → хобби'",
			"🚫 'CapCut подождет! Сначала трудоустройство!'",
			"⚡ 'Каждый день с Go - инвестиция в будущее!'",
			"🧠 'Лучше потратить время на LeetCode, чем на Netflix!'",
			"🏆 'Когда устроюсь программистом, буду монтировать видео по выходным!'",
			"💡 'Select в жизни: выбираю учебу вместо развлечений!'",
			"🌟 'Моя цель - тёплый офис, а не беготня с коробками по оледеневшим улицам!'",
		}

		for i := 0; i < 6; i++ {
			time.Sleep(time.Duration(rand.Intn(3000)+2000) * time.Millisecond)
			decision := decisions[rand.Intn(len(decisions))]
			channels.DecisionChan <- decision
			stats.DecisionsMade++
		}
		close(channels.DecisionChan)
	}()

	// Главная горутина: алгоритм жизни через select
	wg.Add(1)
	go func() {
		defer wg.Done()

		output <- "\n🧬 ЗАПУСК АЛГОРИТМА ЖИЗНИ:"
		output <- "Подработка → Учеба → Работа программистом → Хобби(видеосъёмка/видеомонтаж)\n"

		workActive := true
		studyActive := true
		temptationActive := true
		decisionActive := true

		// Обработка событий с помощью select
		for workActive || studyActive || temptationActive || decisionActive {
			select {
			case job, ok := <-channels.WorkChan:
				if !ok {
					workActive = false
					channels.WorkChan = nil
				} else {
					// Обработка работы
					time.Sleep(job.Duration / 10) // Ускоренная симуляция
					stats.Money += job.Reward
					stats.CourierJobs++
					output <- fmt.Sprintf("✅ ВЫПОЛНЕНА РАБОТА: %s. Баланс: %d руб.",
						job.Description, stats.Money)

					// После работы - учеба (приоритет)
					channels.ProgressChan <- 10
				}

			case study, ok := <-channels.StudyChan:
				if !ok {
					studyActive = false
					channels.StudyChan = nil
				} else {
					// Обработка учебы
					time.Sleep(study.Duration / 10)
					stats.XP += study.Reward
					stats.StudySessions++
					output <- fmt.Sprintf("🎓 ВЫПОЛНЕНО УЧЕБНОЕ ЗАДАНИЕ: %s. XP: %d",
						study.Description, stats.XP)

					// Прогресс обучения
					channels.ProgressChan <- study.Reward
				}

			case temptation, ok := <-channels.TempationChan:
				if !ok {
					temptationActive = false
					channels.TempationChan = nil
				} else {
					// Искушение - нужен выбор
					output <- fmt.Sprintf("⚖️  ПЕРЕД ВЫБОРОМ: %s", temptation)

					// Имитация внутренней борьбы
					time.Sleep(500 * time.Millisecond)

					// 80% шанс устоять
					if rand.Intn(100) < 80 {
						stats.TemptationsResisted++
						output <- "✅ УСТОЯЛ ПЕРЕД ИСКУШЕНИЕМ! +10 к силе воли"
						channels.ProgressChan <- 5
					} else {
						output <- "⚠️  НА МИНУТУ ОТВЛЁКСЯ... но быстро вернулся к учебе"
					}
				}

			case decision, ok := <-channels.DecisionChan:
				if !ok {
					decisionActive = false
					channels.DecisionChan = nil
				} else {
					// Принятие решения укрепляет
					output <- fmt.Sprintf("🧠 РЕШЕНИЕ: %s", decision)
					stats.XP += 20
				}

			case <-time.After(2 * time.Second):
				// Таймаут - время на размышления
				if workActive || studyActive {
					output <- "⏱️  Размышляю о выборе между работой и учебой..."
				}
			}

			// Проверка на достижение цели
			if stats.XP >= 3500 {
				output <- "\n🎉 ЦЕЛЬ ДОСТИГНУТА: 3500+ XP по Go!"
				output <- "🏢 Можно готовиться к собеседованиям на программиста!"
				break
			}
		}

		// Завершение алгоритма
		close(channels.ProgressChan)
	}()

	// Горутина: отслеживание прогресса
	wg.Add(1)
	go func() {
		defer wg.Done()
		totalProgress := 0
		for progress := range channels.ProgressChan {
			totalProgress += progress
			if totalProgress%100 == 0 {
				output <- fmt.Sprintf("📈 ПРОГРЕСС ОБУЧЕНИЯ: %d/3500 XP", stats.XP)
			}
		}
	}()

	// Вывод событий в реальном времени
	go func() {
		wg.Wait()
		close(output)
	}()

	// Отображение событий
	fmt.Println("📖 ЖИЗНЬ В РЕАЛЬНОМ ВРЕМЕНИ:")
	fmt.Println(strings.Repeat("-", 70))

	for event := range output {
		fmt.Println(event)
		time.Sleep(300 * time.Millisecond)
	}

	// Итоги дня
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("📊 ИТОГИ ДНЯ 87:")
	fmt.Printf("💰 Заработано: %d руб. (курьерская работа)\n", stats.Money)
	fmt.Printf("🧠 Накоплено XP: %d/3500 (изучение Go)\n", stats.XP)
	fmt.Printf("📦 Выполнено заказов: %d\n", stats.CourierJobs)
	fmt.Printf("📚 Учебных сессий: %d\n", stats.StudySessions)
	fmt.Printf("🚫 Устоял перед искушениями: %d раз\n", stats.TemptationsResisted)
	fmt.Printf("🎯 Принято решений: %d\n", stats.DecisionsMade)

	// Геймификационные достижения
	fmt.Println("\n🏆 ДОСТИЖЕНИЯ ДНЯ:")
	if stats.StudySessions >= 3 {
		fmt.Println("   🥇 'Неутомимый студент' - 3+ учебные сессии")
	}
	if stats.TemptationsResisted >= 5 {
		fmt.Println("   🥈 'Железная воля' - устоял перед 5+ искушениями")
	}
	if stats.XP-stats.StudySessions*50 > 100 {
		fmt.Println("   🥉 'Эффективный learner' - высокий КПД обучения")
	}

	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("💡 ФИЛОСОФИЯ SELECT В ЖИЗНИ:")
	fmt.Println("Select statement в Go учит нас:")
	fmt.Println("1. Слушать несколько каналов событий одновременно")
	fmt.Println("2. Выбирать самое важное в данный момент")
	fmt.Println("3. Не блокироваться на одном деле, если есть другие возможности")
	fmt.Println("4. Использовать таймауты для предотвращения вечного ожидания")
	fmt.Println("5. Закрывать каналы, когда работа завершена")

	fmt.Println("\n🚀 ЗАВТРА: День 88 - Context и отмена горутин!")
	fmt.Println(strings.Repeat("=", 70))
}
