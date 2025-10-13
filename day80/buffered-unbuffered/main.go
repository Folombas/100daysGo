package main

import (
	"fmt"
	"time"
)

// Mushroom представляет гриб
type Mushroom struct {
	Type string
	Size int
}

// MushroomPicker представляет грибника
type MushroomPicker struct {
	Name string
}

// NewMushroomPicker создает нового грибника
func NewMushroomPicker(name string) *MushroomPicker {
	return &MushroomPicker{Name: name}
}

// DemonstrateUnbufferedChannels демонстрирует небуферизованные каналы
func DemonstrateUnbufferedChannels() {
	fmt.Println("🎯 НЕБУФЕРИЗОВАННЫЕ КАНАЛЫ (ручная передача):")
	fmt.Println("============================================")

	unbufferedChan := make(chan Mushroom)
	picker := NewMushroomPicker("Василий")

	// Грибник собирает и передает грибы
	go func() {
		mushrooms := []Mushroom{
			{"белый", 15}, {"подосиновик", 12}, {"лисичка", 8},
		}

		for _, mushroom := range mushrooms {
			fmt.Printf("   🍄 %s нашел %s (%dсм)\n", picker.Name, mushroom.Type, mushroom.Size)
			fmt.Printf("   🤲 Передает %s... (ожидает получателя)\n", mushroom.Type)
			unbufferedChan <- mushroom
			fmt.Printf("   ✅ %s передан!\n", mushroom.Type)
			time.Sleep(500 * time.Millisecond)
		}
		close(unbufferedChan)
	}()

	// Получатель принимает грибы
	go func() {
		for mushroom := range unbufferedChan {
			fmt.Printf("   📦 Получен %s (%dсм)\n", mushroom.Type, mushroom.Size)
			time.Sleep(800 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("   🎉 Все грибы переданы из рук в руки!\n")
}

// DemonstrateBufferedChannels демонстрирует буферизованные каналы
func DemonstrateBufferedChannels() {
	fmt.Println("🚚 БУФЕРИЗОВАННЫЕ КАНАЛЫ (грузовик с кузовом):")
	fmt.Println("============================================")

	bufferedChan := make(chan Mushroom, 3)
	picker := NewMushroomPicker("Петр")

	// Грибник собирает грибы и складывает в "кузов"
	go func() {
		mushrooms := []Mushroom{
			{"масленок", 10}, {"рыжик", 9}, {"волнушка", 7}, {"груздь", 14}, {"сыроежка", 6},
		}

		for idx, mushroom := range mushrooms {
			fmt.Printf("   🍄 %s нашел %s (%dсм)\n", picker.Name, mushroom.Type, mushroom.Size)

			select {
			case bufferedChan <- mushroom:
				fmt.Printf("   🚚 Положил %s в кузов (место занято: %d/%d)\n",
					mushroom.Type, len(bufferedChan), cap(bufferedChan))
			default:
				fmt.Printf("   ⚠️  Кузов полен! Не могу положить %s\n", mushroom.Type)
			}

			time.Sleep(300 * time.Millisecond)

			// Ранний выход если это последняя итерация
			if idx == len(mushrooms)-1 {
				close(bufferedChan)
			}
		}
	}()

	// Разгрузка "кузова"
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("   🏁 Начинаем разгрузку кузова...")

		for mushroom := range bufferedChan {
			fmt.Printf("   📦 Разгружаем %s (%dсм) (осталось: %d)\n",
				mushroom.Type, mushroom.Size, len(bufferedChan))
			time.Sleep(1 * time.Second)
		}
	}()

	time.Sleep(6 * time.Second)
	fmt.Println("   🎉 Грузовик разгружен!\n")
}

// ComparePerformance сравнивает производительность
func ComparePerformance() {
	fmt.Println("📊 СРАВНЕНИЕ ПРОИЗВОДИТЕЛЬНОСТИ:")
	fmt.Println("===============================")

	// Тест небуферизованного канала
	start := time.Now()
	unbuffered := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			unbuffered <- i
		}
		close(unbuffered)
	}()

	go func() {
		for range unbuffered {
			// Читаем значения
		}
	}()

	time.Sleep(100 * time.Millisecond)
	unbufferedTime := time.Since(start)

	// Тест буферизованного канала
	start = time.Now()
	buffered := make(chan int, 5)

	go func() {
		for i := 0; i < 5; i++ {
			buffered <- i
		}
		close(buffered)
	}()

	go func() {
		for range buffered {
			// Читаем значения
		}
	}()

	time.Sleep(100 * time.Millisecond)
	bufferedTime := time.Since(start)

	fmt.Printf("   ⏱️  Небуферизованный: %v\n", unbufferedTime)
	fmt.Printf("   ⏱️  Буферизованный: %v\n", bufferedTime)
	fmt.Printf("   📈 Разница: %v\n\n", unbufferedTime-bufferedTime)
}

// RealWorldExample показывает реальный пример использования
func RealWorldExample() {
	fmt.Println("🌍 РЕАЛЬНЫЙ ПРИМЕР (обработка грибов):")
	fmt.Println("====================================")

	collected := make(chan Mushroom, 10)
	cleaned := make(chan Mushroom, 8)
	packaged := make(chan Mushroom, 5)

	// Стадия 1: Сбор грибов
	go func() {
		types := []string{"белый", "подберезовик", "лисичка", "опенок"}
		for j := 0; j < 12; j++ {
			mushroom := Mushroom{
				Type: types[j%len(types)],
				Size: 8 + j%7,
			}
			collected <- mushroom
			fmt.Printf("   🍄 Собран: %s\n", mushroom.Type)
			time.Sleep(100 * time.Millisecond)
		}
		close(collected)
	}()

	// Стадия 2: Очистка грибов
	go func() {
		for mushroom := range collected {
			fmt.Printf("   🧹 Очищаем: %s\n", mushroom.Type)
			time.Sleep(200 * time.Millisecond)
			cleaned <- mushroom
		}
		close(cleaned)
	}()

	// Стадия 3: Упаковка грибов
	go func() {
		for mushroom := range cleaned {
			fmt.Printf("   📦 Упаковываем: %s\n", mushroom.Type)
			time.Sleep(150 * time.Millisecond)
			packaged <- mushroom
		}
		close(packaged)
	}()

	// Финальная стадия: Подсчет результатов
	count := 0
	for range packaged {
		count++
	}

	fmt.Printf("   ✅ Обработано грибов: %d\n\n", count)
}

// DeadlockExample показывает пример deadlock
func DeadlockExample() {
	fmt.Println("💀 ПРИМЕР DEADLOCK (чего избегать):")
	fmt.Println("=================================")

	fmt.Println("   💡 Мораль: всегда обеспечивайте и отправителя, и получателя!")
}

func main() {
	fmt.Println("🍄 Day 80: Buffered vs Unbuffered Channels - Грибная логистика!")
	fmt.Println("============================================================")

	DemonstrateUnbufferedChannels()
	DemonstrateBufferedChannels()
	ComparePerformance()
	RealWorldExample()
	DeadlockExample()

	fmt.Println("🎯 КЛЮЧЕВЫЕ ВЫВОДЫ:")
	fmt.Println("   • Небуферизованные: синхронные, передача из рук в руки")
	fmt.Println("   • Буферизованные: асинхронные, временное хранение данных")
	fmt.Println("   • Буфер уменьшает блокировки, но использует больше памяти")
	fmt.Println("   • Выбор зависит от требований к синхронизации")

	fmt.Println("\n💪 Отлично! Теперь ты разбираешься в каналах как опытный грибник!")
}
