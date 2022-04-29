package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Swtich cases in golang")

	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("none of the above")
	}

	//-------

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("none of the above")
	}

	var i interface{} = 1 // 1.0  ,  "1",  [2]int{},  [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("i is an int")
	case float64:
		fmt.Println("i is a float64")
	case string:
		fmt.Println("i is a string")
	case [2]int:
		fmt.Println("i is a [2]int")
	default:
		fmt.Println("i is another type")
	}

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
		fallthrough
	case 5:
		fmt.Println("You got 5")
	case 6:
		fmt.Println("You got 6")
	}
}
