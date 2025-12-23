# TUGAS-BESAR-ALPRO-TOPIK-11

## Aplikasi Manajemen Stok Bahan Makanan

Aplikasi command-line sederhana untuk mengelola stok bahan makanan dengan fitur sorting dan pencarian.

### Fitur Utama

1. **Tambah Bahan Makanan** - Menambahkan bahan makanan baru ke dalam sistem
2. **Edit Data** - Mengedit informasi bahan makanan (nama, jumlah, satuan, kategori, tanggal kadaluarsa, hari)
3. **Tampilkan Bahan Makanan** - Menampilkan semua bahan makanan dengan status kadaluarsa
4. **Tampilkan Bahan Makanan (Default)** - Menampilkan bahan makanan terurut berdasarkan tanggal masuk
5. **Urutkan Bahan Makanan** - Mengurutkan bahan dengan Selection Sort atau Insertion Sort
6. **Hapus Bahan Makanan** - Menghapus bahan yang sudah digunakan
7. **Cari Bahan Makanan** - Mencari bahan berdasarkan nama, jumlah, satuan, kategori, atau tanggal kadaluarsa

### Informasi Bahan Makanan

Setiap bahan makanan menyimpan informasi:
- Nama bahan
- Jumlah
- Satuan (kg, g, butir, liter, dll)
- Kategori (sayuran, buah, protein, pemanis, lemak, karbohidrat, dll)
- Tanggal kadaluarsa
- Tanggal masuk

### Sorting Algorithms

- **Selection Sort** - Mengurutkan berdasarkan nama, jumlah, atau sisa hari kadaluarsa
- **Insertion Sort** - Mengurutkan berdasarkan nama, jumlah, atau sisa hari kadaluarsa

### Cara Menjalankan

```bash
go run "tubes 2.go"
```

### Format Input Tanggal

Masukkan tanggal dalam format: `dd mm yyyy` (pisahkan dengan spasi)

Contoh: `25 12 2025`