# Coverage & NoFap Challenge: От Coverage кода к Coverage жизни

## О модуле
Финальный день 100-дневного марафона! Этот модуль демонстрирует тестирование покрытия (coverage) в Go через метафору личной трансформации. Гоша начинает новый челлендж "Ноябрь-Недрочабрь", заменяя порнозависимость на изучение Go.

## Что внутри
- `main.go`: Программа-трекер челленджа с метафорой coverage жизни
- `main_test.go`: Тесты с измерением coverage кода
- Демонстрация связи coverage кода и "coverage" личного развития

## Как использовать
```bash
# Запуск программы
go run main.go

# Запуск тестов
go test

# Запуск тестов с coverage
go test -cover

# Детальный отчет coverage
go test -coverprofile=coverage.out
go tool cover -html=coverage.out

# Coverage с деталями
go test -cover -v
