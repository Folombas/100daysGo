package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	startDateStr   = "2025-10-25"
	challengeDays  = 100
	maxLevelXP     = 1000
	cigaretteCost  = 15.0
	beerBottleCost = 120.0
)

type (
	Stats struct {
		Days, XP, Level, CodingPower, Streak int
		Willpower, MentalState, Mood         string
	}
	Growth struct {
		CigarettesSkipped, BeerBottlesSkipped, PartyNightsSkipped int
		DigitalDetoxHours, StudyHours                             float64
		Confidence, Anxiety                                       float64
		MomPride, DaysClean                                       int
		AbandonedGames, AbandonedSeries                           int
		MoneySaved                                                float64
	}
	Event struct{ Emoji, Desc string }
	Achievement struct {
		Emoji, Name, Desc string
		Day               int
		Rarity            string
	}
	Quest struct {
		Name, Desc string
		Day        int
		Done       bool
	}
)

var (
	currentDay      = daysSince(startDateStr)
	stats           = initStats()
	growth          = initGrowth()
	motivation      = randomItem(motivations)
	dailyFact       = randomItem(goFacts)
	neuroQuirk      = getNeuroQuirk()
	achievements    = initAchievements()
	quests          = initQuests()
	r               = initRand()
	dailyTopic      = "Slices: Growth" // Тема дня
)

func main() {
	printHeroCard()
	printProgress()
	printDailyStruggle()
	printDigitalDetox()
	printGrowth()
	printNeuroUniqueness()
	printAchievements()
	printFuture()
	printFooter()
	interactiveLineCounter()
}

func initAchievements() []Achievement {
	return []Achievement{
		{"🌱", "Первый рассвет без дымки", "Выжил первую ночь без сигарет и бутылки пива", 1, "common"},
		{"🎮➡️💻", "Цифровая детоксикация", "Отказался от 24 часов игр и сериалов ради Go", 5, "uncommon"},
		{"🔥", "Неделя чистого кода", "7 дней без вечеринок, игр и сериалов — только горутины", 7, "uncommon"},
		{"💎", "Алмазная трезвость", "30 дней без цифрового мусора — только чистый Go-код", 30, "rare"},
		{"⚡", "Половина пути к свету", "50 дней без табачного тумана и сериалов в голове", 50, "epic"},
		{"🏆", "Мастер Go", "100 дней без цифрового мусора — победа!", 100, "legendary"},
	}
}

func initQuests() []Quest {
	return []Quest{
		{"Day 1", "Написать 'Hello, трезвый мир!'", 1, false},
		{"Day 10", "Заменить 4 часа сериалов на изучение Go", 10, false},
		{"Day 30", "Удалить 3 игры и 1 стриминговый сервис", 30, false},
		{"Day 100", "Получить оффер в Биг-Тех и купить маме дачу", 100, false},
	}
}

func initStats() Stats {
	xp := 100 + currentDay*10
	level := 1 + xp/maxLevelXP

	return Stats{
		Days:        currentDay,
		XP:          xp,
		Level:       level,
		CodingPower: clamp(10+currentDay*5, 0, 1000),
		Streak:      currentDay,
		Willpower:   willpowerLevel(),
		MentalState: mentalState(),
		Mood:        todayMood(),
	}
}

func initGrowth() Growth {
	return Growth{
		CigarettesSkipped:  currentDay * 15,
		BeerBottlesSkipped: currentDay * 3,
		PartyNightsSkipped: currentDay,
		DigitalDetoxHours:  float64(currentDay * 4), // 4 часа в день вместо сериалов
		StudyHours:         float64(currentDay) * 1.8,
		Confidence:         clampF(float64(currentDay)*1.5, 0, 100),
		Anxiety:            clampF(100-float64(currentDay)*2.5, 0, 100),
		MomPride:           clamp(currentDay*2, 0, 100),
		DaysClean:          currentDay,
		AbandonedGames:     7,  // Количество удалённых игр
		AbandonedSeries:    12, // Количество брошенных сериалов
		MoneySaved: float64(currentDay*15)*cigaretteCost +
			float64(currentDay*3)*beerBottleCost +
			float64(currentDay*4)*35, // Экономия на подписках
	}
}

