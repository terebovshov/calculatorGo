// calculate_roman.go
package main

// Функция для выполнения операции с римскими числами.
func calculateRomanResult(num1 int, operator string, num2 int) string {
	result := 0
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		// Проверка деления нацело.
		if num2 == 0 || num1%num2 != 0 {
			panic("невозможно выполнить деление римских чисел")
		}
		result = num1 / num2
	default:
		panic("неверный оператор: " + operator)
	}

	// Проверка на отрицательный результат.
	if result <= 0 {
		panic("результат римского выражения не может быть отрицательным или нулевым")
	}

	// Преобразование арабского результата в римское число.
	return arabicToRomanNumber(result)
}
