package main

import "fmt"

func main() {

	// always mention array size

	// var fruitlist []string		----> this is wrong

	var fruitlist [6]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Mango"
	fruitlist[5] = "Peach"

	fmt.Println(fruitlist)
	fmt.Println(len(fruitlist)) // memory reserved by array

<<<<<<< HEAD
=======
	fmt.Printf("\ntype is %T ", fruitlist)

>>>>>>> 6636f1edf19918ed01db65d69ef936237c997c43
}
