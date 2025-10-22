1. Gorm "./gorm.go"
- Point penting
    - Gorm memungkinkan mendefinisikan relasi langsung ke dalam struct
    - Jenis relasi paling umum adalah One-to-Many, satu User memiliki banyak Task
    - Hubungan ini disebut foreign key (Kunci asing), yaitu kolom di satu table (misal, UserID di table Tasks) yang merujuk ke primary key di table lain (kolom id di table users)
    - Tujuan CRUD, dengan relasi, kita bisa melakukan query yang powerful, seperti "ambil semua tugas milik Budi" atau "saat menghapus user Budi, hapus juga semua tugasnya"

