package book

import "fmt"

// DockerBook представляет книгу по Docker
type DockerBook struct {
	Title      string
	TotalPages int
	IsPhysical bool
}

// NewDockerBook создает новую книгу по Docker
func NewDockerBook(title string, pages int, isPhysical bool) *DockerBook {
	return &DockerBook{
		Title:      title,
		TotalPages: pages,
		IsPhysical: isPhysical,
	}
}

// Open открывает книгу
func (b *DockerBook) Open() string {
	format := "электронную"
	if b.IsPhysical {
		format = "бумажную"
	}
	
	return fmt.Sprintf("Открываю %s книгу: «%s»", format, b.Title)
}

// Info возвращает информацию о книге
func (b *DockerBook) Info() string {
	format := "электронная"
	if b.IsPhysical {
		format = "бумажная"
	}
	
	return fmt.Sprintf("%s (%d страниц, %s)", b.Title, b.TotalPages, format)
}

// CalculateReadingScore рассчитывает очки за чтение
func (b *DockerBook) CalculateReadingScore(pagesRead int) int {
	baseScore := pagesRead * 2
	
	// Бонус за бумажную книгу
	formatBonus := 0
	if b.IsPhysical {
		formatBonus = 20
	}
	
	// Бонус за тему Docker (потому что написан на Go)
	topicBonus := 25
	
	return baseScore + formatBonus + topicBonus
}
