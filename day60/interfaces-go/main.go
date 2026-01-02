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
	Brand       string
	Model       string
	Battery     int
	Connected   bool
	Temperature float64
	OS          string
}

func (s *Smartphone) Connect() string {
	if s.Battery <= 0 {
		return fmt.Sprintf("🔋 %s %s разряжен, подключение невозможно", s.Brand, s.Model)
	}
	s.Connected = true
	return fmt.Sprintf("📱 %s %s (%s) подключен, заряд: %d%%",
		s.Brand, s.Model, s.OS, s.Battery)
}

func (s *Smartphone) Disconnect() string {
	s.Connected = false
	return fmt.Sprintf("📱 %s %s отключен", s.Brand, s.Model)
}

func (s *Smartphone) GetStatus() (string, bool) {
	status := fmt.Sprintf("%s %s: заряд %d%%, %s, темпер. %.1f°C",
		s.Brand, s.Model, s.Battery, s.OS, s.Temperature)
	return status, s.Connected
}

// 💻 Структура ноутбука
type Laptop struct {
	Brand       string
	Model       string
	SSD         int // ГБ
	RAM         int // ГБ
	Processor   string
	Connected   bool
	ActiveTasks int
}

func (l *Laptop) Connect() string {
	l.Connected = true
	return fmt.Sprintf("💻 %s %s запущен (%s, SSD: %dГБ, RAM: %dГБ)",
		l.Brand, l.Model, l.Processor, l.SSD, l.RAM)
}

func (l *Laptop) Disconnect() string {
	l.Connected = false
	return fmt.Sprintf("💻 %s %s выключен, завершено задач: %d",
		l.Brand, l.Model, l.ActiveTasks)
}

func (l *Laptop) GetStatus() (string, bool) {
	status := fmt.Sprintf("%s %s: %s, SSD %dГБ, RAM %dГБ, задач: %d",
		l.Brand, l.Model, l.Processor, l.SSD, l.RAM, l.ActiveTasks)
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

func (dm *DeviceManager) DisconnectAll() {
	fmt.Println("\n🔌 Отключаю все устройства...")
	for _, device := range dm.Devices {
		fmt.Println("  " + device.Disconnect())
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
	fmt.Println(strings.Repeat("═", 70))
	fmt.Println("🚀 Day60: Интерфейсы в Go (100daysGo Challenge)")
	fmt.Printf("📅 Дата: %s\n", time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC).Format("2 January 2006"))
	fmt.Println("🎯 Тема: Практическое применение интерфейсов")
	fmt.Println("📱 Тестируем 3 смартфона и 3 ноутбука разных брендов")
	fmt.Println(strings.Repeat("═", 70))

	// Создаём 3 разных смартфона
	iphone := &Smartphone{
		Brand:       "Apple",
		Model:       "iPhone 16 Pro Max",
		Battery:     85,
		Temperature: 36.5,
		OS:          "iOS 20",
	}

	samsung := &Smartphone{
		Brand:       "Samsung",
		Model:       "Galaxy S25 Ultra",
		Battery:     92,
		Temperature: 38.2,
		OS:          "Android 15",
	}

	pixel := &Smartphone{
		Brand:       "Google",
		Model:       "Pixel 9 Pro",
		Battery:     78,
		Temperature: 37.1,
		OS:          "Android 15",
	}

	// Создаём 3 разных ноутбука
	macbook := &Laptop{
		Brand:       "Apple",
		Model:       "MacBook Pro M3",
		SSD:         1024,
		RAM:         32,
		Processor:   "Apple M3 Pro",
		ActiveTasks: 12,
	}

	dell := &Laptop{
		Brand:       "Dell",
		Model:       "XPS 15",
		SSD:         512,
		RAM:         16,
		Processor:   "Intel Core i9-13900H",
		ActiveTasks: 8,
	}

	lenovo := &Laptop{
		Brand:       "Lenovo",
		Model:       "ThinkPad X1 Carbon",
		SSD:         1024,
		RAM:         32,
		Processor:   "Intel Core i7-1365U",
		ActiveTasks: 15,
	}

	// Демонстрация работы интерфейса для каждого устройства
	fmt.Println("\n1️⃣ Тестируем отдельные устройства:")

	fmt.Println("\n--- Смартфоны ---")
	TestConnection(iphone)
	TestConnection(samsung)
	TestConnection(pixel)

	fmt.Println("\n--- Ноутбуки ---")
	TestConnection(macbook)
	TestConnection(dell)
	TestConnection(lenovo)

	// Работа через менеджер
	fmt.Println("\n" + strings.Repeat("─", 70))
	fmt.Println("2️⃣ Работа через DeviceManager (все 6 устройств):")

	manager := &DeviceManager{}

	// Добавляем смартфоны
	manager.AddDevice(iphone)
	manager.AddDevice(samsung)
	manager.AddDevice(pixel)

	// Добавляем ноутбуки
	manager.AddDevice(macbook)
	manager.AddDevice(dell)
	manager.AddDevice(lenovo)

	manager.ConnectAll()
	manager.ShowStatus()
	manager.DisconnectAll()

	// Мотивационная часть
	fmt.Println("\n" + strings.Repeat("═", 70))
	fmt.Println("💡 ПОЧЕМУ ИНТЕРФЕЙСЫ В GO ТАК ВАЖНЫ?")
	fmt.Println(strings.Repeat("─", 70))
	fmt.Println("✅ Позволяют писать общий код для разных типов (6 устройств, 1 интерфейс)")
	fmt.Println("✅ Упрощают тестирование (можно использовать моки)")
	fmt.Println("✅ Реализуют полиморфизм без наследования")
	fmt.Println("✅ Используются во всей стандартной библиотеке Go")
	fmt.Println("✅ Ключевой элемент архитектуры чистых приложений")
	fmt.Println("✅ Легко добавлять новые типы устройств без изменения существующего кода")

	// Статистика
	fmt.Println("\n" + strings.Repeat("─", 70))
	fmt.Println("📊 СТАТИСТИКА ТЕСТИРОВАНИЯ:")
	fmt.Printf("   📱 Смартфонов: %d (Apple, Samsung, Google)\n", 3)
	fmt.Printf("   💻 Ноутбуков: %d (Apple, Dell, Lenovo)\n", 3)
	fmt.Printf("   🚀 Всего устройств: %d\n", 6)
	fmt.Printf("   🔧 Универсальных методов: %d (Connect, Disconnect, GetStatus)\n", 3)

	fmt.Println("\n🏆 День 60/100 пройден!")
	fmt.Println("🎯 Параллельный челлендж Go365 (2 января 2026, день 2/365)")
	fmt.Println("🚀 Изучено интерфейсов: 1 (Device), типов: 2 (Smartphone, Laptop)")
	fmt.Println(strings.Repeat("═", 70))
}
