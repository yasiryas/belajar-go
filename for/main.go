package main

import "fmt"

func main() {
	// for dasar
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Bilangkan ke = ", i)
	// }

	//pengganti while
	// i := 1
	// for i <= 10 {
	// 	fmt.Println("Bilalangan = ", i)
	// 	i++
	// }

	//infinity loop
	// i:= 1
	// for {
	// 	fmt.Println("Bilangan = ", i)
	// 	i++

	// 	if i > 50 {
	// 		break
	// 	}
	// }

	//skip salah satu kondisi
	// for i := 1; i <= 20; i++ {
	// 	if i%2 != 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Bilangan ganjil = ", i)
	// }

	//dengan input
	var input int
	fmt.Print("Mau mengkalikan angka berapa? ")
	fmt.Scanln(&input)

	for i := 1; i<=input; i++{
		fmt.Println("Perkalian ", i, " x ", input, " = ", i*input)
	}

}