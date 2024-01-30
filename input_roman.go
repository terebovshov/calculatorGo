// input_roman.go
package main

import (
	"strconv"
	"strings"
)

// Функция для проверки наличия одновременно арабских и римских чисел в выражении.
func containsArabicAndRoman(input string) bool {
	// Разбиваем выражение на части.
	parts := strings.Fields(input)

	// Инициализируем флаги для обнаружения арабских и римских чисел.
	arabicFound := false
	romanFound := false

	// Проверяем каждую часть выражения.
	for _, part := range parts {
		// Если часть - арабское число.
		if _, err := strconv.Atoi(part); err == nil {
			arabicFound = true
		}
		// Если часть - римское число.
		if isRoman(part) {
			romanFound = true
		}
	}

	// Возвращаем результат проверки.
	return arabicFound && romanFound
}

// Функция для проверки наличия римских чисел в выражении.
func containsRoman(input string) bool {
	// Разбиваем выражение на части.
	parts := strings.Fields(input)
	// Проверяем, является ли хотя бы одна часть римским числом.
	for _, part := range parts {
		if isRoman(part) {
			return true
		}
	}
	return false
}
