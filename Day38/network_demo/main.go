package main

import (
	"fmt"
	"network_demo/client"
	"network_demo/server"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "http-server":
		server.StartHTTPServer()
	case "http-client":
		client.StartHTTPClient()
	case "tcp-server":
		server.StartTCPServer()
	case "tcp-client":
		client.StartTCPClient()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Использование: network_demo [команда]")
	fmt.Println("Команды:")
	fmt.Println("  http-server - запустить HTTP сервер на :8080")
	fmt.Println("  http-client - запустить HTTP клиент")
	fmt.Println("  tcp-server  - запустить TCP сервер на :8081")
	fmt.Println("  tcp-client  - запустить TCP клиент")
	fmt.Println("Пример:")
	fmt.Println("  go run main.go http-server")
}