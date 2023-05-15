package main

import (
	"github.com/rsingla/go-learn/singleton"
)

func main() {
	idService := singleton.NewIdService()
	println(idService.Next())
	println(idService.Next())
	println(idService.Next())
}
