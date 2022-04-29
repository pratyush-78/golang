package main

import "fmt"

type greeter struct {
	greeting string
	name     string
}

// func (g greeter) greet() {
// 	fmt.Println("hello", g.greeting, g.name)
// 	g.name = "Bob"
// 	fmt.Println(g.name)
// }

func (g *greeter) greet() {
	fmt.Println("hello", g.greeting, g.name) // implicit dereferencing ptr working for us,
	// so we don't have to dereference it ourselves
	g.name = "Bob"
	fmt.Println(g.name)
}

func main() {
	fmt.Println("Welcome to methods in go")

	// g := greeter{"Good Morning,", "Alex"}
	// OR

	g := greeter{
		greeting: "Good Morning,",
		name:     "Alex",
	}

	g.greet()
	fmt.Println("The new name, after change, is ", g.name)
}
