package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Демонстрация go fmt
func demoGoFmt() {
	fmt.Println("\n1. Демонстрация 'go fmt':")
	fmt.Println("-------------------------")
	
	// Создаем плохо отформатированный файл
	badCode := `package main
import "fmt"
func main(){
x:=10
y:=20
fmt.Printf("Сумма: %d",x+y)
}`
	
	badFilePath := filepath.Join("examples", "formatter", "bad_format.go")
	os.MkdirAll(filepath.Dir(badFilePath), 0755)
	os.WriteFile(badFilePath, []byte(badCode), 0644)
	
	fmt.Println("Создан плохо отформатированный файл:")
	fmt.Println(badCode)
	fmt.Println("\nЗапуск 'go fmt':")
	runCommand("go", "fmt", badFilePath)
	
	// Читаем и показываем отформатированный файл
	if formatted, err := os.ReadFile(badFilePath); err == nil {
		fmt.Println("После форматирования:")
		fmt.Println(string(formatted))
	}
}

// Демонстрация go vet
func demoGoVet() {
	fmt.Println("\n2. Демонстрация 'go vet':")
	fmt.Println("-------------------------")
	
	// Создаем файл с потенциальной ошибкой
	vetCode := `package main

import "fmt"

func main() {
	// Потенциальная ошибка: форматирование без аргументов
	fmt.Printf("Сумма: %d")
	
	// Неиспользуемая переменная
	unused := 42
	
	fmt.Println("Привет, мир!")
}`
	
	vetFilePath := filepath.Join("examples", "linter", "vet_issue.go")
	os.MkdirAll(filepath.Dir(vetFilePath), 0755)
	os.WriteFile(vetFilePath, []byte(vetCode), 0644)
	
	fmt.Println("Запуск 'go vet':")
	runCommand("go", "vet", vetFilePath)
}

// Демонстрация go test
func demoGoTest() {
	fmt.Println("\n3. Демонстрация 'go test':")
	fmt.Println("--------------------------")
	
	// Создаем тестовый файл
	testCode := `package main

import "testing"

func Sum(a, b int) int {
	return a + b
}

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5
	
	if result != expected {
		t.Errorf("Ожидалось %d, получено %d", expected, result)
	}
}`
	
	testFilePath := filepath.Join("examples", "test", "math_test.go")
	os.MkdirAll(filepath.Dir(testFilePath), 0755)
	os.WriteFile(testFilePath, []byte(testCode), 0644)
	
	fmt.Println("Запуск 'go test':")
	runCommand("go", "test", "-v", filepath.Dir(testFilePath))
}

// Демонстрация go mod
func demoGoMod() {
	fmt.Println("\n4. Демонстрация 'go mod':")
	fmt.Println("-------------------------")
	
	fmt.Println("Текущий модуль:")
	runCommand("go", "mod", "edit", "-json")
	
	fmt.Println("Зависимости модуля:")
	runCommand("go", "list", "-m", "all")
}

// Демонстрация go get
func demoGoGet() {
	fmt.Println("\n5. Демонстрация 'go get':")
	fmt.Println("-------------------------")
	
	fmt.Println("Добавление популярной библиотеки:")
	runCommand("go", "get", "github.com/gorilla/mux")
	
	fmt.Println("Удаление ненужной зависимости:")
	runCommand("go", "mod", "tidy")
}

// Демонстрация go run
func demoGoRun() {
	fmt.Println("\n6. Демонстрация 'go run':")
	fmt.Println("-------------------------")
	
	runCode := `package main

import "fmt"

func main() {
	fmt.Println("Запуск с помощью 'go run'")
	fmt.Println("Эта программа не была скомпилирована заранее!")
}`
	
	runFilePath := filepath.Join("examples", "run_demo.go")
	os.WriteFile(runFilePath, []byte(runCode), 0644)
	
	fmt.Println("Запуск 'go run':")
	runCommand("go", "run", runFilePath)
	
	// Удаляем временный файл
	os.Remove(runFilePath)
}

// Демонстрация go build
func demoGoBuild() {
	fmt.Println("\n7. Демонстрация 'go build':")
	fmt.Println("---------------------------")
	
	buildCode := `package main

import "fmt"

func main() {
	fmt.Println("Эта программа была скомпилирована с помощью 'go build'")
}`
	
	buildFilePath := filepath.Join("examples", "build_demo.go")
	os.WriteFile(buildFilePath, []byte(buildCode), 0644)
	
	fmt.Println("Запуск 'go build':")
	runCommand("go", "build", "-o", "bin/build_demo", buildFilePath)
	
	fmt.Println("Запуск скомпилированной программы:")
	runCommand("./bin/build_demo")
	
	// Удаляем временный файл
	os.Remove(buildFilePath)
}