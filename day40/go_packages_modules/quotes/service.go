package quotes

import (
	"math/rand"
	"time"
	"fmt"
)

// Quote представляет мотивационную цитату
type Quote struct {
	Text     string
	Author   string
	Category string
}

// GetRandomQuote возвращает случайную мотивационную цитату
func GetRandomQuote() Quote {
	rand.Seed(time.Now().UnixNano())
	
	quotes := []Quote{
		{
			Text:     "Программирование — это не о том, чтобы знать все ответы, а о том, чтобы знать, где их найти.",
			Author:   "Неизвестный разработчик",
			Category: "Программирование",
		},
		{
			Text:     "Лучший способ предсказать будущее — создать его.",
			Author:   "Абрахам Линкольн",
			Category: "Мотивация",
		},
		{
			Text:     "Учиться — это как грести против течения: как только перестанешь, тебя отбрасывает назад.",
			Author:   "Китайская пословица",
			Category: "Обучение",
		},
		{
			Text:     "Код — это поэзия, которая исполняется.",
			Author:   "Неизвестный разработчик",
			Category: "Программирование",
		},
		{
			Text:     "Успех — это способность идти от неудачи к неудаче, не теряя энтузиазма.",
			Author:   "Уинстон Черчилль",
			Category: "Успех",
		},
		{
			Text:     "Единственный способ делать великие дела — любить то, что ты делаешь.",
			Author:   "Стив Джобс",
			Category: "Карьера",
		},
	}
	
	return quotes[rand.Intn(len(quotes))]
}

// String возвращает форматированное представление цитаты
func (q Quote) String() string {
	return fmt.Sprintf("«%s»\n— %s\n[%s]", q.Text, q.Author, q.Category)
}

// GetQuotesByCategory возвращает все цитаты определенной категории
func GetQuotesByCategory(category string) []Quote {
	allQuotes := []Quote{
		// ... тот же список цитат, что и выше
	}
	
	var filtered []Quote
	for _, quote := range allQuotes {
		if quote.Category == category {
			filtered = append(filtered, quote)
		}
	}
	
	return filtered
}