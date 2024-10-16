package main

import (
	"fmt"
)

func main() {
	var inp1, inp2 int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&inp1)
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&inp2)

	for i := inp1; i <= inp2; i++ {
		fmt.Print(" ", i)
	}
}
