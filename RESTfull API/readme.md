1. Routing "./routing.go"
- Routing adalah proses memetakan sebuah Method HTTP + URL Path ke sebuah Fungsi atau (handler)
- Ini adalah inti atau jantung dari API, tanpa Routing, semua request akan tersesat
- Framework seperti Gin membuat ini sangat mudah dan deklaratif
- Bisa menangkap bagian dinamis dari URL (seperti 123) menggunakan path parameter (ditulis :id di Gin)
- Tujuan CRUD, Routing adalah cara mendefinisikan kelima endpoint dasar CRUD (Create, Reade, Update, Delete)