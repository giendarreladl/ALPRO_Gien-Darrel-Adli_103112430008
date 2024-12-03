package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	for {
		fmt.Print("Masukkan kata: ")
		fmt.Scan(&input)

		if strings.ToLower(input) == "telkom" {
			fmt.Println("Program selesai.")
			break
		} else {
			fmt.Println("Anda mengetik:", input)
		}
	}
}
