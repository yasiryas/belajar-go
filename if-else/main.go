package main

import "fmt"

func main() {

	fmt.Print("Masukkan umur Anda: ")
	var umur int
	fmt.Scanln(&umur)

	if umur < 18 {
		fmt.Println("Anda belum cukup umur.")
	} else {
		fmt.Println("Anda sudah cukup umur.")
	}

}