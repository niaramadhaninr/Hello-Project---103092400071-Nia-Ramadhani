package main

import "fmt"

type tNegara struct {
	Nama     string
	Emas     int
	Perak    int
	Perunggu int
}

const NMAX int = 100

type tabNegara [NMAX]tNegara

var daftarNegara tabNegara
var jumlahNegara int = 0

func tambahNegara(A *tabNegara, n *int) {
	var jumlah int
	fmt.Print("MASUKKAN JUMLAH NEGARA: ")
	fmt.Scanln(&jumlah)

	i := 0
	for i < jumlah {
		var nama string
		fmt.Println("Masukkan nama negara ke-", i+1, ": ")
		fmt.Scanln(&nama)

		j := 0
		for j < *n && A[j].Nama != nama {
			j++
		}

		if j < *n {
			fmt.Println("NEGARA SUDAH TERDAFTAR")
		} else if *n < NMAX {
			A[*n].Nama = nama
			A[*n].Emas = 0
			A[*n].Perak = 0
			A[*n].Perunggu = 0
			*n++
			fmt.Println("NEGARA BERHASIL DITAMBAHKAN")
			i++
		}
	}
}

func ubahNegara(A *tabNegara, n *int) {
	var nama string
	fmt.Print("MASUKKAN NEGARA YANG INGIN DIRUBAH: ")
	fmt.Scanln(&nama)

	for i := 0; i < *n; i++ {
		if A[i].Nama == nama {
			var namaBaru string
			fmt.Print("MASUKKAN NEGARA BARU: ")
			fmt.Scanln(&namaBaru)
			A[i].Nama = namaBaru
			fmt.Println("NEGARA BERHASIL DIRUBAH")
			return
		}
	}

	fmt.Println("NEGARA TIDAK ADA!")
}

func hapusNegara(A *tabNegara, n *int) {
	var nama string
	fmt.Print("MASUKKAN NEGARA YANG INGIN DIHAPUS: ")
	fmt.Scanln(&nama)

	for i := 0; i < *n; i++ {
		if A[i].Nama == nama {
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			fmt.Println("NEGARA SUDAH DIHAPUS")
			return
		}
	}

	fmt.Println("NEGARA TIDAK ADA!")
}

func tambahMedali(A *tabNegara, n *int) {
	var nama string
	fmt.Print("NAMA NEGARA: ")
	fmt.Scanln(&nama)

	for i := 0; i < *n; i++ {
		if A[i].Nama == nama {
			var emas, perak, perunggu int
			fmt.Print("MASUKKAN JUMLAH MEDALI EMAS: ")
			fmt.Scanln(&emas)
			fmt.Print("MASUKKAN JUMLAH MEDALI PERAK: ")
			fmt.Scanln(&perak)
			fmt.Print("MASUKKAN JUMLAH MEDALI PERUNGGU: ")
			fmt.Scanln(&perunggu)

			A[i].Emas += emas
			A[i].Perak += perak
			A[i].Perunggu += perunggu

			fmt.Println("MEDALI BERHASIL DITAMBAH")
			return
		}
	}

	fmt.Println("NEGARA TIDAK ADA")
}

func urutkanDanTampilkan(A *tabNegara, n *int) {
	var temp_element tNegara
	var i, pass int

	for pass = 1; pass < *n; pass++ {
		temp_element = A[pass]
		i = pass

		for i > 0 &&
			(A[i-1].Emas < temp_element.Emas ||
				(A[i-1].Emas == temp_element.Emas && A[i-1].Perak < temp_element.Perak) ||
				(A[i-1].Emas == temp_element.Emas && A[i-1].Perak == temp_element.Perak && A[i-1].Perunggu < temp_element.Perunggu)) {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp_element
	}

	fmt.Println("\nDaftar Peringkat:")
	fmt.Println("--------------------------------------------------")
	fmt.Println("No. | Negara       | Emas | Perak | Perunggu | Total")
	fmt.Println("--------------------------------------------------")
	for i := 0; i < *n; i++ {
		total := A[i].Emas + A[i].Perak + A[i].Perunggu
		fmt.Printf("%-3d | %-12s | %-4d | %-5d | %-8d | %-5d\n",
			i+1, A[i].Nama, A[i].Emas, A[i].Perak, A[i].Perunggu, total)
	}
	fmt.Println("--------------------------------------------------")
}

func tampilkanTop3(A *tabNegara, n *int) {
	if *n == 0 {
		fmt.Println("TIDAK ADA DATA NEGARA")
		return
	}

	urutkanDanTampilkan(A, n)

	fmt.Println("\nJUARA 1, 2, DAN 3:")
	fmt.Println("--------------------------------------------------")
	fmt.Println("No. | Negara       | Emas | Perak | Perunggu | Total")
	fmt.Println("--------------------------------------------------")

	limit := 3
	if *n < 3 {
		limit = *n
	}

	for i := 0; i < limit; i++ {
		total := A[i].Emas + A[i].Perak + A[i].Perunggu
		fmt.Printf("%-3d | %-12s | %-4d | %-5d | %-8d | %-5d\n",
			i+1, A[i].Nama, A[i].Emas, A[i].Perak, A[i].Perunggu, total)
	}

	fmt.Println("--------------------------------------------------")
}

func tampilkanMenu() {
	fmt.Println("\nAplikasi Seagames Manager")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Ubah Nama Negara")
	fmt.Println("3. Hapus Negara")
	fmt.Println("4. Tambah/Update Medali")
	fmt.Println("5. Tampilkan Peringkat")
	fmt.Println("6. Tampilkan Juara 1, 2, dan 3")
	fmt.Println("7. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var pilihan int

	for pilihan != 7 {
		tampilkanMenu()
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahNegara(&daftarNegara, &jumlahNegara)
		} else if pilihan == 2 {
			ubahNegara(&daftarNegara, &jumlahNegara)
		} else if pilihan == 3 {
			hapusNegara(&daftarNegara, &jumlahNegara)
		} else if pilihan == 4 {
			tambahMedali(&daftarNegara, &jumlahNegara)
		} else if pilihan == 5 {
			urutkanDanTampilkan(&daftarNegara, &jumlahNegara)
		} else if pilihan == 6 {
			tampilkanTop3(&daftarNegara, &jumlahNegara)
		} else if pilihan == 7 {
			fmt.Println("TERIMA KASIH")
		} else {
			fmt.Println("Pilihan tidak valid")
		}
	}
}
