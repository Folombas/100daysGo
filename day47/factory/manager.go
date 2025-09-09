package factory

import (
	"fmt"
	"sync"
	"time"
)

// Manager управляет работой фабрики
type Manager struct {
	workers     []*Worker
	conveyor    *Conveyor
	productType string
}

// NewManager создает нового менеджера
func NewManager(productType string, numWorkers int) *Manager {
	conveyor := NewConveyor()
	workers := make([]*Worker, numWorkers)

	for i := range workers {
		workers[i] = NewWorker(i+1, conveyor)
	}

	return &Manager{
		workers:     workers,
		conveyor:    conveyor,
		productType: productType,
	}
}

// Start запускает производство
func (m *Manager) Start(products int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("🏭 Менеджер: Запускаем производство %d единиц %s\n", 
		products, m.productType)

	// Запускаем рабочих
	for _, worker := range m.workers {
		go worker.Work()
	}

	// Отправляем продукты на конвейер
	for i := 1; i <= products; i++ {
		product := fmt.Sprintf("%s №%d", m.productType, i)
		m.conveyor.AddProduct(product)
		time.Sleep(100 * time.Millisecond) // Имитация времени производства
	}

	// Закрываем конвейер после завершения
	m.conveyor.Close()
	fmt.Println("🏭 Менеджер: Все продукты отправлены на конвейер")
}

// Monitor отслеживает прогресс производства
func (m *Manager) Monitor() {
	for {
		select {
		case product, ok := <-m.conveyor.Processed():
			if !ok {
				fmt.Println("🏭 Менеджер: Конвейер завершил работу")
				return
			}
			fmt.Printf("🏭 Менеджер: Получен обработанный %s\n", product)
		}
	}
}