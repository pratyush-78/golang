package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Welcome to web requests")

	url := "http://lco.dev"

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close() // coder's resonsibility to close connection

	fmt.Printf("the response and err are of types %T and %T\n ", response, err)

	databyte, err := ioutil.ReadAll(response.Body)

	fmt.Println("The response is : ", string(databyte))
}
