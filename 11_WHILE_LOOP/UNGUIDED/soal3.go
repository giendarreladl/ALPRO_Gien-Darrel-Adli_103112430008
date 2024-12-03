package main

import "fmt"

func main() {
	var x, y int
	fmt.Println("Masukkan dua bilangan bulat positif x dan y (pisahkan dengan spasi):")
	fmt.Scanln(&x, &y)

	if x < y || y <= 0 {
		fmt.Println("Input tidak valid. Pastikan x >= y dan y > 0.")
		return
	}

	hasil := 0

	for x >= y {
		x -= y
		hasil++
	}

	fmt.Println(hasil)
}
