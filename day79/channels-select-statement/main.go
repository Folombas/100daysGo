package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fisherman представляет рыбака с несколькими удочками
type Fisherman struct {
	Name string
}

// NewFisherman создает нового рыбака
func NewFisherman(name string) *Fisherman {
	return &Fisherman{Name: name}
}

// Fish представляет пойманную рыбу
type Fish struct {
	Type string
	Size int
}

// FishWithRod симулирует рыбалку с одной удочкой
func (f *Fisherman) FishWithRod(rodNumber int, catch chan<- Fish, quit <-chan bool) {
	fmt.Printf("🎣 %s закинул удочку #%d\n", f.Name, rodNumber)

	for {
		select {
		case <-quit:
			fmt.Printf("   🛑 %s убрал удочку #%d\n", f.Name, rodNumber)
			return
		default:
			// Симуляция времени ожидания поклевки
			waitTime := time.Duration(rand.Intn(3000)+1000) * time.Millisecond
			time.Sleep(waitTime)

			// Поймали рыбу!
			fishTypes := []string{"карп", "щука", "окунь", "лещ", "сом"}
			fish := Fish{
				Type: fishTypes[rand.Intn(len(fishTypes))],
				Size: rand.Intn(50) + 10, // размер в см
			}

			fmt.Printf("   🐟 Удочка #%d: поймал %s (%dсм)\n", rodNumber, fish.Type, fish.Size)
			catch <- fish
		}
	}
}

// StartFishing начинает рыбалку с несколькими удочками
func (f *Fisherman) StartFishing(rods int, duration time.Duration) []Fish {
	fmt.Printf("\n🚤 %s начинает рыбалку с %d удочками на %v\n", f.Name, rods, duration)

	catch := make(chan Fish, 10) // Буферизованный канал для улова
	quit := make(chan bool)      // Канал для остановки
	var caught []Fish

	// Запускаем горутины для каждой удочки
	for i := 1; i <= rods; i++ {
		go f.FishWithRod(i, catch, quit)
	}

	// Таймер для всей рыбалки
	timer := time.NewTimer(duration)

	// Собираем улов с помощью select
FishingLoop:
	for {
		select {
		case fish := <-catch:
			caught = append(caught, fish)
			fmt.Printf("   📦 Добавлен в ведро: %s\n", fish.Type)

		case <-timer.C:
			fmt.Printf("\n⏰ Время рыбалки вышло! %s заканчивает.\n", f.Name)
			break FishingLoop
		}
	}

	// Останавливаем все удочки
	for i := 0; i < rods; i++ {
		quit <- true
	}

	time.Sleep(500 * time.Millisecond) // Даем время завершиться горутинам
	return caught
}

// CookUha готовит уху из пойманной рыбы
func CookUha(fishes []Fish, done chan<- string) {
	fmt.Printf("\n🍲 Начинаем готовить уху из %d рыб...\n", len(fishes))

	// Симуляция готовки
	stages := []string{"чистка рыбы", "варка бульона", "добавление овощей", "добавление грибов", "томление"}

	for _, stage := range stages {
		time.Sleep(1 * time.Second)
		fmt.Printf("   👨‍🍳 %s...\n", stage)
	}

	time.Sleep(2 * time.Second)
	done <- fmt.Sprintf("🍜 Уха готова! Из рыб: %d", len(fishes))
}

// SelectWithTimeout демонстрирует select с таймаутом
func SelectWithTimeout() {
	fmt.Println("\n⏰ SELECT С ТАЙМАУТОМ:")
	fmt.Println("====================")

	workChan := make(chan string)
	timeout := time.After(3 * time.Second)

	go func() {
		time.Sleep(5 * time.Second) // Работа дольше таймаута
		workChan <- "работа завершена"
	}()

	select {
	case result := <-workChan:
		fmt.Printf("   ✅ %s\n", result)
	case <-timeout:
		fmt.Printf("   ❌ Таймаут! Работа не завершена вовремя\n")
	}
}

// SelectWithDefault демонстрирует select с default
func SelectWithDefault() {
	fmt.Println("\n⚡ SELECT С DEFAULT:")
	fmt.Println("===================")

	tick := time.Tick(500 * time.Millisecond)
	boom := time.After(3 * time.Second)

	fmt.Println("   Ожидание сигналов...")

	for {
		select {
		case <-tick:
			fmt.Printf("   💥 тик\n")
		case <-boom:
			fmt.Printf("   💣 БУМ! Время вышло!\n")
			return
		default:
			// Выполняется, если другие каналы не готовы
			fmt.Printf("   😴 спим...\n")
			time.Sleep(200 * time.Millisecond)
		}
	}
}

// MultipleChannelsSelect демонстрирует работу с несколькими каналами
func MultipleChannelsSelect() {
	fmt.Println("\n🎯 SELECT С НЕСКОЛЬКИМИ КАНАЛАМИ:")
	fmt.Println("================================")

	channel1 := make(chan string)
	channel2 := make(chan string)
	channel3 := make(chan string)

	// Запускаем горутины, которые отправляют в разные каналы
	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "сообщение из канала 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "сообщение из канала 2"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		channel3 <- "сообщение из канала 3"
	}()

	// Ожидаем сообщения из всех каналов
	for i := 0; i < 3; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Printf("   📨 %s\n", msg1)
		case msg2 := <-channel2:
			fmt.Printf("   📨 %s\n", msg2)
		case msg3 := <-channel3:
			fmt.Printf("   📨 %s\n", msg3)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("🎣 Day 79: Select Statement - Рыбалка в мире каналов!")
	fmt.Println("====================================================")

	// Основная рыбалка
	fisherman := NewFisherman("Семён")
	caughtFish := fisherman.StartFishing(3, 8*time.Second)

	// Готовим уху параллельно с выводом статистики
	uhaDone := make(chan string)

	fmt.Printf("\n📊 СТАТИСТИКА УЛОВА:\n")
	fmt.Printf("   • Всего поймано: %d рыб\n", len(caughtFish))

	fishCount := make(map[string]int)
	for _, fish := range caughtFish {
		fishCount[fish.Type]++
	}

	for fishType, count := range fishCount {
		fmt.Printf("   • %s: %d шт\n", fishType, count)
	}

	// Готовим уху в отдельной горутине
	go CookUha(caughtFish, uhaDone)

	// Демонстрация различных вариантов select
	SelectWithTimeout()
	SelectWithDefault()
	MultipleChannelsSelect()

	// Ждем, пока уха сварится
	select {
	case result := <-uhaDone:
		fmt.Printf("\n%s\n", result)
	case <-time.After(10 * time.Second):
		fmt.Printf("\n⏰ Уха не сварилась вовремя!\n")
	}

	// Итоги обучения
	fmt.Println("\n🎯 ЧТО МЫ ИЗУЧИЛИ:")
	fmt.Println("   • Select - ожидание нескольких каналов одновременно")
	fmt.Println("   • Таймауты с time.After()")
	fmt.Println("   • Тикеры с time.Tick()")
	fmt.Println("   • Default case для неблокирующих операций")
	fmt.Println("   • Паттерны отмены с quit-каналами")
	fmt.Println("   • Обработка множества каналов в цикле")

	fmt.Println("\n💪 Отлично! Теперь твой select ловит сообщения как опытный рыбак!")
}

