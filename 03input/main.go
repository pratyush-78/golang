package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome user\n. Enter your name: ")

	reader := bufio.NewReader(os.Stdin) // created a obj/type for input operations from stdin
	// reader is a reference to

	input, _ := reader.ReadString('\n') // input to store stdin, read input until '\n'

	// input, err
	// _, err

	fmt.Println("You entered : ", input)

}
