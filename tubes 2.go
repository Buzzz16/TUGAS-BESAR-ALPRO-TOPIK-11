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
var pemberitahuan string
var sisaHari int

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
			tampilkanItem()
			tunggu()
		case 4:
			tampilkanItemDefault()
			tunggu()
		case 5:
			urutkanBahan()
		case 6:
			bahanDigunakan()
		case 7:
			cariBahan()
		case 8:
			reset()
		case 9:
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
	fmt.Println(" ▄▄▄·  ▄▄▄·▄▄▌  ▪  ▄ •▄  ▄▄▄· .▄▄ · ▪      • ▌ ▄ ·.  ▄▄▄·  ▐ ▄  ▄▄▄·  ▐▄▄▄▄▄▄ .• ▌ ▄ ·. ▄▄▄ . ▐ ▄     ▄▄▄▄·  ▄▄▄·  ▄ .▄ ▄▄▄·  ▐ ▄     ")
	fmt.Println("▐█ ▀█ ▐█ ▄███•  ██ █▌▄▌▪▐█ ▀█ ▐█ ▀. ██     ·██ ▐███▪▐█ ▀█ •█▌▐█▐█ ▀█   ·██▀▄.▀··██ ▐███▪▀▄.▀·•█▌▐█    ▐█ ▀█▪▐█ ▀█ ██▪▐█▐█ ▀█ •█▌▐█    ")
	fmt.Println("▄█▀▀█  ██▀·██▪  ▐█·▐▀▀▄·▄█▀▀█ ▄▀▀▀█▄▐█·    ▐█ ▌▐▌▐█·▄█▀▀█ ▐█▐▐▌▄█▀▀█ ▪▄ ██▐▀▀▪▄▐█ ▌▐▌▐█·▐▀▀▪▄▐█▐▐▌    ▐█▀▀█▄▄█▀▀█ ██▀▐█▄█▀▀█ ▐█▐▐▌    ")
	fmt.Println("▐█ ▪▐▌▐█▪·•▐█▌▐▌▐█▌▐█.█▌▐█ ▪▐▌▐█▄▪▐█▐█▌    ██ ██▌▐█▌▐█ ▪▐▌██▐█▌▐█ ▪▐▌▐▌▐█▌▐█▄▄▌██ ██▌▐█▌▐█▄▄▌██▐█▌    ██▄▪▐█▐█ ▪▐▌██▌▐▀▐█ ▪▐▌██▐█▌    ")
	fmt.Println(" ▀  ▀ .▀   .▀▀▀ ▀▀▀·▀  ▀ ▀  ▀  ▀▀▀▀ ▀▀▀    ▀▀  █▪▀▀▀ ▀  ▀ ▀▀ █▪ ▀  ▀  ▀▀▀• ▀▀▀ ▀▀  █▪▀▀▀ ▀▀▀ ▀▀ █▪    ·▀▀▀▀  ▀  ▀ ▀▀▀ · ▀  ▀ ▀▀ █▪    ")
	fmt.Println("                         ▄▄▄  ▄• ▄▌• ▌ ▄ ·.  ▄▄▄·  ▄ .▄    ▄▄▄▄▄ ▄▄▄·  ▐ ▄  ▄▄ •  ▄▄ •  ▄▄▄·                                          ")
	fmt.Println("                         ▀▄ █·█▪██▌·██ ▐███▪▐█ ▀█ ██▪▐█    •██  ▐█ ▀█ •█▌▐█▐█ ▀ ▪▐█ ▀ ▪▐█ ▀█                                          ")
	fmt.Println("                         ▐▀▀▄ █▌▐█▌▐█ ▌▐▌▐█·▄█▀▀█ ██▀▐█     ▐█.▪▄█▀▀█ ▐█▐▐▌▄█ ▀█▄▄█ ▀█▄▄█▀▀█                                          ")
	fmt.Println("                         ▐█•█▌▐█▄█▌██ ██▌▐█▌▐█ ▪▐▌██▌▐▀     ▐█▌·▐█ ▪▐▌██▐█▌▐█▄▪▐█▐█▄▪▐█▐█ ▪▐▌                                         ")
	fmt.Println("                         .▀  ▀ ▀▀▀ ▀▀  █▪▀▀▀ ▀  ▀ ▀▀▀ ·     ▀▀▀  ▀  ▀ ▀▀ █▪·▀▀▀▀ ·▀▀▀▀  ▀  ▀                                          ")
	fmt.Println()
	fmt.Println("1. Tambah Bahan Makanan")
	fmt.Println("2. Edit")
	fmt.Println("3. Tampilkan Bahan Makanan")
	fmt.Println("4. Tampilkan Bahan Makanan (Default)")
	fmt.Println("5. Urutkan Bahan Makanan")
	fmt.Println("6. Bahan Makanan yang akan Digunakan (Hapus)")
	fmt.Println("7. Cari bahan")
	fmt.Println("8. Reset Data")
	fmt.Println("9. Keluar")
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
func pemberitahuanKadaluarsa(barang item, i int, pemberitahuan *string) {
	var sisaHari int
	sisaHari = (barang[i].tahunkadaluarsa*365 + barang[i].bulankadaluarsa*30 + barang[i].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
	if sisaHari <= 0 {
		sisaHari = 0
		*pemberitahuan = "Bahan sudah kadaluarsa"
	} else if sisaHari >= 1 && sisaHari <= 3 {
		*pemberitahuan = "Mendekati Kadaluarsa"
	} else if sisaHari > 3 {
		*pemberitahuan = "Bahan Masih Aman"
	}
}
func HitungSisaHari(barang item, i int, sisaHari *int) {
	*sisaHari = (barang[i].tahunkadaluarsa*365 + barang[i].bulankadaluarsa*30 + barang[i].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
	if *sisaHari <= 0 {
		*sisaHari = 0
	}
}

func tampilkanItem() {
	hapus()

	if banyakData == 0 {
		fmt.Println("Tidak ada data bahan makanan yang ditambahkan.")
	} else {
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("                         ~~~~~~~~~~~~~~~~~~~~| TANGGAL HARI INI ", tglSekarang, "-", blnSekarang, "-", thnSekarang, " |~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------------")
		fmt.Println("No. | Nama           | Jumlah     | Satuan       | Kategori      |  tanggal masuk |  Kedaluwarsa |  Sisa Hari  |        pemberitahuan       |")
		fmt.Println("----|----------------|------------|--------------|---------------|----------------|--------------|-------------|----------------------------|")

		for i := 0; i < banyakData; i++ {

			pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
			HitungSisaHari(barang, i, &sisaHari)
			fmt.Printf("%-3d | %-14s | %-10d | %-12s | %-13s | %02d/%02d/%-7d  | %02d/%02d/%-4d   | %-11d | %-26s |\n",
				i+1, barang[i].nama, barang[i].jumlah, barang[i].satuan, barang[i].jenis, barang[i].tanggalmasuk, barang[i].bulanmasuk, barang[i].tahunmasuk,
				barang[i].tanggalkadaluarsa, barang[i].bulankadaluarsa, barang[i].tahunkadaluarsa, sisaHari, pemberitahuan)
		}
	}
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
	barang[0] = itemDapur{"Apel", 1, "kg", "Buah", 11, 12, 2025, 12, 12, 2024}
	barang[1] = itemDapur{"Gula Pasir", 2, "kg", "Pemanis", 3, 12, 2024, 13, 9, 2024}
	barang[2] = itemDapur{"Beras", 10, "kg", "Karbohidrat", 9, 12, 2024, 3, 1, 2024}
	barang[3] = itemDapur{"Daging Sapi", 5, "kg", "Protein", 2, 11, 2024, 4, 1, 2024}
	barang[4] = itemDapur{"Minyak Goreng", 2, "liter", "Lemak", 1, 1, 2025, 5, 1, 2024}
	barang[5] = itemDapur{"Telur", 12, "butir", "Protein", 10, 11, 2024, 6, 1, 2024}
	banyakData = 6
}
func hariIni() {
	tglSekarang, blnSekarang, thnSekarang = 1, 12, 2024
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
	if banyakData == 0 {
		fmt.Println("Data Tidak tersedia ")
		fmt.Println("Tambahkan Data terlebih Dahulu")
		tunggu()
		return
	}
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
	var idx int
	tampilkanItem()
	fmt.Print("Masukkan nomor bahan yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx > banyakData {
		fmt.Println("Data tidak ditemukan")
		tunggu()
		return
	}
	fmt.Printf("Nama bahan saat ini: %s\n", barang[idx].nama)
	fmt.Print("Masukkan nama bahan baru: ")
	fmt.Scan(&barang[idx].nama)
	fmt.Println("Nama bahan berhasil diubah!")
	tunggu()
}
func editJumlah() {
	hapus()
	var idx int
	tampilkanItem()
	fmt.Print("Masukkan nomor bahan yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx > banyakData {
		fmt.Println("Data tidak ditemukan")
		tunggu()
		return
	}
	fmt.Printf("Jumlah bahan saat ini: %d\n", barang[idx].jumlah)
	fmt.Print("Masukkan nama bahan baru: ")
	fmt.Scan(&barang[idx].jumlah)
	fmt.Println("jumlah bahan berhasil diubah!")
	tunggu()
}
func editSatuan() {
	hapus()
	var idx int
	tampilkanItem()
	fmt.Print("Masukkan nomor bahan yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx > banyakData {
		fmt.Println("Data tidak ditemukan")
		tunggu()
		return
	}
	fmt.Printf("Nama satuan saat ini: %s\n", barang[idx].satuan)
	fmt.Print("Masukkan nama bahan baru: ")
	fmt.Scan(&barang[idx].satuan)
	fmt.Println("Nama Satuan berhasil diubah!")
	tunggu()
}
func editKategori() {
	hapus()
	var idx int
	tampilkanItem()
	fmt.Print("Masukkan nomor bahan yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx > banyakData {
		fmt.Println("Data tidak ditemukan")
		tunggu()
		return
	}
	fmt.Printf("Nama kategori saat ini: %s\n", barang[idx].jenis)
	fmt.Print("Masukkan nama bahan baru: ")
	fmt.Scan(&barang[idx].jenis)
	fmt.Println("Nama kategori berhasil diubah!")
	tunggu()
}
func editKadaluarsa() {
	hapus()
	var idx int
	tampilkanItem()
	fmt.Print("Masukkan nomor bahan yang ingin diedit: ")
	fmt.Scan(&idx)
	idx--
	if idx > banyakData {
		fmt.Println("Data tidak ditemukan")
		tunggu()
		return
	}
	fmt.Printf("Nama bahan saat ini: %d/%d/%d\n", barang[idx].tanggalkadaluarsa, barang[idx].bulankadaluarsa, barang[idx].tahunkadaluarsa)
	fmt.Print("Masukkan tanggal baru (1-31) : ")
	fmt.Scan(&barang[idx].tanggalkadaluarsa)
	fmt.Print("Masukkan bulan  baru (1-12) : ")
	fmt.Scan(&barang[idx].bulankadaluarsa)
	fmt.Print("Masukkan tahun  baru (1-12) : ")
	fmt.Scan(&barang[idx].tahunkadaluarsa)
	fmt.Println("tanggal, bulan dan tahun kadaluarsa bahan berhasil diubah!")
	tunggu()
}
func urutkanBahan() {
	hapus()
	var pilih int
	fmt.Println("Urutkan Dengan :")
	fmt.Println("1. Selection Sort")
	fmt.Println("2. Insertion Sort")
	fmt.Println("3. Kembali ke menu")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		SelectionSort()
	case 2:
		InsertionSort()
	case 3:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 3)      ~~~")
	}
	tunggu()
}
func SelectionSort() {
	hapus()
	var pilih int
	fmt.Println("Urutkan Dengan :")
	fmt.Println("1. Nama")
	fmt.Println("2. Jumlah")
	fmt.Println("3. Sisa Hari")
	fmt.Println("4. Kembali")
	fmt.Println("5. Kembali ke menu utama")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		ururtNamaSLC()
	case 2:
		urutjumlahSLC()
	case 3:
		urutSisaHariSLC()
	case 4:
		urutkanBahan()
	case 5:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 5)      ~~~")
	}
}
func InsertionSort() {
	hapus()
	var pilih int
	fmt.Println("Urutkan Dengan :")
	fmt.Println("1. Nama")
	fmt.Println("2. Jumlah")
	fmt.Println("3. Sisa Hari")
	fmt.Println("4. kembali")
	fmt.Println("5. kembali ke menu utama")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	switch pilih {
	case 1:
		urutNamaISR()
	case 2:
		urutjumlahISR()
	case 3:
		urutSisaHariISR()
	case 4:
		urutkanBahan()
	case 5:
		hapus()
		return
	default:
		hapus()
		fmt.Println("~~~           Opsi tidak valid!           ~~~")
		fmt.Println("~~~       Silahkan coba lagi (1 - 5)      ~~~")
	}
}
func urutjumlahISR() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Print("Pilih opsi  : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
		for i := 1; i < banyakData; i++ {
			key := barang[i]
			j := i - 1
			if pilih == 1 {
				for j >= 0 && barang[j].jumlah > key.jumlah {
					barang[j+1] = barang[j]
					j--
				}
			} else if pilih == 2 {
				for j >= 0 && barang[j].jumlah < key.jumlah {
					barang[j+1] = barang[j]
					j--
				}
			}
			barang[j+1] = key
		}
		fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
		tampilkanItem()
	} else if pilih == 3 {
		InsertionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}
