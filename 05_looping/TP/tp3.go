package main

import (
	"fmt"
)

func main() {
	// Mencetak bilangan genap dari 1 hingga 50
	fmt.Println("Bilangan genap dari 1 hingga 50 adalah:")
	for i := 1; i <= 50; i++ {
		// Mengecek apakah angka genap
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
