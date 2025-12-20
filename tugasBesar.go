// Nama Kelompok : Kelompok 17
// Anggota Kelompok : Fathuridza Akmal Hafidz Priyambodo, Yamamo Juan Alterico Situmorang
// NIM Kelompok : 103032500048, 103032530028
// Kelas : IT-49-03

package main

import "fmt"
import "time"

// konversi
func kmToMeter(km float64) float64 { return km * 1000 }
func hourToSecond(hour float64) float64 { return hour * 3600 }
func msToKmh(ms float64) float64 { return ms * 3.6 }

// hitung kecepatan
func hitungKecepatan() {
	var jarak, waktu float64
	var satuanJarak, satuanWaktu int

	fmt.Println("Satuan Jarak: 1.Km  2.Meter")
	fmt.Print("Pilih: ")
	fmt.Scan(&satuanJarak)
	fmt.Print("Masukkan jarak: ")
	fmt.Scan(&jarak)

	fmt.Println("\nSatuan Waktu: 1.Jam  2.Detik")
	fmt.Print("Pilih: ")
	fmt.Scan(&satuanWaktu)
	fmt.Print("Masukkan waktu: ")
	fmt.Scan(&waktu)

	if jarak <= 0 || waktu <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	if satuanJarak == 1 {
		jarak = kmToMeter(jarak)
	}
	if satuanWaktu == 1 {
		waktu = hourToSecond(waktu)
	}

	kecepatan := jarak / waktu
	fmt.Printf("\nKecepatan: %.2f km/h atau %.2f m/s\n", msToKmh(kecepatan), kecepatan,)	
}

// hitung jarak
func hitungJarak() {
	var kecepatan, waktu float64

	fmt.Print("Kecepatan (km/h): ")
	fmt.Scan(&kecepatan)
	fmt.Print("Waktu (jam): ")
	fmt.Scan(&waktu)

	if kecepatan <= 0 || waktu <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	jarak := kecepatan * waktu
	fmt.Printf("\nJarak: %.2f km atau %.0f meter\n", jarak, kmToMeter(jarak))

}

// hitung waktu
func hitungWaktu() {
	var jarak, kecepatan float64

	fmt.Print("Jarak (km): ")
	fmt.Scan(&jarak)
	fmt.Print("Kecepatan (km/h): ")
	fmt.Scan(&kecepatan)

	if jarak <= 0 || kecepatan <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	totalJam := jarak / kecepatan

	jam := int(totalJam)
	sisaMenit := (totalJam - float64(jam)) * 60
	menit := int(sisaMenit)
	detik := int((sisaMenit - float64(menit)) * 60)

	fmt.Printf("\nWaktu: %d jam, %d menit, %d detik\n", jam, menit, detik)
}


func main() {
	var menu, ulang int

	fmt.Println("dibuat dengan penuh semangat oleh kelompok 17")
	time.Sleep(500 * time.Millisecond)

	for {
		fmt.Println("\nAplikasi Penghitung Kecepatan, Jarak, dan Waktu")
		fmt.Println("1. Hitung Kecepatan")
		fmt.Println("2. Hitung Jarak")
		fmt.Println("3. Hitung Waktu")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&menu)

		fmt.Println("--------------------")

		switch menu {
		case 1:
			hitungKecepatan()
		case 2:
			hitungJarak()
		case 3:
			hitungWaktu()
		default:
			fmt.Println("Menu tidak tersedia")
			continue
		}

		fmt.Println("\nProgram selesai!")
		fmt.Println("1. Hitung ulang")
		fmt.Println("2. Keluar")
		fmt.Print("Pilih: ")
		fmt.Scan(&ulang)

		if ulang == 2 {
			fmt.Println("Terima kasih!")
			break
		}
	}
}
