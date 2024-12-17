package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan sebuah bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("bukan prima")
		return
	}

	isPrime := true
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		fmt.Println("prima")
	} else {
		fmt.Println("bukan prima")
	}
}
