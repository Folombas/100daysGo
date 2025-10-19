package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

// DairyPipeline представляет конвейер производства молочной продукции
type DairyPipeline struct {
	mu sync.Mutex
}

type MilkProduct struct {
	base       string
	processed  string
	packaged   string
	delivered  string
	stageTimes []time.Duration
}

func main() {
	fmt.Println("🏡 СМЕШАННАЯ ВСЕЛЕННАЯ: Простоквашино + Смешарики + Барбоскины!")
	fmt.Println("==============================================================")

	pipeline := &DairyPipeline{}

	fmt.Println("🥛 МОЛОЧНЫЙ КОНВЕЙЕР: Создаем производственную линию!")
	fmt.Println("🔄 ПАТТЕРН: Pipeline - последовательная обработка через этапы!")

	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Запускаем конвейер
	fmt.Println("\n🎯 ЗАПУСК PIPELINE КОНВЕЙЕРА...")
	fmt.Println("📦 ЭТАПЫ: Молоко → Обработка → Упаковка → Доставка")

	// Создаем каналы для каждого этапа конвейера
	rawMilkCh := pipeline.generateRawMilk(ctx)
	processedCh := pipeline.processMilk(ctx, rawMilkCh)
	packagedCh := pipeline.packageProducts(ctx, processedCh)
	deliveredCh := pipeline.deliverProducts(ctx, packagedCh)

	// Собираем финальные результаты
	pipeline.collectResults(deliveredCh)

	fmt.Println("\n🎉 КОНВЕЙЕР ЗАВЕРШЕН! Все продукты доставлены!")
	fmt.Println("🏆 Простоквашино, Смешарики и Барбоскины - идеальная команда!")
}

// generateRawMilk - этап 1: Добыча сырого молока (Простоквашино)
func (d *DairyPipeline) generateRawMilk(ctx context.Context) <-chan MilkProduct {
	out := make(chan MilkProduct)

	go func() {
		defer close(out)

		sources := []string{
			"Молоко от коровы Мурки",
			"Молоко от бычка Бурёнки",
			"Молоко от козы Рогатки",
			"Молоко от овечки Белянки",
		}

		for i, source := range sources {
			select {
			case <-ctx.Done():
				return
			default:
				// Имитация времени добычи
				time.Sleep(1 * time.Second)

				product := MilkProduct{
					base:       source,
					stageTimes: []time.Duration{1 * time.Second},
				}

				fmt.Printf("   🐄 [ПРОСТОКВАШИНО] %s добыто (%d/4)\n", source, i+1)
				out <- product
			}
		}
	}()

	return out
}

// processMilk - этап 2: Обработка молока (Смешарики)
func (d *DairyPipeline) processMilk(ctx context.Context, in <-chan MilkProduct) <-chan MilkProduct {
	out := make(chan MilkProduct)

	go func() {
		defer close(out)

		processors := []string{"Крош", "Ёжик", "Бараш", "Нюша"}
		processIndex := 0

		for product := range in {
			select {
			case <-ctx.Done():
				return
			default:
				// Имитация времени обработки
				time.Sleep(2 * time.Second)

				processor := processors[processIndex%len(processors)]
				processIndex++

				product.processed = fmt.Sprintf("Обработано %s", processor)
				product.stageTimes = append(product.stageTimes, 2*time.Second)

				fmt.Printf("   🦔 [СМЕШАРИКИ] %s: %s → %s\n",
					processor, product.base, strings.Split(product.base, " ")[2])
				out <- product
			}
		}
	}()

	return out
}

// packageProducts - этап 3: Упаковка продуктов (Барбоскины)
func (d *DairyPipeline) packageProducts(ctx context.Context, in <-chan MilkProduct) <-chan MilkProduct {
	out := make(chan MilkProduct)

	go func() {
		defer close(out)

		packagers := []string{"Гена", "Лиза", "Роза", "Дружок"}
		packages := []string{"в бутылки", "в пакеты", "в банки", "в тетрапаки"}
		packageIndex := 0

		for product := range in {
			select {
			case <-ctx.Done():
				return
			default:
				// Имитация времени упаковки
				time.Sleep(1 * time.Second)

				packager := packagers[packageIndex%len(packagers)]
				packageType := packages[packageIndex%len(packages)]
				packageIndex++

				product.packaged = fmt.Sprintf("Упаковано %s %s", packager, packageType)
				product.stageTimes = append(product.stageTimes, 1*time.Second)

				fmt.Printf("   🏠 [БАРБОСКИНЫ] %s: %s\n", packager, packageType)
				out <- product
			}
		}
	}()

	return out
}

// deliverProducts - этап 4: Доставка продуктов (совместные усилия)
func (d *DairyPipeline) deliverProducts(ctx context.Context, in <-chan MilkProduct) <-chan MilkProduct {
	out := make(chan MilkProduct)

	go func() {
		defer close(out)

		deliverers := []string{"Почтальон Печкин", "Кар-Карыч", "Совунья", "Пин"}
		deliverIndex := 0

		for product := range in {
			select {
			case <-ctx.Done():
				return
			default:
				// Имитация времени доставки
				time.Sleep(1 * time.Second)

				deliverer := deliverers[deliverIndex%len(deliverers)]
				deliverIndex++

				product.delivered = fmt.Sprintf("Доставлено %s", deliverer)
				product.stageTimes = append(product.stageTimes, 1*time.Second)

				fmt.Printf("   🚚 [ДОСТАВКА] %s: продукт доставлен потребителю\n", deliverer)
				out <- product
			}
		}
	}()

	return out
}

// collectResults - сбор и отображение результатов
func (d *DairyPipeline) collectResults(results <-chan MilkProduct) {
	fmt.Println("\n📊 ОТЧЁТ О РАБОТЕ КОНВЕЙЕРА:")
	fmt.Println("============================")

	totalProducts := 0
	totalTime := time.Duration(0)

	for product := range results {
		totalProducts++

		productTime := time.Duration(0)
		for _, t := range product.stageTimes {
			productTime += t
		}
		totalTime += productTime

		fmt.Printf("\n   📦 ПРОДУКТ #%d:\n", totalProducts)
		fmt.Printf("      🥛 %s\n", product.base)
		fmt.Printf("      🦔 %s\n", product.processed)
		fmt.Printf("      🏠 %s\n", product.packaged)
		fmt.Printf("      🚚 %s\n", product.delivered)
		fmt.Printf("      ⏱️  Общее время: %v\n", productTime)
	}

	fmt.Printf("\n📈 ИТОГО: Обработано %d продуктов за общее время %v\n",
		totalProducts, totalTime)
	fmt.Printf("📊 СРЕДНЕЕ: %v на продукт\n", totalTime/time.Duration(totalProducts))
}
