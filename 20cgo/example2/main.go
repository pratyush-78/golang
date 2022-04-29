package main

//#include <string.h>
//#include <time.h>
//#include <math.h>
//#include "file.c"
//#include <stdio.h>
//#include <stdlib.h>
import "C"
import "fmt"

func main() {

	// cstring
	fmt.Println("Second example of cgo")

	goString := "Hello, this is a go string" // 26 chars

	strC := C.CString(goString)

	lenC := C.strlen(strC)

	fmt.Println(int(lenC))

	//----------time

	tC := C.time(nil)

	fmt.Println("The time in seconds from Jan 1 1970 is: ", int32(tC), "seconds")

	//------------ date

	// convert char* to GO string

	dtPtrC := C.ctime(&tC)

	g := C.GoString(dtPtrC)

	fmt.Println(g)
}
