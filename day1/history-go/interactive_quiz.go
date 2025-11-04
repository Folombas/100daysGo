package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StartInteractiveQuiz() {
	fmt.Println("\n🎮 ИНТЕРАКТИВНЫЙ КВИЗ ПО ИСТОРИИ GO:")
	fmt.Println("=====================================")

	reader := bufio.NewReader(os.Stdin)
	score := 0

	questions := []struct {
		question string
		answer   string
		hint     string
	}{
		{"В каком году началась разработка Go?", "2007", "До выхода iPhone"},
		{"Кто из этих разработчиков НЕ был создателем Go? (1-Деннис Ритчи, 2-Роб Пайк, 3-Кен Томпсон)", "1", "Он создал C язык"},
		{"Какая версия Go принесла Generics? (1-1.18, 2-1.20, 3-1.22)", "1", "2022 год"},
		{"Как называлась система зависимостей до Go Modules?", "GOPATH", "GOPATH/src"},
	}

	for i, q := range questions {
		fmt.Printf("\n❓ Вопрос %d: %s\n", i+1, q.question)
		fmt.Printf("💡 Подсказка: %s\n", q.hint)
		fmt.Print("📝 Ваш ответ: ")

		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, q.answer) {
			fmt.Println("✅ Правильно!")
			score++
		} else {
			fmt.Printf("❌ Неверно. Правильный ответ: %s\n", q.answer)
		}
	}

	fmt.Printf("\n🎯 Ваш результат: %d/%d\n", score, len(questions))

	switch {
	case score == len(questions):
		fmt.Println("🏆 Отлично! Ты настоящий эксперт по истории Go!")
	case score >= len(questions)/2:
		fmt.Println("👍 Хорошо! Продолжай изучать историю!")
	default:
		fmt.Println("💪 Не сдавайся! История запоминается со временем!")
	}
}
