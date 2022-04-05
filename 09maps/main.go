package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang") //		Just like SLICES, the underlying data is passed by reference

	myMap := make(map[string]string)

	myMap["JS"] = "Javascript"
	myMap["GO"] = "Golang"
	myMap["PY"] = "Python"

	fmt.Println(myMap)

	delete(myMap, "GO") // map/slice, key

	fmt.Printf("\ntype is %T \n\n", myMap)

	for _, value := range myMap {
		fmt.Println("the element are: ", value)
	}

	// to check if a value truely exists in a map

	pop, ok := myMap["JS"]
	fmt.Println(pop, ok) // 	--> true

	pop2, ok := myMap["php"]
	fmt.Println(pop2, ok) // 	--> false

	//	myArrMap := make(map[[]int]string) 		// error

	//	mySliceMap := make(map[[3]int]string) 	// ok

	//-----------------------------------------------------------------IMP
	// just like SLICES, the underlying data is passed by reference

	statepopulation := map[string]int{
		"California": 100000,
		"Texas":      200000,
		"Florida":    300000,
		"NewYork":    400000,
		"Ohio":       500000,
	}

	fmt.Println(statepopulation)

	sp := statepopulation

	delete(sp, "Ohio")

	fmt.Println(sp)
	fmt.Println(statepopulation) // also gets deleted from 'statepopulation'

}
