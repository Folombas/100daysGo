package train

import (
	"fmt"
	"math/rand"
	"time"
)

// Journey представляет поездку в электричке
type Journey struct {
	From          string
	To            string
	DepartureHour int
	DepartureMinute int
	Duration      float64 // в часах
}

// NewJourney создает новую поездку
func NewJourney(from, to string, hour, minute int) *Journey {
	// Случайная длительность от 1.5 до 3 часов
	rand.Seed(time.Now().UnixNano())
	duration := 1.5 + rand.Float64()*1.5
	
	return &Journey{
		From:          from,
		To:            to,
		DepartureHour: hour,
		DepartureMinute: minute,
		Duration:      duration,
	}
}

// StartJourney начинает поездку
func (j *Journey) StartJourney() string {
	return fmt.Sprintf("Отправление: %s → %s в %02d:%02d", 
		j.From, j.To, j.DepartureHour, j.DepartureMinute)
}

// EndJourney завершает поездку
func (j *Journey) EndJourney() string {
	arrivalHour := j.DepartureHour + int(j.Duration)
	arrivalMinute := j.DepartureMinute + int((j.Duration-float64(int(j.Duration)))*60)
	
	if arrivalMinute >= 60 {
		arrivalHour++
		arrivalMinute -= 60
	}
	
	return fmt.Sprintf("Прибытие в %s в %02d:%02d", j.To, arrivalHour, arrivalMinute)
}

// Route возвращает маршрут
func (j *Journey) Route() string {
	return fmt.Sprintf("%s → %s (%.1f часа)", j.From, j.To, j.Duration)
}

// CalculateJourneyScore рассчитывает очки за поездку
func (j *Journey) CalculateJourneyScore() int {
	baseScore := 40
	durationBonus := int(j.Duration * 10)
	
	// Бонус за дальнюю поездку
	distanceBonus := 25
	
	return baseScore + durationBonus + distanceBonus
}
