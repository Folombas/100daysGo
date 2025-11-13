package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
	"unsafe"
)

func main() {
	fmt.Println("🚀 100daysGo: Эпическая Перезагрузка 🚀")
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Printf("👤 Айти-Студент: Гоша, 37 лет | СДВГ+ОКР+Социофобия\n")
	fmt.Printf("🎯 Миссия: Превратиться из курьера в Go-разработчика\n")
	fmt.Printf("📚 Сегодня 13 ноября 2025 года мы изучаем тему: Data Types: Numeric Types\n")
	fmt.Println()

	fmt.Println("🔢 ЧИСЛОВЫЕ ТИПЫ В GO")
	fmt.Println("====================")

	// Целые числа
	fmt.Println("🎯 ЦЕЛЫЕ ЧИСЛА:")

	// int8
	var int8Var int8 = 127
	fmt.Printf("int8: %d (диапазон: %d до %d, размер: %d байт)\n",
		int8Var, int8(math.MinInt8), int8(math.MaxInt8),
		unsafe.Sizeof(int8Var))

	// int16
	var int16Var int16 = 32767
	fmt.Printf("int16: %d (диапазон: %d до %d, размер: %d байт)\n",
		int16Var, int16(math.MinInt16), int16(math.MaxInt16),
		unsafe.Sizeof(int16Var))

	// int32
	var int32Var int32 = 2147483647
	fmt.Printf("int32: %d (диапазон: %d до %d, размер: %d байт)\n",
		int32Var, int32(math.MinInt32), int32(math.MaxInt32),
		unsafe.Sizeof(int32Var))

	// int64
	var int64Var int64 = 9223372036854775807
	fmt.Printf("int64: %d (диапазон: %d до %d, размер: %d байт)\n",
		int64Var, int64(math.MinInt64), int64(math.MaxInt64),
		unsafe.Sizeof(int64Var))

	// int (архитектурозависимый)
	var intVar int = 2147483647 // или 9223372036854775807 на 64-битной системе
	fmt.Printf("int: %d (архитектурозависимый, размер: %d байт)\n",
		intVar, unsafe.Sizeof(intVar))

	// uint8 (byte)
	var uint8Var uint8 = 255
	fmt.Printf("uint8 (byte): %d (диапазон: 0 до %d, размер: %d байт)\n",
		uint8Var, uint8(math.MaxUint8),
		unsafe.Sizeof(uint8Var))

	// uint16
	var uint16Var uint16 = 65535
	fmt.Printf("uint16: %d (диапазон: 0 до %d, размер: %d байт)\n",
		uint16Var, uint16(math.MaxUint16),
		unsafe.Sizeof(uint16Var))

	// uint32
	var uint32Var uint32 = 4294967295
	fmt.Printf("uint32: %d (диапазон: 0 до %d, размер: %d байт)\n",
		uint32Var, uint32(math.MaxUint32),
		unsafe.Sizeof(uint32Var))

	// uint64
	var uint64Var uint64 = 18446744073709551615
	fmt.Printf("uint64: %d (диапазон: 0 до %d, размер: %d байт)\n",
		uint64Var, uint64(math.MaxUint64),
		unsafe.Sizeof(uint64Var))

	// uint (архитектурозависимый)
	var uintVar uint = 4294967295 // или 18446744073709551615 на 64-битной системе
	fmt.Printf("uint: %d (архитектурозависимый, размер: %d байт)\n",
		uintVar, unsafe.Sizeof(uintVar))

	fmt.Println()

	// Числа с плавающей точкой
	fmt.Println("🎯 ЧИСЛА С ПЛАВАЮЩЕЙ ТОЧКОЙ:")

	// float32
	var float32Var float32 = 3.14159265358979323846
	fmt.Printf("float32: %.10f (точность: ~6-9 цифр, размер: %d байт)\n",
		float32Var, unsafe.Sizeof(float32Var))

	// float64
	var float64Var float64 = 3.14159265358979323846
	fmt.Printf("float64: %.15f (точность: ~15-17 цифр, размер: %d байт)\n",
		float64Var, unsafe.Sizeof(float64Var))

	fmt.Println()

	// Комплексные числа
	fmt.Println("🎯 КОМПЛЕКСНЫЕ ЧИСЛА:")

	// complex64
	var complex64Var complex64 = complex(3.0, 4.0)
	fmt.Printf("complex64: %v (размер: %d байт)\n",
		complex64Var, unsafe.Sizeof(complex64Var))

	// complex128
	var complex128Var complex128 = complex(3.141592653589793, 2.718281828459045)
	fmt.Printf("complex128: %v (размер: %d байт)\n",
		complex128Var, unsafe.Sizeof(complex128Var))

	fmt.Println()

	// Специальные типы
	fmt.Println("🎯 СПЕЦИАЛЬНЫЕ ТИПЫ:")

	// rune (alias для int32)
	var runeVar rune = 'A'
	fmt.Printf("rune: %c (alias для int32, диапазон: %d до %d, размер: %d байт)\n",
		runeVar, int32(math.MinInt32), int32(math.MaxInt32),
		unsafe.Sizeof(runeVar))

	// byte (alias для uint8)
	var byteVar byte = 'B'
	fmt.Printf("byte: %c (alias для uint8, диапазон: 0 до %d, размер: %d байт)\n",
		byteVar, uint8(math.MaxUint8),
		unsafe.Sizeof(byteVar))

	fmt.Println()

	// Демонстрация арифметических операций
	fmt.Println("🎯 АРИФМЕТИЧЕСКИЕ ОПЕРАЦИИ:")

	a := 10
	b := 3

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("Сложение: a + b = %d\n", a+b)
	fmt.Printf("Вычитание: a - b = %d\n", a-b)
	fmt.Printf("Умножение: a * b = %d\n", a*b)
	fmt.Printf("Деление: a / b = %d\n", a/b)
	fmt.Printf("Остаток: a %% b = %d\n", a%b)

	// Деление с плавающей точкой
	floatA := 10.0
	floatB := 3.0
	fmt.Printf("Деление с плавающей точкой: %.2f / %.2f = %.2f\n", floatA, floatB, floatA/floatB)

	fmt.Println()

	// Приведение типов
	fmt.Println("🎯 ПРИВЕДЕНИЕ ТИПОВ:")

	var intVal int = 42
	var floatVal float64 = float64(intVal)
	var stringVal string = fmt.Sprintf("%d", intVal)

	fmt.Printf("int %d -> float64: %.2f\n", intVal, floatVal)
	fmt.Printf("int %d -> string: %s\n", intVal, stringVal)

	// Примеры специфичных значений
	fmt.Println()
	fmt.Println("🎯 СПЕЦИФИЧНЫЕ ЗНАЧЕНИЯ:")

	fmt.Printf("Максимальное значение float64: %g\n", math.MaxFloat64)
	fmt.Printf("Минимальное положительное значение float64: %g\n", math.SmallestNonzeroFloat64)
	fmt.Printf("Бесконечность: %t\n", math.IsInf(math.Inf(1), 1))
	fmt.Printf("Не число (NaN): %t\n", math.IsNaN(math.NaN()))

	fmt.Println()

	// Информация о системе
	fmt.Printf("🎯 ИНФОРМАЦИЯ О СИСТЕМЕ:\n")
	fmt.Printf("Архитектура: %s\n", runtime.GOARCH)
	fmt.Printf("Операционная система: %s\n", runtime.GOOS)
	fmt.Printf("Количество логических процессоров: %d\n", runtime.NumCPU())

	fmt.Println()
	fmt.Println("══════════════════════════════════════════════════")
	fmt.Println("💡 Напутствие: Понимание числовых типов - основа эффективного программирования!")
	fmt.Println("   Знай, какой тип данных использовать в какой ситуации!")
	fmt.Println("   Ты не просто учишь Go - ты переписываешь свою судьбу!")
	fmt.Println()

	// Статистика по типам
	fmt.Println("🎯 СТАТИСТИКА ПО ТИПАМ:")

	types := []interface{}{
		int8(0), int16(0), int32(0), int64(0), int(0),
		uint8(0), uint16(0), uint32(0), uint64(0), uint(0),
		float32(0), float64(0), complex64(0), complex128(0),
		rune(0), byte(0),
	}

	for _, t := range types {
		fmt.Printf("Тип: %s, Размер: %d байт\n", reflect.TypeOf(t).String(), unsafe.Sizeof(t))
	}
}
