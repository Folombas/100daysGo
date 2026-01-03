package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	hundredDaysStart = "2025-11-03"
	go365Start       = "2026-01-01"
	hundredDaysTotal = 100
	go365TotalDays   = 365
	maxLevelXP       = 1000
	codeLinesPerDay  = 67.3 // Увеличенное количество строк с фокусом
	deletedGames     = 9    // Удалено больше игр для фокуса
)

type Person struct {
	Name, Background, Goal string
	Age                    int
}

type Progress struct {
	HundredDaysCount, HundredDaysXP, HundredDaysLevel int
	Go365DaysCount, Go365XP, Go365Level               int
	CodeLines                                         float64
	FocusDepthLevel                                   int // Уровень глубины изучения
	MonthsWithoutDistractions                         int // Месяцы без распыления
}

type Achievement struct {
	Emoji, Name, Desc string
	Unlocked          bool
}

type App struct {
	gosha        Person
	currentDate  time.Time
	progress     Progress
	theme        string
	rng          *rand.Rand
	motivations  []string
	achievements []Achievement
	dailyThemes  []string
	dailyEvents  []string
	focusQuotes  []string
}

func NewApp() *App {
	now := time.Now()
	hundredDays := calculateDaysSince(hundredDaysStart)

	// Исправленный расчёт Go365 с правильным отсчётом
	baseGo365 := calculateDaysSince(go365Start)
	go365Days := max(1, baseGo365+1) // Начинаем с дня 1

	// Расчёт месяцев без распыления (с 1 января 2026)
	monthsWithoutDistractions := 0
	if go365Days > 0 {
		monthsWithoutDistractions = go365Days / 30
		if monthsWithoutDistractions == 0 && go365Days > 0 {
			monthsWithoutDistractions = 1
		}
	}

	return &App{
		gosha: Person{
			Name:       "Гоша",
			Age:        38,
			Background: "Экс-распылитель (Python/Java/C#/C++/JS) → Глубокий исследователь Go",
			Goal:       "Стать Go-специалистом экстра-класса. 2026: ПОГРУЖЕНИЕ В ГЛУБИНЫ GO",
		},
		currentDate: now,
		progress: Progress{
			HundredDaysCount:          hundredDays,
			HundredDaysXP:             min(hundredDays*15, hundredDaysTotal*15), // Больше XP за фокус
			HundredDaysLevel:          1 + hundredDays*15/maxLevelXP,
			Go365DaysCount:            go365Days,
			Go365XP:                   go365Days * 25, // Больше XP за углубление
			Go365Level:                1 + go365Days*25/maxLevelXP,
			CodeLines:                 float64(hundredDays+go365Days) * codeLinesPerDay,
			FocusDepthLevel:           go365Days / 7, // Уровень глубины (новые уровни каждую неделю)
			MonthsWithoutDistractions: monthsWithoutDistractions,
		},
		theme: "2026: СТРАТЕГИЯ ПОГРУЖЕНИЯ В ГЛУБИНЫ GO | УРОВЕНЬ ПОГРУЖЕНИЯ: %d | Тема: Интерфейсы — абстракция высшего порядка",
		rng:   rand.New(rand.NewPCG(uint64(now.UnixNano()), uint64(now.Unix()))),
		motivations: []string{
			"В 2025 ты распылялся на 5 языков. В 2026 ты углубляешься в 1 язык на 500%.",
			"Глубина изучения Go сегодня = глубина понимания проблем завтра. Копай глубже!",
			"Твой мозг — шахта знаний. Каждый день Go — новый туннель к драгоценным камням мастерства.",
			"Широта знаний создаёт дилетантов. Глубина знаний создаёт экспертов. Выбирай глубину.",
			"Интерфейсы в Go — это не просто типы. Это туннели в глубину системы. Иди по ним.",
			"Каждый отказ от другого языка — +100 к фокусу на Go. Каждая книга по Go — +100 к глубине.",
			"Твой GTX 1060 не рендерит 3D-миры. Он рендерит 3D-понимание Go: грамматика → семантика → идиомы.",
			"Методы определяют поведение. Функции — действия. Ты — метод. Каждый день — новая реализация интерфейса.",
			"Поверхностное знание 10 языков = 0. Глубокое знание 1 языка = ∞. GO = ∞.",
			"Ты не изучаешь Go. Ты ПОГРУЖАЕШЬСЯ в Go. Каждый день — новое измерение глубины.",
			"Компилятор Go не прощает поверхностности. Ты тоже не прощай себе распыления.",
			"Сборщик мусора убирает неиспользуемое. Ты убираешь лишние языки. Оставляй только Go.",
			"Горутины — это не потоки. Это уровни глубины понимания параллелизма. Запускай их глубже.",
			"Каждый день без C# — это день с Go. Каждый день без Java — это день ближе к мастерству.",
			"Глубина — это когда ты знаешь не только КАК работает defer, но и ПОЧЕМУ в таком порядке.",
			"Ты больше не прыгаешь с языка на язык. Ты роешь туннель в ядро Go. С каждым днем глубже.",
		},
		achievements: []Achievement{
			{"🔱", "Манифест глубины", "Полный отказ от всех языков кроме Go. Погружение началось.", false},
			{"🧠", "Нейронная перестройка", "Мозг перестроен с распыления на глубокую фокусировку", false},
			{"⚡", "Экстремальный фокус", "30 дней без единого отвлечения на другие технологии", false},
			{"⛏️", "Шахтёр знаний", "Достигнут уровень глубины 5: понимание компилятора Go", false},
			{"🧬", "Архитектор систем", "Спроектирована и реализована сложная система на чистом Go", false},
			{"🎯", "Стрелок по целям", "100% фокус на Go в течение 100 дней", false},
			{"🏊", "Ныряльщик в код", "Прочитано и понято 10000 строк исходного кода Go", false},
			{"🔍", "Детектив багов", "Найдены и исправлены баги в собственном понимании Go", false},
			{"🚀", "Вертикальный взлёт", "Переход от синтаксиса к философии языка", false},
		},
		dailyThemes: []string{
			"Погружение в интерфейсы: от поверхностного понимания к глубинному",
			"Глубина вместо ширины: почему 1 язык на 100% лучше 10 на 10%",
			"Стратегия погружения: как углубиться в Go, не распыляясь",
			"Туннелирование в ядро Go: от пользователя к контрибьютеру",
			"Архитектурное мышление на Go: глубина проектирования",
			"Компилятор Go как объект изучения: копаем глубже стандартной библиотеки",
			"Параллельные вселенные Go: глубокое понимание горутин",
			"Системное программирование на Go: от поверхностного к глубокому",
			"Производительность на уровне наносекунд: углубление в профилирование",
			"Философия Go: от кода к мировоззрению",
		},
		dailyEvents: []string{
			"Утром: 2 часа глубокого изучения интерфейсов Go (не поверхностно!)",
			"Днём: Анализ исходного кода стандартной библиотеки Go",
			"Вечером: Написание кода с тремя уровнями абстракции (глубина!)",
			"Ночью: Изучение внутренностей компилятора Go (погружение в глубину)",
			"Утром: Рефакторинг кода с увеличением глубины абстракций",
			"Днём: Написание бенчмарков для понимания производительности на глубоком уровне",
			"Вечером: Изучение работы сборщика мусора Go (не поверхностно!)",
			"Ночью: Анализ байт-кода Go для глубокого понимания",
			"Утром: Практика с продвинутыми паттернами Go (глубинные знания)",
			"Днём: Оптимизация алгоритмов с фокусом на глубину понимания",
			"Вечером: Написание документации, объясняющей глубинные концепции",
			"Ночью: Ментальное моделирование сложных систем на Go",
		},
		focusQuotes: []string{
			"«Глубина знаний в одном превосходит ширину знаний во многом» — Законы Гофера",
			"«Не распыляйся — погружайся. Не прыгай — копай. Не скользи по поверхности — ныряй в глубину»",
			"«Один язык, изученный на 1000%, лучше десяти языков на 10%»",
			"«Мастерство — это не знание многих вещей, а глубокое знание одной вещи»",
			"«Гофер не прыгает с ветки на ветку. Он роет глубокую нору к знаниям»",
			"«Интерфейсы Go — это не поверхностные контракты. Это туннели в глубину системы»",
			"«Каждая новая фича Go, изученная глубоко, — это новая глубина твоего мастерства»",
			"«Распыление создаёт дилетантов. Фокусировка создаёт экспертов. Погружение создаёт мастеров»",
			"«Ты больше не турист в мире языков. Ты исследователь глубин одного языка»",
			"«Поверхностное знание — иллюзия. Глубокое понимание — реальность. GO — реальность»",
		},
	}
}

