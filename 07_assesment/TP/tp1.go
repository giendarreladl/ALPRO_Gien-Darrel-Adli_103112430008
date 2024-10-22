package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	fmt.Println("Hasil kuadrat dari 1 sampai", N, ":")
	for i := 1; i <= N; i++ {
		fmt.Println(i, "kuadrat =", i*i)
	}
}
