// Nomor Kelompok: Kelompok 17
// Anggota Kelompok: Fathuridza Akmal Hafidz Priyambodo, Yamamo Juan Alterico Situmorang
// Kelas: IT-49-03

package main
import "fmt"

// konversi

func kmToM(km float64) float64 { return km * 1000 }
func hToS(h float64) float64  { return h * 3600 }
func msToKmh(ms float64) float64 { return ms * 3.6 }

// kecepatan

func hitungKecepatan() {
	var jarak, waktu float64
	var sj, sw int

	fmt.Println("Satuan Jarak: 1.Km  2.Meter")
	fmt.Print("Pilih: ")
	fmt.Scan(&sj)
	fmt.Print("Masukkan jarak: ")
	fmt.Scan(&jarak)

	fmt.Println("\nSatuan Waktu: 1.Jam  2.Detik")
	fmt.Print("Pilih: ")
	fmt.Scan(&sw)
	fmt.Print("Masukkan waktu: ")
	fmt.Scan(&waktu)

	if jarak <= 0 || waktu <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	if sj == 1 {
		jarak = kmToM(jarak)
	}
	if sw == 1 {
		waktu = hToS(waktu)
	}

	ms := jarak / waktu
	fmt.Printf("\nKecepatan: %.2f m/s\n", ms)
	fmt.Printf("Kecepatan: %.2f km/h\n", msToKmh(ms))
}

// jarak

func hitungJarak() {
	var v, t float64

	fmt.Print("Kecepatan (km/h): ")
	fmt.Scan(&v)
	fmt.Print("Waktu (jam): ")
	fmt.Scan(&t)

	if v <= 0 || t <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	km := v * t
	fmt.Printf("\nJarak: %.2f km\n", km)
	fmt.Printf("Jarak: %.2f meter\n", kmToM(km))
}

// waktu 

func hitungWaktu() {
	var s, v float64

	fmt.Print("Jarak (km): ")
	fmt.Scan(&s)
	fmt.Print("Kecepatan (km/h): ")
	fmt.Scan(&v)

	if s <= 0 || v <= 0 {
		fmt.Println("Input tidak valid")
		return
	}

	h := s / v
	fmt.Printf("\nWaktu: %.2f jam\n", h)
	fmt.Printf("Waktu: %.2f detik\n", hToS(h))
}

// utama

func main() {
	var menu int

	fmt.Println("Aplikasi Jarak, Waktu, Kecepatan")
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
	}
}
