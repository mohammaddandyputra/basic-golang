package main

import "fmt"

func main() {
	var (
		name1 string = "Dandy"
		name2 string = "Dandy"
	)

	//Mengecek apakah variable name1 dengan name2 memiliki nilai yang sama ?
	var result bool = name1 == name2
	fmt.Println(result)
}
