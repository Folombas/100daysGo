package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	startDateStr     = "2025-11-03"
	challengeDays    = 100
	maxLevelXP       = 1000
	codeLinesPerDay  = 42.5
	cigsPerDay       = 15
	beerPerDay       = 3
	studyHoursPerDay = 2.5
	cigCost          = 15
	beerCost         = 120
	dailySavings     = 35
	baseSalary       = 80000
	salaryIncrease   = 1700
)

type Person struct {
	Name       string
	Age        int
	Background string
	Goal       string
}

type Progress struct {
	Days, XP, Level, Streak, Confidence int
	CodeLines                           float64
}

type Achievement struct {
	Emoji, Name, Desc string
	Day               int
}

type App struct {
	gosha        Person
	currentDay   int
	progress     Progress
	theme        string
	prng         *rand.Rand
	motivations  []string
	facts        []string
	achievements []Achievement
}

func NewApp() *App {
	currentDay := daysSince(startDateStr)
	progress := calculateProgress(currentDay)
	return &App{
		gosha: Person{
			Name:       "Гоша",
			Age:        38,
			Background: "Курьер с 20-летним стажем (бывший рэпер)",
			Goal:       "Стать Go-разработчиком и выйти из курьерского цикла",
		},
		currentDay: currentDay,
		progress:   progress,
		theme:      "Get a Brief Overview",
		prng:       rand.New(rand.NewSource(time.Now().UnixNano())),
		motivations: []string{
			"20 лет назад ты носил коробки за 250₽. Сегодня — за те же 250₽. С учётом инфляции это в 4 РАЗА меньше реальной стоимости! В IT зарплата растёт с каждым проектом.",
			"В курьерских приложениях опыт не ценится: новичок с iPhone 15 отхватит твой заказ за секунды. В IT твой опыт — это твоя защита и рост.",
			"Летом — жара +40°C, зимой — сугробы -25°C, осенью — лужи по колено. А плата за доставку не меняется 15 лет. Программист же получает +30% к зарплате за каждую новую технологию!",
			"Не бросай обучение! Каждый вечер кода — это шаг от холодных улиц к тёплому офису с кондиционером и зарплатой 250,000 ₽.",
			"Курьерская работа: чем больше опыт — тем больше износ суставов. IT-работа: чем больше опыт — тем выше зарплата и свобода.",
			"Твой рэп научил тебя ритму и структуре. Теперь примени это к коду — здесь опыт умножает твою ценность, а не разрушает тело.",
			"Сегодня ты возишь чужие подарки. Через год будешь получать свой подарок — первую зарплату разработчика, где опыт = деньги.",
		},
		facts: []string{
			"Карьерная лестница курьера: от 'бегущего по снегу' до 'бегущего по снегу с артритом'. В IT: Стажёр (80k) → Junior (150k) → Middle (250k) → Senior (400k) → Tech Lead (600k+)",
			"В курьерских агрегаторах алгоритм даёт хорошие заказы тому, кто быстрее нажал кнопку. В IT алгоритм отбора — твоё резюме и GitHub.",
			"Цена доставки не растёт с инфляцией. Зарплата разработчика растёт на 20-30% ежегодно. Твой опыт должен работать НА ТЕБЯ, а не против тебя!",
			"1 горутина в Go = 1 поток выполнения. 1 вечер обучения = 1 шаг к карьере, где опыт = деньги, а не боль в суставах.",
			"В IT ты конкурируешь со своим прошлым 'я'. В курьерке — с 20-летними подростками с айфончиками.",
			"Go создан для решения реальных проблем. И твоя проблема — выбраться из курьерского цикла — самая реальная из всех.",
		},
		achievements: []Achievement{
			{"🌱", "Новое начало!", "Первые 24 часа без иллюзий о курьерском будущем", 1},
			{"⚔️", "Цифровой Гуру", "7 дней чистого кода вместо ожидания заказов", 7},
			{"💎", "Сердце чемпиона", "Ты прошел четверть пути! 25 дней перемен!", 25},
			{"🚀", "Наполовину к звёздам", "50 дней без оглядки назад — только вперёд!", 50},
			{"🌟", "Полный круг", "100 дней нового Гоши — легенда в мире кода!", 100},
		},
	}
}

func main() {
	app := NewApp()
	app.printHeader()
	app.printProgress()
	app.printToday()
	app.printStats()
	app.printAchievements()
	app.printFuture()
	app.printFooter()
	app.interactiveLineCounter()
}

func daysSince(dateStr string) int {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		panic(fmt.Sprintf("invalid start date: %v", err))
	}
	return int(time.Since(t).Hours() / 24)
}

func calculateProgress(days int) Progress {
	xp := 100 + days*10
	return Progress{
		Days:       days,
		XP:         xp,
		Level:      1 + xp/maxLevelXP,
		Streak:     days,
		Confidence: min(100, days*2),
		CodeLines:  float64(days) * codeLinesPerDay,
	}
}

