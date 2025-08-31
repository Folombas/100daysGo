package server

import (
	"bufio"
	"fmt"
	"net"
)

func StartTCPServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Printf("Ошибка запуска TCP сервера: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Println("🔌 TCP сервер запущен на localhost:8081")
	fmt.Println("📨 Ожидание подключений...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка принятия соединения: %v\n", err)
			continue
		}

		go handleTCPConnection(conn)
	}
}

func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("✅ Новое подключение: %s\n", clientAddr)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Printf("📩 Получено от %s: %s\n", clientAddr, message)

		response := fmt.Sprintf("Сервер получил: %s (длина: %d символов)\n", message, len(message))
		if _, err := conn.Write([]byte(response)); err != nil {
			fmt.Printf("Ошибка отправки ответа: %v\n", err)
			return
		}

		if message == "exit" {
			fmt.Printf("🔌 Соединение с %s закрыто\n", clientAddr)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения: %v\n", err)
	}
}