package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

// GoSquirrel — суслик-гофер, который копает ГЛУБЖЕ, а не шире
type GoSquirrel struct {
	Name           string
	DigDepth       int    // Глубина "копания" в Go (метры)
	CurrentProject string // Текущий проект
	Distractions   []string // Отвлечения, которые он игнорирует
}

// VALUE RECEIVER: Показывает статус (не меняет состояние)
func (s GoSquirrel) StatusReport() string {
	return fmt.Sprintf("🐿️ %s копает на глубине %dм в проекте '%s'. Отвлечений проигнорировано: %d",
		s.Name, s.DigDepth, s.CurrentProject, len(s.Distractions))
}

// POINTER RECEIVER: Глубже копает в Go (меняет состояние!)
func (s *GoSquirrel) DigDeeper(meters int) {
	s.DigDepth += meters
	log.Printf("⛏️ %s углубился на %dм! Теперь глубина: %dм", s.Name, meters, s.DigDepth)

	// Новый уровень — новое понимание Go
	if s.DigDepth >= 100 {
		s.CurrentProject = "Ядро компилятора Go"
		log.Println("✨ ДОСТИГНУТ УРОВЕНЬ МАСТЕРА: Работа над ядром компилятора!")
	}
}

// POINTER RECEIVER: Игнорирует отвлечения (другие языки!)
func (s *GoSquirrel) IgnoreDistraction(lang string) {
	s.Distractions = append(s.Distractions, lang)
	log.Printf("🚫 %s проигнорировал %s! Фокус только на Go.", s.Name, lang)

	// Правило 2026: никакого распыления!
	if len(s.Distractions) >= 5 {
		log.Println("🔥 Челлендж 'Go365' активирован: В 2026 году НИ ОДНОГО отвлечения!")
	}
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	// Создаём Гошу-суслика на пороге Нового Года
	gosha := GoSquirrel{
		Name:           "Гоша",
		DigDepth:       57, // 57 дней марафона!
		CurrentProject: "курьерский трекер на Go",
	}

	log.Println("🎉 31 декабря 2025. Последние минуты уходящего года. Заснеженная Москва.")
	log.Println("🎯 Тема дня: Pointer Receivers — как глубоко копать в одном направлении (Go), а не распыляться!")
	log.Println("📜 Клятва Гоши: '2026 — Год Go. Никаких PHP/Java/Python. Только Go и его эко-система.'")

	// Демонстрируем VALUE RECEIVER (статус без изменений)
	log.Println(gosha.StatusReport())

	// POINTER RECEIVER: Копаем глубже к мечте
	time.Sleep(1 * time.Second)
	gosha.DigDeeper(43) // До 100м — уровня мастерства!

	// POINTER RECEIVER: Игнорируем отвлечения (другие языки)
	time.Sleep(1 * time.Second)
	distractions := []string{"Python", "Java", "PHP", "Rust", "C++"}
	for _, lang := range distractions {
		gosha.IgnoreDistraction(lang)
		time.Sleep(500 * time.Millisecond)
	}

	// Финал: Проверяем статус после изменений (благодаря указателям!)
	log.Println("\n" + gosha.StatusReport())

	log.Println("\n⏰ 00:00 1 января 2026...")
	log.Println("🚀 Новый Год начинается с вызова:")
	log.Println("   go run --focus=only main.go")
	log.Println("🔥 Гоша-суслик надеется: 'Скоро я найду работу программистом на Go!'")
}
