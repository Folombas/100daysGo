package main

import "fmt"

func demoAdvancedMaps() {
	printSeparator()

	fmt.Println("🏗️ Map со структурами:")
	type Person struct {
		Name string
		Age  int
		City string
	}

	people := map[int]Person{
		1: {"Алексей 🧑‍💻", 28, "Москва"},
		2: {"Мария 👩‍🔬", 32, "Санкт-Петербург"},
		3: {"Иван 🧑‍🚀", 25, "Казань"},
	}

	for id, person := range people {
		fmt.Printf("👤 ID %d: %s, %d лет, %s\n", id, person.Name, person.Age, person.City)
	}

	fmt.Println("\n🌐 Вложенные Map:")
	university := map[string]map[string]int{
		"Физтех": {
			"студентов": 1500,
			"преподавателей": 200,
		},
		"МГУ": {
			"студентов": 40000,
			"преподавателей": 5000,
		},
	}

	for uni, stats := range university {
		fmt.Printf("🎓 %s: %d студентов, %d преподавателей\n",
			uni, stats["студентов"], stats["преподавателей"])
	}

	fmt.Println("\n🎯 Map с функциями:")
	operations := map[string]func(int, int) int{
		"➕": func(a, b int) int { return a + b },
		"➖": func(a, b int) int { return a - b },
		"✖️": func(a, b int) int { return a * b },
		"➗": func(a, b int) int { return a / b },
	}

	a, b := 10, 5
	for op, fn := range operations {
		if op == "➗" && b == 0 {
			fmt.Printf("%s Деление на ноль! Пропускаем\n", op)
			continue
		}
		fmt.Printf("%s %d %s %d = %d\n", op, a, op, b, fn(a, b))
	}
}
