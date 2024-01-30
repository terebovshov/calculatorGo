// input.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

// Функция для получения ввода от пользователя.
func getUserInput() string {
	// Вывод приглашения для ввода выражения.
	fmt.Print("Введите выражение (или 'exit' для выхода): ")

	// Инициализация сканера для считывания ввода.
	scanner := bufio.NewScanner(os.Stdin)
	// Считывание строки ввода.
	scanner.Scan()

	// Возвращение введенной строки.
	return scanner.Text()
}
