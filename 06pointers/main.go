package main

import "fmt"

func main() {

	var a int = 42
	var b *int = &a

	fmt.Println(a, *b) // same value 42
	fmt.Println(&a, b) // same value (addr of var)
	a = 27
	fmt.Println(a, *b) // both value change

	*b = 14
	fmt.Println(a, *b) // both value change

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

	arr := [3]int{1, 2, 3}

	a0 := &arr[0]
	a1 := &arr[1] // +4

	fmt.Println(a0, a1)

	//----------------------------------------------------------------

	var ms *myStruct // default value of ms pointer is nil
	ms = &myStruct{foo: 20, foo2: 13.02}
	fmt.Println(ms) // prints with &{ , }	means ms contains the addr of an obj which has given values

	// (*ms)	means dereference ms pointer

	fo := (*ms).foo // means ms->foo	in context with c/c++, ms is a ptr pointing to a structure that has a field foo
	fo1 := ms.foo   // means ms.foo		in context with c/c++, ms has a field foo,

	// But here, we define ms as a POINTER. So it does not have a foo field, So how is it working fine ?
	// This is GO compiler helping us, bcoz it understands, so go deferences the ms ptr for us

	fmt.Println(fo, fo1)
}

type myStruct struct {
	foo  int
	foo2 float32
}
