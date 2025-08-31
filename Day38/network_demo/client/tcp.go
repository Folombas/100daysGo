package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func StartTCPClient() {
	fmt.Println("🔌 Запуск TCP клиента...")
	fmt.Println("Подключение к localhost:8081")
	fmt.Println("Введите сообщения для отправки (или 'exit' для выхода):")

	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Printf("❌ Ошибка подключения: %v\n", err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		if message == "" {
			continue
		}

		// Отправка сообщения
		fmt.Fprintf(conn, message+"\n")

		// Чтение ответа
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("❌ Ошибка чтения ответа: %v\n", err)
			break
		}

		fmt.Printf("📥 Ответ сервера: %s", response)

		if message == "exit" {
			break
		}
	}

	fmt.Println("🔌 Отключение от TCP сервера")
}