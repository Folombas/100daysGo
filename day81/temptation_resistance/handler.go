package main

import (
	"errors"
	"fmt"
)

// TaskManager управляет задачами программирования
type TaskManager struct {
	FocusLevel   int
	Dopamine     int
	Temptations  int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		FocusLevel:  85,
		Dopamine:    200,
		Temptations: 0,
	}
}

func (tm *TaskManager) CompleteProgrammingTask(task string) error {
	fmt.Printf("\n🎯 Начинаем задачу: %s\n", task)

	// Симуляция возникновения искушения
	if tm.Temptations > 2 {
		// Создаём цепочку ошибок
		baseErr := ErrCapCutTemptation
		distractionErr := &DistractionError{
			Distraction: "видеомонтаж отпускных видео из тропических стран 🏝️",
			InnerErr:    baseErr,
		}
		willpowerErr := &WillpowerError{
			Action:    "изучение Go",
			InnerErr:  distractionErr,
			Resisted:  false,
		}

		return fmt.Errorf("критическая ошибка фокуса: %w", willpowerErr)
	}

	tm.Dopamine += 100
	tm.FocusLevel += 5
	return nil
}

// IsTemptation проверяет, является ли ошибка искушением
func IsTemptation(err error) bool {
	var temptationErr *TemptationError
	return errors.As(err, &temptationErr)
}

// ResistTemptation обрабатывает искушение
func ResistTemptation(err error) {
	fmt.Println("\n🛡️  АКТИВАЦИЯ ЗАЩИТЫ:")
	fmt.Println("1. Вспоминаем цель: устроиться программистом")
	fmt.Println("2. Вспоминаем мотивационные фразы")
	fmt.Println("3. Добавляем +50 к силе воли")
	fmt.Println("4. Продолжаем изучение Go!")
}

// PrintErrorChain печатает цепочку ошибок
func PrintErrorChain(err error) {
	for err != nil {
		fmt.Printf("  → %v\n", err)
		err = errors.Unwrap(err)
	}
}