func main() {
	app := NewApp()
	app.unlockAchievements()
	app.renderUI()
}

func (a *App) renderUI() {
	// Динамическое обновление темы с уровнем погружения
	a.theme = fmt.Sprintf(a.theme, a.progress.FocusDepthLevel)

	a.printHeader()
	a.printProgressSection()
	a.printDepthMetrics()
	a.printDailyInsight()
	a.printStatsSection()
	a.printAchievementsSection()
	a.printFutureSection()
	a.printFooter()
	a.interactiveCheck()
}

// --- СЕКЦИИ ИНТЕРФЕЙСА ---

func (a *App) printHeader() {
	// Исправленный вызов printTitle
	a.printTitle("⚡ 2026: СТРАТЕГИЯ ПОГРУЖЕНИЯ В GO | УРОВЕНЬ ГЛУБИНЫ: %d ⚡", "36", a.progress.FocusDepthLevel)
	a.printLine("▰", 75)

	a.printfColored("👤 %s | %d лет | %s\n", "36", a.gosha.Name, a.gosha.Age, a.gosha.Background)
	a.printfColored("🎯 %s\n", "32;1", a.gosha.Goal)

	a.printf("📅 %s | 100daysGo: День %d/%d | Go365: День %d/%d\n",
		a.currentDate.Format("02.01.2006"),
		a.progress.HundredDaysCount, hundredDaysTotal,
		a.progress.Go365DaysCount, go365TotalDays)

	a.printfColored("🧠 Тема погружения: %s\n", "34;1", a.theme)

	// Цитата фокуса дня
	focusQuote := a.getRandomItem(a.focusQuotes)
	a.printfColored("💬 %s\n", "33", focusQuote)

	a.printfColored("⚡ Уровень глубины изучения: %d | Месяцев без распыления: %d\n", "35;1",
		a.progress.FocusDepthLevel,
		a.progress.MonthsWithoutDistractions)
}

