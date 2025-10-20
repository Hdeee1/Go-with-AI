1. Goroutine "./goroutine.go"
- Goroutine adalah sebuah fungsi yang berjalan secara independen dan bersamaan dengan fungsi lainnya
- Dimulai dengan kata kunci go di depan pemanggilan fungsi
- Sangat efisien dan murah. Bisa menjalankan Goroutine tanpa masalah.
- Penggunaan di proyek nyata, melakukan banyak panggilan API atau query database secara bersamaan tanpa harus mengunggu satu persatu selesai.

2. Channel "./channel.go"
- Channel adalah jalur komunikasi yang aman antar Goroutine
- Mengirim data ke Channel menggunakan operator <- (misal: ch <- data)
- Menerima data dari Channel juga menggunakan <- (misal: data := <- ch)
- Proses pengiriman dan permintaan akan memblokir sampai pasangannya siap. Ini adalah mekanisme singkronisasi yang kuat
- Penggunaan di proyek nyata, sebuah Goroutine mengambil data dari antrean pesan (seperti RabbitMQ) dan mengirimkannya melalui channel ke sekelompok Goroutine lain yang bertugas memproses data tersebut.

3. Pointer "./pointer.go"
- Gunakan & untuk mendpatkan alamat sebuah variable (misal: &namaVariable)
- Gunakan * untuk mendapatkan nilai yang ada di dalam alamat tersebut (disebut dereferencing)
- Sangat berguna untuk mengubah nilai variable asli disebuah fungsi
- Penggunaan di peoyek nyata, saat mengupdate data di sebuah fungsi (misal: funsdi UpdateUser), kamu mengirim pointer ke struct User agar perubahannya lansung terjadi pada data asli bukan pada salinannya

