package main

import "fmt"

/*

#cgo CFLAGS: -g -Wall
#include "helo.c"
#include "sum.c"
#include "helo1.c"
#include "helo2.c"
#include <time.h>
*/
import "C"

func main() {
	fmt.Println("Welcome to cgo")
	goAnother() // not executing ?

	goHello()

	// C.myPrint()
	// C.myPrint2()

	res, err := makeSum(10, 0) // changing the args refreshes/corrects the output

	if err != nil {
		panic(err)
	}

	fmt.Println("The sum is: ", res)

	printTime()
	fmt.Println("End")

}

func goHello() {
	C.hello()
}

func goAnother() {
	// _, err := C.myPrint()

	// if err != nil {
	// 	fmt.Println(err)
	// }

	C.myPrint()
}

func makeSum(a int, b int) (int, error) {

	aC := C.int(a)
	bC := C.int(b)

	resC, err := C.sum(aC, bC)

	return int(resC), err
}

func printTime() {
	C.time(C.NULL)
}