func daysSince(dateStr string) int {
	t, _ := time.Parse(time.DateOnly, dateStr)
	t = t.UTC()
	now := time.Now().UTC()
	days := int(now.Sub(t).Hours() / 24)
	return clamp(days, 0, challengeDays)
}

func willpowerLevel() string {
	levels := []string{
		"Хрупкий (как аккаунт в утечке данных)",
		"Неустойчивый (как Wi-Fi в метро)",
		"Стабильный (как хороший алгоритм)",
		"Железный (как сервер в дата-центре)",
		"Алмазный (как чистый код после рефакторинга)",
	}
	return levels[min(currentDay/20, len(levels)-1)]
}

func mentalState() string {
	states := []string{
		"Туман от сериалов и пива",
		"Борьба с игровой зависимостью как с багами",
		"Чистый код вместо багов жизни",
		"Глубокий сон вместо жёсткого похмелья",
		"Поток ясного сознания как горутина",
	}
	return states[min(currentDay/20, len(states)-1)]
}

func todayMood() string {
	moods := []string{
		"Ностальгия по прокрастинации",
		"Сопротивление соблазну загуглить 'как быстро выучить Go'",
		"Уверенность в каждом if",
		"Гордость за закрытый issue в реальной жизни",
		"Свобода от цифрового мусора — как от legacy кода",
	}
	return moods[min(currentDay/20, len(moods)-1)]
}

func printHeroCard() {
	fmt.Printf("\n🔥 100 ДНЕЙ КОДА VS 20 ЛЕТ ЦИФРОВОГО АДА 🔥\n")
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("👤 Гоша | 38 лет | Бывший гламурный рэпер MC Fool\n")
	fmt.Printf("🎯 Миссия: %sстать Go-разработчиком%s\n",
		color("1;32"), color("0"))
	fmt.Printf("📅 %s | День %d/%d | Тема: %s\n",
		time.Now().UTC().Format("02.01.2006"),
		currentDay, challengeDays, dailyTopic)
	fmt.Printf("💻 %sСегодня вместо 3 часов сериалов — 2 часа Go!%s\n",
		color("1;35"), color("0"))
}

func printProgress() {
	percent := float64(currentDay) / challengeDays * 100
	fmt.Printf("\n🔥 ПРОГРЕСС ПЕРЕРОЖДЕНИЯ ИЗ СТЁБНОГО ФРИКА В ТРЕЗВОГО GOLANG-РАЗРАБОТЧИКА: %.0f%%\n",
		percent)
	fmt.Println(progressBar(percent, 30))

	xpNeeded := stats.Level * maxLevelXP
	fmt.Printf("🏆 Lvl %d (%d/%d XP) | 💪 %s\n",
		stats.Level, stats.XP, xpNeeded, stats.Willpower)
	fmt.Printf("🧠 %s | 😌 %s\n", stats.MentalState, stats.Mood)
}

func printDailyStruggle() {
	fmt.Printf("\n⚡ ЕЖЕДНЕВНАЯ БОРЬБА ЗА ФОКУС:\n")
	fmt.Printf("   💬 %s%s (%s): %s\"%s\"%s\n",
		color("1;33"), "Женя", "Go-энтузиаст", color("1;35"), "Гоша, твои 24 часа без сериалов — это как первый коммит в продакшен!", color("0"))
	fmt.Printf("   👵 Мама: %s\"%s\"%s\n", color("1;33"), "Сынок, вместо сериалов ты теперь разбираешься в типах данных. Это настоящее чудо!", color("0"))
	fmt.Printf("   🧙 Ментор: %s\"%s\"%s %s\n",
		color("1;34"), "Твоя зависимость от сериалов — не твоя личность. Каждая строка кода — шаг к свободе", color("0"), "🧠")
	fmt.Printf("   💫 Мотивация: %s%s%s\n",
		color("1;35"), motivation, color("0"))
	fmt.Printf("   🎲 Факт о Go: %s%s%s\n",
		color("1;36"), dailyFact, color("0"))
}