func (a *App) printProgressSection() {
	a.printSectionHeader("🚀 ПРОГРЕСС ПОГРУЖЕНИЯ", "34;1")

	hundredDaysPercent := a.progress.HundredDaysCount * 100 / hundredDaysTotal
	go365Percent := a.progress.Go365DaysCount * 100 / go365TotalDays

	// 100daysGo с акцентом на глубину
	a.printfColored("▸ 100daysGo (фундамент глубины): %.0f%% | Ур. глубины: %d | XP: %d/%d\n", "36",
		float64(hundredDaysPercent),
		a.progress.HundredDaysLevel,
		a.progress.HundredDaysXP,
		hundredDaysTotal*15)
	a.printDepthProgressBar(hundredDaysPercent, "█")

	// Go365 с акцентом на погружение
	a.printfColored("▸ Go365 (стратегия погружения): %.1f%% | Ур. глубины: %d | XP: %d/%d\n", "32;1",
		float64(go365Percent),
		a.progress.Go365Level,
		a.progress.Go365XP,
		go365TotalDays*25)
	a.printDepthProgressBar(go365Percent, "▓")

	a.printSectionHeader("📊 МЕТРИКИ УГЛУБЛЕНИЯ", "36")
	a.printf("   • Средняя глубина изучения: %.0f строк/день (не поверхностно!)\n", codeLinesPerDay)
	a.printf("   • Часы глубокого погружения: 3.5 часа/день (интенсив)\n")
	a.printf("   • Уровень концентрации: %s\n", a.getFocusDepthLevel())
	a.printf("   • Распыление в прошлом: %d языков → Сейчас: 1 язык (GO)\n", 5)
}