func urutNamaISR() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Print("Pilih opsi  : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
		for i := 1; i < banyakData; i++ {
			key := barang[i]
			j := i - 1
			if pilih == 1 {
				for j >= 0 && barang[j].nama > key.nama {
					barang[j+1] = barang[j]
					j--
				}
			} else if pilih == 2 {
				for j >= 0 && barang[j].nama < key.nama {
					barang[j+1] = barang[j]
					j--
				}
			}
			barang[j+1] = key
		}
		fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
		tampilkanItem()
	} else if pilih == 3 {
		InsertionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}
func urutSisaHariISR() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Print("Pilih opsi  : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
		for i := 1; i < banyakData; i++ {
			key := barang[i]
			j := i - 1
			sisaHariI := (barang[i].tahunkadaluarsa*365 + barang[i].bulankadaluarsa*30 + barang[i].tanggalkadaluarsa) - (thnSekarang*365 + blnSekarang*30 + tglSekarang)
			if pilih == 1 {
				for j >= 0 && sisaHariI > (barang[j].tahunkadaluarsa*365+barang[j].bulankadaluarsa*30+barang[j].tanggalkadaluarsa)-(thnSekarang*365+blnSekarang*30+tglSekarang) {
					barang[j+1] = barang[j]
					j--
				}
			} else if pilih == 2 {
				for j >= 0 && sisaHariI < (barang[j].tahunkadaluarsa*365+barang[j].bulankadaluarsa*30+barang[j].tanggalkadaluarsa)-(thnSekarang*365+blnSekarang*30+tglSekarang) {
					barang[j+1] = barang[j]
					j--
				}
			}
			barang[j+1] = key
		}
		tampilkanItem()
		fmt.Println("Data berhasil diurutkan berdasarkan jumlah.")
	} else if pilih == 3 {
		InsertionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}
