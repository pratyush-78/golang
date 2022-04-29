package main

import "fmt"

// func main() {
// 	greeting := "Hello"
// 	name := "Stacey"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name) // name changed bcoz ptr is used
// }

// func sayGreeting(greetingPtr, namePtr *string) {
// 	fmt.Println(*greetingPtr, *namePtr)
// 	*namePtr = "Jacob"
// 	fmt.Println(*namePtr)
// }

//---------------------------------------------------------------------------

// func main() {
// 	s := sum("The sum is", 1, 2, 3, 4, 5)
// 	fmt.Println(s)
// }

// //using variadic parameters
// func sum(msg string, values ...int) int {
// 	fmt.Println(values)
// 	res := 0

// 	for _, v := range values {
// 		res += v
// 	}
// 	return res
// }

//--------------------------------------------------------------------------

// Another feature in GO thats pretty rare in other langs is
// to return a local var as a ptr

// How it works is : when GO runtime recons that you are returning a ptr value thats generated on local stack,
// it promotes the value to the shared memory in computer ( also called HEAP )

// func main() {
// 	s := sum("The sum is", 1, 2, 3, 4, 5)
// 	fmt.Println(*s)
// }

// //using variadic parameters
// func sum(msg string, values ...int) *int {
// 	fmt.Println(values)
// 	res := 0

// 	for _, v := range values {
// 		res += v
// 	}
// 	return &res
// }

// --------------------------------------------------------------------------

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println(s)
}

//using variadic parameters
func sum(values ...int) (result int) { // can also use this
	fmt.Println(values)
	// result = 0

	for _, v := range values {
		result += v
	}
	return
}
