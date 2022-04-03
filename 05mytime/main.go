package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()

	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006"))                 // no other date, it has to be jan-02-2006
	fmt.Println(presentTime.Format("01-02-2006 Monday"))          // always Monday, no other day
	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05")) // always 15:04:05

	createDate := time.Date(2020, time.August, 13, 20, 40, 53, 105, time.UTC)

	fmt.Println(createDate.Format("01-02-2006 Monday 15:04:05"))
}
