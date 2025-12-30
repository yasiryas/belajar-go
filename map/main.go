package main

import (
	"fmt"
)

func firstMap() {
	var person map[string]string
	person = make(map[string]string)
	person["name"] = "Yasyes"
	person["age"] = "26"

	println("Name:", person["name"])
	println("Age:", person["age"])
}

func secondMap() {
	data := map[string]string{
		"brand": "Toyota",
		"model": "Avanza",
		"year":  "2009",
	}

	fmt.Println("Brand :", data["brand"])
	fmt.Println("Model :", data["model"])
	fmt.Println("Year :", data["year"])
}

func threeMap(){
	// umur, nilai := map[string]string{
	// 	"Nama":"Yadi",
	// 	"Nilai":"90",
	// }

	// fmt.Println("Nama :", nilai["Nama"])
	// fmt.Println("Nilai :", nilai["Nilai"])

	data := map[string]string{
		"Nama": "Andy",
	}

	umur := make(map[string]int)
	umur["Andy"] = 21
	umur["Dedy"] = 34
	umur["Vika"] = 44


	//

	for umur, value := range data{
		fmt.Println(umur, "=", value)
	}
}
func main() {
	// firstMap()
	// secondMap()
	threeMap()
}