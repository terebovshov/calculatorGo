// main.go
package main

func main() {
	for {
		input := getUserInput()
		if isExitCommand(input) {
			break
		}
		num1, operator, num2 := parseExpression(input)
		result := calculateResult(num1, operator, num2)
		printResult(result)
	}
}
