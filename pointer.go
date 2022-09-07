package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	address2 := &address1

	address2.City = "Bandung"

	*address2 = Address{"Tegal", "Jawa Tengah", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
