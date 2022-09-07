package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr." + man.Name
	fmt.Println(man.Name)
}

func main() {
	dandy := Man{"Dandy"}
	dandy.Married()

	fmt.Println(dandy.Name)
}