func printDigitalDetox() {
	fmt.Printf("\n🎮➡️💻 ЦИФРОВАЯ ДЕТОКСИКАЦИЯ:\n")
	fmt.Printf("   📵 Отказ от сериалов: %.0f часов (хватит на %d серий «Игры престолов»)\n",
		growth.DigitalDetoxHours, int(growth.DigitalDetoxHours/1.2))
	fmt.Printf("   🎮 Удалено игр: %d (включая «CyberPunk 2077» и «Call of Duty»)\n",
		growth.AbandonedGames)
	fmt.Printf("   📺 Отписка от стриминговых сервисов: Netflix → GitHub\n")
	fmt.Printf("   💸 Экономия на подписках: %.0f ₽ (хватит на механическую клавиатуру)\n",
		float64(currentDay)*35)
}

func printGrowth() {
	fmt.Printf("\n🌱 ПЕРЕЗАГРУЗКА ЖИЗНИ:\n")
	fmt.Printf("   🚭 Пропущено сигарет: %s%d%s (достаточно, чтобы %sобкуриться до безумства%s)\n",
		color("1;33"), growth.CigarettesSkipped, color("0"),
		color("1;36"), color("0"))
	fmt.Printf("   🍺 Пропущено пива: %s%d%s бутылок (весом с %sмаленького ослика%s)\n",
		color("1;33"), growth.BeerBottlesSkipped, color("0"),
		color("1;36"), color("0"))
	fmt.Printf("   💻 Часов обучения вместо сериалов: %.1f\n",
		growth.StudyHours)
	fmt.Printf("   💰 Сэкономлено всего: %.0f ₽ (из них %.0f ₽ — на цифровых наркотиках)\n",
		growth.MoneySaved, float64(currentDay)*35)
	fmt.Printf("   😊 Уверенность: %.0f/100 | 👵 Гордость мамы: %d/100\n",
		growth.Confidence, growth.MomPride)
}

func printNeuroUniqueness() {
	fmt.Printf("\n🧠 НЕЙРО-СУПЕРСИЛА СЕГОДНЯ:\n")
	fmt.Printf("   %s→ %s%s\n", neuroQuirk.Emoji, neuroQuirk.Desc, color("0"))
}

func printAchievements() {
	unlocked := countUnlocked(achievements)
	fmt.Printf("\n🏆 ДОСТИЖЕНИЯ: %d/%d\n",
		unlocked, len(achievements))

	for _, a := range achievements {
		if currentDay >= a.Day {
			rarityColor := rarityColor(a.Rarity)
			fmt.Printf("   %s%s %s: %s%s\n",
				color(rarityColor), a.Emoji, a.Name, a.Desc, color("0"))
		}
	}

	active := countActiveQuests()
	fmt.Printf("\n📜 КВЕСТЫ: %d активно\n", active)
	for _, q := range quests {
		if !q.Done && currentDay >= q.Day {
			fmt.Printf("   ➤ %s%s: %s%s\n",
				color("1;33"), q.Name, q.Desc, color("0"))
		}
	}
}

func printFuture() {
	targetSalary := 250000
	currentSalary := 80000 + currentDay*1700
	daysToJob := max(0, 45-currentDay)

	fmt.Printf("\n💰 ЦИФРОВОЕ БУДУЩЕЕ:\n")
	fmt.Printf("   💸 Зарплата: ~%s%d₽%s → %s%d₽/мес%s\n",
		color("1;33"), currentSalary, color("0"),
		color("1;32"), targetSalary, color("0"))
	fmt.Printf("   👨‍💻 Работа в Биг-Техе Go-Разработчиком (через %s%d дней%s)\n",
		color("1;34"), daysToJob, color("0"))
	fmt.Printf("   🏠 Мечта: квартира без экранов в спальне (через 6 месяцев)\n")
}

func printFooter() {
	fmt.Println(strings.Repeat("═", 50))
	fmt.Printf("💡 ФИЛОСОФИЯ ДНЯ:\n")
	fmt.Println("   \"Строка кода сильнее 5 серий сериала. Баг в коде исправить легче,")
	fmt.Println("   чем зависимость от цифрового мусора. Сегодня я выбрал Go —")
	fmt.Println("   завтра он выберет меня в Senior-разработчики.\"")
	fmt.Printf("\n🌟 СЕГОДНЯ Я УДАЛИЛ ЕЩЁ 1 ИГРУ И НАПИСАЛ ЕЩЁ 42 СТРОКИ КОДА!🌟\n")
}

