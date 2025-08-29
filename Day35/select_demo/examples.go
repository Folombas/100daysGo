package main

import (
	"fmt"
	"time"
)

// basicSelectExamples демонстрирует базовое использование select
func basicSelectExamples() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- "сообщение из канала 1"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- "сообщение из канала 2"
	}()

	// select ожидает первый доступный канал
	select {
	case msg1 := <-ch1:
		fmt.Printf("Получено: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("Получено: %s\n", msg2)
	}
}

// defaultExamples демонстрирует использование default в select
func defaultExamples() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Printf("Получено сообщение: %s\n", msg)
	default:
		fmt.Println("Сообщений нет, выполняется default")
	}

	// Неблокирующая отправка
	select {
	case ch <- "сообщение":
		fmt.Println("Сообщение отправлено")
	default:
		fmt.Println("Никто не готов принять сообщение")
	}
}

// timeoutExamples демонстрирует использование таймаутов с select
func timeoutExamples() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- "результат операции"
	}()

	select {
	case res := <-ch:
		fmt.Printf("Операция завершена: %s\n", res)
	case <-time.After(time.Second * 2):
		fmt.Println("Таймаут операции")
	}
}

// multiplexingExamples демонстрирует мультиплексирование каналов
func multiplexingExamples() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 500)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i * 10
			time.Sleep(time.Millisecond * 300)
		}
		close(ch2)
	}()

	fmt.Println("Мультиплексирование каналов:")
	for i := 0; i < 6; i++ {
		select {
		case val, ok := <-ch1:
			if ok {
				fmt.Printf("Канал 1: %d\n", val)
			}
		case val, ok := <-ch2:
			if ok {
				fmt.Printf("Канал 2: %d\n", val)
			}
		}
	}
}

// infiniteLoopExamples демонстрирует бесконечные циклы с select
func infiniteLoopExamples() {
	ch := make(chan int)
	quit := make(chan bool)

	// Генератор данных
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 400)
		}
		quit <- true
	}()

	fmt.Println("Бесконечный цикл с select:")
	for {
		select {
		case val := <-ch:
			fmt.Printf("Получено значение: %d\n", val)
		case <-quit:
			fmt.Println("Завершение работы")
			return
		case <-time.After(time.Second * 1):
			fmt.Println("Таймаут ожидания данных")
			return
		}
	}
}

// advancedSelect демонстрирует продвинутые техники select
func advancedSelect() {
	ch := make(chan int)
	ticker := time.NewTicker(time.Millisecond * 500)
	defer ticker.Stop()

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Millisecond * 800)
		}
		close(ch)
	}()

	fmt.Println("Продвинутый select с тикером:")
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Printf("Данные: %d\n", val)
		case <-ticker.C:
			fmt.Println("Тик")
		}
	}
}

// selectWithPriority демонстрирует приоритизацию каналов
func selectWithPriority() {
	highPriority := make(chan string)
	lowPriority := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			lowPriority <- fmt.Sprintf("низкий приоритет %d", i)
			time.Sleep(time.Millisecond * 300)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			highPriority <- fmt.Sprintf("высокий приоритет %d", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	fmt.Println("Приоритизация каналов:")
	for i := 0; i < 6; i++ {
		select {
		case msg := <-highPriority:
			fmt.Printf("Обработано: %s\n", msg)
		case msg := <-lowPriority:
			fmt.Printf("Обработано: %s\n", msg)
		}
	}
}

// randomSelect демонстрирует случайный выбор при множестве готовых каналов
func randomSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// Заполняем оба канала одновременно
	go func() {
		ch1 <- "канал 1"
	}()
	go func() {
		ch2 <- "канал 2"
	}()

	// Go случайным образом выберет готовый канал
	select {
	case msg := <-ch1:
		fmt.Printf("Выбран: %s\n", msg)
	case msg := <-ch2:
		fmt.Printf("Выбран: %s\n", msg)
	}
}