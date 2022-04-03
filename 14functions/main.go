package main

import "fmt"

// canNOT define func INSIDE a func
// for unknown numbr of args, use  . . . <type>
// can return multiple values

func proAdder(args ...int) int {
	ret := 0

	for _, v := range args {
		ret += v
	}
	return ret
}

func multiReturn(args ...int) (int, string) {
	ret := 0
	for _, v := range args {
		ret += v
	}

	var msg string
	if ret%2 == 0 {
		msg = "EVEN"
	} else {
		msg = "ODD"
	}

	return ret, msg

}

func main() {
	fmt.Println("Welcome to functions in Golang")
	greeter() // call the func

	sum := adder(3, 4)
	fmt.Println(sum)

	sum2 := proAdder(3, 4, 5, 6, 7)
	fmt.Println(sum2)

	sum3, message := multiReturn(1, 2, 3, 4, 5, 7)

	fmt.Printf("Sum of the above values is %v is  %v \n", sum3, message)

}

func greeter() {
	fmt.Println("Hello , How are you?")
}

func adder(i int, j int) int {
	return i + j
}
