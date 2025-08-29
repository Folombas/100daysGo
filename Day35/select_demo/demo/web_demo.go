package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Регистрируем обработчики
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/basic", basicHandler)
	http.HandleFunc("/default", defaultHandler)
	http.HandleFunc("/timeout", timeoutHandler)
	http.HandleFunc("/multiplex", multiplexHandler)

	fmt.Println("Веб-сервер запущен на http://localhost:8080")
	fmt.Println("Прервать работу: Ctrl+C")
	
	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
	<title>Select и Default в Go</title>
	<style>
		body { font-family: Arial, sans-serif; margin: 40px; }
		.container { max-width: 800px; margin: 0 auto; }
		.example { background: #f5f5f5; padding: 20px; margin: 20px 0; border-radius: 8px; }
		.result { background: #e0f7fa; padding: 15px; border-radius: 5px; }
	</style>
</head>
<body>
	<div class="container">
		<h1>Демонстрация select и default в Go</h1>
		
		<div class="example">
			<h2>Базовый select</h2>
			<p><a href="/basic">Запустить базовый пример</a></p>
		</div>
		
		<div class="example">
			<h2>Default в select</h2>
			<p><a href="/default">Запустить пример с default</a></p>
		</div>
		
		<div class="example">
			<h2>Таймауты</h2>
			<p><a href="/timeout">Запустить пример с таймаутом</a></p>
		</div>
		
		<div class="example">
			<h2>Мультиплексирование</h2>
			<p><a href="/multiplex">Запустить мультиплексирование</a></p>
		</div>
	</div>
</body>
</html>`
	
	fmt.Fprint(w, html)
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Millisecond * 500)
		ch1 <- "сообщение из быстрого канала"
	}()

	go func() {
		time.Sleep(time.Millisecond * 1000)
		ch2 <- "сообщение из медленного канала"
	}()

	var result string
	select {
	case msg := <-ch1:
		result = msg
	case msg := <-ch2:
		result = msg
	}

	fmt.Fprintf(w, `<div class="result">Результат: %s</div><p><a href="/">Назад</a></p>`, result)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)

	var result string
	select {
	case msg := <-ch:
		result = msg
	default:
		result = "default: канал пуст"
	}

	fmt.Fprintf(w, `<div class="result">Результат: %s</div><p><a href="/">Назад</a></p>`, result)
}

func timeoutHandler(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	
	go func() {
		time.Sleep(time.Millisecond * 1500)
		ch <- "результат операции"
	}()

	var result string
	select {
	case res := <-ch:
		result = res
	case <-time.After(time.Millisecond * 1000):
		result = "таймаут операции"
	}

	fmt.Fprintf(w, `<div class="result">Результат: %s</div><p><a href="/">Назад</a></p>`, result)
}

func multiplexHandler(w http.ResponseWriter, r *http.Request) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch1 <- i
			time.Sleep(time.Millisecond * 300)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			ch2 <- i * 10
			time.Sleep(time.Millisecond * 500)
		}
	}()

	result := "<h3>Мультиплексирование каналов:</h3>"
	for i := 0; i < 6; i++ {
		select {
		case val := <-ch1:
			result += fmt.Sprintf("<p>Канал 1: %d</p>", val)
		case val := <-ch2:
			result += fmt.Sprintf("<p>Канал 2: %d</p>", val)
		}
	}

	fmt.Fprintf(w, `<div class="result">%s</div><p><a href="/">Назад</a></p>`, result)
}