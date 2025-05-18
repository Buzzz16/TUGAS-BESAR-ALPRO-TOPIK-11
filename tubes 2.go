package main

import "fmt"

const maxCapacity = 30

type itemDapur struct {
	nama              string
	jumlah            int
	satuan            string
	jenis             string
	tanggalkadaluarsa int
	bulankadaluarsa   int
	tahunkadaluarsa   int
	tanggalmasuk      int
	bulanmasuk        int
	tahunmasuk        int
}

type item [maxCapacity]itemDapur

var barang item
var banyakData int
var tglSekarang, blnSekarang, thnSekarang int

func main() {
	var n int
	n = 1
	baseTersedia()
	hariIni()
	for n != 0 {
		listmenu()
		fmt.Scan(&n)
		switch n {
		case 1:
			tambahItem(&barang, &banyakData)
			tunggu()
		case 2:
			edit()
		case 3:
			tampilkanItemDefault()
			tunggu()
		case 4:
			urutkanBahan()
		case 5:
			bahanDigunakan()
		case 6:
			cariBarang()
		case 7:
			reset()
		case 8:
			n = 0
			fmt.Println("Terima kasih telah menggunakan Aplikasi Manajemen Stok Bahan Makanan!")
		default:
			hapus()
			fmt.Println("~~~           Opsi tidak valid!           ~~~")
			fmt.Println("~~~       Silahkan coba lagi (1 - 8)      ~~~")
		}
	}
}

func listmenu() {
	hapus()
	fmt.Println()
	fmt.Println("=== APLIKASI MANAJEMEN STOK BAHAN MAKANAN ===")
	fmt.Println()
	fmt.Println("1. Tambah Bahan Makanan")
	fmt.Println("2. edit")
	fmt.Println("3. Tampilkan Bahan Makanan (Default)")
	fmt.Println("4. Urutkan Bahan Makanan")
	fmt.Println("5. Bahan Makanan yang akan Digunakan")
	fmt.Println("6. cari barang")
	fmt.Println("7. Reset Data")
	fmt.Println("8. Keluar")
	fmt.Println()
	fmt.Println("=============================================")
	fmt.Print("Pilih Opsi:  ")
}

func tambahItem(barang *item, banyakData *int) {
	if *banyakData >= maxCapacity {
		fmt.Println("Kapasitas maksimum telah tercapai!")
		return
	}
	hapus()
	fmt.Println("Tambah Bahan Makanan")
	fmt.Print("Masukkan nama bahan             : ")
	fmt.Scan(&barang[*banyakData].nama)
	fmt.Print("Masukkan jumlah bahan           : ")
	fmt.Scan(&barang[*banyakData].jumlah)
	fmt.Print("Satuan (kg, g, butir, dll)      : ")
	fmt.Scan(&barang[*banyakData].satuan)
	fmt.Print("Kategori (sayuran, buah, protein, dll)   : ")
	fmt.Scan(&barang[*banyakData].jenis)
	fmt.Print("Tanggal kadaluarsa (dd-mm-yyyy)       : ")
	fmt.Scan(&barang[*banyakData].tanggalkadaluarsa, &barang[*banyakData].bulankadaluarsa, &barang[*banyakData].tahunkadaluarsa)
	fmt.Print("Tanggal masuk (dd-mm-yyyy))           : ")
	fmt.Scan(&barang[*banyakData].tanggalmasuk, &barang[*banyakData].bulanmasuk, &barang[*banyakData].tahunmasuk)
	*banyakData = *banyakData + 1
}

