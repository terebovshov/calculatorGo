// parse.go
package main

import (
	"strconv"
	"strings"
)

// Функция для разбора выражения на операнды и оператор.
func parseExpression(expression string) (int, string, int) {
	// Разделение строки на части с использованием пробелов.
	parts := strings.Fields(expression)

	// Проверка на корректное количество частей в выражении.
	if len(parts) != 3 {
		// В случае ошибки паника с сообщением.
		panic("неверный формат выражения")
	}

	// Извлечение операндов и оператора из частей выражения.
	operand1, operator, operand2 := parts[0], parts[1], parts[2]

	// Преобразование операндов в числа.
	num1, err := strconv.Atoi(operand1)
	// Проверка на ошибки преобразования и допустимый диапазон чисел.
	if err != nil || num1 < 1 || num1 > 10 {
		panic("неверное число: " + operand1)
	}

	num2, err := strconv.Atoi(operand2)
	// Проверка на ошибки преобразования и допустимый диапазон чисел.
	if err != nil || num2 < 1 || num2 > 10 {
		panic("неверное число: " + operand2)
	}

	// Возвращение операндов и оператора.
	return num1, operator, num2
}
