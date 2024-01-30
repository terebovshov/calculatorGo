// calculate_roman.go
package main

import "fmt"

// Функция для вывода результата с поддержкой римских чисел.
func printResultRoman(result string) {
	fmt.Println("Результат:", result)
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
