package main

import "fmt"

func swap(a, b *int) {
    *a, *b = *b, *a
}

func demoSwap() {
    x, y := 10, 20
    fmt.Printf("\n2. Swap до:\nx=%d y=%d\n", x, y)
    
    swap(&x, &y)
    fmt.Printf("Swap после:\nx=%d y=%d\n", x, y)
}