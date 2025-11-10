package main

import "fmt"

func main() {
	GreetUser()
	Calculator(2, 3)
	Goodbye()
}

func GreetUser() {
	fmt.Println("Hello, welcome to our Go app!")
}

func Calculator(a, b int) {
	fmt.Printf("The sum of %d and %d is %d\n", a, b, a+b)
}

func Goodbye() {
	fmt.Println("Goodbye! See you soon!")
}
