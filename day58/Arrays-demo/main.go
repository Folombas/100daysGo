package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := setupRoutes()

	fmt.Printf("🚀 Сервер запущен на http://localhost:%s\n", port)
	fmt.Println("✨ Откройте браузер и наслаждайтесь изучением массивов!")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
