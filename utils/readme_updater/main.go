package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/template"
	"time"
	"unicode/utf8"
)

// Конфигурация проекта
const (
	TotalDays        = 100
	ProgressBarSize  = 40
	DataFilePath     = "utils/readme_updater/progress.json"
	TemplateFilePath = "utils/readme_updater/template.md"
	ReadmeFilePath   = "README.md"
	StartDate        = "2025-07-25T00:00:00Z"
)

// Структуры данных
type ProgressData struct {
	StartDate        string
	CurrentDay       int
	Streak           int
	ProgressTable    string
	ProgressPercent  int
	ProgressPadding  string
	ProgressBar      string
	DaysWithoutGames int
	CommitCount      string
	LinesOfCode      int
	NextMilestone    int
	LastUpdated      string
}

type DayRecord struct {
	Number     int    `json:"number"`
	Date       string `json:"date"`
	Title      string `json:"title"`
	KeyInsight string `json:"keyInsight"`
	LinesCode  int    `json:"linesCode"`
}

type ProjectData struct {
	StartDate string      `json:"startDate"`
	Days      []DayRecord `json:"days"`
}

func main() {
	startTime := time.Now()
	fmt.Println("🚀 Запуск генератора README...")

	// Загрузка данных проекта
	projectData := loadProjectData()
	currentDay := calculateCurrentDay(projectData.StartDate)

	// Автоматически добавляем новый день, если сегодня еще не добавлен
	addNewDayIfNeeded(&projectData)

	// Подготовка данных для шаблона
	templateData := ProgressData{
		StartDate:        formatDate(projectData.StartDate),
		CurrentDay:       currentDay,
		Streak:           calculateStreak(projectData.Days),
		ProgressTable:    generateProgressTable(projectData.Days),
		ProgressPercent:  calculateProgressPercent(currentDay),
		ProgressPadding:  calculateProgressPadding(currentDay),
		ProgressBar:      generateProgressBar(currentDay),
		DaysWithoutGames: currentDay,
		CommitCount:      getCommitCount(),
		LinesOfCode:      countLinesOfCode(),
		NextMilestone:    calculateNextMilestone(currentDay),
		LastUpdated:      time.Now().Format("02.01.2006 15:04:05"),
	}

	// Генерация README
	if err := generateReadme(templateData); err != nil {
		fmt.Printf("❌ Ошибка генерации README: %v\n", err)
		os.Exit(1)
	}

	// Сохраняем обновленные данные прогресса
	saveProjectData(projectData)

	fmt.Printf("✅ README успешно обновлен за %v\n", time.Since(startTime))
	fmt.Printf("📊 Прогресс: %d%% | 🚀 Стрик: %d дней | 💾 Коммитов: %s | 📜 Строк кода: %d\n",
		templateData.ProgressPercent, templateData.Streak, templateData.CommitCount, templateData.LinesOfCode)
}

func loadProjectData() ProjectData {
	data, err := os.ReadFile(DataFilePath)
	if err != nil {
		fmt.Println("⚠️ Файл прогресса не найден, используются данные по умолчанию")
		return ProjectData{
			StartDate: StartDate,
			Days:      []DayRecord{},
		}
	}

	var projectData ProjectData
	if err := json.Unmarshal(data, &projectData); err != nil {
		fmt.Printf("❌ Ошибка чтения данных: %v\n", err)
		return ProjectData{
			StartDate: StartDate,
			Days:      []DayRecord{},
		}
	}
	return projectData
}

func saveProjectData(data ProjectData) {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("⚠️ Ошибка сериализации данных: %v\n", err)
		return
	}

	if err := os.WriteFile(DataFilePath, file, 0644); err != nil {
		fmt.Printf("⚠️ Ошибка записи в файл: %v\n", err)
	}
}

func calculateCurrentDay(startDate string) int {
	start, err := time.Parse(time.RFC3339, startDate)
	if err != nil {
		fmt.Printf("⚠️ Ошибка формата даты: %v\n", err)
		start, _ = time.Parse(time.RFC3339, StartDate)
	}

	days := int(time.Since(start).Hours() / 24)
	return days // возвращаем количество дней без +1
}

func calculateStreak(days []DayRecord) int {
	if len(days) == 0 {
		return 0
	}

	// Создаем карту дней для быстрого поиска
	dateMap := make(map[string]bool)
	for _, day := range days {
		dateMap[day.Date] = true
	}

	// Рассчитываем стрик
	streak := 0
	currentDate := time.Now()

	for {
		dateStr := currentDate.Format("02.01.2006")
		if !dateMap[dateStr] {
			break
		}
		streak++
		currentDate = currentDate.AddDate(0, 0, -1)
	}
	return streak
}

