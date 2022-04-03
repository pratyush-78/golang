package main

import "fmt"

func main() {
	fmt.Println("Welcome to struct in golang")

	// no inheritance	: 	no super, no parent

	pratyush := Employee{"Pratyush", 21, 10000000.50, "IT"}

	fmt.Println(pratyush)

	fmt.Printf("\nDetails of user is : %+v\n", pratyush)
	fmt.Printf("\nDepartment of %v is %v, \n", pratyush.Name, pratyush.Dept)
}

type Employee struct {
	Name   string
	Empid  int
	Salary float64
	Dept   string
}
