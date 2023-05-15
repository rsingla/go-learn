package main

import (
	"fmt"
)

func main() {
	height := 10

	if height > 5 {
		fmt.Println("You are tall")
	} else {
		fmt.Println("You are short")
	}

	switch height {
	case 10:
		fmt.Println("You are tall")
	case 5:
		fmt.Println("You are short")
	default:
		fmt.Println("You are average")
	}

	email := "dunesh@apicode.io"

	if length := getLength(email); length > 20 {
		fmt.Println("That's a long email")
	} else {
		fmt.Println("That's a short email")
	}

}

func getLength(email string) int {
	return len(email)
}
