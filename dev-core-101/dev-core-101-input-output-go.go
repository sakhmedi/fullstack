package main

import (
    "fmt"
    "time"
)

func main() {
    // Объявляем переменные для имени и возраста
    var name string
    var age int

    // Ввод данных от пользователя
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)

    fmt.Print("Enter your age: ")
    fmt.Scanln(&age)

    // Получаем текущий год
    currentYear := time.Now().Year()

    // Рассчитываем год, когда пользователю исполнится 100 лет
    yearHundred := currentYear + (100 - age)

    // Выводим результат
    fmt.Printf("Hello, %s! You will turn 100 years old in the year %d.\n", name, yearHundred)
}
