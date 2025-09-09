package examples

import (
	"fmt"
	"strings"
	"time"
)

// DemoPipeline демонстрирует конвейерную обработку
func DemoPipeline() {
	// Создаем этапы конвейера
	rawProducts := make(chan string, 5)
	stage1 := make(chan string, 5)
	stage2 := make(chan string, 5)
	finished := make(chan string, 5)

	// Запускаем этапы обработки
	go assembleStage(rawProducts, stage1)
	go paintStage(stage1, stage2)
	go packageStage(stage2, finished)

	// Отправляем сырье на конвейер
	products := []string{"Деталь A", "Деталь B", "Деталь C", "Деталь D"}
	for _, product := range products {
		fmt.Printf("📦 Отправка на конвейер: %s\n", product)
		rawProducts <- product
		time.Sleep(200 * time.Millisecond)
	}

	// Закрываем каналы
	close(rawProducts)

	// Получаем готовую продукцию
	for i := 0; i < len(products); i++ {
		result := <-finished
		fmt.Printf("🎁 Готовая продукция: %s\n", result)
	}
}

// assembleStage - этап сборки
func assembleStage(in <-chan string, out chan<- string) {
	for product := range in {
		processed := fmt.Sprintf("Собранный %s", product)
		fmt.Printf("🔧 Этап сборки: %s → %s\n", product, processed)
		out <- processed
	}
	close(out)
}

// paintStage - этап покраски
func paintStage(in <-chan string, out chan<- string) {
	for product := range in {
		processed := fmt.Sprintf("Покрашенный %s", strings.ToLower(product))
		fmt.Printf("🎨 Этап покраски: %s → %s\n", product, processed)
		out <- processed
	}
	close(out)
}

// packageStage - этап упаковки
func packageStage(in <-chan string, out chan<- string) {
	for product := range in {
		processed := fmt.Sprintf("Упакованный %s", strings.ToLower(product))
		fmt.Printf("📦 Этап упаковки: %s → %s\n", product, processed)
		out <- processed
	}
	close(out)
}