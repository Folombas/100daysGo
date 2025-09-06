package main

import (
	"math/rand"
	"strings"
	"time"

	"gopkg.in/telebot.v3"
)

// createMainMenu создает главное меню с кнопками
func createMainMenu() *telebot.ReplyMarkup {
	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

	// Создаем ряды кнопок
	row1 := menu.Row(
		menu.Text("📚 Факт о Go"),
		menu.Text("💻 Пример кода"),
	)
	row2 := menu.Row(
		menu.Text("🔗 Полезные ссылки"),
		menu.Text("🌤️ Погода"),
	)
	row3 := menu.Row(
		menu.Text("❓ Помощь"),
		menu.Text("⚙️ Настройки"),
	)

	menu.Reply(row1, row2, row3)
	return menu
}

// createInlineKeyboard создает инлайн-кнопки
func createInlineKeyboard() *telebot.ReplyMarkup {
	inlineKeys := &telebot.ReplyMarkup{}

	// Создаем строки кнопок и сразу добавляем их в клавиатуру
	inlineKeys.Inline(
		inlineKeys.Row(
			inlineKeys.URL("Официальный сайт", "https://golang.org"),
			inlineKeys.URL("Документация", "https://pkg.go.dev"),
		),
		inlineKeys.Row(
			inlineKeys.URL("Tour of Go", "https://tour.golang.org"),
			inlineKeys.URL("Awesome Go", "https://awesome-go.com"),
		),
	)

	return inlineKeys
}

// handleStart обрабатывает команду /start
func handleStart(c telebot.Context) error {
	menu := createMainMenu()

	message := `👋 Привет! Я бот-помощник для изучения Go!

Я помогу тебе изучить язык программирования Go через интерактивные уроки, примеры кода и полезные ресурсы.

Выбери действие с помощью кнопок ниже или используй команды:`

	return c.Send(message, menu)
}

// handleHelp обрабатывает команду /help
func handleHelp(c telebot.Context) error {
	menu := createMainMenu()

	message := `📖 Справка по боту:

Я создан чтобы помогать в изучении языка программирования Go. Вот что я умею:

• Рассказывать о возможностях Go
• Показывать примеры кода
• Делиться полезными ресурсами
• Отвечать на простые вопросы

Используй кнопки для быстрого доступа к функциям!`

	return c.Send(message, menu)
}

// handleGoFact обрабатывает команду /go и кнопку "Факт о Go"
func handleGoFact(c telebot.Context) error {
	facts := []string{
		"Go был создан в Google в 2009 году Робом Пайком, Кеном Томпсоном и Робертом Гризмером.",
		"Go компилируется в один бинарный файл без зависимостей - это упрощает развертывание!",
		"Горутины - это легковесные потоки, которые позволяют легко писать конкурентный код.",
		"В Go есть сборщик мусора, но при этом язык предлагает низкоуровневый контроль над памятью.",
		"Стандартная библиотека Go очень богатая и включает HTTP-сервер, шифрование и многое другое.",
		"Go используется в Docker, Kubernetes, Terraform и других популярных проектах.",
		"Интерфейсы в Go реализуются неявно - это делает код более гибким и расширяемым.",
	}

	rand.Seed(time.Now().UnixNano())
	fact := facts[rand.Intn(len(facts))]

	menu := createMainMenu()
	return c.Send("📚 Факт о Go:\n\n"+fact, menu)
}

// handleCodeExample обрабатывает команду /code и кнопку "Пример кода"
func handleCodeExample(c telebot.Context) error {
	examples := []string{
		`// Простой HTTP-сервер на Go
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Привет, мир!")
    })
    
    fmt.Println("Сервер запущен на :8080")
    http.ListenAndServe(":8080", nil)
}`,

		`// Горутины и каналы
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Printf("Воркер %d начал задачу %d\n", id, j)
        time.Sleep(time.Second)
        fmt.Printf("Воркер %d завершил задачу %d\n", id, j)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    for j := 1; j <= 5; j++ {
        jobs <- j
    }
    close(jobs)
    
    for a := 1; a <= 5; a++ {
        <-results
    }
}`,

		`// Обработка ошибок в Go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("деление на ноль")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Ошибка:", err)
        return
    }
    fmt.Println("Результат:", result)
}`,
	}

	rand.Seed(time.Now().UnixNano())
	example := examples[rand.Intn(len(examples))]

	menu := createMainMenu()
	return c.Send("💻 Пример кода на Go:\n\n```go\n"+example+"\n```", menu, telebot.ModeMarkdownV2)
}

