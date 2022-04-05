package main

import "fmt"

func main() {
	fmt.Println("Inline structs or Anonymous structs")

	// short living structs

	aDoctor := struct {
		name string
		age  int
	}{name: "John Petwee"}

	// anotherDoc := aDoctor // independent copy
	anotherDoc := &aDoctor // 	reference
	anotherDoc.name = "Tom Baker"

	fmt.Println(aDoctor)
	fmt.Println(anotherDoc)
}
