// input.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func getUserInput() string {
	fmt.Print("Введите выражение (или 'exit' для выхода): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
