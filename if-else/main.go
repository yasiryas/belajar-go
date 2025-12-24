package main

import "fmt"

func ifElse1() {

	fmt.Print("Masukkan umur Anda: ")
	var umur int
	fmt.Scanln(&umur)

	if umur < 18 {
		fmt.Println("Anda belum cukup umur.")
	} else {
		fmt.Println("Anda sudah cukup umur.")
	}

}

func ifElse2() {
	fmt.Print("Berapa nilai kamu? ")
	var nilai int
	fmt.Scanln(&nilai)

	if nilai >=90{
		fmt.Println("Nilai Kamu A")
	}else if nilai >=80{
		fmt.Println("Nilai Kamu B")
	}else if nilai >= 70{
		fmt.Println("Nilai Kamu C")
	}else if nilai<70 {
		fmt.Println("Nilai Kamu D")
	}

}

func main(){
	ifElse1()
	ifElse2()
}