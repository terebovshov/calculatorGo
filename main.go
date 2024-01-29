package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	// Инициализация сканера для считывания ввода с консоли.
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Вывод приглашения для ввода выражения или команды для выхода.
		fmt.Print("Введите выражение (или 'exit' для выхода): ")

		// Считывание строки ввода.
		scanner.Scan()
		input := scanner.Text()

		// Проверка на команду выхода.
		if strings.ToLower(input) == "exit" {
			break
		}

		// Вычисление результата выражения и обработка ошибок.
		result, err := calculateExpression(input)
		if err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func calculateExpression(input string) (int, error) {
	// Проверка на соответствие формату математической операции.
	match := regexp.MustCompile(`^(\d{1,2})\s*([+\-*/])\s*(\d{1,2})$`).FindStringSubmatch(input)
	if match == nil {
		panic("неверный формат выражения")
	}

	// Извлечение операндов и оператора из совпадения.
	operand1, operator, operand2 := match[1], match[2], match[3]

	// Перевод операндов в числа.
	num1, err := strconv.Atoi(operand1)
	if err != nil || num1 < 1 || num1 > 10 {
		panic("неверное число: " + operand1)
	}

	num2, err := strconv.Atoi(operand2)
	if err != nil || num2 < 1 || num2 > 10 {
		panic("неверное число: " + operand2)
	}

	// Выполнение математической операции и возврат результата.
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		// Проверка деления на ноль.
		if num2 == 0 {
			panic("деление на ноль")
		}
		return num1 / num2, nil
	default:
		panic("неверный оператор: " + operator)
	}
}
