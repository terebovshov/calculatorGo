// exit.go
package main

import "strings"

func isExitCommand(input string) bool {
	return strings.ToLower(input) == "exit"
}
