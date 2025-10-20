package main

import (
	"fmt"
	"math"
)

type Bentuk interface {
	hitungLuas()	float64
}

type Persegi struct {
	sisi	float64
}

func (p Persegi) hitungLuas() float64 {
	return p.sisi * p.sisi
}

type Lingkaran struct {
	radius	float64
}

func (l Lingkaran) hitungLuas() float64 {
	return math.Pi * l.radius * l.radius
}

func cetakInfoLuas(b Bentuk) {
	fmt.Printf("Luas bentuk ini adalah %0.2f\n", b.hitungLuas())
}

func main() {
	persegi := Persegi{sisi: 4}
	lingkaran := Lingkaran{radius: 5}

	cetakInfoLuas(persegi)
	cetakInfoLuas(lingkaran)
}