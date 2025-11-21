package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ввод имени с проверкой 
	var name string
	for {
		fmt.Print("Введите ваше имя: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if isValidName(input) {
			name = input
			break
		} else {
			fmt.Println("Имя должно состоять только из букв. Попробуйте снова.\n")
		}
	}

	fmt.Printf("Привет, %s! Добро пожаловать в мини-викторину!\n", name)

	// Вопросы 
	questions := []string{
		"Столица Казахстана?",
		"Сколько планет в Солнечной системе?",
		"Сколько континентов на Земле?",
	}

	answers := [][]string{
		{"Астана"},
		{"8", "восемь"},
		{"7", "семь"},
	}

	score := 0

	// Основной цикл
	for i, question := range questions {
		fmt.Printf("\nВопрос %d: %s\nВаш ответ: ", i+1, question)
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		correct := false
		for _, ans := range answers[i] {
			if strings.EqualFold(userAnswer, ans) {
				correct = true
				break
			}
		}

		if correct {
			fmt.Println("Верно!")
			score++
		} else {
			fmt.Printf("Неверно. Правильный ответ: %s\n", answers[i][0])
		}
	}

	// Итог 
	fmt.Printf("\n%s, вы ответили правильно на %d из %d вопросов.\n", name, score, len(questions))
}

// Функция проверяет, что в имени ТОЛЬКО буквы и пробелы
func isValidName(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsSpace(r)) {
			return false
		}
	}
	return len(strings.TrimSpace(s)) > 0
}
