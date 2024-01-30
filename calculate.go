// calculate.go
package main

// Функция для выполнения математической операции и получения результата.
func calculateResult(num1 int, operator string, num2 int) int {
	// Выбор операции в зависимости от оператора.
	switch operator {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
	case "/":
		// Проверка деления на ноль.
		if num2 == 0 {
			panic("деление на ноль")
		}
		return num1 / num2
	default:
		// В случае неверного оператора паника с сообщением.
		panic("неверный оператор: " + operator)
	}
}
