package main

import "fmt"

func main() {
	const firstname string = "Dandy"

	// Tipe data constant tidak bisa di ubah karena nilainya tetap,
	// jika diubah seperti kode dibawah, akan muncul error yang menandakan bahwa tipe data ini constant
	// firstname = "test"

	const (
		lastname = "Putra"
		email    = "putradandy@gmail.com"
	)
	fmt.Println(lastname)
	fmt.Println(email)
}
