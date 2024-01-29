package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для получения ввода от пользователя.
func getUserInput() string {
	fmt.Print("Введите выражение (или 'exit' для выхода): ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

// Функция для проверки является ли строка командой выхода.
func isExitCommand(input string) bool {
	return strings.ToLower(input) == "exit"
}

// Функция для разбора выражения на операнды и оператор.
func parseExpression(expression string) (int, string, int) {
	parts := strings.Fields(expression)

	if len(parts) != 3 {
		panic("неверный формат выражения")
	}

	operand1, operator, operand2 := parts[0], parts[1], parts[2]

	num1, err := strconv.Atoi(operand1)
	if err != nil || num1 < 1 || num1 > 10 {
		panic("неверное число: " + operand1)
	}

	num2, err := strconv.Atoi(operand2)
	if err != nil || num2 < 1 || num2 > 10 {
		panic("неверное число: " + operand2)
	}

	return num1, operator, num2
}

// Функция для выполнения математической операции.
func calculateResult(num1 int, operator string, num2 int) int {
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		if num2 == 0 {
			panic("деление на ноль")
		}
		return num1 / num2
	default:
		panic("неверный оператор: " + operator)
	}
}

// Функция для вывода результата.
func printResult(result int) {
	fmt.Println("Результат:", result)
}

func main() {
	for {
		// Получение ввода от пользователя.
		input := getUserInput()

		// Проверка на команду выхода.
		if isExitCommand(input) {
			break
		}

		// Разбор выражения.
		num1, operator, num2 := parseExpression(input)

		// Выполнение операции и вывод результата.
		result := calculateResult(num1, operator, num2)
		printResult(result)
	}
}
