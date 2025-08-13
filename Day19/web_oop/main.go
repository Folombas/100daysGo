package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    port := ":8080"
    controller := &WebController{}
    
    // Статические файлы
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    // Роуты
    http.HandleFunc("/", controller.HomeHandler)
    http.HandleFunc("/animal", controller.AnimalHandler)
    
    fmt.Printf("Сервер запущен на http://localhost%s\n", port)
    fmt.Println("Доступные животные: кошка, собака, корова, утка, петух, свинья, мышь, лягушка, ворона, кукушка")
    log.Fatal(http.ListenAndServe(port, nil))
}