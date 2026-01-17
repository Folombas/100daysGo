package delivery

import "fmt"

// RouteOptimizer представляет оптимизатор маршрута доставки
type RouteOptimizer struct {
    UsedStops      []string
    UnusedStops    []string
    EfficientRoute bool
}

// NewOptimizer создает новый оптимизатор маршрута
func NewOptimizer() *RouteOptimizer {
    return &RouteOptimizer{
        UsedStops:      []string{"Химки", "Ховрино", "Смоленская", "Раменки"},
        UnusedStops:    []string{"Перово (химчистка)"}, // Отмененная точка
        EfficientRoute: false,
    }
}

// Optimize выполняет "go mod tidy" для маршрута
func (o *RouteOptimizer) Optimize() {
    fmt.Println("🗺  Оптимизация маршрута доставки...")
    fmt.Printf("Используемые остановки: %v\n", o.UsedStops)
    fmt.Printf("Удаляем неиспользуемые: %v\n", o.UnusedStops)
    
    // Очищаем неиспользуемые остановки
    o.UnusedStops = []string{}
    o.EfficientRoute = true
    
    fmt.Println("✅ Маршрут оптимизирован! Как после go mod tidy")
}
