1. Routing "./routing.go"
- Routing adalah proses memetakan sebuah Method HTTP + URL Path ke sebuah Fungsi atau (handler)
- Ini adalah inti atau jantung dari API, tanpa Routing, semua request akan tersesat
- Framework seperti Gin membuat ini sangat mudah dan deklaratif
- Bisa menangkap bagian dinamis dari URL (seperti 123) menggunakan path parameter (ditulis :id di Gin)
- Tujuan CRUD, Routing adalah cara mendefinisikan kelima endpoint dasar CRUD (Create, Reade, Update, Delete)

2. JSON Handling "./json-handling.go"
- Membaca JSON (Request Biding): Saat klien (misal: frontend) ingin membuat user baru, mereka akan mengirim data JSON di body request POST. Tugas backend adalah mengikat (bind) data JSON yang masuk itu ke dalam struct Go agar mudah di olah
- Menulis JSON (Response): Saat backend ingin mengirim data ke klien (misal: detail user), backend tidak mengirim struct Go. Backend menyuruh Gin untuk merubah struct atau map menjadi format JSON, lalu mengirimnya sebagai balasan.
- Point penting:
    - Backend wajib membuat struct untuk mempresentasikan data yang diterima atau kirim. Contoh, struct User, sturct Product
    - Gunakan struct tags (seperti `json:"nama_field"`) untuk mengontrol nama field di JSON
    - Binding adalah proses validasi dan konversi otomatis dari JSON request ke struct GO. Ini menghemat banyak kode
    - Tujuan CRUD, Create dan Update butuh membaca JSON. Read (satu atau semua) butuh menulis JSON

4. Middleware "./middleware.go"
- Logging: Mencatat siapa saja yang masuk dan jam berapa (logger)
- Authentikasi: Memeriksa kartu identitas atau token (Auth). Jika tidak punya, langsung diusir
- Keamanan: Memeriksa barang bawaan berbahaya (Security Header, CORS)
- Recovery: Jika ada tamu yang panik dan "error" di dialam, satpam ini yang menangani agar gedung tidak ikut "Crash" (panik)
- Point penting: 
    - Middleware adalah fungsi yang dieksekusi sebelum handler utamamu.
    - Sangat penting untuk logika yang ingin diterapkan di banyak route (don't repeat your code!)
    - Gin (gin.Default()) sudah otomatis menyertakan middleware Logger dan Recovery.
    - Kita bisa custom sendiri middleware (misal untuk cek API Key)
    - Tujuan Prodection-Ready, API mu tidak protection-ready tanpa middleware, terutama untuk logging, error handling, dan authentikasi
