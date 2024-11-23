package main

import (
	"fmt"
)

func main() {
	var tanaman string
	fmt.Print("Input: ")
	fmt.Scan(&tanaman)

	switch tanaman {
	case "nepenthes", "drosera":
		fmt.Println("Tanaman Karnivora")
		fmt.Println("Asli Indonesia")
	case "venus", "sarracenia":
		fmt.Println("Bukan Tanaman Karnivora")
		fmt.Println("Bukan Asli Indonesia")
	default:
		fmt.Print("Tidak Termasuk Tanaman")
	}
}
