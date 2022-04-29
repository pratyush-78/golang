package main

import "fmt"

func main() {

	// 1.	open the resource
	// 2.	check for error
	// 3.	close the resource
	fmt.Println("Welcome to class of defers")

	// 2 rules:
	// defer puts the execution of the line at the end of func() with the initial arguments, (at the time when defer is written)
	// works like LIFO or ( stack )

	a := "start"
	defer fmt.Println(a) // will print "start"
	a = "end"

	defer fmt.Println("Last line")
	defer fmt.Println("2nd Last line")

	fmt.Println("Hello")

	myfunc()
}

// use stack to visualize defer stck in main() [Lastline, 2nd lastline]

// Hello , myfunc(), 	main defer stck

// exec :->		Hello, myfunc, 	2nd lastline, 	lastline

func myfunc() {
	fmt.Println("In this func., we loop and print  values of i:")

	for i := 0; i < 5; i++ {
		defer fmt.Printf("i is  %d\n", i)
	}

	// defer stck of myfunc : [0,1,2,3,4]
	// printline, forloop
	// so exec : line   ,   4,3,2,1,0
}
