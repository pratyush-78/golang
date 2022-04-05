package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   // 			embedding, bird has a animal
	SpeedKPH float32
	CanFly   bool
}

func main() {
	fmt.Println("Alternative of Inheritance")

	fmt.Println("** Embedding through COMPOSITION **")

	fmt.Println("has a relationship")

	// in Normal Inheritance , its "is a" relationship
	// child is a parent
	// eg:		Bird is a Animal

	// But in Composition, its "has a" relationship
	// eg:		Bird has a Animal

	b := Bird{}

	b.Name = "Eagle"
	b.Origin = "Australia"
	b.CanFly = true
	b.SpeedKPH = 80.0

	fmt.Println(b)
}
