package main

import (
	"fmt"
)

func main() {
	fmt.Println("Testing...")
	var c complex64 = 3 + 4i
	var d complex64 = 8 + 10i

	fmt.Printf("their sum is %v \n", c+d)
	fmt.Printf("their type is %T and %T\n", real(c), imag(c))

	var num complex128 = complex(10, 12)

	fmt.Printf("the value is %v and type is %T\n", num, num)

	//	fmt.Println(a / b) , error
}
