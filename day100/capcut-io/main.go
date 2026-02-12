package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const (
	resistanceFile = "capcut_resistance.json"
	startDate      = "2026-01-18"
)

// ResistanceData хранит данные о сопротивлении
type ResistanceData struct {
	StartDate            string   `json:"start_date"`
	LastCheck            string   `json:"last_check"`
	DaysCount            int      `json:"days_count"`
	UnlockedAchievements []string `json:"unlocked_achievements"`
	TotalXP              int      `json:"total_xp"`
}

var achievements = map[int]string{
	1:   "🚀 Первые 24 часа",
	7:   "🛡️ Неделя стойкости",
	14:  "⚡ Go-туннель виден",
	30:  "🏆 Месяц фокуса",
	60:  "👑 Мастер конвейеров",
	100: "🎯 Легенда 100 дней",
}

func main() {
	fmt.Println("❄️ ДЕНЬ 100: ФИНАЛЬНЫЙ РУБЕЖ ❄️")
	fmt.Println("=================================")

	data := loadOrCreateData()
	today := time.Now().Format("2006-01-02")

	// Вычисляем дни с даты старта
	start, _ := time.Parse("2006-01-02", data.StartDate)
	current, _ := time.Parse("2006-01-02", today)
	days := int(current.Sub(start).Hours() / 24)

	// Обновляем данные
	data.DaysCount = days
	data.LastCheck = today
	checkAchievements(data)
	saveData(data)

	// Вывод легенды и статистики
	printSaga(data)
	printMotivation(data)
	printDisclaimer()
}

func loadOrCreateData() *ResistanceData {
	file, err := os.Open(resistanceFile)
	if os.IsNotExist(err) {
		// Первый запуск — создаём новый файл
		data := &ResistanceData{
			StartDate:            startDate,
			LastCheck:            time.Now().Format("2006-01-02"),
			DaysCount:            0,
			UnlockedAchievements: []string{},
			TotalXP:              0,
		}
		saveData(data)
		return data
	} else if err != nil {
		panic(err)
	}
	defer file.Close()

	var data ResistanceData
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		panic(err)
	}
	return &data
}

func saveData(data *ResistanceData) {
	file, err := os.Create(resistanceFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		panic(err)
	}
}

func checkAchievements(data *ResistanceData) {
	for days, name := range achievements {
		if data.DaysCount >= days && !contains(data.UnlockedAchievements, name) {
			data.UnlockedAchievements = append(data.UnlockedAchievements, name)
			data.TotalXP += days * 10
			fmt.Printf("🏅 ДОСТИЖЕНИЕ РАЗБЛОКИРОВАНО: %s (+%d XP)\n", name, days*10)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func printSaga(data *ResistanceData) {
	saga := fmt.Sprintf(`
╔══════════════════════════════════════════════════════════╗
║              🧠  ЛЕГЕНДА ДНЯ 100  🧠                    ║
╠══════════════════════════════════════════════════════════╣
║ 12 февраля 2026. Снег за окном. Гоша на кухне.          ║
║ Чай, вафли, бисквитный рулетик. Планшет в руках.        ║
║ Стратегия «Мозговой штурм»: курсы, конференции, подкасты.║
║ — всё о Go. Дворники чистят снег, а Гоша чистит код.    ║
║                                                          ║
║ Капля за каплей, коммит за коммитом.                   ║
║ CapCut был удалён пару недель назад. Сегодня — юбилей.   ║
╚══════════════════════════════════════════════════════════╝

📅 Старт сопротивления: %s
📆 Сегодня: %s
🔥 Дней без CapCut: %d
⭐ Всего XP: %d
🏆 Разблокировано достижений: %d
`, data.StartDate, data.LastCheck, data.DaysCount, data.TotalXP, len(data.UnlockedAchievements))

	fmt.Println(saga)
}

func printMotivation(data *ResistanceData) {
	phrases := []string{
		"1. 100 дней — это 1/3 года. Ты прошёл эту треть как профи.",
		"2. CapCut подождёт. Go не ждёт. Твой мозг уже компилирует быстрее.",
		"3. Депрессия отступает перед файлом `go.mod`. Каждый модуль — шаг вперёд в обучении Go и Go-стека",
		"4. Вафли и чай — топливо. Код — продукт. Ты — фабрика.",
		"5. Не отвлекайся на монтаж — монтируй пакеты и структуры.",
		"6. Снег растает, а твои навыки останутся навсегда.",
		"7. Ты изучаешь Go без видеоигр, баров и сериалов. Вместо этого — ты выиграл главную игру.",
		"8. Каждая строчка Go — это строка в твоём резюме.",
		"9. Сегодня заканчивается 100daysGo. Завтра продолжается Go365. Бесконечный рефакторинг себя.",
		"10. Ты не просто учишь язык. Ты строишь личность разработчика. Поздравляю, Гоша!",
	}
	fmt.Println("💬 10 МОТИВАЦИОННЫХ КОДОВ НА ФИНИШЕ:")
	for _, p := range phrases {
		fmt.Println(p)
	}
	fmt.Println()
}

func printDisclaimer() {
	fmt.Println("=== DISCLAIMER ===")
	fmt.Println("Все персонажи «Гошиных Daily Code Life Story» выдуманы.")
	fmt.Println("Сюжеты созданы исключительно для мотивации и метафор в учебном процессе.")
	fmt.Println("Любые совпадения с реальными людьми или событиями случайны.")
	fmt.Println("CapCut — отличный редактор, но сейчас не время его устанавливать.")
	fmt.Println("Вафли «К чаю» — настоящие, их можно купить в любом универсаме.")
	fmt.Println("===================")
}