func (a *App) getFocusDepthLevel() string {
	depthLevel := a.progress.FocusDepthLevel
	switch {
	case depthLevel >= 20:
		return "🔱 ЭКСТРЕМАЛЬНАЯ ГЛУБИНА (уровень контрибьютера Go)"
	case depthLevel >= 15:
		return "🏊 ГЛУБОКОЕ ПОГРУЖЕНИЕ (уровень архитектора)"
	case depthLevel >= 10:
		return "⚡ СИЛЬНАЯ ФОКУСИРОВКА (уровень senior)"
	case depthLevel >= 5:
		return "🎯 УМЕРЕННОЕ ПОГРУЖЕНИЕ (уровень middle)"
	default:
		return "🌱 НАЧАЛО ПОГРУЖЕНИЯ (уровень junior)"
	}
}

func (a *App) printDepthMetrics() {
	a.printSectionHeader("⛏️ МЕТРИКИ ГЛУБИНЫ ИЗУЧЕНИЯ GO", "35;1")

	a.printBlock(60, func() {
		a.printf("📈 УРОВНИ ПОГРУЖЕНИЯ В GO:\n")
		a.printDepthLevel(1, "Синтаксис и базовые конструкции", a.progress.FocusDepthLevel >= 1)
		a.printDepthLevel(2, "Стандартная библиотека (50% изучено)", a.progress.FocusDepthLevel >= 2)
		a.printDepthLevel(3, "Продвинутые концепции (горутины, каналы)", a.progress.FocusDepthLevel >= 3)
		a.printDepthLevel(4, "Архитектурные паттерны и best practices", a.progress.FocusDepthLevel >= 4)
		a.printDepthLevel(5, "Внутренности компилятора и рантайма", a.progress.FocusDepthLevel >= 5)
		a.printDepthLevel(6, "Оптимизация и высоконагруженные системы", a.progress.FocusDepthLevel >= 6)
		a.printDepthLevel(7, "Контрибьютинг в open-source проекты", a.progress.FocusDepthLevel >= 7)
		a.printDepthLevel(8, "Создание собственных инструментов и фреймворков", a.progress.FocusDepthLevel >= 8)
		a.printDepthLevel(9, "Экспертиза в нишевых областях (системное програмирование)", a.progress.FocusDepthLevel >= 9)
		a.printDepthLevel(10, "Мастерство уровня создателей языка", a.progress.FocusDepthLevel >= 10)
	})
}

func (a *App) printDepthLevel(level int, description string, achieved bool) {
	status := "🔒"
	color := "37"
	if achieved {
		status = "✅"
		color = "32"
	}
	a.printfColored("   %s Уровень %d: %s\n", color, status, level, description)
}

