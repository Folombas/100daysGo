package main

import (
	"fmt"
	"strings"
)

// ==================== ТИПЫ ДАННЫХ ====================

// TrollMessage - сообщение от тролля
type TrollMessage string

func (tm TrollMessage) TrollInfo() string {
	return fmt.Sprintf("Тролль: %q", string(tm))
}

func (tm TrollMessage) Block() DopamineReward {
	return DopamineReward{
		Dopamine: 50,
		Action:   "Заблокирован на срок N-дней",
		Regret:   false,
	}
}

// Message - обычное сообщение
type Message string

func (m Message) Content() string {
	return string(m)
}

func (m Message) Process() DopamineReward {
	return DopamineReward{
		Dopamine: 10,
		Action:   "Прочитано и сохранено",
		Regret:   false,
	}
}

// DopamineReward - награда в дофамине
type DopamineReward struct {
	Dopamine int
	Action   string
	Regret   bool
}

// AlpacaFarm - ферма альпак (нереализованное событие)
type AlpacaFarm struct {
	Location string
	Cuteness int
}

func (af AlpacaFarm) TrollInfo() string {
	return fmt.Sprintf("Ферма альпак в %s (милота: %d/10)", af.Location, af.Cuteness)
}

func (af AlpacaFarm) Block() DopamineReward {
	return DopamineReward{
		Dopamine: 0,
		Action:   "Нельзя заблокировать ферму, только тролля",
		Regret:   true, // Сожаление, что не поехал
	}
}

// ==================== TYPE CONSTRAINTS ====================

// TrollConstraint - ограничение для троллей
type TrollConstraint interface {
	~string // Разрешены string и типы с underlying type string (как TrollMessage)
	TrollInfo() string
	Block() DopamineReward
}

// MessageConstraint - ограничение для сообщений
type MessageConstraint interface {
	~string // Разрешены string и типы с underlying type string (как Message)
	Content() string
	Process() DopamineReward
}

// ==================== GENERIC ФУНКЦИИ ====================

// GenericProcessWithConstraint[T TrollConstraint] - обработка с ограничением для троллей
func GenericProcessWithConstraint[T TrollConstraint](troll T) DopamineReward {
	fmt.Printf("🔧 Type Constraint: Обрабатываем %T\n", troll)
	fmt.Printf("   Информация: %s\n", troll.TrollInfo())
	return troll.Block()
}

// GenericMessageProcess[T MessageConstraint] - обработка сообщений
func GenericMessageProcess[T MessageConstraint](msg T) DopamineReward {
	fmt.Printf("📨 Обработка сообщения типа: %T\n", msg)
	fmt.Printf("   Содержимое: %s\n", msg.Content())
	return msg.Process()
}

// GenericRegretAnalyzer[T any] - анализ сожаления (любой тип)
func GenericRegretAnalyzer[T any](item T, description string) (bool, int) {
	// Проверяем, содержит ли описание ключевые слова сожаления
	regretKeywords := []string{"альпак", "ферм", "мил", "пушист", "скучно"}
	hasRegret := false

	for _, keyword := range regretKeywords {
		if strings.Contains(strings.ToLower(description), keyword) {
			hasRegret = true
			break
		}
	}

	if hasRegret {
		return true, -20 // Минус дофамин за сожаление
	}

	return false, 0
}

// ==================== ГЕЙМИФИКАЦИЯ ====================

type DayStats struct {
	WakeUpTime    string
	OutdoorTime   float64
	TrollsBlocked int
	DopamineTotal int
	TopicsLearned []string
	RegretLevel   int
}

func NewDayStats() *DayStats {
	return &DayStats{
		WakeUpTime:    "10:00",
		OutdoorTime:   1.5,
		TrollsBlocked: 0,
		DopamineTotal: 100, // Базовый дофамин за пробуждение
		TopicsLearned: []string{},
		RegretLevel:   0,
	}
}

