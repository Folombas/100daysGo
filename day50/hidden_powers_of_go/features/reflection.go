package features

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string `json:"name" tagExample:"value"`
    Age     int    `json:"age"`
    Address string `json:"address,omitempty"`
}

func DemoReflection() {
    fmt.Println("🔍 Работа с рефлексией в Go")
    fmt.Println("---------------------------")

    p := Person{Name: "Анна", Age: 30, Address: "Москва"}

    // Получаем тип и значение структуры
    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)

    fmt.Printf("Тип: %s\n", t.Name())
    fmt.Printf("Количество полей: %d\n", t.NumField())

    // Итерируем по полям структуры
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)

        fmt.Printf("  Поле: %s\n", field.Name)
        fmt.Printf("    Тип: %s\n", field.Type)
        fmt.Printf("    Значение: %v\n", value.Interface())
        fmt.Printf("    Тег json: %s\n", field.Tag.Get("json"))
        fmt.Printf("    Тег example: %s\n", field.Tag.Get("tagExample"))
    }

    // Изменение значений через рефлексию
    p2 := &Person{Name: "Иван", Age: 25}
    v2 := reflect.ValueOf(p2).Elem()

    if v2.CanSet() {
        nameField := v2.FieldByName("Name")
        if nameField.IsValid() && nameField.CanSet() {
            nameField.SetString("Петр")
        }
    }

    fmt.Printf("Измененное имя: %s\n\n", p2.Name)
}
