package advanced

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name    string `json:"name" db:"user_name"`
    Age     int    `json:"age" db:"user_age"`
    Address string `json:"address,omitempty" db:"user_address"`
}

func DemoReflection() {
    fmt.Println("\n3. Работа с рефлексией")
    fmt.Println("---------------------")

    p := Person{Name: "Анна", Age: 30, Address: "Москва"}

    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)

    fmt.Printf("Тип: %s\n", t.Name())
    fmt.Printf("Количество полей: %d\n", t.NumField())

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)

        fmt.Printf("\nПоле %d:\n", i+1)
        fmt.Printf("  Имя: %s\n", field.Name)
        fmt.Printf("  Тип: %s\n", field.Type)
        fmt.Printf("  Значение: %v\n", value.Interface())
        fmt.Printf("  JSON тег: %s\n", field.Tag.Get("json"))
        fmt.Printf("  DB тег: %s\n", field.Tag.Get("db"))
    }

    // Изменение значений через рефлексию
    p2 := &Person{Name: "Иван", Age: 25}
    v2 := reflect.ValueOf(p2).Elem()

    if v2.CanSet() {
        nameField := v2.FieldByName("Name")
        if nameField.IsValid() && nameField.CanSet() {
            nameField.SetString("Петр")
            fmt.Printf("\nИмя изменено на: %s\n", p2.Name)
        }
    }

    // Вызов методов через рефлексию
    fmt.Println("\nМетоды типа Person:")
    for i := 0; i < t.NumMethod(); i++ {
        method := t.Method(i)
        fmt.Printf("  %s\n", method.Name)
    }
}
