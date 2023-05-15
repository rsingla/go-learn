package main

type Car struct {
	Make  string
	Model string
	Year  int
	Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

func (c Car) Drive() string {
	return c.Make + " " + c.Model + " is driving"
}

func main() {

	car := Car{Make: "Honda", Model: "Civic", Year: 2018, Wheel: Wheel{Radius: 20, Material: "Aluminum"}}

	println(car.Drive())

	car2 := Car{Make: "Mercedes", Model: "ML", Year: 2014, Wheel: Wheel{Radius: 26, Material: "Aluminum"}}

	println(car2.Drive())
}
