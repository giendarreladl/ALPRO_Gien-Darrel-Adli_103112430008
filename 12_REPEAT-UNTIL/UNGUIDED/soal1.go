package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scanln(&angka)

	hitungDigit := 0
	temp := angka

	for {
		hitungDigit++
		temp /= 10
		if temp == 0 {
			break
		}
	}

	fmt.Println("Jumlah digit dari", angka, "adalah : ", hitungDigit)
}
