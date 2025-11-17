// dev-core-101-quiz-go.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Создаём ввод с клавиатуры
	reader := bufio.NewReader(os.Stdin)

	// Приветствие и ввод имени пользователя
	fmt.Print("Введите ваше имя: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Printf("Привет, %s! Добро пожаловать в мини-викторину!\n", name)

	// Список вопросов и правильных ответов
	questions := []string{
		"Столица Казахстана?",
		"Сколько планет в Солнечной системе?",
		"Сколько континентов на Земле?",
	}

	// Для каждого вопроса можно указать несколько допустимых вариантов ответа
	answers := [][]string{
		{"Астана"},
		{"8", "восемь"},
		{"7", "семь"}, // допускаем цифру и слово
	}

	score := 0 // счетчик правильных ответов

	// Основной цикл викторины
	for i, question := range questions {
		fmt.Printf("\nВопрос %d: %s\nВаш ответ: ", i+1, question)
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer)

		// Проверка ответа: совпадает ли с одним из допустимых вариантов
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

	// Итоговый результат
	fmt.Printf("\n%s, вы ответили правильно на %d из %d вопросов.\n", name, score, len(questions))
}
