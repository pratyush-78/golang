package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Panic in Go")

	// to deal with exceptions that can crash an application
	// panics happen AFTER defer statements are executed

	// 1.	exec the func
	// 2.	exec the defers
	// 3.	then handle panics
	// 4.	then handle return value

	// recover from a panic 3:57:00

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go"))
	})

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		panic(err)
	}
}
