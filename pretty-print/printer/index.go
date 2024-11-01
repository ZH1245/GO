package p

import (
	"fmt"
	"strings"
)

func WelcomeMessage(name string) {
	fmt.Printf("Welcome to Tech Palace, %s\n", strings.ToUpper(name))
}
func PrintStars(numberOfStars int) {
	for range numberOfStars {
		fmt.Printf("*")
	}
	fmt.Println()
}

func AddBorder(message string, numberOfStars int) {
	PrintStars(numberOfStars)
	fmt.Println(message)
	PrintStars(numberOfStars)
}
func CleanUpMessage(message string) string {
	var cleanedMessage string
	for i := range message {
		if message[i] == '*' || message == "\n" || message == "\t" {
			continue
		} else {
			cleanedMessage += string(message[i])
		}
	}
	return cleanedMessage
}
