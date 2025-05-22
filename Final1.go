package main

import (
	"fmt"
	"strings"
	"time"
	"os"
)


type Sampah struct {
	Jenis      string
	Berat      float64
	Tanggal    string
	Recyclable bool
}

const maxData = 1000 

func main() {
	
	var data [maxData]Sampah
	var currentIndex int

	for {
		tampilkanMenu()
		pilihMenu(&data, &currentIndex)
	}
}


func tampilkanMenu() {
	fmt.Println("\n=== APLIKASI PENGELOLAAN SAMPAH ===")
	fmt.Println("1. Input Data Sampah")
	fmt.Println("2. Tampilkan Statistik")
	fmt.Println("3. Tampilkan Rekomendasi")
	fmt.Println("4. Tampilkan Riwayat")
	fmt.Println("5. Edit dan Hapus Data") 
	fmt.Println("6. Keluar")
	fmt.Print("Pilih menu: ")
}



func pilihMenu(data *[maxData]Sampah, currentIndex *int) {
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		inputData(data, currentIndex)
	case 2:
		tampilkanStatistik(data, *currentIndex)
	case 3:
		beriRekomendasi(data, *currentIndex)
	case 4:
		tampilkanRiwayat(data, *currentIndex)
	case 5:
		editData(data, currentIndex)
	case 6:
		fmt.Println("Terima kasih!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid!")
	}
}



func inputData(data *[maxData]Sampah, idx *int) {
    if *idx >= maxData {
        fmt.Println("Penyimpanan penuh!")
        return
    }

    var jenis, tanggal string
    var berat float64

    fmt.Print("Jenis sampah: ")
    fmt.Scanln(&jenis)

    for {
        fmt.Print("Berat (kg): ")
        fmt.Scanln(&berat)
        if berat > 0 {
            break
        }
        fmt.Println("Berat harus lebih dari 0!")
    }

    fmt.Print("Tanggal (DD-MM-YYYY): ")
    fmt.Scanln(&tanggal)

    recyclable := cekDaurUlang(jenis)
    (*data)[*idx] = Sampah{jenis, berat, tanggal, recyclable}
    *idx++
    fmt.Println("Data tersimpan!")
}


func cekDaurUlang(jenis string) bool {
	s := strings.ToLower(jenis)
	switch s {
	case "plastik", "kertas", "kaca", "elektronik":
		return true
	default:
		return false
	}
}


func tampilkanStatistik(data *[maxData]Sampah, count int) {
    today := time.Now().Format("02-01-2006") // Ubah format ke DD-MM-YYYY
    var totalHari, totalMinggu, recHari, recMinggu float64
    year, week := time.Now().ISOWeek()

    for i := 0; i < count; i++ {
        entry := data[i]
        if entry.Tanggal == today {
            totalHari += entry.Berat
            if entry.Recyclable {
                recHari += entry.Berat
            }
        }

        t, err := time.Parse("02-01-2006", entry.Tanggal)
        if err != nil {
            continue
        }
        y, w := t.ISOWeek()
        if y == year && w == week {
            totalMinggu += entry.Berat
            if entry.Recyclable {
                recMinggu += entry.Berat
            }
        }
    }

    fmt.Println("\n=== STATISTIK ===")
    fmt.Printf("Hari ini: %.1f kg (%.1f%% recyclable)\n", totalHari, hitungPersen(recHari, totalHari))
    fmt.Printf("Minggu ini: %.1f kg (%.1f%% recyclable)\n", totalMinggu, hitungPersen(recMinggu, totalMinggu))
}

func hitungPersen(part, total float64) float64 {
	if total == 0 { return 0 }
	return (part / total) * 100
}


func beriRekomendasi(data *[maxData]Sampah, count int) {
	var total, rec float64
	for i := 0; i < count; i++ {
		total += data[i].Berat
		if data[i].Recyclable {
			rec += data[i].Berat
		}
	}

	persenNon := 100 - hitungPersen(rec, total)
	if persenNon > 50 {
		fmt.Println("\nREKOMENDASI: Kurangi sampah non-recyclable!")
		fmt.Println("Gunakan wadah reusable dan kurangi sampah organik.")
	} else {
		fmt.Println("\nBagus! Pertahankan daur ulang.")
	}
}