func (ds *DayStats) AddDopamine(amount int, reason string) {
	ds.DopamineTotal += amount
	fmt.Printf("   💊 Дофамин %+d: %s\n", amount, reason)
}

func (ds *DayStats) BlockTroll(trollName string) {
	ds.TrollsBlocked++
	ds.AddDopamine(50, fmt.Sprintf("Заблокирован тролль %s", trollName))
	ds.AddDopamine(30, "Освобождение от токсичности")
}

func (ds *DayStats) LearnTopic(topic string) {
	ds.TopicsLearned = append(ds.TopicsLearned, topic)
	ds.AddDopamine(100, fmt.Sprintf("Изучение темы: %s", topic))
	ds.AddDopamine(50, "Type Constraints освоены")
}

func (ds *DayStats) AddRegret(amount int, reason string) {
	ds.RegretLevel += amount
	ds.AddDopamine(amount, reason) // Отрицательный дофамин
}

func (ds *DayStats) PrintStats() {
	fmt.Println("\n📊 ИТОГИ ДНЯ 68:")
	fmt.Println("════════════════════════════════════")
	fmt.Printf("   🕗 Подъём: %s (провалялся +3 часа)\n", ds.WakeUpTime)
	fmt.Printf("   🚶 Прогулка: %.1f часа свежего воздуха\n", ds.OutdoorTime)
	fmt.Printf("   🚫 Заблокировано троллей: %d\n", ds.TrollsBlocked)
	fmt.Printf("   📚 Изучено тем: %d\n", len(ds.TopicsLearned))
	fmt.Printf("   😔 Уровень сожаления: %d%%\n", ds.RegretLevel)
	fmt.Printf("   🎯 Общий дофамин: %d\n", ds.DopamineTotal)

	if ds.RegretLevel > 50 {
		fmt.Println("   ⚠️  Слишком много сожаления! Фокус на коде!")
	} else {
		fmt.Println("   ✅ Сожаления под контролем!")
	}
}

// ==================== ОСНОВНАЯ ПРОГРАММА ====================

