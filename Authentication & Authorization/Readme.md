1. JWT (Json Web Token)
- Header: Berisi informasi tentang token itu sendiri (misal: algoritma enkripsi yang dipakai)
- Payload: Berisi "klaim" atau data tentang pengguna (misal: user_id, username, role: admin, dan waktu kadaluarsa token). Ini tidak di enkripsi, jadi jangan taruh data sensitif seperti password
- Signature: Tanda tangan digital yang dibuat dari header, payload, dan sebuah kunci rahasia (secret key) yang hanya diketahui oleh server. Ini adalah bagian terpenting untuk keamanan
- Point penting:
    a. Stateless: Server tidak perlu menyimpan informasi token di sisinya. Semua data yang dibutuhkan ada di dalam token itu sendiri. Ini membuat API sangat mudah di scale
    b. Verivikasi Tanda Tangan: Saat server menerima JWT dari klien, ia akan membuat ulang tanda tangan menggunakan header, payload, dan secret key yang ia miliki. Jika tanda tangan yang dibuat ulang cocok dengan yang ada di token, berarti token itu asli dan tidak dirubah di tengah jalan
    c. Kadaluarsa: JWT wajib memiliki waktu kadaluarsa (exp) untuk membatasi masa yang berlaku jika token dicuri
    d. Tujuan Sistem Login: Saat user berhasil login, server akan membuat (generate) sebuah JWT dan mengirimkannya kembali ke klient (misal: aplikasi web/mobile). Klien kemudian menyimpan token ini (misal: di Local Storage atau Cookie) 
- Contoh Alur Kerja
    1) User mengirim email & password ke endpoint "/login"
    2) Server memeriksa email & password ke database
    3) Jika valid, Server membuat JWT yang berisi user_id dan exp (misal: berlaku 1 jam), lalu menandatanganinya dengan secret key
    4) Server mengirim JWT kembali ke User
    5) Untuk request selanjutnya (misal: ke "/profile"), User menyertakan JWT ini di header Authorization (contoh: Authorization: Bearer <token_panjang_jwt>)
    6) Server menerima request, memverifikasi tanda tangan JWT, dan jika valid, mengizinkan akses ke "/profile"

