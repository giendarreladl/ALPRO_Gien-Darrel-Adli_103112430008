package main

import (
	"fmt"
)

func main() {
	var bil int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bil)

	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(i, " ")
		}
	}
}