func main() {
	fmt.Println("🦙 DAY 68: Generics in Go - Type Constraints")
	fmt.Println("════════════════════════════════════")
	fmt.Println("Сюжет: Блокировка тролля Рокки и изучение ограничений типов")

	// Статистика дня
	stats := NewDayStats()

	// 1. Утренний ритуал
	fmt.Println("🌅 УТРЕННИЙ РИТУАЛ:")
	fmt.Println("   • Проснулся в 7:00, посмотрел на сугробы")
	fmt.Println("   • Провалялся до 10:00 (нейтральный дофамин)")
	fmt.Println("   • Прогулка на свежем воздухе 1.5 часа")
	stats.AddDopamine(30, "Свежий воздух и размышления")

	// 2. Сообщение от тролля Рокки
	fmt.Println("\n📱 СООБЩЕНИЕ ОТ ТРОЛЛЯ РОККИ:")
	trollMsg := TrollMessage("Гошик, а зря ты с нами не поехал на джипе гулять сегодня на ферму альпак в Подмосковье, они такие милые и пушистые!")

	// Используем generic функцию с constraint для троллей
	reward := GenericProcessWithConstraint(trollMsg)
	fmt.Printf("   🎯 Результат: %s\n", reward.Action)
	fmt.Printf("   💊 Дофамин: %d\n", reward.Dopamine)

	// Блокируем тролля
	stats.BlockTroll("Рокки")

	// 3. Анализ сожаления с использованием generic функции
	fmt.Println("\n😔 АНАЛИЗ СОЖАЛЕНИЯ:")
	regret, dopamineChange := GenericRegretAnalyzer(trollMsg, "альпаки такие милые и пушистые")
	if regret {
		fmt.Println("   ⚠️  Обнаружено скрытое сожаление!")
		stats.AddRegret(dopamineChange, "Сожаление об альпаках")
	} else {
		fmt.Println("   ✅ Сожаления не обнаружено!")
	}

	// 4. Изучение Type Constraints
	fmt.Println("\n📚 ИЗУЧЕНИЕ TYPE CONSTRAINTS:")
	stats.LearnTopic("Generics Type Constraints in Go")

	// Демонстрация работы с разными типами через constraints
	fmt.Println("\n🔧 ДЕМОНСТРАЦИЯ CONSTRAINTS:")

	// Работа с обычным сообщением
	normalMsg := Message("Привет, как дела?")
	msgReward := GenericMessageProcess(normalMsg)
	fmt.Printf("   📨 Результат: %s (дофамин: %d)\n", msgReward.Action, msgReward.Dopamine)
	stats.AddDopamine(msgReward.Dopamine, "Обработка нормального сообщения")

	// Попытка обработать ферму альпак (вызовет ошибку компиляции, если раскомментировать)
	// farm := AlpacaFarm{Location: "Подмосковье", Cuteness: 9}
	// farmReward := GenericProcessWithConstraint(farm) // Ошибка: AlpacaFarm не реализует TrollConstraint
	// fmt.Println("   Эта строка не скомпилируется - ферма не тролль!")

	// 5. Преимущества блокировки
	fmt.Println("\n🎯 ПРЕИМУЩЕСТВА БЛОКИРОВКИ ТРОЛЛЯ:")
	advantages := []string{
		"Больше нет токсичных сообщений",
		"Экономия времени на чтение ерунды",
		"Фокус на изучении Go",
		"Меньше негатива в ленте",
		"Меньше поводов для сожаления",
		"Улучшение ментального здоровья",
		"Повышение продуктивности кодинга",
		"Освобождение от постоянной проверки уведомлений",
		"Больше времени для реальных проектов",
		"Снижение уровня стресса и тревоги",
		"Возможность сосредоточиться на глубокой работе",
		"Защита от манипуляций и газлайтинга",
		"Сохранение энергии для творчества",
		"Укрепление личных границ в цифровом пространстве",
		"Развитие навыка игнорирования нерелевантного",
	}

	for i, advantage := range advantages {
		stats.AddDopamine(15, advantage)
		fmt.Printf("   ✅ %d. %s (+15 дофамин)\n", i+1, advantage)
	}

	// 6. Лечение сожаления через код
	fmt.Println("\n💊 ЛЕЧЕНИЕ СОЖАЛЕНИЯ КОДОМ:")
	if stats.RegretLevel > 0 {
		fmt.Println("   Обнаружено сожаление! Лечим кодом...")
		codeTherapy := []struct {
			action   string
			dopamine int
		}{
			{"Написать generic функцию", 40},
			{"Протестировать type constraints", 35},
			{"Сделать коммит", 30},
			{"Спроектировать интерфейс", 25},
		}

		for _, therapy := range codeTherapy {
			stats.AddDopamine(therapy.dopamine, therapy.action)
			fmt.Printf("   💻 %s (+%d дофамин)\n", therapy.action, therapy.dopamine)
		}

		stats.RegretLevel = 0 // Обнуляем сожаление
	}

	// 7. Итоги
	stats.PrintStats()

	fmt.Println("\n════════════════════════════════════")
	fmt.Println("💡 ВЫВОД ДНЯ 68:")
	fmt.Println("   Type Constraints в Go — как блокировка троллей:")
	fmt.Println("   1. Определяешь, что допустимо (интерфейсы)")
	fmt.Println("   2. Запрещаешь несовместимое (компилятор не пропустит)")
	fmt.Println("   3. Получаешь безопасность типов и спокойствие")
	fmt.Println("\n   Заблокировал тролля? Не жалей!")
	fmt.Println("   Теперь у тебя больше времени на изучение Go.")
	fmt.Println("   А альпаки... они подождут, пока ты станешь senior.")
	fmt.Println("\n🚀 КОД > ТРОЛЛИ. CONSTRAINTS > СОЖАЛЕНИЙ.")
}
