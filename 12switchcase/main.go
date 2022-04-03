package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Swtich cases in golang")

	fmt.Println("Welcomoe to dice game")

	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(6) + 1

	switch num {
	case 1:
		fmt.Println("You got one")
	case 2:
		fmt.Println("You got 2")
	case 3:
		fmt.Println("you got 3")
	case 4:
		fmt.Println("You got 4")
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	}
}
