package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	for j := 1; j <= n; j++ {
		for k := 1; k <= n; k++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
}
