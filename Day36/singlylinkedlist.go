package main

import "fmt" // Добавлен импорт fmt

// Узел списка
type Node struct {
	Data interface{}
	Next *Node
}

// Односвязный список
type SinglyLinkedList struct {
	Head *Node
	Size int
}

// Создание нового пустого списка
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{
		Head: nil,
		Size: 0,
	}
}

// Добавление элемента в начало списка
func (list *SinglyLinkedList) AddToFront(data interface{}) {
	newNode := &Node{Data: data, Next: list.Head}
	list.Head = newNode
	list.Size++
}

// Добавление элемента в конец списка
func (list *SinglyLinkedList) AddToEnd(data interface{}) {
	newNode := &Node{Data: data, Next: nil}

	if list.Head == nil {
		list.Head = newNode
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}

	list.Size++
}

// Удаление первого элемента списка
func (list *SinglyLinkedList) RemoveFromFront() (interface{}, bool) {
	if list.Head == nil {
		return nil, false
	}

	data := list.Head.Data
	list.Head = list.Head.Next
	list.Size--

	return data, true
}

// Поиск элемента в списка
func (list *SinglyLinkedList) Find(value interface{}) (*Node, bool) {
	current := list.Head

	for current != nil {
		if current.Data == value {
			return current, true
		}
		current = current.Next
	}

	return nil, false
}

// Получение размера списка
func (list *SinglyLinkedList) GetSize() int {
	return list.Size
}

// Вывод списка в виде строки
func (list *SinglyLinkedList) String() string {
	if list.Head == nil {
		return "Список пуст"
	}

	result := ""
	current := list.Head

	for current != nil {
		result += fmt.Sprintf("%v", current.Data)
		if current.Next != nil {
			result += " -> "
		}
		current = current.Next
	}

	return result
}

// Реверс списка
func (list *SinglyLinkedList) Reverse() {
	var prev, next *Node
	current := list.Head

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
}