func progressBar(percent float64, width int) string {
	bar := strings.Builder{}
	bar.Grow(width)
	filled := int(percent / 100 * float64(width))

	for i := 0; i < width; i++ {
		if i < filled {
			bar.WriteString(colorBar(i, width))
		} else {
			bar.WriteString(color("0") + "░")
		}
	}
	return bar.String()
}

func colorBar(index, total int) string {
	switch {
	case index < total/3:
		return color("31") + "█" // Красный
	case index < 2*total/3:
		return color("33") + "█" // Жёлтый
	default:
		return color("32") + "█" // Зелёный
	}
}

func rarityColor(rarity string) string {
	colors := map[string]string{
		"common":    "1;37",
		"uncommon":  "1;32",
		"rare":      "1;34",
		"epic":      "1;35",
		"legendary": "1;33",
	}
	return colors[rarity]
}

func getNeuroQuirk() Event {
	quirks := []Event{
		{"⚡", "СДВГ-гиперфокус: 4 часа кода вместо игры"},
		{"🧩", "ОКР помогает писать идеальный код без «технического долга»"},
		{"💡", "Социофобия: предпочитаю общение через GitHub Issues"},
		{"🎯", "Нейротипичное мышление: вижу паттерны в данных вместо сюжетов сериалов"},
	}
	return quirks[currentDay%len(quirks)]
}

func initRand() *rand.Rand {
	return rand.New(rand.NewPCG(
		uint64(time.Now().UnixNano()),
		uint64(time.Now().UnixNano()>>32),
	))
}

func randomItem[T any](items []T) T {
	return items[r.IntN(len(items))]
}

func clamp(value, minVal, maxVal int) int {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}

func clampF(value, minVal, maxVal float64) float64 {
	if value < minVal {
		return minVal
	}
	if value > maxVal {
		return maxVal
	}
	return value
}

func color(code string) string {
	return "\033[" + code + "m"
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func countUnlocked(achs []Achievement) int {
	count := 0
	for _, a := range achs {
		if currentDay >= a.Day {
			count++
		}
	}
	return count
}

func countActiveQuests() int {
	count := 0
	for _, q := range quests {
		if !q.Done && currentDay >= q.Day {
			count++
		}
	}
	return count
}

func countGoLines(dirPath string) (int, error) {
	totalLines := 0
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" && !strings.HasPrefix(line, "//") {
					totalLines++
				}
			}
		}
		return nil
	})
	return totalLines, err
}

func interactiveLineCounter() {
	fmt.Println("\n" + strings.Repeat("═", 50))
	fmt.Printf("📊 ХОТИТЕ УЗНАТЬ РЕАЛЬНОЕ КОЛИЧЕСТВО СТРОК В ДРУГОМ ДНЕ?\n")
	fmt.Print("Введите день челленджа (например: day19): ")

	var dayInput string
	fmt.Scanln(&dayInput)

	if !strings.HasPrefix(dayInput, "day") || len(dayInput) < 4 {
		fmt.Printf("%s⚠️ Ошибка: используйте формат 'dayXX'%s\n",
			color("1;31"), color("0"))
		return
	}

	dirPath := fmt.Sprintf("../%s", dayInput)
	lines, err := countGoLines(dirPath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("%s📁 Директория %s не существует!%s\n",
				color("1;33"), dayInput, color("0"))
		} else {
			fmt.Printf("%s❌ Ошибка при подсчёте: %v%s\n",
				color("1;31"), err, color("0"))
		}
		return
	}

	printLineCountResult(dayInput, lines)
}

func printLineCountResult(dayInput string, lines int) {
	emoji := map[bool]string{
		lines > 100: "🔥",
		lines < 10:  "😴",
		true:        "✅", // default
	}[true]

	fmt.Printf("%s%s 💻 В директории %s найдено %d программных строк кода%s\n",
		color("1;32"), emoji, dayInput, lines, color("0"))

	advice := map[bool]string{
		lines == 0:  "Запустите 'git checkout %s' чтобы увидеть код этого дня",
		lines < 50:  "Добавьте комментарии и рефакторинг для глубины изучения темы",
		lines >= 50: "Вы молодец! %d строк — это серьёзный прогресс для одного дня!",
	}[true]

	fmt.Printf("%s💡 %s%s\n",
		color("1;34"),
		fmt.Sprintf(advice, dayInput, lines),
		color("0"))
}