func tampilkanItem() {
	hapus()

	if banyakData == 0 {
		fmt.Println("Tidak ada data bahan makanan yang ditambahkan.")
	} else {
		fmt.Println("-----------------------------------------------------------------------------------------------")
		fmt.Println("        ~~~~~~~~~~~~~~~~~~~~| TANGGAL HARI INI ", tglSekarang, "-", blnSekarang, "-", thnSekarang, " |~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("---------------------------------------------------------------------------------------------")
		fmt.Println("No. | Nama           | Jumlah     | Satuan       | Kategori      | Kedaluwarsa |  Sisa Hari  |")
		fmt.Println("----|----------------|------------|--------------|---------------|-------------|-------------|")
		for i := 0; i < banyakData; i++ {
			sisaHari := (barang[i].tahunkadaluarsa*365 + barang[i].bulankadaluarsa*30 + barang[i].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
			if sisaHari < 0 {
				sisaHari = 0
			}
			fmt.Printf("%-3d | %-14s | %-10d | %-12s | %-13s | %02d/%02d/%-4d  | %-11d |\n",
				i+1, barang[i].nama, barang[i].jumlah, barang[i].satuan, barang[i].jenis,
				barang[i].tanggalkadaluarsa, barang[i].bulankadaluarsa, barang[i].tahunkadaluarsa, sisaHari)
		}
	}
	fmt.Println()
}

func tampilkanItemDefault() {
	hapus()
	var pilih int
	pilih = 1
	for i := 0; i < banyakData-1; i++ {
		idx := i
		for j := i + 1; j < banyakData; j++ {
			if pilih == 1 && barang[j].tanggalmasuk > barang[idx].tanggalmasuk {
				idx = j
			}
		}
		if idx != i {
			barang[i], barang[idx] = barang[idx], barang[i]
		}
	}
	tampilkanItem()
}

func baseTersedia() {
	barang[0] = itemDapur{"Apel", 1, "kg", "Buah", 27, 12, 2026, 1, 12, 2025}
	barang[1] = itemDapur{"Gula Pasir", 2, "kg", "Pemanis", 15, 6, 2025, 2, 1, 2025}
	barang[2] = itemDapur{"Beras", 10, "kg", "Karbohidrat", 12, 11, 2024, 3, 1, 2023}
	barang[3] = itemDapur{"Daging Sapi", 5, "kg", "Protein", 1, 12, 2024, 4, 1, 2023}
	barang[4] = itemDapur{"Minyak Goreng", 2, "liter", "Lemak", 8, 11, 2024, 5, 1, 2023}
	barang[5] = itemDapur{"Telur", 12, "butir", "Protein", 3, 12, 2024, 6, 1, 2023}
	banyakData = 6
}
func hariIni() {
	tglSekarang, blnSekarang, thnSekarang = 1, 1, 2024
}
func ubahHari() {
	fmt.Print("masukan tanggal hari ini (1-31) : ")
	fmt.Scan(&tglSekarang)
	fmt.Print("masukan bulan hari ini (1-12) : ")
	fmt.Scan(&blnSekarang)
	fmt.Print("masukan tahun hari ini: ")
	fmt.Scan(&thnSekarang)
}
func edit() {
	hapus()
	var pilih int
	fmt.Println("1. Edit Nama Bahan Makanan")
	fmt.Println("2. Edit Jumlah")
	fmt.Println("3. Edit Satuan")
	fmt.Println("4. Edit Kategori")
	fmt.Println("5. Edit Tanggal Kadaluarsa")
	fmt.Println("6. Edit Tanggal Hari Ini")
	fmt.Println("7. Kembali ke Menu Utama")
	fmt.Print("Pilih Opsi: ")

	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		hapus()
		editNamaBahan()
	case 2:
		hapus()
		editJumlah()
	case 3:
		hapus()
		editSatuan()
	case 4:
		hapus()
		editKategori()
	case 5:
		hapus()
		editKadaluarsa()
	case 6:
		hapus()
		ubahHari()
	case 7:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 7)      ~~~")
	}
}
func editNamaBahan() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func editJumlah() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func editSatuan() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func editKategori() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func editKadaluarsa() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func urutkanBahan() {
	hapus()
	var pilih int
	fmt.Println("Urutkan berdasarkan :")
	fmt.Println("1. jumlah")
	fmt.Println("2. nama")
	fmt.Println("3. sisa hari kadaluarsa")
	fmt.Println("4. kembali ke menu")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		urutjumlah()
	case 2:
		ururtNama()
	case 3:
		urutSisaHari()
	case 4:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 3)      ~~~")
	}
	tunggu()
}
func urutjumlah() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Print("Pilih opsi  : ")
	fmt.Scan(&pilih)
	for i := 0; i < banyakData-1; i++ {
		idx := i
		for j := i + 1; j < banyakData; j++ {
			if pilih == 1 && barang[j].jumlah > barang[idx].jumlah {
				idx = j
			} else if pilih == 2 && barang[j].jumlah < barang[idx].jumlah {
				idx = j
			}
		}
		if idx != i {
			barang[i], barang[idx] = barang[idx], barang[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
	tampilkanItem()
}

func ururtNama() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Scan(&pilih)
	for i := 0; i < banyakData-1; i++ {
		idx := i
		for j := i + 1; j < banyakData; j++ {
			if pilih == 1 && barang[j].nama > barang[idx].nama {
				idx = j
			} else if pilih == 2 && barang[j].nama < barang[idx].nama {
				idx = j
			}
		}
		if idx != i {
			barang[i], barang[idx] = barang[idx], barang[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
	tampilkanItem()
}
func urutSisaHari() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Scan(&pilih)
	for i := 0; i < banyakData-1; i++ {
		idx := i
		for j := i + 1; j < banyakData; j++ {
			sisaHariI := (barang[i].tahunkadaluarsa*365 + barang[i].bulankadaluarsa*30 + barang[i].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
			sisaHariJ := (barang[j].tahunkadaluarsa*365 + barang[j].bulankadaluarsa*30 + barang[j].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
			if pilih == 1 && sisaHariJ > sisaHariI {
				idx = j
			} else if pilih == 2 && sisaHariJ < sisaHariI {
				idx = j
			}
		}
		if idx != i {
			barang[i], barang[idx] = barang[idx], barang[i]
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
	tampilkanItem()
}
func bahanDigunakan() {
	hapus()
	var pilih int
	tampilkanItem()
	fmt.Println("cari bahan yang akan digunakan hari ini berdasarkan")
	fmt.Println("1. nama")
	fmt.Println("2. jumlah")
	fmt.Println("3. satuan")
	fmt.Println("4. kategori")
	fmt.Println("5. tanggal kadaluarsa")
	fmt.Println("6. tanggal sisa kadaluarsa")
	fmt.Println("7. kembali ke menu utama")
	fmt.Println("Pilih Opsi: ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		hapus()
		cariNama()
	case 2:
		hapus()
		cariJumlah()
	case 3:
		hapus()
		cariSatuan()
	case 4:
		hapus()
		cariKategori()
	case 5:
		hapus()
		cariKadaluarsa()
	case 6:
		hapus()
		cariSisaHari()
	case 7:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 7)      ~~~")
	}
	fmt.Println()
}
func cariNama() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariJumlah() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariSatuan() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariKategori() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariKadaluarsa() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariSisaHari() {
	hapus()
	fmt.Println("Fungsi urutkan belum diimplementasikan.")
	tunggu()
}
func cariBarang() {
	hapus()
	fmt.Print("belom")
}
func reset() {
	hapus()
	banyakData = 0
	fmt.Println("Semua data telah dihapus.")
	tunggu()
}

func hapus() {
	fmt.Print("\033[H\033[2J")
	for i := 0; i < 10; i++ {
		fmt.Println()
	}
}

func tunggu() {
	fmt.Print("Ketik 'ok' untuk melanjutkan: ")
	var s string
	fmt.Scan(&s)
	for s != "ok" {
		fmt.Print("Silahkan ketik 'ok' untuk melanjutkan: ")
		fmt.Scan(&s)
	}
	hapus()
}
