package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza: ")

	reader := bufio.NewReader(os.Stdin)

	inp, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", inp)

	// numPlus, err := strconv.ParseFloat(inp, 64)
	numPlus, err := strconv.ParseFloat(strings.TrimSpace(inp), 64) // to trim space/blankspace/newline

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Added 1 to rating, becomes", numPlus+1)

	}

	fmt.Println("Thanks for rating: ", numPlus+1)
	fmt.Printf("Rating type %T\n", numPlus)
}
