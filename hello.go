package main

import (
	"bufio"
	"fmt"
	"os"
)

// var Obj = "hello"
// var obe = "yo"

// var num = 10

func main() {

	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")

	// comma ok OR err ok

	input, _ := reader.ReadString('\n') // read till \n

	fmt.Println("You entered :", input)
	fmt.Printf("You entered : %T\n", input)

}
