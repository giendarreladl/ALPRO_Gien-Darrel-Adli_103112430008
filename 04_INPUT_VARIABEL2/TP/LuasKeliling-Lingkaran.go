package main

import "fmt"

func main() {
	var radius float64
	var pi = 3.14159

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	luas := pi * radius * radius

	keliling := 2 * pi * radius

	fmt.Println("Luas lingkaran:", luas)
	fmt.Println("Keliling lingkaran: ", keliling)
}
