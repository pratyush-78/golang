package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops ingolang")
	fmt.Printf("\nGolang has only FOR loop with variations\n\n")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	// for i := 0; i < len(days); i++ {			// dangerous for arrays, ok for slices
	// 	fmt.Println(days[i])
	// }

	// for d := range days {					 // ok for both arrays & slices
	// 	fmt.Println(days[d])
	// }

	// for index, element := range days {
	// 	fmt.Printf("the day at index %v is --->  %v \n", index, element)
	// }

	// for _, day := range days {
	// 	fmt.Println("the day is    ", day)
	// }

	//   WHILE loop

	num := 3

	// for num < 10 {								// while loop
	// 	fmt.Printf("the numbr is  %v \n", num)
	// 	num++
	// }

	// Break

	// for num < 10 {
	// 	if num == 6 {
	// 		fmt.Println("-----")
	// 		num++
	// 		continue
	// 	}

	// 	fmt.Println(num)
	// 	num++
	// }

	// GOTO  ==   same as c, use ( goto )

	for num < 10 {
		if num == 7 {
			goto mySpot
		}

		fmt.Println(num)
		num++
	}

mySpot:
	fmt.Println("We jumped to this spot")

}
