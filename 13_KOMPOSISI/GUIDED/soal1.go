package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan ganjil dari 1 hingga", n, "adalah:")
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
	}
}
