package main

import (
	"fmt"
	"unsafe"
)

// Демонстрация управления памятью
func demoMemoryManagement() {
	fmt.Println("Управление памятью:")

	// Размеры типов (как sizeof в C)
	fmt.Println("Размеры базовых типов:")
	fmt.Printf("int: %d байт\n", unsafe.Sizeof(int(0)))
	fmt.Printf("int32: %d байт\n", unsafe.Sizeof(int32(0)))
	fmt.Printf("float64: %d байт\n", unsafe.Sizeof(float64(0)))
	fmt.Printf("bool: %d байт\n", unsafe.Sizeof(false))

	// Выравнивание структур (как в C)
	type AlignedStruct struct {
		A bool    // 1 байт + 7 байт выравнивания
		B float64 // 8 байт
		C int32   // 4 байта + 4 байта выравнивания
	}

	var s AlignedStruct
	fmt.Printf("Размер структуры: %d байт\n", unsafe.Sizeof(s))
	fmt.Printf("Выравнивание структуры: %d байт\n", unsafe.Alignof(s))

	// Работа с памятью через unsafe (аналогично C)
	fmt.Println("\nРабота с unsafe (аналогично C):")
	bytes := []byte{65, 66, 67, 0} // "ABC" + null terminator
	str := (*string)(unsafe.Pointer(&bytes))
	fmt.Printf("Строка из байт: %s\n", *str)
}
