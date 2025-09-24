package main

import (
	"fmt"
	"strings"
)

func demoMapPatterns() {
	printSeparator()

	fmt.Println("📊 Подсчет частот слов:")
	text := "привет мир привет гофер привет программирование"
	words := strings.Fields(text)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("📝 Текст: '%s'\n", text)
	fmt.Println("📈 Частоты слов:")
	for word, count := range wordCount {
		fmt.Printf("  '%s': %d раз\n", word, count)
	}

	fmt.Println("\n🎮 Группировка данных:")
	type Game struct {
		Name     string
		Genre    string
		Platform string
	}

	games := []Game{
		{"The Witcher 3", "RPG", "PC/PS/Xbox"},
		{"Cyberpunk 2077", "RPG", "PC/PS/Xbox"},
		{"Mario Kart", "Racing", "Switch"},
		{"FIFA 23", "Sports", "PC/PS/Xbox"},
		{"Zelda", "Adventure", "Switch"},
	}

	// Группировка по жанру
	genreGroups := make(map[string][]string)
	for _, game := range games {
		genreGroups[game.Genre] = append(genreGroups[game.Genre], game.Name)
	}

	fmt.Println("🎯 Группировка по жанру:")
	for genre, gameList := range genreGroups {
		fmt.Printf("  %s: %v\n", genre, gameList)
	}

	// Кэширование результатов
	fmt.Println("\n💾 Кэширование (мемоизация):")
	fibCache := make(map[int]int)

	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		if result, exists := fibCache[n]; exists {
			return result
		}
		result := fibonacci(n-1) + fibonacci(n-2)
		fibCache[n] = result
		return result
	}

	fmt.Printf("🔢 Числа Фибоначчи: ")
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
	fmt.Println()

	fmt.Printf("📚 Кэш содержит %d элементов: %v\n", len(fibCache), fibCache)
}
