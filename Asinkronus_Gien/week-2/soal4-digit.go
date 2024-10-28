/*
soal 4-digit : Sebuah program digunakan untuk menentukan tiga digit nilai yang terdapat pada suatu bilangan bulat positif x.
Masukan berupa bilangan bulat positif x yang kurang atau sama dengan 999.
Keluaran terdiri dari dari tiga bilangan d1, d2, dan d3 yang menyatakan digit pertama, kedua dan ketiga dari x.
Petunjuk: satuan dapat diperoleh apabila bilangan apapun dimodulo dengan 10


Program: Ekstraksi Digit Bilangan

Kamus:
    bilangan: integer
    Satuan: integer
    Puluhan: integer
    Ratusan: integer

Algoritma:
    input(bilangan bulat positif)

    jika bilangan < 0 maka
        tampilkan "Bilangan harus positif!"
        keluar dari program
    akhir jika

    Satuan := bilangan % 10
    bilangan := bilangan / 10
    Puluhan := bilangan % 10
    Ratusan := bilangan / 10

    output(Ratusan, Puluhan, Satuan)

EndProgram
*/

package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	Satuan := bilangan % 10
	bilangan /= 10
	Puluhan := bilangan % 10
	Ratusan := bilangan / 10

	fmt.Println(Ratusan, Puluhan, Satuan)

}
