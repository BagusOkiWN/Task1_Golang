package main

import (
	"fmt"
	"os"
)

// Fungsi Penjumlahan
func penjumlahan(bilangan1, bilangan2 int) int {
	return bilangan1 + bilangan2
}

// Fungsi Pengurangan
func pengurangan(bilangan1, bilangan2 int) int {
	return bilangan1 - bilangan2
}

// Fungsi Perkalian
func perkalian(bilangan1, bilangan2 int) int {
	return bilangan1 * bilangan2
}

// Fungsi Pembagian
func pembagian(bilangan1, bilangan2 int) (int, error) {
	if bilangan2 == 0 {
		return 0, fmt.Errorf("Tidak bisa melakukan pembagian dengan nol")
	}
	return bilangan1 / bilangan2, nil
}

func main() {

	// Deklarasi Variabel
	var bilangan1 int
	var bilangan2 int

	// Input User
	fmt.Print("Masukkan nilai bilangan pertama: ")
	_, err := fmt.Scan(&bilangan1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print("Masukkan nilai bilangan kedua: ")
	_, err = fmt.Scan(&bilangan2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Invoke Fungsi Penjumlahan
	hasilJumlah := penjumlahan(bilangan1, bilangan2)
	// Invoke Fungsi Pengurangan 
	hasilKurang := pengurangan(bilangan1, bilangan2)
	// Invoke Fungsi Perkalian
	hasilKali := perkalian(bilangan1, bilangan2)

	// Mencetak Hasil Penjumlahan
	fmt.Printf("Hasil penjumlahan: %d\n", hasilJumlah)
	// Mencetak Hasil Pengurangan
	fmt.Printf("Hasil pengurangan: %d\n", hasilKurang)
	// Mencetak Hasil Perkalian
	fmt.Printf("Hasil perkalian: %d\n", hasilKali)

	// Invoke Fungsi Pembagian dan mencetaknya
	hasilBagi, err := pembagian(bilangan1, bilangan2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Hasil pembagian: %d\n", hasilBagi)
	}
}
