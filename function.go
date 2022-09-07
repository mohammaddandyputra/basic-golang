package main

import "fmt"

//Function tanpa parameter
func sayHello() {
	fmt.Println("Hello")
}

//Function dengan parameter
func sayHelloTo(firstname string, lastname string) {
	fmt.Println("Hello", firstname, lastname)
}

//Function dengan return value
func sayHelloToPerson(firstname string) string {
	return "Hello " + firstname
}

func main() {
	sayHello()
	sayHelloTo("Dandy", "Putra")

	hello := sayHelloToPerson("dandy")
}
