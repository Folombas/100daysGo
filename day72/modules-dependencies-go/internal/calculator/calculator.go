package calculator

// ScoreCalculator рассчитывает и хранит очки
type ScoreCalculator struct {
	scores map[string]int
}

// NewScoreCalculator создает новый калькулятор очков
func NewScoreCalculator() *ScoreCalculator {
	return &ScoreCalculator{
		scores: make(map[string]int),
	}
}

// AddScore добавляет очки в категорию
func (c *ScoreCalculator) AddScore(category string, score int) {
	c.scores[category] += score
}

// GetCategoryScore возвращает очки по категории
func (c *ScoreCalculator) GetCategoryScore(category string) int {
	return c.scores[category]
}

// GetTotalScore возвращает общее количество очков
func (c *ScoreCalculator) GetTotalScore() int {
	total := 0
	for _, score := range c.scores {
		total += score
	}
	return total
}

// CalculateLevel определяет уровень по количеству очков
func (c *ScoreCalculator) CalculateLevel(score int) string {
	switch {
	case score >= 300:
		return "🚀 Гуру организации кода"
	case score >= 250:
		return "💪 Мастер модулей"
	case score >= 200:
		return "⭐ Эксперт зависимостей"
	case score >= 150:
		return "📚 Ученик Go"
	default:
		return "👶 Начинающий"
	}
}