func tampilkanRiwayat(data *[maxData]Sampah, count int) {
    if count == 0 {
        fmt.Println("Belum ada data")
        return
    }

    slice := make([]Sampah, count)
    copy(slice, data[:count])

    fmt.Println("\nPilih sorting:")
    fmt.Println("1. Jenis (A-Z)")
    fmt.Println("2. Tanggal (Terbaru)")
    fmt.Println("3. Berat (Terbesar)")
    fmt.Print("Pilihan: ")

    var pil int
    fmt.Scanln(&pil)

    switch pil {
    case 1:
        selectionSortJenis(slice)
    case 2:
        insertionSortTanggal(slice)
    case 3:
        selectionSortBerat(slice)
    }

    fmt.Println("\n=== RIWAYAT ===")
    fmt.Printf("%-15s %-8s %-12s %-10s\n", "Jenis", "Berat", "Tanggal", "Recyclable")
    for _, s := range slice {
        status := "Tidak"
        if s.Recyclable {
            status = "Ya"
        }
        fmt.Printf("%-15s %-8.1f %-12s %-10s\n", s.Jenis, s.Berat, s.Tanggal, status)
    }
}
func selectionSortJenis(arr []Sampah) { /* ... */ }
func insertionSortTanggal(arr []Sampah) { /* ... */ }
func selectionSortBerat(arr []Sampah)  { /* ... */ }


func editData(data *[maxData]Sampah, currentIndex *int) {
	if *currentIndex == 0 {
		fmt.Println("Belum ada data untuk diedit atau dihapus")
		return
	}

	
	fmt.Println("\n=== DAFTAR DATA SAMPAH ===")
    for i := 0; i < *currentIndex; i++ {
        s := (*data)[i]
        status := "Tidak"
        if s.Recyclable {
            status = "Ya"
        }
        fmt.Printf("%d. %s | %.1f kg | %s | Recyclable: %s\n", i+1, s.Jenis, s.Berat, s.Tanggal, status)
    }

	
	fmt.Print("Pilih nomor data yang ingin diedit atau dihapus: ")
	var no int
	fmt.Scanln(&no)
	if no < 1 || no > *currentIndex {
		fmt.Println("Nomor tidak valid!")
		return
	}
	idx := no - 1

	fmt.Println("Pilih aksi:")
	fmt.Println("1. Edit data")
	fmt.Println("2. Hapus data")
	fmt.Print("Masukkan pilihan: ")
	var aksi int
	fmt.Scanln(&aksi)

	switch aksi {
	case 1:
	
		fmt.Printf("Edit Jenis (sekarang: %s): ", (*data)[idx].Jenis)
		fmt.Scanln(&(*data)[idx].Jenis)
		fmt.Printf("Edit Berat (sekarang: %.1f): ", (*data)[idx].Berat)
		fmt.Scanln(&(*data)[idx].Berat)
		(*data)[idx].Recyclable = cekDaurUlang((*data)[idx].Jenis)
		fmt.Println("Data berhasil diperbarui!")

	case 2:

		for i := idx; i < *currentIndex-1; i++ {
			(*data)[i] = (*data)[i+1]
		}
		*currentIndex--
		fmt.Println("Data berhasil dihapus!")

	default:
		fmt.Println("Pilihan aksi tidak valid!")
	}
}


func findIndexByTanggal(data *[maxData]Sampah, count int, tgl string) int {
	for i := 0; i < count; i++ {
		if data[i].Tanggal == tgl {
			return i
		}
	}
	return -1
}




func findIndexByJenis(data *[maxData]Sampah, count int, jenis string) int {
	for i := 0; i < count; i++ {
		if strings.EqualFold(data[i].Jenis, jenis) {
			return i
		}
	}
	return -1
}