func (a *App) printDailyInsight() {
	monthNames := []string{"ЯНВАРЯ", "ФЕВРАЛЯ", "МАРТА", "АПРЕЛЯ", "МАЯ", "ИЮНЯ",
		"ИЮЛЯ", "АВГУСТА", "СЕНТЯБРЯ", "ОКТЯБРЯ", "НОЯБРЯ", "ДЕКАБРЯ"}
	dateLabel := fmt.Sprintf("%02d %s %d",
		a.currentDate.Day(),
		monthNames[a.currentDate.Month()-1],
		a.currentDate.Year())

	theme := a.getRandomItem(a.dailyThemes)
	motivation := a.getRandomItem(a.motivations)
	events := a.getRandomItems(a.dailyEvents, a.rng.IntN(3)+1)

	a.printSectionHeader(fmt.Sprintf("💡 СУТЬ %s: %s", dateLabel, theme), "31;1")
	a.printBlock(65, func() {
		a.printf("❌ ПРОШЛОЕ (РАСПЫЛЕНИЕ 2023-2025):\n")
		a.printBullet("Языковая шизофрения: Python → Java → C# → C++ → JavaScript")
		a.printBullet("Поверхностное изучение: 'Hello World' на 10 языках, мастерство на 0")
		a.printBullet("Энергия распылялась, а не концентрировалась. Результат = 0")
		a.printBullet("ГТХ 1060 страдала от Unreal Engine, а не росла в Go-мастерстве")

		a.printf("\n✅ НАСТОЯЩЕЕ (ПОГРУЖЕНИЕ 2026):\n")
		for _, event := range events {
			a.printBullet(event)
		}
	})

	a.printSectionHeader("✨ МОТИВАЦИЯ ПОГРУЖЕНИЯ", "35;1")
	a.printf("💬 %s\n", motivation)

	a.printSectionHeader("☁️ ПРОГНОЗ ДЛЯ ПОГРУЖЕНИЯ", "36")
	a.printf("   • Ментальная погода: %s\n", a.getMentalWeather())
	a.printf("   • Рекомендуемое действие для углубления: %s\n", a.getDeepAction())
}

func (a *App) getMentalWeather() string {
	weather := []string{
		"Ясное сознание для глубокого погружения в интерфейсы",
		"Гроза сложных концепций (идеально для прорыва в понимании)",
		"Туман неопределенности (прояснится после 3 часов кода)",
		"Снегопат новых знаний (копай глубже, пока не найдешь золото)",
		"Ураган продуктивности (держи курс на глубину!)",
	}
	return weather[a.rng.IntN(len(weather))]
}

func (a *App) getDeepAction() string {
	actions := []string{
		"Изучи исходный код интерфейсов в стандартной библиотеке",
		"Напиши код с тремя уровнями абстракции (не поверхностно!)",
		"Проанализируй байт-код своей программы для глубинного понимания",
		"Оптимизируй алгоритм, уменьшив сложность на один порядок",
		"Прочитай документацию по компилятору Go (уровень погружения +1)",
		"Реализуй сложный паттерн без подглядывания в примеры",
		"Напиши тесты, покрывающие все пограничные случаи",
	}
	return actions[a.rng.IntN(len(actions))]
}

func (a *App) printStatsSection() {
	totalDays := a.progress.HundredDaysCount + a.progress.Go365DaysCount
	learningHours := float64(totalDays) * 3.5 // Больше часов для глубины
	freedomHours := float64(deletedGames) * 4.0

	a.printSectionHeader("📊 СТАТИСТИКА ПОГРУЖЕНИЯ", "36;1")
	a.printBullet(fmt.Sprintf("Удалено игр и отвлекалок: %d (+%.1f часов глубины/день)", deletedGames, freedomHours))
	a.printBullet(fmt.Sprintf("Написано строк кода с фокусом: %.0f (глубокое изучение)", a.progress.CodeLines))
	a.printBullet(fmt.Sprintf("Часов погружения в Go: %.1f | Интенсивность: 3.5 часа/день", learningHours))
	a.printBullet(fmt.Sprintf("Репозиториев глубины: 2 (100daysGo + Go365) | Уровень: %d", a.progress.FocusDepthLevel))
	a.printBullet("Заблокировано навсегда: Unity, IntelliJ, Unreal Engine, Blender, VS (кроме Code)")

	a.printSectionHeader("📈 СТАТИСТИКА СЕГОДНЯШНЕГО ПОГРУЖЕНИЯ", "36")
	a.printBullet(fmt.Sprintf("Дата: %s", a.currentDate.Format("02.01.2006")))
	a.printBullet(fmt.Sprintf("День Go365: %d | Уровень глубины: %d", a.progress.Go365DaysCount, a.progress.FocusDepthLevel))
	a.printBullet(fmt.Sprintf("День 100daysGo: %d", a.progress.HundredDaysCount))
	a.printBullet(fmt.Sprintf("Строк кода с фокусом: %.0f (не поверхностно!)", codeLinesPerDay))
	a.printBullet(fmt.Sprintf("Экономия энергии: %.1f часа (вместо распыления)", freedomHours))
	a.printBullet(fmt.Sprintf("Месяцев без распыления: %d", a.progress.MonthsWithoutDistractions))
}

