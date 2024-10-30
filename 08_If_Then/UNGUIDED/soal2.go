package main

import "fmt"

func main() {
	var a int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&a)

	if a < 0 && a%2 == 0 {
		fmt.Print("Genap Negatif")
	} else {
		fmt.Print("Bukan")
	}
}
