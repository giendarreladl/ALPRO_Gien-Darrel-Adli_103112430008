/*
Soal nomer 3 BMI
Sebuah program digunakan untuk menyimpan data persegi panjang yang berisi panjang, lebar, luas, dan keliling persegi panjang.
Masukan: terdiri dari dua bilangan riil yang menyatakan panjang dan lebar dari persegi panjang.
Keluaran: terdiri dari dua bilangan riil yang menyatakan luas dan keliling dari persegi panjang.
Catatan: Gunakan tipe bentukan untuk menyimpan data persegi panjang tersebut.

#Pseucode

program bilangan
kamus
  	var panjang, lebar float64

algoritma
    input(panjang)
    input(lebar)
    luas = panjang * lebar
    keliling =  2 * (panjang + lebar)
    output (panjang,lebar,luas,keliling)

endprogram
*/

package main

import "fmt"

func main() {
	var panjang, lebar float64

	fmt.Print("Masukkan panjang: ")
	fmt.Scan(&panjang)
	fmt.Print("Masukkan lebar: ")
	fmt.Scan(&lebar)

	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	fmt.Println("Informasi Persegi Panjang:")
	fmt.Printf("Panjang: %.1f\n", panjang)
	fmt.Printf("Lebar: %.1f\n", lebar)
	fmt.Printf("Luas: %.1f\n", luas)
	fmt.Printf("Keliling: %.1f\n", keliling)
}