func urutjumlahSLC() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Print("Pilih opsi  : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
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
	} else if pilih == 3 {
		SelectionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}

func ururtNamaSLC() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. Kembali")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
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
	} else if pilih == 3 {
		SelectionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}
func urutSisaHariSLC() {
	hapus()
	var pilih int
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Print("Pilih Opsi : ")
	fmt.Scan(&pilih)
	if pilih == 1 || pilih == 2 {
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
	} else if pilih == 3 {
		SelectionSort()
	} else {
		fmt.Println("opsi tidak valid")
		fmt.Println("pilih opsi dari 1-3")
	}
}
func bahanDigunakan() {
	hapus()
	var pilih, idx int

	n := banyakData
	if n == 0 {
		fmt.Println("Tidak ada data yang tersedia")
		return
	}

	tampilkanItem()
	fmt.Println("Data Bahan Berapa yang ingin di hapus? (1 -", n, ")")
	fmt.Println("0. kembali")
	fmt.Scan(&pilih)
	if pilih < 0 || pilih > n {
		fmt.Println("Data Tidak Tersedia")
		return
	} else if pilih == 0 {
		hapus()
		return
	} else {

		idx = pilih - 1

		for i := idx; i < n-1; i++ {
			barang[i] = barang[i+1]
		}
		banyakData--
		fmt.Println("Data Berhasil Dihapus")
	}
}
func headerTable() {
	fmt.Println("=======================================================================================================")
	fmt.Printf("%-3s | %-14s | %-10s | %-12s | %-13s | %-10s  | %-11s | %-26s |\n", "No", "Nama Bahan", "Jumlah", "Satuan", "Kategori", "Tanggal kadaluarsa", "Sisa hari", "pemberitahuan")
	fmt.Println("-------------------------------------------------------------------------------------------------------")
}
func isiTableCariBahan(barang item, i int, sisaHari int, pemberitahuan string, cek int) {
	fmt.Printf("%-3d | %-14s | %-10d | %-12s | %-13s | %02d/%02d/%-4d          | %-11d | %-26s |\n",
		cek+1, barang[i].nama, barang[i].jumlah, barang[i].satuan, barang[i].jenis,
		barang[i].tanggalkadaluarsa, barang[i].bulankadaluarsa, barang[i].tahunkadaluarsa, sisaHari, pemberitahuan)
}
func cariBahan() {
	hapus()
	var nama, satuan, kategori, pemberitahuan string
	var jumlah, pilih, tanggal, bulan, tahun, pilihOp int
	var cek int = 0
	if banyakData == 0 {
		fmt.Print("Data Tidak Tersedia ")
		fmt.Print("Tambahkan bahan terlebih dahulu")
		tunggu()
		return
	} else {
		fmt.Println("cari bahan berdasarkan")
		fmt.Println("1. nama")
		fmt.Println("2. jumlah")
		fmt.Println("3. satuan")
		fmt.Println("4. kategori")
		fmt.Println("5. tanggal kadaluarsa")
		fmt.Println("6. kembali ke menu utama")
		fmt.Print("Pilih Opsi: ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			hapus()
			fmt.Print("masukan nama bahan yang mau dicari : ")
			fmt.Scan(&nama)
			for i := 0; i < banyakData; i++ {
				if barang[i].nama == nama {
					if cek == 0 {
						headerTable()
					}

					HitungSisaHari(barang, i, &sisaHari)
					pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
					isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
					cek += 1
				}
			}
			if cek == 0 {
				fmt.Println("Bahan tidak ditemukan")
			}
		} else if pilih == 2 {
			hapus()
			fmt.Print("masukan jumlah bahan yang mau dicari : ")
			fmt.Scan(&jumlah)
			for i := 0; i < banyakData; i++ {
				if barang[i].jumlah == jumlah {
					if cek == 0 {
						headerTable()
					}

					HitungSisaHari(barang, i, &sisaHari)
					pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
					isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
					cek += 1
				}
			}
			if cek == 0 {
				fmt.Println("Bahan tidak ditemukan")
			}
		} else if pilih == 3 {
			hapus()
			fmt.Print("masukan jumlah bahan yang mau dicari : ")
			fmt.Scan(&satuan)
			for i := 0; i < banyakData; i++ {
				if barang[i].satuan == satuan {
					if cek == 0 {
						headerTable()
					}

					HitungSisaHari(barang, i, &sisaHari)
					pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
					isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
					cek += 1
				}
			}
			if cek == 0 {
				fmt.Println("Bahan tidak ditemukan")
			}
		} else if pilih == 4 {
			hapus()
			fmt.Print("masukan jumlah bahan yang mau dicari : ")
			fmt.Scan(&kategori)
			for i := 0; i < banyakData; i++ {
				if barang[i].jenis == kategori {
					if cek == 0 {
						headerTable()
					}

					HitungSisaHari(barang, i, &sisaHari)
					pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
					isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
					cek += 1
				}
			}
			if cek == 0 {
				fmt.Println("Bahan tidak ditemukan")
			}
		} else if pilih == 5 {
			hapus()
			fmt.Println("Cari berdasarkan :")
			fmt.Println("1. tanggal")
			fmt.Println("2. bulan")
			fmt.Println("3. tahun")
			fmt.Print("pilih opsi : ")
			fmt.Scan(&pilihOp)
			if pilihOp == 1 {
				hapus()
				fmt.Print("masukan tanggal kadaluarsa bahan yang mau dicari : ")
				fmt.Scan(&tanggal)
				for i := 0; i < banyakData; i++ {
					if barang[i].tanggalkadaluarsa == tanggal {
						if cek == 0 {
							headerTable()
						}

						HitungSisaHari(barang, i, &sisaHari)
						pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
						isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
						cek += 1
					}
				}
				if cek == 0 {
					fmt.Println("Bahan tidak ditemukan")
				}
			} else if pilihOp == 2 {
				hapus()
				fmt.Print("masukan tanggal kadaluarsa bahan yang mau dicari : ")
				fmt.Scan(&bulan)
				for i := 0; i < banyakData; i++ {
					if barang[i].bulankadaluarsa == bulan {
						if cek == 0 {
							headerTable()
						}

						HitungSisaHari(barang, i, &sisaHari)
						pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
						isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
						cek += 1
					}
				}
				if cek == 0 {
					fmt.Println("Bahan tidak ditemukan")
				}
			} else if pilihOp == 3 {
				hapus()
				fmt.Print("masukan tanggal kadaluarsa bahan yang mau dicari : ")
				fmt.Scan(&tahun)
				for i := 0; i < banyakData; i++ {
					if barang[i].tahunkadaluarsa == tahun {
						if cek == 0 {
							headerTable()
						}

						HitungSisaHari(barang, i, &sisaHari)
						pemberitahuanKadaluarsa(barang, i, &pemberitahuan)
						isiTableCariBahan(barang, i, sisaHari, pemberitahuan, cek)
						cek += 1
					}
				}
				if cek == 0 {
					fmt.Println("Bahan tidak ditemukan")
				}
			}
		} else if pilih == 6 {
			hapus()
			tunggu()
			return
		} else {
			fmt.Println("opsi tidak valid")
		}
	}
	tunggu()
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
