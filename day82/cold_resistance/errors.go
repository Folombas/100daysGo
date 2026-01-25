package main

import (
	"errors"
	"fmt"
)


// Sentinel Errors - глобальные ошибки-стражи
var (
	// Ошибки окружающей среды
	ErrColdRoom    = errors.New("sentinel: температура комнаты ниже 0°C")
	ErrHotRoom     = errors.New("sentinel: температура комнаты выше 25°C")
	ErrNoisyRoom   = errors.New("sentinel: слишком шумно для концентрации")

	// Ошибки-искушения
	ErrTVTemptation     = errors.New("sentinel: телевизор отвлекает")
	ErrSocialMedia      = errors.New("sentinel: соцсети требуют внимания")
	ErrVideoGames       = errors.New("sentinel: видеоигры зовут")
	ErrCapCutTemptation = errors.New("sentinel: хочется монтировать видео")
	ErrWarmBed          = errors.New("sentinel: кровать манит обратно")
	ErrComfortZone      = errors.New("sentinel: слишком комфортно для роста")

	// Ошибки физического состояния
	ErrHunger        = errors.New("sentinel: чувство голода")
	ErrTiredness     = errors.New("sentinel: усталость накатывает")
	ErrThirst        = errors.New("sentinel: жажда отвлекает")

	// Ошибки ментального состояния
	ErrProcrastination = errors.New("sentinel: прокрастинация атакует")
	ErrDoubt          = errors.New("sentinel: сомнения в своих силах")
	ErrLoneliness     = errors.New("sentinel: чувство одиночества")

	// Ошибки обучения
	ErrComplexTopic   = errors.New("sentinel: тема слишком сложная")
	ErrBoringSyntax   = errors.New("sentinel: синтаксис кажется скучным")
	ErrDebuggingHell  = errors.New("sentinel: отладка затянулась")
)

// SentinelError - кастомный тип для sentinel ошибок
type SentinelError struct {
	Err      error
	Severity int // 1-10, где 10 - максимальная опасность
	Category string
}

func (se *SentinelError) Error() string {
	return fmt.Sprintf("[%s:%d] %v", se.Category, se.Severity, se.Err)
}

func NewSentinelError(err error, severity int, category string) *SentinelError {
	return &SentinelError{
		Err:      err,
		Severity: severity,
		Category: category,
	}
}

// Проверка является ли ошибка sentinel ошибкой
func IsSentinelError(err error) bool {
	_, ok := err.(*SentinelError)
	return ok
}

// Проверка конкретной sentinel ошибки
func IsSpecificSentinel(err error, target error) bool {
	return errors.Is(err, target)
}
