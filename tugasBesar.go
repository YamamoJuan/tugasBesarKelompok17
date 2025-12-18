package main

import "fmt"

func main() {
	var pilihan int

	fmt.Println("===================================")
	fmt.Println(" PROGRAM PERHITUNGAN GERAK LURUS ")
	fmt.Println("===================================")
	fmt.Println("1. Hitung Kecepatan")
	fmt.Println("2. Hitung Jarak")
	fmt.Println("3. Hitung Waktu")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scan(&pilihan)

	fmt.Println("-----------------------------------")

	if pilihan != 1 && pilihan != 2 && pilihan != 3 {
		fmt.Println("KESALAHAN: Pilihan menu tidak tersedia.")
		fmt.Println("Program dihentikan.")
	} else {

		if pilihan == 1 {
			var jarak, waktu float64

			fmt.Print("Masukkan jarak (meter): ")
			fmt.Scan(&jarak)
			fmt.Print("Masukkan waktu (detik): ")
			fmt.Scan(&waktu)

			fmt.Println("-----------------------------------")

			if jarak <= 0 || waktu <= 0 {
				fmt.Println("KESALAHAN INPUT")
				fmt.Println("Jarak dan waktu harus bernilai lebih dari nol.")
			} else {
				kecepatan := jarak / waktu
				fmt.Println("HASIL PERHITUNGAN")
				fmt.Printf("Kecepatan = %.2f m/s\n", kecepatan)
			}

		} else if pilihan == 2 {
			var kecepatan, waktu float64

			fmt.Print("Masukkan kecepatan (m/s): ")
			fmt.Scan(&kecepatan)
			fmt.Print("Masukkan waktu (detik): ")
			fmt.Scan(&waktu)

			fmt.Println("-----------------------------------")

			if kecepatan <= 0 || waktu <= 0 {
				fmt.Println("KESALAHAN INPUT")
				fmt.Println("Kecepatan dan waktu harus bernilai lebih dari nol.")
			} else {
				jarak := kecepatan * waktu
				fmt.Println("HASIL PERHITUNGAN")
				fmt.Printf("Jarak = %.2f meter\n", jarak)
			}

		} else if pilihan == 3 {
			var jarak, kecepatan float64

			fmt.Print("Masukkan jarak (meter): ")
			fmt.Scan(&jarak)
			fmt.Print("Masukkan kecepatan (m/s): ")
			fmt.Scan(&kecepatan)

			fmt.Println("-----------------------------------")

			if jarak <= 0 || kecepatan <= 0 {
				fmt.Println("KESALAHAN INPUT")
				fmt.Println("Jarak dan kecepatan harus bernilai lebih dari nol.")
			} else {
				waktu := jarak / kecepatan
				fmt.Println("HASIL PERHITUNGAN")
				fmt.Printf("Waktu = %.2f detik\n", waktu)
			}
		}
	}

	fmt.Println("===================================")
	fmt.Println("Program selesai dijalankan.")
	fmt.Println("===================================")
}
