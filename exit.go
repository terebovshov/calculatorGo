// exit.go
package main

import "strings"

// Функция для проверки является ли строка командой выхода.
func isExitCommand(input string) bool {
	// Приведение строки к нижнему регистру и сравнение с командой выхода.
	return strings.ToLower(input) == "exit"
}
