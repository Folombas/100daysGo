package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

// Friend - интерфейс, представляющий "друга"
type Friend interface {
	BeFriend() string
	GetName() string
}

// RealFriend - настоящий друг
type RealFriend struct {
	Name   string
	Mood   string
	Invite bool
}

func (rf RealFriend) BeFriend() string {
	if rf.Invite {
		return fmt.Sprintf("%s: 'Привет, Гоша! Едем во Владимир, собирайся!'", rf.Name)
	}
	return fmt.Sprintf("%s: 'Как дела с Go? Помогу, если что!'", rf.Name)
}

func (rf RealFriend) GetName() string {
	return rf.Name
}

// Troll - тролль-энергетический вампир
type Troll struct {
	Name       string
	Location   string
	PhotoCount int
}

func (t Troll) BeFriend() string {
	return fmt.Sprintf("%s из %s: Смотри мои %d фото! (но не зовет тебя)",
		t.Name, t.Location, t.PhotoCount)
}

func (t Troll) GetName() string {
	return t.Name
}

// Depression - состояние депрессии Гоши
type Depression struct {
	Level     int
	StartTime time.Time
}

func (d Depression) BeFriend() string {
	return fmt.Sprintf("Депрессия уровня %d/10: 'Остался только Go...'", d.Level)
}

func (d Depression) GetName() string {
	return "Зимняя депрессия"
}

// Проверка друга с помощью type assertion
func checkFriendType(f Friend) {
	fmt.Printf("🔍 Проверяю: %s\n", f.GetName())

	// Type assertion с проверкой
	if troll, ok := f.(Troll); ok {
		fmt.Printf("   🚫 ТРОЛЛЬ обнаружен! Локация: %s\n", troll.Location)
		fmt.Println("   💡 Применяю правило: 'Не корми тролля!'")
		fmt.Println("   ✅ Решение: Удаляю фото, блокирую, иду писать код на Go")
		return
	}

	if friend, ok := f.(RealFriend); ok {
		if friend.Invite {
			fmt.Printf("   ✅ НАСТОЯЩИЙ ДРУГ! %s\n", friend.BeFriend())
			fmt.Println("   🎉 Гоша счастлив! Но... 'Спасибо, ребята, я на Go марафоне!'")
		} else {
			fmt.Printf("   🤝 Друг-программист: %s\n", friend.BeFriend())
		}
		return
	}

	if _, ok := f.(Depression); ok {
		fmt.Println("   😔 Обнаружена депрессия...")
		fmt.Println("   💊 Лечение: 100 строк кода на Go три раза в день")
		return
	}

	// Если тип неизвестен
	fmt.Printf("   ❓ Неизвестный тип: %T\n", f)
	fmt.Println("   ⚠️  Нужно добавить обработку этого типа!")
}

// Type switch - альтернативный способ
func checkFriendWithSwitch(f Friend) {
	fmt.Printf("\n🎲 Type switch проверка для: %s\n", f.GetName())

	switch v := f.(type) {
	case Troll:
		fmt.Printf("   🧌 %s - тролль уровня %d фото\n", v.Name, v.PhotoCount)
		fmt.Println("   🗑️  Фото удалены, энергия сохранена")
	case RealFriend:
		fmt.Printf("   👨‍💻 %s - коллега по Go\n", v.Name)
		if v.Mood != "" {
			fmt.Printf("   📊 Настроение: %s\n", v.Mood)
		}
	case Depression:
		fmt.Printf("   🌧️  %s с %v\n", v.GetName(), v.StartTime.Format("15:04"))
		fmt.Printf("   📈 Уровень: %d/10\n", v.Level)
	default:
		fmt.Printf("   🔮 Неожиданный тип: %T\n", v)
	}
}

func main() {
	fmt.Println("================================")
	fmt.Println("   FRIEND TYPE ASSERTION SIMULATOR")
	fmt.Println("   День 64: Interfaces - Type Assertions")
	fmt.Println("   Легенда: Январская депрессия и тролли")
	fmt.Println("================================")

	rand.Seed(time.Now().UnixNano())

	// Создаем "друзей" разных типов
	friends := []Friend{
		Troll{Name: "Вася", Location: "Суздаль", PhotoCount: 42},
		Troll{Name: "Петя", Location: "Владимир", PhotoCount: 23},
		RealFriend{Name: "Alex Gopher", Mood: "сфокусированный", Invite: false},
		RealFriend{Name: "Go Mentor", Mood: "", Invite: true},
		Depression{Level: 7, StartTime: time.Date(2026, 1, 6, 13, 30, 0, 0, time.Local)},
		RealFriend{Name: "100DaysGo", Mood: "упорный", Invite: false},
	}

	fmt.Println("\n📱 Лента друзей (интерфейс Friend):")
	for _, f := range friends {
		fmt.Printf("   %s\n", f.BeFriend())
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("🚀 Запускаем Type Assertion проверку...")

	// Проверяем каждого с помощью type assertion
	for i, f := range friends {
		fmt.Printf("\n%d. ", i+1)
		checkFriendType(f)
		checkFriendWithSwitch(f)
	}

	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("📊 ИТОГИ ДНЯ 64:")
	fmt.Println("   Type Assertions - это как проверка:")
	fmt.Println("   'Кто ты на самом деле?'")
	fmt.Println("")
	fmt.Println("   🎯 Научился:")
	fmt.Println("   • Делать утверждения типа: value, ok := interface.(ConcreteType)")
	fmt.Println("   • Использовать type switch для категоризации")
	fmt.Println("   • Обрабатывать неизвестные типы")
	fmt.Println("")
	fmt.Println("   💡 Применение в Легенде:")
	fmt.Println("   Тролль != Друг")
	fmt.Println("   Депрессия != Вечность")
	fmt.Println("   Go == Спасение")
	fmt.Println("")
	fmt.Println("   🚫 Правило: 'Не корми тролля'")
	fmt.Println("   ✅ Альтернатива: 'Корми Go кодом'")

	// Демонстрация паники при неправильном assertion
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Println("⚠️  ОПАСНЫЙ ПРИМЕР (panic):")

	// Этот код вызовет панику, если assertion неверный
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("   💥 Поймана паника:", r)
				fmt.Println("   ✅ Вывод: Всегда используйте 'ok' проверку!")
			}
		}()

		fmt.Println("   Пытаюсь сделать assertion без проверки...")
		// Это вызовет панику, так как первый friend - Troll
		// realFriend := friends[0].(RealFriend)
		// fmt.Printf("   %s\n", realFriend.BeFriend())
		fmt.Println("   (код закомментирован для безопасности)")
	}()
}
