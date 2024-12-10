package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&n)

	fmt.Println("Bilangan prima dari 1 hingga", n, "adalah:")
	for i := 2; i <= n; i++ {
		prime := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
