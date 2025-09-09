package factory

import (
	"fmt"
	"sync"
)

// Conveyor представляет конвейер на фабрике
type Conveyor struct {
	products       chan string
	processed      chan string
}

// NewConveyor создает новый конвейер
func NewConveyor() *Conveyor {
	return &Conveyor{
		products:  make(chan string, 5),  // Буферизированный канал
		processed: make(chan string, 10), // Буферизированный канал
	}
}

// AddProduct добавляет продукт на конвейер
func (c *Conveyor) AddProduct(product string) {
	c.products <- product
	fmt.Printf("📦 Конвейер: принял %s\n", product)
}

// Products возвращает канал с продуктами
func (c *Conveyor) Products() <-chan string {
	return c.products
}

// SendProcessed отправляет обработанный продукт
func (c *Conveyor) SendProcessed(product string) {
	c.processed <- product
}

// Processed возвращает канал с обработанными продуктами
func (c *Conveyor) Processed() <-chan string {
	return c.processed
}

// Close закрывает каналы конвейера
func (c *Conveyor) Close() {
	close(c.products)
	close(c.processed)
}

// DemoFactory демонстрирует работу фабрики
func DemoFactory() {
	var wg sync.WaitGroup
	wg.Add(1)

	// Создаем менеджера для производства телефонов
	manager := NewManager("Телефон", 3) // 3 рабочих

	// Запускаем мониторинг в отдельной горутине
	go manager.Monitor()

	// Запускаем производство
	go manager.Start(10, &wg) // Произвести 10 телефонов

	// Ожидаем завершения
	wg.Wait()
}