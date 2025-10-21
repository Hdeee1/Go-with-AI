1. Routing "./routing.go"
- Routing adalah proses memetakan sebuah Method HTTP + URL Path ke sebuah Fungsi atau (handler)
- Ini adalah inti atau jantung dari API, tanpa Routing, semua request akan tersesat
- Framework seperti Gin membuat ini sangat mudah dan deklaratif
- Bisa menangkap bagian dinamis dari URL (seperti 123) menggunakan path parameter (ditulis :id di Gin)
- Tujuan CRUD, Routing adalah cara mendefinisikan kelima endpoint dasar CRUD (Create, Reade, Update, Delete)

2. JSON Handling "./json-handling.go"
- Membaca JSON (Request Biding): Saat klien (misal: frontend) ingin membuat user baru, mereka akan mengirim data JSON di body request POST. Tugas backend adalah mengikat (bind) data JSON yang masuk itu ke dalam struct Go agar mudah di olah
- Menulis JSON (Response): Saat backend ingin mengirim data ke klien (misal: detail user), backend tidak mengirim struct Go. Backend menyuruh Gin untuk merubah struct atau map menjadi format JSON, lalu mengirimnya sebagai balasan.

3. Point penting:
- Backend wajib membuat struct untuk mempresentasikan data yang diterima atau kirim. Contoh, struct User, sturct Product
- Gunakan struct tags (seperti `json:"nama_field"`) untuk mengontrol nama field di JSON
- Binding adalah proses validasi dan konversi otomatis dari JSON request ke struct GO. Ini menghemat banyak kode
- Tujuan CRUD, Create dan Update butuh membaca JSON. Read (satu atau semua) butuh menulis JSON