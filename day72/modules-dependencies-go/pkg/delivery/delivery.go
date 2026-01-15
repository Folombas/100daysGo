package delivery

import "fmt"

// Delivery представляет доставку заказа
type Delivery struct {
	From    string
	To      string
	Count   int
	Type    string // "жирный", "обычный", и т.д.
}

// NewDelivery создает новый экземпляр Delivery
func NewDelivery(from, to string, count int, deliveryType string) *Delivery {
	return &Delivery{
		From:  from,
		To:    to,
		Count: count,
		Type:  deliveryType,
	}
}

// Start начинает доставку
func (d *Delivery) Start() string {
	return fmt.Sprintf("Начало доставки: %s → %s (%d %s заказ)", 
		d.From, d.To, d.Count, d.Type)
}

// Details возвращает детали доставки
func (d *Delivery) Details() string {
	return fmt.Sprintf("Доставка из %s в %s, %d заказ(ов), тип: %s", 
		d.From, d.To, d.Count, d.Type)
}

// CalculateScore рассчитывает очки за доставку
func (d *Delivery) CalculateScore() int {
	baseScore := 50
	typeBonus := 0
	
	if d.Type == "жирный" {
		typeBonus = 30
	}
	
	distanceBonus := 20 // За дальнюю поездку
	
	return baseScore + typeBonus + distanceBonus
}
