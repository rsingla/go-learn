package main

import (
	"log"
	"strings"
	"time"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	arr := strings.Split(msg, " ")

	for _, val := range arr {
		word := ""
		for i, char := range val {
			for j := 0; j <= i; j++ {
				word += string(char)
			}
		}
		print(word)
	}
}

func main() {
	msg := "Time to enjoy the learning in programming language of Go!!!!~~~~~~~~~~~~~~~~~"
	slowDown(msg)
}
