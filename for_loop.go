package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Mohammad", "Dandy", "Putra"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	//NOTE di golang itu jika ada 1 variable yang tidak dipakai maka akan error, maka dari itu
	//kode diatas diubah variable i menjadi _ yang fungsinya memberi tahu kepada golang bahwa
	//variable ini tidak digunakan
	for _, value := range slice {
		fmt.Println(value)
	}

	person := make(map[string]string)
	person["name"] = "Dandy"
	person["title"] = "Backend"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
