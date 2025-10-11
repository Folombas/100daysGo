package main

import (
	"fmt"
	"sync"
	"time"
)

// BrazilianDancer представляет танцора самбы
type BrazilianDancer struct {
	Name     string
	Style    string
	Energy   int
}

// NewDancer создает нового танцора
func NewDancer(name, style string) *BrazilianDancer {
	return &BrazilianDancer{
		Name:   name,
		Style:  style,
		Energy: 100,
	}
}

// Dance заставляет танцора танцевать - работает в горутине!
func (d *BrazilianDancer) Dance(wg *sync.WaitGroup, moves chan string) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		d.Energy -= 10
		move := fmt.Sprintf("%s исполняет %s (движение %d, энергия: %d%%)",
			d.Name, d.Style, i, d.Energy)
		moves <- move
		time.Sleep(time.Millisecond * 500) // Имитация времени на движение
	}

	fmt.Printf("💃 %s завершил танец!\n", d.Name)
}

// DanceGroup представляет группу танцоров
type DanceGroup struct {
	Name    string
	Dancers []*BrazilianDancer
}

// StartDance начинает групповой танец с использованием горутин
func (dg *DanceGroup) StartDance() {
	fmt.Printf("\n🎭 Группа '%s' начинает самбу!\n", dg.Name)

	var wg sync.WaitGroup
	moves := make(chan string, 10) // Буферизованный канал для движений

	// Запускаем всех танцоров в отдельных горутинах
	for _, dancer := range dg.Dancers {
		wg.Add(1)
		go dancer.Dance(&wg, moves)
	}

	// Горутина для чтения движений из канала
	go func() {
		wg.Wait()
		close(moves)
	}()

	// Читаем и выводим движения
	for move := range moves {
		fmt.Printf("   🎵 %s\n", move)
	}

	fmt.Printf("✅ Группа '%s' завершила танец!\n\n", dg.Name)
}

// SyncDance демонстрирует синхронизированный танец
func SyncDance() {
	fmt.Println("🔄 СИНХРОНИЗИРОВАННЫЙ ТАНЕЦ:")
	fmt.Println("==========================")

	group := DanceGroup{
		Name: "Синхронные звезды",
		Dancers: []*BrazilianDancer{
			NewDancer("Карлос", "самба-де-гафиейра"),
			NewDancer("Мария", "самба-акробатика"),
			NewDancer("Жуан", "пагоде"),
		},
	}

	group.StartDance()
}

// AsyncDance демонстрирует асинхронный танец
func AsyncDance() {
	fmt.Println("⚡ АСИНХРОННЫЙ ТАНЕЦ:")
	fmt.Println("====================")

	dancers := []*BrazilianDancer{
		NewDancer("Антонио", "фрево"),
		NewDancer("Изабелла", "маракату"),
		NewDancer("Педро", "коку"),
	}

	var wg sync.WaitGroup
	results := make(chan string, 3)

	for _, dancer := range dancers {
		wg.Add(1)
		go func(d *BrazilianDancer) {
			defer wg.Done()
			for i := 1; i <= 3; i++ {
				result := fmt.Sprintf("%s танцует %s - движение %d",
					d.Name, d.Style, i)
				results <- result
				time.Sleep(time.Millisecond * 300)
			}
		}(dancer)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("   🎭 %s\n", result)
	}
	fmt.Println("✅ Асинхронный танец завершен!\n")
}

// ChannelDance демонстрирует работу с каналами
func ChannelDance() {
	fmt.Println("📡 ТАНЕЦ С КАНАЛАМИ:")
	fmt.Println("===================")

	danceMoves := make(chan string, 5)
	done := make(chan bool)

	// Горутина-танцор
	go func() {
		moves := []string{"кружение", "шаг самбы", "волна", "прыжок", "финал"}
		for _, move := range moves {
			danceMoves <- move
			time.Sleep(time.Millisecond * 400)
		}
		close(danceMoves)
	}()

	// Горутина-зритель
	go func() {
		for move := range danceMoves {
			fmt.Printf("   👀 Зрители видят: %s\n", move)
		}
		done <- true
	}()

	<-done
	fmt.Println("✅ Танец с каналами завершен!\n")
}

// SelectDance демонстрирует select с каналами
func SelectDance() {
	fmt.Println("🎯 SELECT В ТАНЦЕ:")
	fmt.Println("=================")

	dance1 := make(chan string)
	dance2 := make(chan string)

	go func() {
		time.Sleep(300 * time.Millisecond)
		dance1 <- "Самба"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		dance2 <- "Форро"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-dance1:
			fmt.Printf("   💃 Получено: %s\n", msg1)
		case msg2 := <-dance2:
			fmt.Printf("   🎵 Получено: %s\n", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("   ⏰ Время вышло!")
		}
	}

	fmt.Println("✅ Select танец завершен!\n")
}

// WorkerPoolDance демонстрирует пул воркеров
func WorkerPoolDance() {
	fmt.Println("🏊 ПУЛ ТАНЦОРОВ-ВОРКЕРОВ:")
	fmt.Println("=======================")

	jobs := make(chan int, 10)
	results := make(chan string, 10)

	// Создаем пул воркеров-танцоров
	for w := 1; w <= 3; w++ {
		go dancerWorker(w, jobs, results)
	}

	// Отправляем задания
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// Собираем результаты
	for a := 1; a <= 9; a++ {
		fmt.Printf("   %s\n", <-results)
	}

	fmt.Println("✅ Пул танцоров завершил работу!\n")
}

func dancerWorker(id int, jobs <-chan int, results chan<- string) {
	for j := range jobs {
		time.Sleep(500 * time.Millisecond)
		results <- fmt.Sprintf("Танцор %d исполнил движение %d", id, j)
	}
}

func main() {
	fmt.Println("💃 Day 78: Goroutines - Бразильская самба параллелизма!")
	fmt.Println("======================================================")

	// Демонстрация различных подходов
	SyncDance()      // Синхронизированные горутины
	AsyncDance()     // Асинхронное выполнение
	ChannelDance()   // Работа с каналами
	SelectDance()    // Select с множеством каналов
	WorkerPoolDance() // Пул воркеров

	// Финальное шоу
	fmt.Println("🎉 ФИНАЛЬНОЕ ШОУ ГОРУТИН:")
	fmt.Println("========================")

	finalGroup := DanceGroup{
		Name: "Финальный карнавал",
		Dancers: []*BrazilianDancer{
			NewDancer("Рио", "карнавальная самба"),
			NewDancer("Сан-Паулу", "аши-и-кса"),
			NewDancer("Баия", "аше"),
			NewDancer("Ресифи", "фрево"),
		},
	}

	finalGroup.StartDance()

	// Итоги обучения
	fmt.Println("🎯 ЧТО МЫ ИЗУЧИЛИ:")
	fmt.Println("   • Горутины - легковесные потоки выполнения")
	fmt.Println("   • sync.WaitGroup для ожидания завершения")
	fmt.Println("   • Каналы для связи между горутинами")
	fmt.Println("   • Select для работы с множеством каналов")
	fmt.Println("   • Пул воркеров для ограничения параллелизма")
	fmt.Println("   • Буферизованные и небуферизованные каналы")

	fmt.Println("\n💪 Отлично! Теперь твои горутины танцуют как бразильские профессионалы!")
}

