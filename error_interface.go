package main

// import (
// 	"errors"
// 	"fmt"
// )

// /*
// Note:
// Di golang jika ingin mengecek return dari function apakah ada error atau tidak,
// Maka kita cek dengan menggunakan nil
// */

// func Pembagian(nilai, pembagi int) (int, error) {
// 	if pembagi == 0 {
// 		return 0, errors.New("Pembagian tidak boleh 0")
// 	} else {
// 		result := nilai / pembagi
// 		return result, nil
// 	}
// }

// func main() {
// 	hasil, err := Pembagian(100, 0)
// 	if err == nil {
// 		fmt.Println("Hasil", hasil)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}
// }
