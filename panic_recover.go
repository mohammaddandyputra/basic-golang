package main

import "fmt"

func endApp() {
	//Recover menangkap error jika terjadi panic
	message := recover()
	if message != nil {
		fmt.Println("Error dengan pesan :", message)
	}

	fmt.Println("Aplikasi selesai")
}

func runApp(err bool) {
	defer endApp()
	if !err {
		panic("Aplikasi error")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
}
