// main.go
package main

func main() {
	// Бесконечный цикл для повторения ввода и обработки выражений.
	for {
		// Получение ввода от пользователя.
		input := getUserInput()

		// Проверка на команду выхода.
		if isExitCommand(input) {
			break
		}

		// Разбор выражения на операнды и оператор.
		num1, operator, num2 := parseExpression(input)

		// Выполнение математической операции и получение результата.
		result := calculateResult(num1, operator, num2)

		// Вывод результата.
		printResult(result)
	}
}
