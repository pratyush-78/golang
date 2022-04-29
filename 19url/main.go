package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?course=Reactjs&paymentid=jvnr123002"

func main() {
	fmt.Println("Welcome to URL parsing")
	fmt.Println(myurl)

	// parsing

	result, _ := url.Parse(myurl)

	// fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	// fmt.Println(result.Opaque)
	fmt.Println(result.Path)
	fmt.Println(result.Port())

	fmt.Println(result.RawQuery)
	// alternate way

	q_params := result.Query() // formated rawquery

	fmt.Printf("\n")
	for c, _ := range q_params {
		fmt.Println(c)
	}

}
