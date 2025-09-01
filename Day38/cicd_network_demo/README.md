# CI/CD и работа с сетью в Go

Этот проект демонстрирует интеграцию Continuous Integration/Continuous Deployment (CI/CD) с сетевыми возможностями Go.

## 🚀 Возможности

- HTTP сервер с эндпоинтами для мониторинга
- Нагрузочное тестирование
- Docker контейнеризация
- CI/CD пайплайн с GitHub Actions
- Бенчмарк-тесты производительности
- Деплой в Kubernetes

## 📦 Быстрый старт

```bash
# Запуск сервера
go run main.go server

# Запуск клиента
go run main.go client

# Нагрузочное тестирование
go run main.go loadtest

# Бенчмарк-тесты
go run main.go benchmark