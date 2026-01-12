package delivery

import "fmt"

type Status string

const (
	Pending   Status = "pending"
	InTransit Status = "in_transit"
	Delivered Status = "delivered"
	Cancelled Status = "cancelled"
)

type Order[T any] struct {
	ID     string
	Item   T
	Status Status
}

type DeliverySystem[T any] struct {
	orders []Order[T]
}

func NewSystem[T any]() *DeliverySystem[T] {
	return &DeliverySystem[T]{
		orders: make([]Order[T], 0),
	}
}

func (ds *DeliverySystem[T]) AddOrder(order Order[T]) {
	ds.orders = append(ds.orders, order)
}

func (ds *DeliverySystem[T]) ProcessOrders(processor func(Order[T])) {
	for _, order := range ds.orders {
		processor(order)
	}
}

// Generic function with type constraints
func ProcessItem[T any](item T) string {
	return fmt.Sprintf("Обработано: %v", item)
}

// Generic interface example
type Transport interface {
	Move() string
}

type GenericTransport[T any] struct {
	Name  string
	Speed T
}

func (gt GenericTransport[T]) Move() string {
	return fmt.Sprintf("%s движется со скоростью %v", gt.Name, gt.Speed)
}
