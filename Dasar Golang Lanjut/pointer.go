package main

import "fmt"

func changeValue(val *int) {
	*val = 100 //Ubah nilai yang ada di alamat memori tersebu
}

func main() {
	x := 10
	fmt.Println("Default value:", x)

	changeValue(&x) //kirim alamat memori dari x ke fungsi

	fmt.Println("Value after calling changeValue():", x)
}