func (a *App) printAchievementsSection() {
	unlocked := countUnlocked(a.achievements)
	a.printSectionHeader(fmt.Sprintf("🏆 ДОСТИЖЕНИЯ ПОГРУЖЕНИЯ (%d/%d)", unlocked, len(a.achievements)), "33;1")

	for _, ach := range a.achievements {
		status := "🔒"
		color := "37"
		if ach.Unlocked {
			status = "✅"
			color = "32"
		}
		a.printfColored("   %s %s: %s\n", color, status, ach.Name, ach.Desc)
	}

	nextAchievement := a.getNextAchievement()
	if nextAchievement != "" {
		a.printSectionHeader("🔍 СЛЕДУЮЩИЙ УРОВЕНЬ ПОГРУЖЕНИЯ", "35;1")
		a.printBullet(nextAchievement)
	}
}

func (a *App) getNextAchievement() string {
	for _, ach := range a.achievements {
		if !ach.Unlocked {
			return fmt.Sprintf("%s: %s", ach.Name, ach.Desc)
		}
	}
	return "Все уровни глубины достигнуты! Ты — мастер погружения!"
}

func (a *App) printFutureSection() {
	// Зарплата растет с уровнем глубины
	baseSalary := 120000
	depthBonus := a.progress.FocusDepthLevel * 25000
	currentSalary := baseSalary + depthBonus

	a.printSectionHeader("🔮 БУДУЩЕЕ ПОСЛЕ ГЛУБОКОГО ПОГРУЖЕНИЯ", "35;1")
	a.printf("💼 Go-специалист уровня глубины %d: %s%d ₽/мес → %d ₽/мес%s\n",
		a.progress.FocusDepthLevel,
		ansi("31;1"), currentSalary, 500000, ansi("0"))
	a.printBullet(fmt.Sprintf("Карьера: Junior (сейчас) → Middle (уровень глубины 5) → Senior (уровень 10) → Expert (уровень 15+)"))
	a.printBullet("Свобода: Удалённая работа из любой точки мира + глубокое понимание систем")
	a.printBullet(fmt.Sprintf("GTX 1060: Теперь компилирует глубокие знания, а не рендерит поверхностную графику"))
	a.printBullet(fmt.Sprintf("Финал 100daysGo: %d дней | До экспертного уровня: %d дней",
		hundredDaysTotal-a.progress.HundredDaysCount,
		500-a.progress.Go365DaysCount))

	a.printSectionHeader("📅 ПРОГНОЗ НА ЗАВТРАШНЕЕ ПОГРУЖЕНИЕ", "35")
	a.printBullet(fmt.Sprintf("Дата: %s", a.currentDate.AddDate(0, 0, 1).Format("02.01.2006")))
	a.printBullet(fmt.Sprintf("День Go365: %d | Уровень глубины: %d", a.progress.Go365DaysCount+1, a.progress.FocusDepthLevel))
	a.printBullet(fmt.Sprintf("День 100daysGo: %d", a.progress.HundredDaysCount+1))
	a.printBullet(fmt.Sprintf("Запланированное погружение: %s", a.getRandomItem(a.dailyEvents)))
}

