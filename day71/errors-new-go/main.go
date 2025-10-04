package main

import (
	"errors"
	"fmt"
)

func checkChicoryLevel(chicoryAmount int) error {
	if chicoryAmount <= 0 {
		return errors.New("Бутыль с водой пуста! Пахомыч останется без цикория")
	}
	return nil
}

func shareChicoryWithPakhomych() error {
	chicoryLeft := 0 // Увы, цикорий закончился!
	return checkChicoryLevel(chicoryLeft)
}

func main() {
	fmt.Println("Вечер в бытовке. Запах цикория и надежда на код...")

	if err := shareChicoryWithPakhomych(); err != nil {
		fmt.Printf("Трагедия: %v\n", err)
		fmt.Println("Пахомыч грустно потягивает воду из чайника")
	} else {
		fmt.Println("Цикорий разделен! Пахомыч доволен")
	}
}
