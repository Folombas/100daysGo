package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// CityNews представляет смешанную вселенную Гуфи и Арнольда
type CityNews struct {
	reporters []string
	topics    []string
}

func main() {
	fmt.Println("🌆 СМЕШАННАЯ ВСЕЛЕННАЯ: Гуфи встречает Эй, Арнольд!")
	fmt.Println("===================================================")

	city := &CityNews{
		reporters: []string{
			"Гуфи-репортёр",
			"Арнольд-журналист",
			"Хельга-корреспондент",
			"Гарольд-обозреватель",
			"Фил-комментатор",
		},
		topics: []string{
			"Соревнования по скейтбордингу",
			"Тайна пропавшего пирога",
			"Новый бизнес-план Гарольда",
			"Приключения Гуфи на работе",
			"Школьные новости",
			"Спортивные мероприятия",
			"Культурные события города",
		},
	}

	fmt.Println("📰 ГОРОДСКИЕ НОВОСТИ: Каждый репортёр собирает информацию отдельно...")
	fmt.Println("🔄 НО: Нам нужно объединить ВСЕ новости в один общий выпуск!")

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Запускаем Fan-in паттерн
	fmt.Println("\n🎯 ЗАПУСК FAN-IN ПАТТЕРНА...")
	finalNewsChannel := city.fanInNews(ctx)

	// Читаем объединенные новости
	fmt.Println("\n📻 В ЭФИРЕ: ОБЪЕДИНЕННЫЙ ВЫПУСК НОВОСТЕЙ:")
	fmt.Println("==========================================")

	for news := range finalNewsChannel {
		fmt.Printf("🎙️  %s\n", news)
	}

	fmt.Println("\n🎉 ВЫПУСК ЗАВЕРШЕН! Все новости собраны и обработаны!")
	fmt.Println("🏆 Гуфи и Арнольд становятся лучшей медиа-командой города!")
}

// fanInNews реализует Fan-in паттерн: множество источников -> один канал
func (c *CityNews) fanInNews(ctx context.Context) <-chan string {
	// Создаем отдельные каналы для каждого репортера
	reporterChannels := make([]<-chan string, len(c.reporters))

	for i, reporter := range c.reporters {
		reporterChannels[i] = c.reporterWork(ctx, reporter)
	}

	// Объединяем все каналы в один (FAN-IN)
	return c.mergeChannels(ctx, reporterChannels...)
}

// reporterWork имитирует работу репортера
func (c *CityNews) reporterWork(ctx context.Context, reporter string) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer fmt.Printf("   📋 %s завершил работу\n", reporter)

		// Каждый репортер готовит несколько новостей
		for i := 0; i < 3; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("   ⏰ %s: Время вышло, прекращаю работу!\n", reporter)
				return
			case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
				topic := c.topics[rand.Intn(len(c.topics))]
				news := fmt.Sprintf("%s: %s - эксклюзивный репортаж!", reporter, topic)
				out <- news
			}
		}
	}()

	return out
}

// mergeChannels объединяет несколько каналов в один (ядро Fan-in паттерна)
func (c *CityNews) mergeChannels(ctx context.Context, channels ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	merged := make(chan string)

	// Функция для переноса данных из одного канала в объединенный
	output := func(ch <-chan string) {
		defer wg.Done()
		for news := range ch {
			select {
			case <-ctx.Done():
				return
			case merged <- news:
			}
		}
	}

	wg.Add(len(channels))

	// Запускаем горутины для каждого входного канала
	for _, ch := range channels {
		go output(ch)
	}

	// Закрываем объединенный канал, когда все входные каналы закрыты
	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
