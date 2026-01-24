package main

import (
	"fmt"
)

// TemptationError - базовая ошибка искушения
type TemptationError struct {
	Temptation string
	Strength   int // Сила искушения от 1 до 10
}

func (e *TemptationError) Error() string {
	return fmt.Sprintf("Искушение '%s' (сила: %d/10)", e.Temptation, e.Strength)
}

// DistractionError - ошибка отвлечения
type DistractionError struct {
	Distraction string
	InnerErr    error
}

func (e *DistractionError) Error() string {
	return fmt.Sprintf("Отвлечение на: %s -> %v", e.Distraction, e.InnerErr)
}

func (e *DistractionError) Unwrap() error {
	return e.InnerErr
}

// WillpowerError - ошибка силы воли (обёртка)
type WillpowerError struct {
	Action    string
	InnerErr  error
	Resisted  bool
}

func (e *WillpowerError) Error() string {
	resistStatus := "поддался"
	if e.Resisted {
		resistStatus = "преодолел"
	}
	return fmt.Sprintf("Сила воли: %s %s -> %v", resistStatus, e.Action, e.InnerErr)
}

func (e *WillpowerError) Unwrap() error {
	return e.InnerErr
}

// Создание ошибок-искушений
var (
	ErrCapCutTemptation = &TemptationError{
		Temptation: "Установить CapCut и монтировать видео",
		Strength:   8,
	}

	ErrGameTemptation = &TemptationError{
		Temptation: "Поиграть в видеоигры",
		Strength:   6,
	}

	ErrBarTemptation = &TemptationError{
		Temptation: "Сходить в бар/клуб",
		Strength:   7,
	}

	ErrMovieTemptation = &TemptationError{
		Temptation: "Посмотреть фильмы/сериалы",
		Strength:   5,
	}
)
