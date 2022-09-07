package main

import "fmt"

func main() {
	var (
		panjang = 10
		lebar   = 5
		luas    = panjang * lebar
	)
	fmt.Println(luas)

	//Modulo
	modulo := luas % 4
	fmt.Println(modulo)

	i := 5
	i += 5 // i = i + 5
	fmt.Println(i)

	kondisi := true
	fmt.Println(!kondisi) //Negasi atau kebalikan dari nilai biasanya tipe boolean

}
