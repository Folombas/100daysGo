package main

import (
	"fmt"
	"strings"
	"time"
)

// 🎯 Интерфейс устройства (общий контракт)
type Device interface {
	Connect() string
	Disconnect() string
	GetStatus() (string, bool)
}

// 📱 Структура смартфона
type Smartphone struct {
	Model       string
	Battery     int
	Connected   bool
	Temperature float64
}

func (s *Smartphone) Connect() string {
	if s.Battery <= 0 {
		return "🔋 Смартфон разряжен, подключение невозможно"
	}
	s.Connected = true
	return fmt.Sprintf("📱 %s подключен, заряд: %d%%", s.Model, s.Battery)
}

func (s *Smartphone) Disconnect() string {
	s.Connected = false
	return fmt.Sprintf("📱 %s отключен", s.Model)
}

func (s *Smartphone) GetStatus() (string, bool) {
	status := fmt.Sprintf("📱 %s: заряд %d%%, температура %.1f°C", 
		s.Model, s.Battery, s.Temperature)
	return status, s.Connected
}

// 💻 Структура ноутбука
type Laptop struct {
	Brand       string
	SSD         int // ГБ
	RAM         int // ГБ
	Connected   bool
	ActiveTasks int
}

func (l *Laptop) Connect() string {
	l.Connected = true
	return fmt.Sprintf("💻 %s запущен (SSD: %dГБ, RAM: %dГБ)", 
		l.Brand, l.SSD, l.RAM)
}

func (l *Laptop) Disconnect() string {
	l.Connected = false
	return fmt.Sprintf("💻 %s выключен, завершено задач: %d", 
		l.Brand, l.ActiveTasks)
}

func (l *Laptop) GetStatus() (string, bool) {
	status := fmt.Sprintf("💻 %s: SSD %dГБ, RAM %dГБ, задач: %d", 
		l.Brand, l.SSD, l.RAM, l.ActiveTasks)
	return status, l.Connected
}

// 🏭 Менеджер устройств (работает через интерфейс)
type DeviceManager struct {
	Devices []Device
}

func (dm *DeviceManager) AddDevice(d Device) {
	dm.Devices = append(dm.Devices, d)
	fmt.Println("✅ Устройство добавлено в систему")
}

func (dm *DeviceManager) ConnectAll() {
	fmt.Println("🔌 Подключаю все устройства...")
	for _, device := range dm.Devices {
		fmt.Println("  " + device.Connect())
	}
}

func (dm *DeviceManager) ShowStatus() {
	fmt.Println("\n📊 Статус устройств:")
	connectedCount := 0
	for _, device := range dm.Devices {
		status, connected := device.GetStatus()
		if connected {
			connectedCount++
			fmt.Printf("  ✅ %s\n", status)
		} else {
			fmt.Printf("  ❌ %s (отключено)\n", status)
		}
	}
	fmt.Printf("📈 Всего устройств: %d, подключено: %d\n", 
		len(dm.Devices), connectedCount)
}

// 🔧 Функция, принимающая любой Device
func TestConnection(d Device) {
	fmt.Println("\n🔧 Тестирование устройства:")
	fmt.Println("  " + d.Connect())
	status, _ := d.GetStatus()
	fmt.Println("  " + status)
	fmt.Println("  " + d.Disconnect())
}

// 🎯 Основная функция
func main() {
	fmt.Println(strings.Repeat("═", 60))
	fmt.Println("🚀 Day60: Интерфейсы в Go (100daysGo Challenge)")
	fmt.Printf("📅 Дата: %s\n", time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC).Format("2 January 2006"))
	fmt.Println("🎯 Тема: Практическое применение интерфейсов")
	fmt.Println(strings.Repeat("═", 60))

	// Создаём устройства
	iphone := &Smartphone{
		Model:       "iPhone 16 Pro",
		Battery:     78,
		Temperature: 36.7,
	}

	macbook := &Laptop{
		Brand:       "MacBook M3",
		SSD:         512,
		RAM:         16,
		ActiveTasks: 7,
	}

	// Демонстрация работы интерфейса
	fmt.Println("\n1️⃣ Тестируем отдельные устройства:")
	TestConnection(iphone)
	TestConnection(macbook)

	// Работа через менеджер
	fmt.Println("\n" + strings.Repeat("─", 60))
	fmt.Println("2️⃣ Работа через DeviceManager:")

	manager := &DeviceManager{}
	manager.AddDevice(iphone)
	manager.AddDevice(macbook)

	manager.ConnectAll()
	manager.ShowStatus()

	// Мотивационная часть
	fmt.Println("\n" + strings.Repeat("═", 60))
	fmt.Println("💡 ПОЧЕМУ ИНТЕРФЕЙСЫ В GO ТАК ВАЖНЫ?")
	fmt.Println(strings.Repeat("─", 60))
	fmt.Println("✅ Позволяют писать общий код для разных типов")
	fmt.Println("✅ Упрощают тестирование (можно использовать моки)")
	fmt.Println("✅ Реализуют полиморфизм без наследования")
	fmt.Println("✅ Используются во всей стандартной библиотеке Go")
	fmt.Println("✅ Ключевой элемент архитектуры чистых приложений")

	fmt.Println("\n🏆 День 60/100 пройден!")
	fmt.Println("🎯 Следующий шаг: Параллельный челлендж Go365 (2026 год!)")
	fmt.Println(strings.Repeat("═", 60))
}
