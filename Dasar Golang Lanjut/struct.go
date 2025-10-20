package main

import "fmt"

type Mahasiswa struct {
	NIM		string
	Nama	string
	Jurusan	string
}

func main() {
	mah1 := Mahasiswa{
		NIM: "57248390",
		Nama: "Alya",
		Jurusan: "Teknik Informatika",
	}

	fmt.Println("Nama Mahasiswa: ", mah1.Nama)
	fmt.Println("Jurusan: ", mah1.Jurusan)
}