package main

import (
	"fmt"
)

func salam() {
	fmt.Println("Hallo")
}

func salamDenganNama(name string){
	fmt.Println("Hallo", name)
}

func tambah(a int, b int){
	fmt.Println("Operator penjumlahan =" , a + b)
}

func jumlah(a,b int) int{
	return a + b
}

func hitung (a,b int)(int,int){
	return a+b, a*b
}

func cekUmur(umur int) bool {
	return umur >= 17
}

func perkalian(a,b int )int{
	return a*b
}

func luasPersegi(a int)int{
	return a*a
}

func  nilaiHuruf(a int) string {
	if a >= 90 {
		return "A"
	} else if a >= 80 {
		return "B"
	} else if a >= 70 {
		return "C"
	} else {
		return "D"
	}
}




func main() {
	// salam()
	// salamDenganNama("Yasir")

	// tambah(23,7)

	// hasil := jumlah(10,20)
	// fmt.Println(hasil)

		// jumlah, kali := hitung(4,5)
		// fmt.Println("jumlah =",jumlah )
		// fmt.Println("kali =",kali )

		// _, kali := hitung(2,3)
		// fmt.Println("hasil jumlah =", kali)

	// var umur int
	// fmt.Print("Berapakah umurmu :")
	// fmt.Scan(&umur)

	// if cekUmur(umur){
	// 	fmt.Print("Kamu sudah bisa masuk")
	// } else {
	// 	fmt.Print("Ups, kamu belum bisa masuk")
	// }

	// var angka int
	// fmt.Print("masukkan angka yang ingin dikalikan : ")
	// fmt.Scan(&angka)

	// for i:=1;i<=10;i++{
	// 	fmt.Println("Perkalian ke" , i, " adalah ", i, "x", angka, "=", perkalian(i, angka))
	// }

	// var sisi int
	// fmt.Println("Berapakah sisi persegi? ")
	// fmt.Scan(&sisi)

	// fmt.Println("Luas Persegi adalah : ",luasPersegi(sisi))

	var input int
	fmt.Println("Berapa nilai yang kamu miliki?")
	fmt.Scan(&input)

	grade := nilaiHuruf(input)
	fmt.Println("Kamu memiliki grade :", grade)

}