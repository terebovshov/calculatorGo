// main.go
package main

import (
	"fmt"
)

func main() {
	for {
		// Получение ввода от пользователя.
		input := getUserInput()

		// Проверка на команду выхода.
		if isExitCommand(input) {
			break
		}

		// Проверка наличия римских чисел в выражении.
		if containsRoman(input) {
			// Проверка на использование одновременно арабских и римских чисел.
			if containsArabicAndRoman(input) {
				// Если есть и арабские, и римские числа, выдаем панику и завершаем программу.
				panic("используются одновременно арабские и римские числа")
			}

			// Если есть только римские числа, использовать функции для римских чисел.
			calculateRomanExpression(input)
		} else {
			// Иначе, использовать стандартные функции для арабских чисел.
			num1, operator, num2 := parseExpression(input)
			result := calculateResult(num1, operator, num2)
			printResult(result)
		}
	}
}

// Функция для выполнения операции с римскими числами.
func calculateRomanExpression(input string) {
	// Разбор выражения с поддержкой римских чисел.
	num1, operator, num2 := parseExpressionRoman(input)

	// Выполнение операции и получение результата.
	result := calculateRomanResult(num1, operator, num2)

	// Вывод результата с поддержкой римских чисел.
	printResultRoman(result)
}

// Функция для вывода результата с поддержкой римских чисел.
func printResultRoman(result string) {
	fmt.Println("Результат:", result)
}
