package main

import "fmt"

func main() {
	type (
		NoKTP   string
		Married bool
	)

	var (
		noKtpDandy    NoKTP   = "3201243294234"
		marriedStatus Married = true
	)

	fmt.Println(noKtpDandy)
	fmt.Println(marriedStatus)

}
