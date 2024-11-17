package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	if b <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
	}

	fmt.Print("Faktor: ")
	faktorCount := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
			faktorCount++
		}
	}
	fmt.Println()

	fmt.Printf("Prima: %t\n", faktorCount == 2)
}
