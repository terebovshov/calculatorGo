// parse_roman.go
package main

import (
	"strconv"
	"strings"
)

// Функция для разбора выражения на операнды и оператор с поддержкой римских чисел.
func parseExpressionRoman(expression string) (int, string, int) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		panic("неверный формат выражения")
	}

	operand1, operator, operand2 := parts[0], parts[1], parts[2]

	var num1, num2 int
	var err error

	// Проверка, является ли первый операнд римским числом.
	if isRoman(operand1) {
		// Преобразование римского числа в арабское.
		num1 = romanToArabicNumber(operand1)
	} else {
		// Преобразование арабского числа из строки в число.
		num1, err = strconv.Atoi(operand1)
		if err != nil || num1 < 1 || num1 > 10 {
			panic("неверное число: " + operand1)
		}
	}

	// Проверка, является ли второй операнд римским числом.
	if isRoman(operand2) {
		// Преобразование римского числа в арабское.
		num2 = romanToArabicNumber(operand2)
	} else {
		// Преобразование арабского числа из строки в число.
		num2, err = strconv.Atoi(operand2)
		if err != nil || num2 < 1 || num2 > 10 {
			panic("неверное число: " + operand2)
		}
	}

	return num1, operator, num2
}
