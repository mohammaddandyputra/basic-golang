package main

import "fmt"

//Named return values
func getFullName() (firstname, middlename, lastname string) {
	firstname = "Mohammad"
	middlename = "Dandy"
	lastname = "Putra"

	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a, b, c)
}
