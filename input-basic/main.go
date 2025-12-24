package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	var age int
	var statusKawin bool

	fmt.Print("Masukkan nama lengkap anda: ")
	name, _ = bufio.NewReader(os.Stdin).ReadString('\n')

fmt.Print("Masukan umur anda: ")
fmt.Scanln(&age)

fmt.Print("Apakah anda sudah menikah (true/false):")
fmt.Scanln(&statusKawin)


fmt.Printf("Nama saya %s, umur saya %d tahun, status kawin: %t\n", name, age, statusKawin)


}