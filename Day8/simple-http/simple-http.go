package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет, Golang-программист Гоша!")
	w.Write([]byte("Сегодня 2 августа 2025 года мы с тобою изучаем обслуживание HTTP-запросов на Go!"))
}

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