var (
	motivations = []string{
		"Каждый раз, когда хочешь включить сериал — напиши 5 строк на Go!",
		"Игровой прогресс даёт иллюзию развития. Go-код даёт реальный результат",
		"Твой потешный рэп-текст станет 'чистым Go-кодом' — без дублей и ошибок",
		"38 лет — идеальный возраст для перезагрузки. Как хороший рефакторинг legacy кода",
		"Код — это твой новый способ выражения. Go делает его мощным и чистым.",
		"Каждый день без игр — шаг к новому миру. Go — твой ключ к нему.",
		"Go — язык, который помогает строить не только приложения, но и свою жизнь.",
		"Ты не просто кодишь. Ты создаешь будущее.",
		"С каждым днём ты становишься сильнее. Go — твой инструмент для силы.",
		"Ты не просто изучаешь язык. Ты меняешь свою жизнь.",
		"Каждый день — это уровень в игре. Ты развиваешься!",
		"Ты не просто программируешь. Ты становишься супергероем IT!",
		"Ты проходишь уровень за уровнем, как в любимой игре!",
		"Ты не просто ученик. Ты мастер Go!",
		"Ты ведешь себя как герой в игре, где каждый день — новое испытание!",
		"Ты не просто кодер. Ты строитель будущего!",
		"Ты — игрок, который побеждает в игре своей жизни!",
		"Ты — геймер, который побеждает в самой большой игре!",
		"Ты — настоящий боец, который побеждает в игре жизни!",
		"Ты — игрок, который выигрывает в игре своего успеха!",
	}

	goFacts = []string{
		"Go создан для решения реальных проблем — как твоя",
		"10k горутин легче, чем 1 ночь прокрастинации в играх",
		"go fmt форматирует код автоматически — пусть и твоя жизнь станет упорядоченной",
		"Go компилируется в один бинарник — как твоя новая жизнь: простая и надёжная",
		"Go используется в Google, Docker, Kubernetes, Uber, Twitch и других крупных проектах",
		"Go — язык для современных системных приложений и микросервисов",
		"Средняя зарплата Go-разработчика в России: от 120,000 до 250,000 ₽",
		"Go — это простота, эффективность и надежность в одном языке",
		"Go позволяет создавать высокопроизводительные приложения с минимальной сложностью",
		"Изучение Go — это инвестиция в будущее вашей карьеры в IT",
		"Go — язык для тех, кто хочет создавать быстрые, надёжные и масштабируемые системы",
		"Go — идеальный выбор для создания облачных сервисов и микросервисов",
		"Go-сообщество активно развивается, предлагая множество библиотек и инструментов",
		"Go имеет встроенную поддержку конкурентности, что делает его идеальным для современных задач",
		"Go — это язык, который помогает сосредоточиться на решении задач, а не на сложностях языка",
		"Go используется для создания таких проектов, как Docker, Kubernetes, Terraform и т.д.",
		"Go — это язык, который делает разработку доступной и продуктивной для всех",
		"Go — это выбор тех, кто хочет быть частью будущего технологий",
		"Go — язык, который поможет тебе достичь финансовой независимости и карьерного роста",
		"Go — язык, который делает тебя сильнее и увереннее в своих возможностях",
		"Go — язык, который делает тебя гением в мире программирования!",
		"Go — язык, который дает тебе силы быть выше всех!",
		"Go — язык, который открывает двери в мир IT!",
		"Go — язык, который делает тебя самым лучшим!",
		"Go — язык, который делает тебя уникальным!",
		"Go — язык, который делает тебя сильнее!",
		"Go — язык, который делает тебя мудрее!",
		"Go — язык, который делает тебя богаче!",
		"Go — язык, который делает тебя счастливее!",
		"Go — язык, который делает тебя успешнее!",
	}
)
