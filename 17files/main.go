package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File handling using go")

	file, err := os.Create("./file1.txt")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	file.WriteString("My first file handling in go...")

	// reading takes more effort

	file2, err := ioutil.ReadFile("file1.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file2))

}
