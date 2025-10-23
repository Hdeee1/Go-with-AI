1. Gorm "./gorm.go"
- Point penting :
    - Gorm memungkinkan mendefinisikan relasi langsung ke dalam struct
    - Jenis relasi paling umum adalah One-to-Many, satu User memiliki banyak Task
    - Hubungan ini disebut foreign key (Kunci asing), yaitu kolom di satu table (misal, UserID di table Tasks) yang merujuk ke primary key di table lain (kolom id di table users)
    - Tujuan CRUD, dengan relasi, kita bisa melakukan query yang powerful, seperti "ambil semua tugas milik Budi" atau "saat menghapus user Budi, hapus juga semua tugasnya"

2. Migrasi "./migrasi.go"
- Migrasi adalah blue print versi terkontrol untuk database
- Setiap mengubah struct (misal: menambahkan kolom Priority di Task), cukup jalankan migrasi lagi, dan Gorm akan menambah kolom tanpa menghapus data yang sudah ada
- Ini memastikan struktur database selalu sinkron dengan aplikasi
- Tujuan CRUD, bisa melakukan CRUD jika tablenya bahkan belum ada. Migrasi adalah langkah nol yang wajib dilakukan sebelum API-mu bisa berinteraksi dengan database

3. Query Builder
- Point penting :
    - Bekerja dengan struct dan method Go, bukan string. Ini mengurangi typo dan error
    - Method-methodnya bisa dirangkai (chaining), membuat query kompleks jadi terbaca, `db.Where(...).Order(...).Limit(...).Find(...)`
    - Secara otomatis melindungi aplikasi dari SQL Injection
    - Tujuan CRUD: 
        - Create: db.Create(&newUser)
        - Read (all): db.Find(&tasks)
        - Read (one): db.First(&task, taskID)
        - Update: db.Model(&task).Updates(updateData)
        - Delete: db.Delete(&task, taskID)