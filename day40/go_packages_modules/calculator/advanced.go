package calculator

// Power возвращает a в степени b
func Power(a, b float64) float64 {
    result := 1.0
    for i := 0; i < int(b); i++ {
        result *= a
    }
    return result
}

// IsPrime проверяет, является ли число простым
func IsPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}