func (a *App) printHeader() {
	fmt.Printf("\n%s🔥 100 ДНЕЙ GО С ГОШЕЙ: ОТ КУРЬЕРСКИХ СУГРОБОВ К IT-ВЕРШИНАМ 🔥%s\n",
		ansi("1;33"), ansi("0"))
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("👤 %s%s%s | %d лет | %s\n",
		ansi("1;36"), a.gosha.Name, ansi("0"), a.gosha.Age, a.gosha.Background)
	fmt.Printf("🎯 %s%s%s\n",
		ansi("1;32"), a.gosha.Goal, ansi("0"))
	fmt.Printf("📅 %s | День: %d/%d | Тема: %s\n",
		time.Now().Format("02.01.2006"), a.currentDay, challengeDays, a.theme)
}

func (a *App) printProgress() {
	percent := float64(a.currentDay) / challengeDays * 100
	xpNeeded := a.progress.Level * maxLevelXP

	fmt.Printf("\n%s🚀 ПРОГРЕСС: %.0f%% завершено (Дней с 30 ноября: %d)%s\n",
		ansi("1;34"), percent, a.currentDay, ansi("0"))
	fmt.Println(progressBar(percent, 40))
	fmt.Printf("🏆 Уровень %d (%d/%d XP) | 💪 Уверенность: %d%%\n",
		a.progress.Level, a.progress.XP, xpNeeded, a.progress.Confidence)
}

func (a *App) printToday() {
	fmt.Printf("\n%s💡 СУТЬ ДНЯ: ПОЧЕМУ НЕЛЬЗЯ БРОСАТЬ САМООБУЧЕНИЕ%s\n", ansi("1;31"), ansi("0"))
	fmt.Println("   ┌───────────────────────────────────────────────────────────────────────┐")
	fmt.Println("   │ ❌ РЕАЛЬНОСТЬ КУРЬЕРА:                                               │")
	fmt.Println("   │   - 2005 год: коробка за 250₽ (нормальная зарплата)                  │")
	fmt.Println("   │   - 2025 год: та же коробка за 250₽ (с учётом инфляции = 60₽!)       │")
	fmt.Println("   │   - Опыт не ценится: новичок на iPhone отхватит твой заказ за 3 сек  │")
	fmt.Println("   │   - Лето: +40°C в асфальтовых джунглях                               │")
	fmt.Println("   │   - Зима: -25°C по колено в сугробах                                 │")
	fmt.Println("   │   - Осень: лужи до колен + мокрые ноги                               │")
	fmt.Println("   │   - Нет карьерного роста: только износ тела                          │")
	fmt.Println("   │                                                                      │")
	fmt.Println("   │ ✅ IT-ПУТЬ:                                                          │")
	fmt.Println("   │   - Стажёр (80,000 ₽) → Junior (150,000 ₽)                           │")
	fmt.Println("   │   → Middle (250,000 ₽) → Senior (400,000 ₽)                          │")
	fmt.Println("   │   → Tech Lead (600,000+ ₽)                                           │")
	fmt.Println("   │   - Каждый новый язык/фреймворк = +20-30% к зарплате                 │")
	fmt.Println("   │   - Опыт = деньги. Всегда. Даже в 38 лет.                            │")
	fmt.Println("   └───────────────────────────────────────────────────────────────────────┘")

	fmt.Printf("\n%s✨ СЕГОДНЯ ГОВОРИТ СЕРДЦЕ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💬 %s\n", a.motivations[a.currentDay%len(a.motivations)])
	fmt.Printf("   💡 %s\n", a.facts[a.currentDay%len(a.facts)])
}

func (a *App) printStats() {
	cigarettes := a.currentDay * cigsPerDay
	beerBottles := a.currentDay * beerPerDay
	studyHours := float64(a.currentDay) * studyHoursPerDay
	moneySaved := float64(cigarettes)*cigCost + float64(beerBottles)*beerCost + float64(a.currentDay)*dailySavings

	fmt.Printf("\n%s📊 СТАТИСТИКА ПЕРЕРОЖДЕНИЯ:%s\n", ansi("1;36"), ansi("0"))
	fmt.Printf("   🚭 Пропущено сигарет: %d | 🍺 Бутылок пива: %d\n", cigarettes, beerBottles)
	fmt.Printf("   💻 Часов обучения: %.1f | 💰 Сэкономлено: %.0f ₽\n", studyHours, moneySaved)
	fmt.Printf("   📝 Написано строк кода: %.0f | 🔥 Удалено игр: %d\n",
		a.progress.CodeLines, 7+a.currentDay/5)
	fmt.Printf("   📉 ПОТЕРИ ОТ КУРЬЕРКИ: %.0f ₽ (реальная стоимость 18 лет опыта)\n",
		float64(a.currentDay)*250*4) // 250₽ за заказ * 4 раза из-за инфляции
}

func (a *App) printAchievements() {
	unlocked := 0
	for _, ach := range a.achievements {
		if a.currentDay >= ach.Day {
			unlocked++
		}
	}

	fmt.Printf("\n%s🏆 ДОСТИЖЕНИЯ (%d/%d):%s\n", ansi("1;33"), unlocked, len(a.achievements), ansi("0"))
	for _, ach := range a.achievements {
		if a.currentDay >= ach.Day {
			fmt.Printf("   %s%s %s%s\n", ansi("1;32"), ach.Emoji, ach.Name, ansi("0"))
		}
	}
}

