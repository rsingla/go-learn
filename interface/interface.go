package main

type AnimalFunc interface {
	Eat()
	Sleep()
}

type Animal struct {
	Species     string
	Name        string
	Age         int
	IsCarnivore bool
}

func (a Animal) Eat() {
	println(a.Name + " is eating")
}

func (a Animal) Sleep() {
	println(a.Name + " is sleeping")
}

type Dog struct {
	Animal
}

func (d Dog) Bark() {
	println(d.Name + " is barking")
}

type Cat struct {
	Animal
}

func (c Cat) Meow() {
	println(c.Name + " is meowing")
}

func main() {
	dog := Dog{Animal{Species: "Canine", Name: "Fido", Age: 5, IsCarnivore: true}}
	cat := Cat{Animal{Species: "Feline", Name: "Fluffy", Age: 3, IsCarnivore: true}}

	dog.Eat()
	dog.Sleep()
	dog.Bark()

	cat.Eat()
	cat.Sleep()
	cat.Meow()
}
