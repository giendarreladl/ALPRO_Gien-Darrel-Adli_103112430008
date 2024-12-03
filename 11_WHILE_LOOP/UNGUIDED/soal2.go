package main

import (
	"fmt"
)

func main() {
	var angka int
	fmt.Println("Masukkan bilangan bulat positif:")
	fmt.Scanln(&angka)

	fmt.Println("Digit dari belakang ke depan:")
	for angka > 0 {
		digit := angka % 10
		fmt.Println(digit)
		angka /= 10
	}
}
