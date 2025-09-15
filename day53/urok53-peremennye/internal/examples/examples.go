package examples

import (
	"fmt"
	"strings"
)

type Example struct {
	Title  string
	Code   string
	Output string
	Tags   []string
}

// GetExamples returns curated examples about variables in Go.
func GetExamples() []Example {
	var examples []Example

	// 1. Объявление переменных и нулевые значения
	{
		code := `package main

import "fmt"

func main() {
	var a int        // нулевое значение 0
	var s string     // нулевое значение ""
	var b bool       // нулевое значение false
	var f float64    // нулевое значение 0
	fmt.Println(a, s, b, f)
}`
		output := fmt.Sprintln(0, "", false, float64(0))
		examples = append(examples, Example{
			Title:  "Нулевые значения",
			Code:   code,
			Output: strings.TrimSpace(output),
			Tags:   []string{"var", "zero values"},
		})
	}

	// 2. Краткое объявление и вывод типов
	{
		code := `package main

import "fmt"

func main() {
	x := 42        // тип int выведен автоматически
	y := 3.14      // тип float64
	z := "Привет" // тип string (UTF-8)
	fmt.Printf("%T %T %T\n", x, y, z)
}`
		output := fmt.Sprintf("%T %T %T", 42, 3.14, "Привет")
		examples = append(examples, Example{
			Title:  "Краткое объявление := и вывод типов",
			Code:   code,
			Output: output,
			Tags:   []string{"short var", "type inference"},
		})
	}

	// 3. Несколько присваиваний и обмен значениями
	{
		code := `package main

import "fmt"

func main() {
	a, b := 1, 2
	a, b = b, a // обмен
	fmt.Println(a, b)
}`
		output := "2 1"
		examples = append(examples, Example{
			Title:  "Множественное присваивание и swap",
			Code:   code,
			Output: output,
			Tags:   []string{"multiple assignment", "swap"},
		})
	}

	// 4. Константы и iota
	{
		code := `package main

import "fmt"

const (
	Red = iota
	Green
	Blue
)

func main() {
	fmt.Println(Red, Green, Blue)
}`
		examples = append(examples, Example{
			Title:  "Константы и iota",
			Code:   code,
			Output: "0 1 2",
			Tags:   []string{"const", "iota"},
		})
	}

	// 5. Указатели: адрес и разыменование
	{
		code := `package main

import "fmt"

func main() {
	v := 10
	p := &v // указатель на v
	*p = 20 // изменяем значение по адресу
	fmt.Println(v)
}`
		output := "20"
		examples = append(examples, Example{
			Title:  "Указатели",
			Code:   code,
			Output: output,
			Tags:   []string{"pointers"},
		})
	}

	// 6. Тени переменных (shadowing)
	{
		code := `package main

import "fmt"

var x = 5

func main() {
	fmt.Println(x) // 5
	x := 10         // новая локальная переменная, скрывает глобальную
	fmt.Println(x) // 10
}`
		output := "5\n10"
		examples = append(examples, Example{
			Title:  "Теневание переменных (shadowing)",
			Code:   code,
			Output: output,
			Tags:   []string{"shadowing", "scope"},
		})
	}

	// 7. Руны и строки (кириллица)
	{
		code := `package main

import "fmt"

func main() {
	s := "Привет, мир!"
	fmt.Println(len(s))        // байты
	fmt.Println([]rune(s))     // руны
}`
		// length in bytes; compute runes slice string representation
		runes := []rune("Привет, мир!")
		output := fmt.Sprintf("%d\n%v", len("Привет, мир!"), runes)
		examples = append(examples, Example{
			Title:  "Строки и руны (UTF-8)",
			Code:   code,
			Output: output,
			Tags:   []string{"rune", "string", "utf-8"},
		})
	}

	// 8. Области видимости: блоки
	{
		code := `package main

import "fmt"

func main() {
	x := 1
	{
		y := x + 1
		fmt.Println(y)
	}
	// fmt.Println(y) // ошибка: y вне области видимости
}`
		output := "2"
		examples = append(examples, Example{
			Title:  "Области видимости",
			Code:   code,
			Output: output,
			Tags:   []string{"scope", "block"},
		})
	}

	return examples
}
