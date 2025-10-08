package main

import (
	"errors"
	"fmt"
)

// FarmError представляет кастомную ошибку для фермерских условий
type FarmError struct {
	Operation string
	Err       error
	Condition string
}

func (e *FarmError) Error() string {
	return fmt.Sprintf("🚜 Сбой операции: %s (условия: %s) - %v", 
		e.Operation, e.Condition, e.Err)
}

func (e *FarmError) Unwrap() error {
	return e.Err
}

// Ошибки в полевых условиях
var (
	ErrNoInternet = errors.New("без стабильного широкополосного городского быстрого Интернета")
	ErrLowBattery = errors.New("низкий заряд батареи телефона")
	ErrBrightSun  = errors.New("солнце слепит экран телефона")
	ErrDistracted = errors.New("отвлекли животные с фермы")
)

// simulateCodingAttempt симулирует попытку программирования в полевых условиях
func simulateCodingAttempt(day int) error {
	if err := findShelter(); err != nil {
		return &FarmError{
			Operation: "поиск укрытия для кодинга",
			Err:       err,
			Condition: "прямые солнечные лучи",
		}
	}

	if err := establishConnection(); err != nil {
		return fmt.Errorf("сбой установки соединения: %w", err)
	}

	if err := writeGoCode(day); err != nil {
		return fmt.Errorf("не удалось написать код на Go для дня %d: %w", day, err)
	}

	fmt.Printf("✅ День %d завершён успешно прямо с сеновала!\n", day)
	return nil
}

func findShelter() error {
	return ErrBrightSun // Всегда солнце слепит!
}

func establishConnection() error {
	return ErrNoInternet // Интернет на ферме - это роскошь
}

func writeGoCode(day int) error {
	if day > 70 {
		return ErrDistracted // После 70 дней сложно концентрироваться
	}
	return nil
}

// analyzeError демонстрирует развёртку ошибок
func analyzeError(err error) {
	fmt.Println("\n🔍 Анализ цепочки ошибок:")
	
	for err != nil {
		fmt.Printf("   → %v\n", err)
		err = errors.Unwrap(err)
	}
}

func main() {
	fmt.Println("🎯 День 75: Обёртка и развёртка ошибок")
	fmt.Println("=======================================")
	fmt.Println("📍 Местоположение: Фермерский сеновал")
	fmt.Println("📱 Устройство: Honor 10x Lite + Termux")
	fmt.Println("💪 Дух: Несломленный!")
	fmt.Println("")

	fmt.Println("🚀 Пытаюсь писать код в полевых условиях...")
	err := simulateCodingAttempt(75)
	
	if err != nil {
		fmt.Printf("\n❌ Сессия кодинга провалилась: %v\n", err)
		analyzeError(err)
		
		// Демонстрация errors.Is
		fmt.Println("\n🔎 Используем errors.Is для проверки типов ошибок:")
		if errors.Is(err, ErrDistracted) {
			fmt.Println("   🐐 Подтверждено: Отвлекли животные с фермы!")
		}
		
		// Демонстрация errors.As
		fmt.Println("\n🎯 Используем errors.As для извлечения FarmError:")
		var farmErr *FarmError
		if errors.As(err, &farmErr) {
			fmt.Printf("   Детали фермерской ошибки: Операция=%s, Условия=%s\n", 
				farmErr.Operation, farmErr.Condition)
		}
	}

	fmt.Println("\n💪 Несмотря на все препятствия, челлендж #100DaysOfGo продолжается!")
	fmt.Println("🚜 Код с фермы на GitHub: Свежий с сеновала!")
}
