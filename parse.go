// parse.go
package main

import (
	"strconv"
	"strings"
)

func parseExpression(expression string) (int, string, int) {
	parts := strings.Fields(expression)
	if len(parts) != 3 {
		panic("неверный формат выражения")
	}
	operand1, operator, operand2 := parts[0], parts[1], parts[2]
	num1, err := strconv.Atoi(operand1)
	if err != nil || num1 < 1 || num1 > 10 {
		panic("неверное число: " + operand1)
	}
	num2, err := strconv.Atoi(operand2)
	if err != nil || num2 < 1 || num2 > 10 {
		panic("неверное число: " + operand2)
	}
	return num1, operator, num2
}
