package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in Golang")

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
}