func (a *App) printFooter() {
	a.printLine("▰", 75)
	a.printSectionHeader("💬 МАНИФЕСТ ПОГРУЖЕНИЯ НА 2026", "34;1")
	a.printBullet("2026: ГОД ПОГРУЖЕНИЯ В GO. НИКАКОГО РАСПЫЛЕНИЯ!")
	a.printBullet("Каждый день — новый уровень глубины. Каждая строка кода — туннель к мастерству.")
	a.printBullet("Мой Гофер не прыгает по веткам. Он роет глубокую нору к ядру языка.")
	a.printBullet("Широта — для дилетантов. Глубина — для мастеров. Я выбираю глубину.")

	a.printSectionHeader(fmt.Sprintf("🎉 %s: ДЕНЬ ПОГРУЖЕНИЯ №%d", a.currentDate.Format("02.01.2006"), a.progress.Go365DaysCount), "33;1")
	for _, event := range a.getRandomItems(a.dailyEvents, 2) {
		a.printBullet(event)
	}

	a.printf("\n%s🚀 ПОМНИ: ГЛУБИНА ЗНАНИЙ В ОДНОМ ПРЕВОСХОДИТ ШИРИНУ ВО МНОГИХ. КОПАЙ ГЛУБЖЕ!%s\n",
		ansi("35;1"), ansi("0"))

	a.printSectionHeader("📅 КАЛЕНДАРЬ ПОГРУЖЕНИЯ GO365", "36")
	a.printBullet(fmt.Sprintf("Дней погружения: %d | До экспертного уровня: %d дней",
		a.progress.Go365DaysCount, 500-a.progress.Go365DaysCount))
	a.printBullet(fmt.Sprintf("Следующий уровень глубины через: %d дней",
		max(0, 7-(a.progress.Go365DaysCount%7))))
	a.printBullet(fmt.Sprintf("Месяцев без распыления: %d | Цель: 12 месяцев чистого Go",
		a.progress.MonthsWithoutDistractions))
}

// --- ВСПОМОГАТЕЛЬНЫЕ МЕТОДЫ ---

func (a *App) unlockAchievements() {
	a.achievements[0].Unlocked = a.progress.Go365DaysCount >= 1
	a.achievements[1].Unlocked = a.progress.Go365DaysCount >= 7
	a.achievements[2].Unlocked = a.progress.Go365DaysCount >= 30
	a.achievements[3].Unlocked = a.progress.FocusDepthLevel >= 5
	a.achievements[4].Unlocked = a.progress.FocusDepthLevel >= 8
	a.achievements[5].Unlocked = a.progress.Go365DaysCount >= 100
	a.achievements[6].Unlocked = a.progress.CodeLines >= 10000
	a.achievements[7].Unlocked = a.progress.HundredDaysCount >= 50
	a.achievements[8].Unlocked = a.progress.FocusDepthLevel >= 3
}

func (a *App) interactiveCheck() {
	a.printLine("▰", 75)
	a.printSectionHeader("🔍 ПРОВЕРИТЬ ГЛУБИНУ ПОГРУЖЕНИЯ", "36")

	fmt.Println("   - Для проверки глубины изучения: введите номер уровня (1-10)")
	fmt.Println("   - Для статистики дня: введите дату (например: 2026-01-01)")
	fmt.Print("   Ваш выбор: ")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return
	}

	// Проверка на уровень глубины
	if strings.Contains(input, "-") {
		dirPath := a.getProgressPath(input)
		lines, err := countCodeLines(dirPath)
		if err != nil {
			a.printfColored("❌ Ошибка глубины: %v\n", "31", err)
			return
		}

		emoji := "✅"
		if lines > 150 {
			emoji = "🔥"
		} else if lines < 30 {
			emoji = "💪"
		}

		a.printfColored("\n%s Глубина погружения за %s: %.0f строк кода!\n", "32;1", emoji, input, lines)
		if lines > 0 {
			a.printfColored("💡 Совет для углубления: Проанализируй каждую строку, а не просто напиши!\n", "34;1")
		}
	} else {
		level, err := strconv.Atoi(input)
		if err == nil && level >= 1 && level <= 10 {
			if a.progress.FocusDepthLevel >= level {
				a.printfColored("\n✅ Уровень глубины %d достигнут! Продолжай погружение!\n", "32", level)
			} else {
				a.printfColored("\n🔒 Уровень глубины %d еще не достигнут. Копай глубже!\n", "31", level)
				a.printfColored("   Необходимо: %d дней непрерывного погружения\n", "33", level*7)
			}
		}
	}
}

