package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const ukuranTetesan = 0.0001

	var jumlahTetesan, daerahA, daerahB, daerahC, daerahD int
	var curahHujanA, curahHujanB, curahHujanC, curahHujanD float64

	fmt.Print("Masukkan jumlah tetesan air hujan: ")
	fmt.Scan(&jumlahTetesan)

	for i := 0; i < jumlahTetesan; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x <= 0.5 && y <= 0.5 {
			daerahA++
		} else if x > 0.5 && y <= 0.5 {
			daerahB++
		} else if x <= 0.5 && y > 0.5 {
			daerahC++
		} else {
			daerahD++
		}
	}

	curahHujanA = float64(daerahA) * ukuranTetesan
	curahHujanB = float64(daerahB) * ukuranTetesan
	curahHujanC = float64(daerahC) * ukuranTetesan
	curahHujanD = float64(daerahD) * ukuranTetesan

	fmt.Printf("Curah hujan daerah A: %.4f milimeter\n", curahHujanA)
	fmt.Printf("Curah hujan daerah B: %.4f milimeter\n", curahHujanB)
	fmt.Printf("Curah hujan daerah C: %.4f milimeter\n", curahHujanC)
	fmt.Printf("Curah hujan daerah D: %.4f milimeter\n", curahHujanD)
}