// handleLinks обрабатывает команду /links и кнопку "Полезные ссылки"
func handleLinks(c telebot.Context) error {
	inlineKeys := createInlineKeyboard()

	message := `🔗 Полезные ресурсы по Go:

• Официальный сайт: https://golang.org
• Документация: https://pkg.go.dev
• Tour of Go: https://tour.golang.org
• Go by Example: https://gobyexample.com
• Awesome Go: https://awesome-go.com
• Go Forum: https://forum.golangbridge.org

📚 Книги:
• "The Go Programming Language" (Donovan & Kernighan)
• "Go in Action" (Kennedy, Ketelsen, St. Martin)
• "Learning Go" (Jon Bodner)`

	return c.Send(message, inlineKeys)
}

// handleWeather обрабатывает команду /weather и кнопку "Погода"
func handleWeather(c telebot.Context) error {
	weatherConditions := []string{
		"☀️ Сегодня солнечно и тепло - отличный день для изучения Go!",
		"🌧️ На улице дождь - самое время устроиться с ноутбуком и почитать про горутины!",
		"❄️ Похолодало - согреемся горячим чаем и тёплым кодом на Go!",
		"🌤️ Легкая облачность - идеальные условия для отладки программы!",
		"🌪️ Погода переменчива, но стабильность Go неизменна!",
	}

	rand.Seed(time.Now().UnixNano())
	weather := weatherConditions[rand.Intn(len(weatherConditions))]

	menu := createMainMenu()
	return c.Send(weather, menu)
}

// handleSettings обрабатывает кнопку "Настройки"
func handleSettings(c telebot.Context) error {
	menu := createMainMenu()

	message := `⚙️ Настройки бота:

Здесь ты можешь настроить работу бота под себя. В будущем здесь появятся:

• Выбор языка интерфейса
• Настройка уведомлений
• Персональные предпочтения обучения
• Темы оформления

Следи за обновлениями! 🚀`

	return c.Send(message, menu)
}

// handleText обрабатывает текстовые сообщения и кнопки
func handleText(c telebot.Context) error {
	text := strings.ToLower(c.Text())
	menu := createMainMenu()

	// Обработка нажатий на кнопки
	switch text {
	case "📚 факт о go":
		return handleGoFact(c)
	case "💻 пример кода":
		return handleCodeExample(c)
	case "🔗 полезные ссылки":
		return handleLinks(c)
	case "🌤️ погода":
		return handleWeather(c)
	case "❓ помощь":
		return handleHelp(c)
	case "⚙️ настройки":
		return handleSettings(c)
	}

	// Обработка обычных текстовых сообщений
	responses := map[string]string{
		"привет":    "👋 Привет! Как твои успехи в изучении Go?",
		"как дела":  "🚀 Отлично! Готов помочь с изучении Go. Что хочешь узнать?",
		"спасибо":   "😊 Всегда рад помочь! Удачи в изучении Go!",
		"go":        "🐹 Go - отличный выбор! Это быстрый, простой и эффективный язык.",
		"горутина":  "🔄 Горутины - это легковесные потоки выполнения в Go. Они дешевле потоков ОС и их можно создавать тысячами!",
		"канал":     "📨 Каналы - это примитивы связи в Go, которые позволяют горутинам общаться друг с другом.",
		"интерфейс": "🔌 Интерфейсы в Go определяют поведение, а не данные. Тип реализует интерфейс неявно, просто имея нужные методы.",
	}

	for keyword, response := range responses {
		if strings.Contains(text, keyword) {
			return c.Send(response, menu)
		}
	}

	// Если не нашли подходящий ответ
	return c.Send("🤔 Не совсем понял вопрос. Используй кнопки ниже или команды:\n/start, /help, /go, /code, /links", menu)
}

// handleCallback обрабатывает нажатия на инлайн-кнопки
func handleCallback(c telebot.Context) error {
	// Обрабатываем callback данные
	data := c.Callback().Data

	switch data {
	case "next_fact":
		return handleGoFact(c)
	case "next_code":
		return handleCodeExample(c)
	default:
		return c.Respond(&telebot.CallbackResponse{
			Text:      "Неизвестная команда",
			ShowAlert: false,
		})
	}
}
