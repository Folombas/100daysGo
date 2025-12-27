package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
    "time"
)

// optimizedSieve - оптимизированное решето Эратосфена
func optimizedSieve(limit int) []int {
    if limit < 2 {
        return []int{}
    }

    isComposite := make([]bool, limit+1)
    
    for i := 2; i*i <= limit; i++ {
        if !isComposite[i] {
            for j := i * i; j <= limit; j += i {
                isComposite[j] = true
            }
        }
    }

    approxPrimes := int(float64(limit) / (1.5 * math.Log(float64(limit))))
    primes := make([]int, 0, approxPrimes)
    
    for i := 2; i <= limit; i++ {
        if !isComposite[i] {
            primes = append(primes, i)
        }
    }
    
    return primes
}

// benchmarkSieve - измеряет время выполнения
func benchmarkSieve() {
    fmt.Println("🚌 Бенчмарк (пишу в автобусе на Honor):")
    fmt.Println(strings.Repeat("─", 40))
    
    testLimits := []int{100, 1000, 10000, 50000}
    
    for _, limit := range testLimits {
        start := time.Now()
        primes := optimizedSieve(limit)
        elapsed := time.Since(start)
        
        fmt.Printf("До %6d: %5d простых | Время: %v\n", 
            limit, len(primes), elapsed)
    }
}

func main() {
    limit := 15
    
    if len(os.Args) > 1 {
        if userLimit, err := strconv.Atoi(os.Args[1]); err == nil && userLimit > 0 {
            limit = userLimit
            fmt.Printf("📱 Лимит из аргументов Termux: %d\n", limit)
        } else {
            fmt.Println("⚠️  Неверный аргумент, использую лимит по умолчанию")
        }
    } else {
        fmt.Println("🚌 Режим 'автобус': лимит по умолчанию (15)")
    }
    
    fmt.Printf("\n🔍 Ищем простые числа до %d...\n", limit)
    
    primes := optimizedSieve(limit)
    
    fmt.Printf("\n✅ Найдено %d простых чисел:\n", len(primes))
    fmt.Println(primes)
    
    if len(primes) > 0 {
        fmt.Printf("\n📊 Первое: %d, Последнее: %d\n", 
            primes[0], primes[len(primes)-1])
    }
    
    if limit >= 1000 {
        fmt.Println("\n" + strings.Repeat("═", 50))
        benchmarkSieve()
    }
    
    fmt.Println("\n🎧 Код написан в движущемся автобусе")
    fmt.Println("💾 Сохранено из Termux на Honor под Ubuntu")
    fmt.Println("🚀 Готово для git push из мобильного терминала!")
}
