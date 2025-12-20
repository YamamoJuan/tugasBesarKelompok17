//Nama Kelompok : Kelompok 17
//Anggota Kelompok : Fathuridza Akmal Hafidz Priyambodo, Yamamo Juan Alterico Situmorang
//NIM Kelompok : 103032500048, 103032530028
//Kelas : IT-49-03

package main

import (
	"fmt"
	"time"
)

//convertion
func kmKeMeter(km float64) float64 {
	return km * 1000
}

func meterKeKm(m float64) float64 {
	return m / 1000
}

func jamKeDetik(jam float64) float64 {
	return jam * 3600
}

func detikKeJam(detik float64) float64 {
	return detik / 3600
}

func msKeKmh(ms float64) float64 {
	return ms * 3.6
}

func kmhKeMs(kmh float64) float64 {
	return kmh / 3.6
}

//main
func main() {

	//delayed text
	fmt.Println("===================================")
	fmt.Println("Nama Kelompok  : Kelompok 17")
	time.Sleep(700 * time.Millisecond)

	fmt.Println("Anggota        :")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("- Fathuridza Akmal Hafidz Priyambodo")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("- Yamamo Juan Alterico Situmorang")
	time.Sleep(700 * time.Millisecond)

	fmt.Println("Kelas          : IT-49-03")
	time.Sleep(700 * time.Millisecond)
	fmt.Println("===================================")
	time.Sleep(1 * time.Second)

	//menu program
	var pilihan int

	fmt.Println(" PROGRAM PERHITUNGAN GERAK LURUS ")
	fmt.Println(" (km & km/h)")
	fmt.Println("===================================")
	fmt.Println("1. Hitung Kecepatan")
	fmt.Println("2. Hitung Jarak")
	fmt.Println("3. Hitung Waktu")
	fmt.Print("Masukkan pilihan (1/2/3): ")
	fmt.Scan(&pilihan)

	fmt.Println("-----------------------------------")

	switch pilihan {

	//velocity measurement
	case 1:
		var jarakKm, waktuJam float64

		fmt.Print("Masukkan jarak (km): ")
		fmt.Scan(&jarakKm)
		fmt.Print("Masukkan waktu (jam): ")
		fmt.Scan(&waktuJam)

		if jarakKm <= 0 || waktuJam <= 0 {
			fmt.Println("KESALAHAN INPUT")
			fmt.Println("Jarak dan waktu harus > 0")
			return
		}

		jarakMeter := kmKeMeter(jarakKm)
		waktuDetik := jamKeDetik(waktuJam)

		kecepatanMS := jarakMeter / waktuDetik
		kecepatanKMH := msKeKmh(kecepatanMS)

		fmt.Println("HASIL PERHITUNGAN")
		fmt.Printf("Kecepatan = %.2f m/s\n", kecepatanMS)
		fmt.Printf("Kecepatan = %.2f km/h\n", kecepatanKMH)

	//distance measurement
	case 2:
		var kecepatanKmh, waktuJam float64

		fmt.Print("Masukkan kecepatan (km/h): ")
		fmt.Scan(&kecepatanKmh)
		fmt.Print("Masukkan waktu (jam): ")
		fmt.Scan(&waktuJam)

		if kecepatanKmh <= 0 || waktuJam <= 0 {
			fmt.Println("KESALAHAN INPUT")
			fmt.Println("Kecepatan dan waktu harus > 0")
			return
		}

		jarakKm := kecepatanKmh * waktuJam
		jarakMeter := kmKeMeter(jarakKm)

		fmt.Println("HASIL PERHITUNGAN")
		fmt.Printf("Jarak = %.2f km\n", jarakKm)
		fmt.Printf("Jarak = %.2f meter\n", jarakMeter)

	//time measurement
	case 3:
		var jarakKm, kecepatanKmh float64

		fmt.Print("Masukkan jarak (km): ")
		fmt.Scan(&jarakKm)
		fmt.Print("Masukkan kecepatan (km/h): ")
		fmt.Scan(&kecepatanKmh)

		if jarakKm <= 0 || kecepatanKmh <= 0 {
			fmt.Println("KESALAHAN INPUT")
			fmt.Println("Jarak dan kecepatan harus > 0")
			return
		}

		waktuJam := jarakKm / kecepatanKmh
		waktuDetik := jamKeDetik(waktuJam)

		fmt.Println("HASIL PERHITUNGAN")
		fmt.Printf("Waktu = %.2f jam\n", waktuJam)
		fmt.Printf("Waktu = %.2f detik\n", waktuDetik)

	default:
		fmt.Println("KESALAHAN: Menu tidak tersedia.")
	}

	fmt.Println("===================================")
	fmt.Println("Program selesai dijalankan.")
	fmt.Println("===================================")
}
