package main

import "fmt"

func main() {
	var (
		nama, nim, kelas,
		Output string
	)
	fmt.Print("Masukan nama: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukan nim: ")
	fmt.Scanln(&nim)
	fmt.Print("Masukan kelas: ")
	fmt.Scanln(&kelas)
	fmt.Print("Masukan resume: ")
	fmt.Scanln(&Output)
	fmt.Println("Identitas", nama, nim, kelas)
	fmt.Println("Resume", Output)
}