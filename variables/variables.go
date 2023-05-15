package main

import (
	"fmt"
	"math"
)

func main() {

	val := math.Pow(2, 16)

	fmt.Printf("The value is: %f, %T\n", val, val)

	a1 := 100
	a2 := 20

	add, sub, multiply, divide := Operations(a1, a2)

	fmt.Printf("The addition is: %d, %T\n", add, add)
	fmt.Printf("The subtraction is: %d, %T\n", sub, sub)
	fmt.Printf("The multiplication is: %d, %T\n", multiply, multiply)
	fmt.Printf("The division is: %f, %T\n", divide, divide)

}

func Operations(a1, a2 int) (int, int, int, float64) {
	return a1 + a2, a1 - a2, a1 * a2, float64(a1 / a2)
}
