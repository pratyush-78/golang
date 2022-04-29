package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("IF-ELSE in golang")

	sp := map[string]int{
		"California": 10000,
		"Texas":      200000,
		"New York":   44400,
		"Ohio":       12334,
	}

	if pop, ok := sp["New York"]; ok {
		fmt.Printf("Yes, New York is in sp map with value %v", pop)
	}

	// scope of pop variable is limited to if block
	// fmt.Println(pop)		--> 	error
	// cannot use pop here
	num := 21

	if num < 21 {
		fmt.Println("Number is less than 21")
	} else if num > 21 {
		fmt.Println("num is greater than 21")
	} else {
		fmt.Println("Number is equal to 21")
	}

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	inp, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// err will have value if error occured, else err == nil
	// err == nil ---> NO ERROR

	if err != nil {
		fmt.Print("error\n")
	} else {
		fmt.Printf("Number entered is %v ", inp)
	}

}
