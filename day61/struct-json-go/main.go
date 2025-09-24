package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"
)

// 🏷️ Структура с тегами для JSON
type User struct {
	ID        int       `json:"id" db:"user_id" validate:"required"`
	Name      string    `json:"name" db:"user_name" validate:"min=3"`
	Email     string    `json:"email,omitempty" db:"email"`
	Password  string    `json:"-" db:"password_hash"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	IsActive  bool      `json:"is_active" db:"is_active"`
	Metadata  Metadata  `json:"metadata" db:"metadata"`
}

// 📊 Дополнительная структура
type Metadata struct {
	Role      string   `json:"role" db:"role"`
	Tags      []string `json:"tags" db:"tags"`
	Settings  Settings `json:"settings" db:"settings"`
}

// ⚙️ Настройки с кастомной маршаллизацией
type Settings struct {
	Theme    string `json:"theme" db:"theme"`
	Language string `json:"language" db:"language"`
}

// 🎯 Метод для демонстрации чтения тегов
func (u User) PrintTags() {
	val := reflect.ValueOf(u)
	typ := val.Type()

	fmt.Println("🔍 АНАЛИЗ ТЕГОВ СТРУКТУРЫ:")
	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		fmt.Printf("🏷️  %s:\n", field.Name)
		fmt.Printf("   JSON: %s\n", field.Tag.Get("json"))
		fmt.Printf("   DB: %s\n", field.Tag.Get("db"))
		fmt.Printf("   Validate: %s\n", field.Tag.Get("validate"))
		fmt.Println("   ──────────────────────")
	}
}

func main() {
	fmt.Println("🎉 ДЕМО СТРУКТУРЫ, ТЕГОВ И JSON В GO!")
	fmt.Println("=========================================")

	// 📝 Создаем экземпляр пользователя
	user := User{
		ID:        1,
		Name:      "Алёша Гоферов",
		Email:     "alex@go.dev",
		Password:  "secret123",
		CreatedAt: time.Now(),
		IsActive:  true,
		Metadata: Metadata{
			Role: "developer",
			Tags: []string{"golang", "backend", "microservices"},
			Settings: Settings{
				Theme:    "dark",
				Language: "ru",
			},
		},
	}

	// 🔍 Показываем теги структуры
	user.PrintTags()
	fmt.Println()

	// 📤 Маршалинг в JSON
	fmt.Println("📤 МАРШАЛИНГ В JSON:")
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatal("❌ Ошибка маршалинга:", err)
	}
	fmt.Println(string(jsonData))
	fmt.Println()

	// 📥 Демаршалинг из JSON
	fmt.Println("📥 ДЕМАРШАЛИНГ ИЗ JSON:")
	var newUser User
	jsonStr := `{
		"id": 2,
		"name": "Мария Кодерова",
		"email": "maria@go.dev",
		"created_at": "2023-10-01T12:00:00Z",
		"is_active": true,
		"metadata": {
			"role": "team lead",
			"tags": ["leadership", "architecture"],
			"settings": {
				"theme": "light",
				"language": "en"
			}
		}
	}`

	if err := json.Unmarshal([]byte(jsonStr), &newUser); err != nil {
		log.Fatal("❌ Ошибка демаршалинга:", err)
	}

	fmt.Printf("✅ Успешно создан пользователь:\n")
	fmt.Printf("   ID: %d\n", newUser.ID)
	fmt.Printf("   Имя: %s\n", newUser.Name)
	fmt.Printf("   Email: %s\n", newUser.Email)
	fmt.Printf("   Пароль: %s\n", newUser.Password)
	fmt.Printf("   Роль: %s\n", newUser.Metadata.Role)
	fmt.Println()

	// 🎭 Демонстрация omitempty и игнорирования поля
	fmt.Println("🎭 ДЕМО OMTEMPTY И ИГНОРИРОВАНИЯ ПОЛЕЙ:")
	userWithoutEmail := User{
		ID:       3,
		Name:    "Петр Тестеров",
		Password: "test123",
		CreatedAt: time.Now(),
		IsActive: false,
	}

	jsonData2, _ := json.MarshalIndent(userWithoutEmail, "", "  ")
	fmt.Println("Пользователь без email (omitempty в действии):")
	fmt.Println(string(jsonData2))
	fmt.Println()

	// 🏆 Итоги
	fmt.Println("🎯 ВЫВОДЫ:")
	fmt.Println("✅ Теги struct управляют JSON маршалингом")
	fmt.Println("✅ json:\"name\" - переименование поля")
	fmt.Println("✅ json:\"-\" - игнорирование поля")
	fmt.Println("✅ json:\",omitempty\" - пропуск пустых значений")
	fmt.Println("✅ Можно использовать несколько тегов для разных целей")
}
