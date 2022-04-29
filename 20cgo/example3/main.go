package main

/*

void doSomething(int *p) {
	*p = 123;
}

*/
import "C"
import "fmt"

func main() {
	fmt.Println("Welcome to calling go func from C")

	var i C.int = 23
	C.doSomething(&i)
	fmt.Println(i)
}
