package main

import "fmt"

const (
	a = iota
	b = 1 << iota
	c
	d
	e
)

const (
	_  = iota
	kb = 1 << (10 * iota)
	mb
	gb
)

func main() {
	fmt.Println("consts enumerators")

	fmt.Printf("%d %d %d %d %d\n", a, b, c, d, e)

	fmt.Printf("%d %d %d\n", kb, mb, gb)

}
