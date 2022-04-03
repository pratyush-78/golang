package main

// similar as vector

import (
	"fmt"
	"sort"
)

func main() {

	// var fruitl = []string{} // you need to initialize it as well
	// either with 0 elements or some elements

	var fruitlist = []string{"Apple", "mango", "Peach"}
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "Banana", "Tomato")
	fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:])
	// fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:3]) // 3-1 = 2 elements
	fmt.Println(fruitlist)

	//--------------- making dynamic slice --------

	highscores := make([]int, 4) // type, size

	highscores[0] = 234
	highscores[1] = 23
	highscores[2] = 2334
	highscores[3] = 2323
	// highscores[4] = 77			// crash

	fmt.Println(highscores)

	// APPEND : reallocates the memory

	highscores = append(highscores, 888, 999, 111, 222) // dont crash

	fmt.Println(highscores)

	// -------- SORT ( pkg )------

	fmt.Println(sort.IntsAreSorted(highscores))

	sort.Ints(highscores)

	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	fmt.Printf("\n")

	// REMOVE values from slices based on index

	var courses = []string{"reactjs", "python", "java", "swift", "kotlin"}

	fmt.Println(courses)

	var ind int = 2
<<<<<<< HEAD
	courses = append(courses[:ind], courses[ind+1:]...)

	fmt.Println(courses)

=======

	courses = append(courses[0:ind], courses[ind+1:]...)

	fmt.Println(courses)
	fmt.Printf("\ntype is %T \n", courses)
>>>>>>> 6636f1edf19918ed01db65d69ef936237c997c43
}
