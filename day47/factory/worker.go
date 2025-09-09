package factory

import (
	"fmt"
	"math/rand"
	"time"
)

// Worker представляет рабочего на фабрике
type Worker struct {
	id      int
	conveyor *Conveyor
}

// NewWorker создает нового рабочего
func NewWorker(id int, conveyor *Conveyor) *Worker {
	return &Worker{
		id:      id,
		conveyor: conveyor,
	}
}

// Work запускает процесс работы
func (w *Worker) Work() {
	for product := range w.conveyor.Products() {
		w.processProduct(product)
	}
}

// processProduct обрабатывает продукт
func (w *Worker) processProduct(product string) {
	// Имитация времени обработки
	processingTime := time.Duration(rand.Intn(300)+100) * time.Millisecond
	time.Sleep(processingTime)

	// Обработка продукта
	processedProduct := fmt.Sprintf("%s (обработан рабочим %d)", product, w.id)
	
	// Отправка обработанного продукта
	w.conveyor.SendProcessed(processedProduct)
	
	fmt.Printf("👷 Рабочий %d: обработал %s за %v\n", 
		w.id, product, processingTime)
}