func (a *App) printFuture() {
	currentSalary := baseSalary + a.currentDay*salaryIncrease
	daysToJob := max(0, 45-a.currentDay)
	courierSalary := 45000 // средняя зарплата курьера в Москве

	fmt.Printf("\n%s🔮 БУДУЩЕЕ ЧЕРЕЗ 100 ДНЕЙ:%s\n", ansi("1;35"), ansi("0"))
	fmt.Printf("   💼 Go-разработчик в Биг-Техе (Текущая: %d ₽/мес → %s250,000 ₽/мес%s)\n",
		currentSalary, ansi("1;32"), ansi("0"))
	fmt.Printf("   📈 Карьера: Стажёр → Junior → Middle → Senior → Tech Lead\n")
	fmt.Printf("   ⚖️ КОНТРАСТ: Курьер (%d ₽/мес) vs IT-разработчик (250,000 ₽/мес)\n",
		courierSalary)
	fmt.Printf("   ❄️ КОНЕЦ СЕЗОННОЙ БОРЬБЫ: Ни сугробов, ни луж, ни жары — только кондиционер и кофе\n")
	fmt.Printf("   🏠 Своя квартира-студия в новом районе у метро (мечта с 30 ноября)\n")
	fmt.Printf("   👵 Родители гордятся тобой (а не тем, как ты 'прославился' в прошлом)\n")
	fmt.Printf("   ⏳ Работа найдется через %d дней. Ты справишься!\n", daysToJob)
}

func (a *App) printFooter() {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Printf("%s💬 ФИЛОСОФИЯ 38-ЛЕТНЕГО ГОШИ:%s\n", ansi("1;34"), ansi("0"))
	fmt.Println("   \"Курьерство — это временная подработка. IT — это инвестиция в будущее.\"")
	fmt.Println("   \"Твой 18-летний курьерский опыт не увеличит зарплату ни на копейку. Твой 1-летний опыт в Go — увеличит в 3 раза.\"")
	fmt.Println("   \"Не обменивай сегодняшний голод на завтрашнюю бедность. Каждый вечер кода — это страховка от вечного бега по сугробам.\"")

	birthdayMessage := "\n%s🎉 НАПОМИНАНИЕ: 30 ноября 2025 года тебе исполнилось 38 лет. " +
		"Это не поздно для старта в IT. Это идеальное время, чтобы превратить опыт жизни в преимущество!%s"
	fmt.Printf(birthdayMessage, ansi("1;33"), ansi("0"))

	fmt.Printf("\n%s🌟 СЕГОДНЯ: УДАЛИЛ 1 ИГРУ + НАПИСАЛ %.0f СТРОКИ КОДА! %s\n",
		ansi("1;32"), a.progress.CodeLines, ansi("0"))
	fmt.Printf("%s🔥 ЗАПОМНИ: %s\n", ansi("1;31"), ansi("0"))
	fmt.Println("   'Ты устал от того, что цена за доставку не растёт 15 лет? ")
	fmt.Println("   Зарплата разработчика растёт КАЖДЫЙ ГОД. Твой опыт = деньги. ")
	fmt.Println("   Не бросай это — это твой единственный выход из цикла сугробов и луж!'")
}

func progressBar(percent float64, width int) string {
	filled := int(percent/100*float64(width) + 0.5)
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func (a *App) interactiveLineCounter() {
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Print("🔍 Проверить прогресс за другой день (например: day25): ")

	var input string
	if _, err := fmt.Scanln(&input); err != nil {
		return
	}

	if !strings.HasPrefix(input, "day") {
		fmt.Printf("%s⚠️ Неправильный формат! Используй: day25%s\n", ansi("1;31"), ansi("0"))
		return
	}

	dirPath := filepath.Join("..", input)

	lines, err := countCodeLines(dirPath)
	if err != nil {
		fmt.Printf("%s❌ Ошибка: %v%s\n", ansi("1;31"), err, ansi("0"))
		return
	}

	emoji := "✅"
	switch {
	case lines > 100:
		emoji = "🔥"
	case lines < 10:
		emoji = "💪"
	}

	fmt.Printf("\n%s%s %s: %.0f строк кода!%s\n",
		ansi("1;32"), emoji, input, lines, ansi("0"))

	if lines > 0 {
		fmt.Printf("%s💡 Совет: Добавь комментарии и рефакторинг!%s\n", ansi("1;34"), ansi("0"))
		fmt.Printf("%s🚀 Напоминание:%s Каждая строка кода увеличивает твою ценность в IT. "+
			"Курьерский опыт не накапливается. Продолжай писать!\n", ansi("1;35"), ansi("0"))
	}
}

func countCodeLines(dir string) (float64, error) {
	var total float64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || filepath.Ext(path) != ".go" {
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
			if line == "" || strings.HasPrefix(line, "//") {
				continue
			}
			total++
		}
		return scanner.Err()
	})
	return total, err
}

func ansi(code string) string {
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
