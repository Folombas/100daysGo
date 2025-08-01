package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func main() {
	// Запуск TCP сервера
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Printf("🚀 Сервер запущен на %s:%s\n", HOST, PORT)
	fmt.Println("⌛ Ожидание подключений...")

	for {
		// Принимаем входящие подключения
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Ошибка подключения: %v", err)
			continue
		}

		// Обработка подключения в отдельной горутине
		go handleRequest(conn)
	}
}
