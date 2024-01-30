// calculate.go
package main

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
