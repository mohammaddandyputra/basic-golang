package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Mohammad Dandy", "Dandy")) // Cek apakah ada nilai string yang sa
	fmt.Println(strings.Trim(" Dandy Putra ", " "))
	fmt.Println(strings.ToLower("DANDY"))
	fmt.Println(strings.ToUpper("dandy"))
	fmt.Println(strings.Split("Mohammad Dandy Putra", " "))
	fmt.Println(strings.ReplaceAll("Dandy Dandy Dandy", "Dandy", "Mohammad"))
}
