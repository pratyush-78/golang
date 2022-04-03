package main

import "fmt"

func main() {

	fmt.Println("Welcome to class of pointers")

	// var ptr *int // *string , etc
	// fmt.Println(ptr)

	mynum := 23

	var ptr = &mynum // not usign :=

	fmt.Println("value of actual ptr is ", ptr)    // prints memory location
	fmt.Println("value of pointing ptr is ", *ptr) // prints the number it is referencing

	*ptr = *ptr + 2
	fmt.Println("value of pointing ptr is ", *ptr) // prints the number it is referencing
	fmt.Println("value of actual ptr is ", ptr)    // prints memory location,     doesn't change

	*ptr += 2
	fmt.Println("value of pointing ptr is ", *ptr) // prints the number it is referencing

}
