package main

import (
	"errors"
	"fmt"
)

// Fungsi ini mengembalikan hasil pembagian (float64) dan error
func bagi(a, b float64) (float64, error) {
	if b == 0 {
		// Jika ada masalah, kembalikan nilai default dan sebuah error
		return 0, errors.New("Tidak bisa membagi dengan 0")
	}

	// Jika berhasil, kembalikan hasil dan error bernilai nil
	return  a / b, nil
}

func main() {
	hasil, err := bagi(10, 4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasilnya adalah:", hasil)
	}

	hasilBaru, errBaru := bagi(104, 0)
	if errBaru != nil {
		fmt.Println("Error:", errBaru)
	} else {
		fmt.Println("Hasilnya adalah:", hasilBaru)
	}
}