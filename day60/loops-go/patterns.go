package main

import (
	"fmt"
	"strings"
	"time"
)

func demoPatterns() {
	printSeparator()

	fmt.Println("🎄 Рисуем ёлку циклом:")
	height := 5
	for i := 1; i <= height; i++ {
		spaces := strings.Repeat(" ", height-i)
		stars := strings.Repeat("🌟", i*2-1)
		fmt.Printf("%s%s\n", spaces, stars)
	}
	trunk := strings.Repeat(" ", height-1) + "🎄"
	fmt.Println(trunk)

	fmt.Println("\n📊 Гистограмма чисел:")
	data := []int{3, 7, 2, 5, 9}
	for _, value := range data {
		bar := strings.Repeat("█", value) + strings.Repeat("░", 10-value)
		fmt.Printf("%2d: %s\n", value, bar)
	}

	fmt.Println("\n🎰 Анимация вращения:")
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	for i := 0; i < 20; i++ {
		frame := frames[i%len(frames)]
		fmt.Printf("\r%s Загрузка... %d%%", frame, i*5)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("\r✅ Загрузка завершена! 100%")

	fmt.Println("\n🔢 Таблица умножения:")
	fmt.Println("   | 1  2  3  4  5  6  7  8  9")
	fmt.Println("---+---------------------------")
	for i := 1; i <= 9; i++ {
		fmt.Printf("%2d |", i)
		for j := 1; j <= 9; j++ {
			fmt.Printf("%2d ", i*j)
		}
		fmt.Println()
	}
}
