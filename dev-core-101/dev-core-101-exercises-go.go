// dev-core-101-exercises-go.go
package main

import (
    "fmt"
    "math"
)

func main() {
    // 1. Конвертация температуры
    var celsius float64
    fmt.Print("Введите температуру в Цельсиях: ")
    fmt.Scan(&celsius)
    fahrenheit := (celsius*9/5) + 32
    fmt.Printf("%.2f°C = %.2f°F\n", celsius, fahrenheit)

    // 2. Чётное или нечётное
    var num int
    fmt.Print("Введите число: ")
    fmt.Scan(&num)
    if num%2 == 0 {
        fmt.Println("Число чётное")
    } else {
        fmt.Println("Число нечётное")
    }

    // 3. Проверка простого числа
    fmt.Print("Введите число для проверки на простоту: ")
    fmt.Scan(&num)
    if isPrime(num) {
        fmt.Println("Число простое")
    } else {
        fmt.Println("Число не является простым")
    }
}

// функция проверки простого числа
func isPrime(n int) bool {
    if n <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}
