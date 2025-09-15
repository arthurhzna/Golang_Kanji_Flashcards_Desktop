# Japanese Flashcard Application

Aplikasi flashcard sederhana untuk belajar kosakata bahasa Jepang dengan tampilan GUI menggunakan Fyne.

## Fitur

- 📚 Menampilkan kosakata Jepang (Kanji, Hiragana, Bacaan, Arti)
- ⏭️ Navigasi maju/mundur antar kartu
- 👁️ Toggle untuk menyembunyikan/menampilkan bacaan dan arti
- 📱 Interface GUI yang simpel dan mudah digunakan
- 📊 Membaca data dari file Excel (.xlsx)

## Prasyarat

### Windows
1. **Install GCC Compiler**
   - Download dan install TDM-GCC dari: https://jmeubank.github.io/tdm-gcc/
   - Atau install MinGW-w64

2. **Install Go**
   - Download dari: https://golang.org/dl/
   - Versi minimum: Go 1.16+

## Instalasi

1. **Clone atau download project ini**
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Siapkan file data.xlsx** dengan format:
   | Kanji | Hiragana | Bacaan | Arti |
   |-------|----------|--------|------|
   | 私 | わたし | watashi | saya |
   | 本 | ほん | hon | buku |

## Cara Menjalankan

### Menjalankan langsung
```bash
go run main.go
```

### Compile menjadi executable
```bash
# Untuk Windows
go build -o flashcard.exe main.go

# Untuk sistem lain
go build -o flashcard main.go
```

Setelah di-compile, Anda tinggal double-click file `flashcard.exe` untuk menjalankannya.

### Cross-compile untuk OS lain
```bash
# Untuk Linux dari Windows
set GOOS=linux
set GOARCH=amd64
go build -o flashcard-linux main.go

# Untuk macOS dari Windows  
set GOOS=darwin
set GOARCH=amd64
go build -o flashcard-mac main.go
```

## Penggunaan

1. Pastikan file `data.xlsx` ada di folder yang sama dengan executable
2. Jalankan aplikasi
3. Gunakan tombol `<` dan `>` untuk navigasi
4. Klik tombol "Hide" untuk menyembunyikan bacaan atau arti
5. Pelajari kosakata dengan cara menyembunyikan jawaban terlebih dahulu

## Format File Excel

File `data.xlsx` harus memiliki struktur:
- **Sheet1** dengan kolom:
  - Kolom A: Kanji
  - Kolom B: Hiragana  
  - Kolom C: Bacaan (romaji)
  - Kolom D: Arti (bahasa Indonesia)
- Baris pertama sebagai header (akan diabaikan)

## Troubleshooting

### Error "gcc not found"
- Install TDM-GCC atau MinGW-w64
- Pastikan gcc ada di PATH

### Error "cannot find data.xlsx"
- Pastikan file data.xlsx ada di folder yang sama dengan executable
- Periksa format dan isi file Excel

### GUI tidak muncul
- Pastikan semua dependencies Fyne terinstall
- Coba jalankan: `go mod tidy`

## Dependencies

- `github.com/xuri/excelize/v2` - untuk membaca file Excel
- `fyne.io/fyne/v2` - untuk GUI framework

## Kontribusi

Silakan buat issue atau pull request untuk perbaikan dan penambahan fitur.

## Lisensi

Project ini bebas digunakan untuk tujuan pembelajaran.
