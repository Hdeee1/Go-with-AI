1. Gorm "./gorm.go"
- Point penting
    - Gorm memungkinkan mendefinisikan relasi langsung ke dalam struct
    - Jenis relasi paling umum adalah One-to-Many, satu User memiliki banyak Task
    - Hubungan ini disebut foreign key (Kunci asing), yaitu kolom di satu table (misal, UserID di table Tasks) yang merujuk ke primary key di table lain (kolom id di table users)
    - Tujuan CRUD, dengan relasi, kita bisa melakukan query yang powerful, seperti "ambil semua tugas milik Budi" atau "saat menghapus user Budi, hapus juga semua tugasnya"

2. Migrasi "./migrasi.go"
- Migrasi adalah blue print versi terkontrol untuk database
- Setiap mengubah struct (misal: menambahkan kolom Priority di Task), cukup jalankan migrasi lagi, dan Gorm akan menambah kolom tanpa menghapus data yang sudah ada
- Ini memastikan struktur database selalu sinkron dengan aplikasi
- Tujuan CRUD, bisa melakukan CRUD jika tablenya bahkan belum ada. Migrasi adalah langkah nol yang wajib dilakukan sebelum API-mu bisa berinteraksi dengan database

