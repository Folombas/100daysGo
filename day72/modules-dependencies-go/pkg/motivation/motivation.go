package motivation

// Motivator предоставляет мотивационные фразы
type Motivator struct {
	phrases []string
	quotes  []string
}

// NewMotivator создает новый мотиватор
func NewMotivator() *Motivator {
	return &Motivator{
		phrases: []string{
			"Каждая строка кода на Go - шаг к Docker!",
			"Модули Go как вагоны поезда - каждый везёт свой груз",
			"Зависимости - это не слабость, а сила организованного кода",
			"Электричка стучит, а код на Go компилируется",
			"Бумажная книга по Docker в поезде > соцсети в метро",
			"Go выбрал Docker, Docker выбрал успех - что выберешь ты?",
			"Дальняя поездка = время для обучения, а не для скуки",
			"Один жирный заказ = много времени на чтение",
			"Изучай Go сегодня - завтра будешь понимать Docker изнутри",
			"Код организуешь как поезд: вагоны-модули, пути-зависимости",
		},
		quotes: []string{
			"Код без организации - как поезд без расписания",
			"Зависимости соединяют модули так же, как рельсы соединяют города",
			"Лучший код - тот, где каждый модуль знает своё место",
			"Изучая Go, ты изучаешь язык будущих технологий",
			"Дорога в Апрелевку длинная, но путь к знанию Go того стоит",
		},
	}
}

// GetMotivationalPhrases возвращает мотивационные фразы
func (m *Motivator) GetMotivationalPhrases(count int) []string {
	if count > len(m.phrases) {
		count = len(m.phrases)
	}
	return m.phrases[:count]
}

// GetInspirationalQuotes возвращает вдохновляющие изречения
func (m *Motivator) GetInspirationalQuotes(count int) []string {
	if count > len(m.quotes) {
		count = len(m.quotes)
	}
	return m.quotes[:count]
}

// GetExtraMotivationalPhrases возвращает дополнительные мотивационные фразы
func (m *Motivator) GetExtraMotivationalPhrases(count int) []string {
	extraPhrases := []string{
		"Не отвлекайся на видеомонтаж! Сначала Go, потом монтаж!",
		"Выходные для видео, будни для кода!",
		"Работа в офисе ждёт того, кто освоит Go!",
		"Docker на Go - твой билет в мир DevOps",
		"Организация кода = организация мыслей = организация жизни",
	}
	
	if count > len(extraPhrases) {
		count = len(extraPhrases)
	}
	return extraPhrases[:count]
}
