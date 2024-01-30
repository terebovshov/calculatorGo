// roman.go
package main

// Карта для преобразования римских цифр в арабские.
var romanToArabic = map[string]int{
	"I":    1,
	"II":   2,
	"III":  3,
	"IV":   4,
	"V":    5,
	"VI":   6,
	"VII":  7,
	"VIII": 8,
	"IX":   9,
	"X":    10,
}

// Карта для преобразования арабских цифр в римские.
var arabicToRoman = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
}

// Функция для проверки, является ли введенная строка римским числом.
func isRoman(input string) bool {
	_, exists := romanToArabic[input]
	return exists
}

// Функция для преобразования римского числа в арабское.
func romanToArabicNumber(roman string) int {
	return romanToArabic[roman]
}

// Функция для преобразования арабского числа в римское.
func arabicToRomanNumber(arabic int) string {
	return arabicToRoman[arabic]
}

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
