package main

func main() {
	x := 5

	increment(&x)

	println(x)
}

func increment(x *int) {
	*x++
}