func (a *App) getProgressPath(input string) string {
	if strings.Contains(input, "-") {
		return filepath.Join("..", "Go365", input)
	}
	return filepath.Join("..", fmt.Sprintf("day%s", input))
}

// --- УНИВЕРСАЛЬНЫЕ УТИЛИТЫ ---

func (a *App) getRandomItem(items []string) string {
	return items[a.rng.IntN(len(items))]
}

func (a *App) getRandomItems(items []string, count int) []string {
	if count > len(items) {
		count = len(items)
	}

	result := make([]string, 0, count)
	used := make(map[int]bool)

	for len(result) < count {
		idx := a.rng.IntN(len(items))
		if !used[idx] {
			used[idx] = true
			result = append(result, items[idx])
		}
	}
	return result
}

func countUnlocked(achievements []Achievement) int {
	count := 0
	for _, a := range achievements {
		if a.Unlocked {
			count++
		}
	}
	return count
}

func calculateDaysSince(dateStr string) int {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return 0
	}
	days := int(time.Since(t).Hours() / 24)
	return max(0, days)
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !isCodeFile(path) {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line != "" && !strings.HasPrefix(line, "//") && !strings.HasPrefix(line, "#") && !strings.HasPrefix(line, "/*") {
				total++
			}
		}
		return scanner.Err()
	})
	return total, err
}

func isCodeFile(path string) bool {
	ext := filepath.Ext(path)
	return ext == ".go" || ext == ".md" || ext == ".txt"
}

// --- ФОРМАТТЕРЫ И ЦВЕТА ---

func (a *App) printTitle(format, color string, args ...any) {
	// Исправленный метод: сначала форматируем текст, затем применяем цвет
	fullText := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s\n", ansi(color+";1"), fullText, ansi("0"))
}

func (a *App) printSectionHeader(text, color string) {
	fmt.Printf("\n%s%s%s\n", ansi(color+";1"), text, ansi("0"))
}

func (a *App) printLine(char string, count int) {
	fmt.Println(strings.Repeat(char, count))
}

func (a *App) printf(format string, args ...any) {
	fmt.Printf(format, args...)
}

func (a *App) printfColored(format, color string, args ...any) {
	formatted := fmt.Sprintf(format, args...)
	fmt.Printf("%s%s%s", ansi(color), formatted, ansi("0"))
}

func (a *App) printBlock(width int, content func()) {
	fmt.Println("   ┌" + strings.Repeat("─", width) + "┐")
	content()
	fmt.Println("   └" + strings.Repeat("─", width) + "┘")
}

func (a *App) printBullet(text string) {
	fmt.Printf("   │   • %s\n", text)
}

func (a *App) printDepthProgressBar(percent int, fillChar string) {
	width := 50
	filled := percent * width / 100

	// Градиент глубины: чем больше заполнено, тем "глубже" символ
	bar := ""
	for i := 0; i < width; i++ {
		if i < filled {
			if i < width/3 {
				bar += "█" // Поверхностный уровень
			} else if i < 2*width/3 {
				bar += "▓" // Средняя глубина
			} else {
				bar += "░" // Глубокая зона
			}
		} else {
			bar += " "
		}
	}

	fmt.Printf("[%s] %d%%\n", bar, percent)
}

func ansi(code string) string {
	return "\033[" + code + "m"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
