package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHeadQuarters
	canSeeFinancials

	canSeeAsia
	canSeeEurope
	canSeeAfrica
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {

	var roles byte = isAdmin | canSeeAsia | canSeeNorthAmerica

	fmt.Println(roles)

	fmt.Printf("is Admin ? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("is canSeeEurope? %v\n", canSeeEurope&roles == canSeeEurope)
}
