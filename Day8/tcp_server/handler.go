package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func handleRequest(conn net.Conn) {
	defer conn.Close()
	
	clientAddr := conn.RemoteAddr().String()
	fmt.Printf("✅ Новое подключение: %s\n", clientAddr)
	
	// Отправляем приветствие
	conn.Write([]byte("Добро пожаловать на TCP-сервер!\n"))
	conn.Write([]byte("Доступные команды:\n"))
	conn.Write([]byte("  time   - текущее время\n"))
	conn.Write([]byte("  echo   - эхо-ответ\n"))
	conn.Write([]byte("  upper  - преобразовать в верхний регистр\n"))
	conn.Write([]byte("  quit   - отключиться\n\n"))
	
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		
		switch strings.ToLower(input) {
		case "quit":
			conn.Write([]byte("👋 До свидания!\n"))
			fmt.Printf("❌ Отключение: %s\n", clientAddr)
			return
			
		case "time":
			currentTime := time.Now().Format("2006-01-02 15:04:05")
			response := fmt.Sprintf("⏱ Текущее время: %s\n", currentTime)
			conn.Write([]byte(response))
			
		case "echo":
			conn.Write([]byte("🔊 Введите текст для эхо-ответа: "))
			scanner.Scan()
			echoText := scanner.Text()
			conn.Write([]byte(fmt.Sprintf("🔔 Эхо: %s\n", echoText)))
			
		case "upper":
			conn.Write([]byte("🔼 Введите текст для преобразования: "))
			scanner.Scan()
			upperText := strings.ToUpper(scanner.Text())
			conn.Write([]byte(fmt.Sprintf("🔠 Результат: %s\n", upperText)))
			
		case "":
			// Игнорируем пустые строки
			continue
			
		default:
			conn.Write([]byte("❌ Неизвестная команда. Попробуйте снова.\n"))
		}
		
		// Добавляем промпт для следующей команды
		conn.Write([]byte("\n> "))
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка чтения: %v\n", err)
	}
}