func generateProgressTable(days []DayRecord) string {
	// Сортируем дни по номеру в обратном порядке (последние дни вверху)
	sort.Slice(days, func(i, j int) bool {
		return days[i].Number > days[j].Number
	})

	var table strings.Builder
	table.WriteString("| День | Дата | Тема | Ключевое понимание |\n")
	table.WriteString("|------|------|------|---------------------|\n")

	for _, day := range days {
		// Ограничиваем длину строк для таблицы
		title := truncate(day.Title, 20)
		insight := truncate(day.KeyInsight, 30)
		table.WriteString(fmt.Sprintf(
			"| Day%d | %s | %s | %s |\n",
			day.Number, day.Date, title, insight,
		))
	}
	return table.String()
}

// Правильное усечение строк с учетом UTF-8
func truncate(s string, max int) string {
	if utf8.RuneCountInString(s) <= max {
		return s
	}
	runes := []rune(s)
	return string(runes[:max-3]) + "..."
}

func calculateProgressPercent(currentDay int) int {
	return currentDay
}

func calculateProgressPadding(percent int) string {
	text := fmt.Sprintf("ПРОГРЕСС: %d%%", percent)
	requiredLength := 30
	padding := requiredLength - utf8.RuneCountInString(text)
	if padding <= 0 {
		return ""
	}
	return strings.Repeat(" ", padding)
}

func generateProgressBar(percent int) string {
	filled := percent * ProgressBarSize / TotalDays
	if filled > ProgressBarSize {
		filled = ProgressBarSize
	}
	empty := ProgressBarSize - filled
	return strings.Repeat("█", filled) + strings.Repeat("░", empty)
}

func getCommitCount() string {
	cmd := exec.Command("git", "rev-list", "--count", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		return "N/A"
	}
	return strings.TrimSpace(string(output))
}

func countLinesOfCode() int {
	totalLines := 0
	fileChan := make(chan string, 100)
	lineCounts := make(chan int, 100)
	done := make(chan struct{})

	// Горутина для обработки файлов
	go func() {
		var wg sync.WaitGroup
		for path := range fileChan {
			wg.Add(1)
			go func(p string) {
				defer wg.Done()
				content, err := os.ReadFile(p)
				if err != nil {
					return
				}
				lines := strings.Split(string(content), "\n")
				lineCounts <- len(lines)
			}(path)
		}
		wg.Wait()
		close(lineCounts)
	}()

	// Горутина для сбора результатов
	go func() {
		for count := range lineCounts {
			totalLines += count
		}
		close(done)
	}()

	// Обход файловой системы
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if info.IsDir() && shouldSkipDir(path) {
			return filepath.SkipDir
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			fileChan <- path
		}
		return nil
	})

	close(fileChan)
	<-done

	if err != nil {
		fmt.Printf("⚠️ Ошибка обхода файлов: %v\n", err)
	}
	return totalLines
}

func shouldSkipDir(path string) bool {
	skipDirs := []string{".git", "vendor", "node_modules", "dist", "bin", "tmp"}
	for _, dir := range skipDirs {
		if strings.Contains(path, dir) {
			return true
		}
	}
	return false
}

func calculateNextMilestone(currentDay int) int {
	if currentDay < 5 {
		return 5
	}
	return ((currentDay / 5) + 1) * 5
}

func formatDate(date string) string {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "25.07.2025"
	}
	return t.Format("02.01.2006")
}

func generateReadme(data ProgressData) error {
	absTemplatePath, _ := filepath.Abs(TemplateFilePath)
	fmt.Printf("⏳ Чтение шаблона из: %s\n", absTemplatePath)

	templateContent, err := os.ReadFile(TemplateFilePath)
	if err != nil {
		return fmt.Errorf("ошибка чтения шаблона: %w", err)
	}

	tmpl, err := template.New("readme").
		Funcs(template.FuncMap{
			"add": func(a, b int) int { return a + b },
		}).
		Parse(string(templateContent))
	if err != nil {
		return fmt.Errorf("ошибка парсинга шаблона: %w", err)
	}

	var output bytes.Buffer
	if err := tmpl.Execute(&output, data); err != nil {
		return fmt.Errorf("ошибка выполнения шаблона: %w", err)
	}

	if err := os.WriteFile(ReadmeFilePath, output.Bytes(), 0644); err != nil {
		return fmt.Errorf("ошибка записи README: %w", err)
	}
	return nil
}

func addNewDayIfNeeded(projectData *ProjectData) {
	today := time.Now().Format("02.01.2006")

	// Проверяем, есть ли сегодняшний день
	for _, day := range projectData.Days {
		if day.Date == today {
			return
		}
	}

	// Создаем новую запись для сегодня
	newDay := DayRecord{
		Number:     len(projectData.Days),
		Date:       today,
		Title:      "В процессе...",
		KeyInsight: "День ещё не завершён",
		LinesCode:  0,
	}

	projectData.Days = append(projectData.Days, newDay)
	fmt.Printf("➕ Добавлен новый день: Day%d (%s)\n", newDay.Number, today)
}
