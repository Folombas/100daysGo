package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type ReadmeData struct {
	StartDate        string
	CurrentDay       int
	Streak           int
	ProgressTable    string
	ProgressPercent  int
	ProgressPadding  string
	ProgressBar      string
	DaysWithoutGames int
	CommitCount      int
	LinesOfCode      int
	NextMilestone    int
	LastDay          int
	LastDayOutput    string
}

func main() {
	data := collectData()
	generateREADME(data)
}

func collectData() ReadmeData {
	// Стартовая дата (25.07.2025)
	startDate := time.Date(2025, time.July, 25, 0, 0, 0, 0, time.UTC)
	currentDate := time.Now()
	daysPassed := int(currentDate.Sub(startDate).Hours()/24) + 1

	// Собираем статистику
	return ReadmeData{
		StartDate:        startDate.Format("02.01.2006"),
		CurrentDay:       daysPassed,
		Streak:           daysPassed,
		ProgressTable:    generateProgressTable(daysPassed),
		ProgressPercent:  daysPassed,
		ProgressPadding:  strings.Repeat(" ", 26-len(strconv.Itoa(daysPassed))),
		ProgressBar:      generateProgressBar(daysPassed),
		DaysWithoutGames: daysPassed,
		CommitCount:      getCommitCount(),
		LinesOfCode:      getLinesOfCode(),
		NextMilestone:    ((daysPassed / 7) + 1) * 7,
		LastDay:          daysPassed - 1,
		LastDayOutput:    getLastDayOutput(daysPassed - 1),
	}
}

func generateProgressBar(day int) string {
	width := 50
	filled := (day * width) / 100
	empty := width - filled
	return "[" + strings.Repeat("█", filled) + strings.Repeat(".", empty) + "]"
}

func getCommitCount() int {
	cmd := exec.Command("git", "rev-list", "--count", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	count, _ := strconv.Atoi(strings.TrimSpace(string(out)))
	return count
}

func getLinesOfCode() int {
	cmd := exec.Command("bash", "-c", "find . -name '*.go' | xargs wc -l | tail -1 | awk '{print $1}'")
	out, err := cmd.Output()
	if err != nil {
		return 0
	}
	lines, _ := strconv.Atoi(strings.TrimSpace(string(out)))
	return lines
}

func getLastDayOutput(day int) string {
	if day < 0 {
		return "// День 0: старт марафона!"
	}

	// Здесь можно добавить реальный запуск программы дня
	// Для примера - заглушка
	return fmt.Sprintf("Результат дня %d:\n- Изучены новые концепции\n- Написано 50 строк кода", day)
}

func generateProgressTable(days int) string {
	var buf bytes.Buffer
	buf.WriteString("| Day | Code | Progress | Key Insight |\n")
	buf.WriteString("|-----|------|----------|-------------|\n")

	for i := 0; i <= days && i < 10; i++ {
		buf.WriteString(fmt.Sprintf(
			"| %d | [day%d](day%d/main.go) | Прогресс дня %d | Ключевое достижение %d |\n",
			i, i, i, i, i,
		))
	}

	if days >= 10 {
		buf.WriteString("| ... | ... | ... | ... |\n")
	}

	return buf.String()
}

func generateREADME(data ReadmeData) {
	tmpl, err := template.ParseFiles("utils/readme_updater/template.md")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("README.md")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = tmpl.Execute(file, data)
	if err != nil {
		panic(err)
	}
}
