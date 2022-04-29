package main

import "fmt"

//---------------------------------------------------------------------------------
// func main() {
// 	greeting := "hello"
// 	name := "Stacey"
// 	sayGreeting(greeting, name)
// 	fmt.Println(name) // here the name remains Stacey
// }

// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)
// 	name = "Ted"
// 	fmt.Println(name)
// }

//---------------------------------------------------------------------------------
// func main() {
// 	greeting := "hello"
// 	name := "Stacey"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name)					// Here the name changes, as we passed the reference
// }

// func sayGreeting(greeting, name *string) {
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

//---------------------------------------------------------------------------------

// func main() {
// 	// here a func is defined as a variable

// 	// "var f func() =" 	OR	 "f :="
// 	f := func() {
// 		fmt.Println("Hello")
// 	}

// 	// calling the func
// 	f()
// }

//---------------------------------------------------------------------------------

func main() {
	var divide func(float64, float64) (float64, error)

	divide = func(a, b float64) (float64, error) {
		if b == 0 {
			return 0.0, fmt.Errorf("Cannot divide by Zero")
		}

		return (a / b), nil
	}
	d, err := divide(5, 2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}
