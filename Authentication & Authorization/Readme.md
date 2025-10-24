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

2. Middleware "./middleware.go"
- Poin Penting:
    - Ini adalah implementasi prinsip DRY (Don't Repeat Yourself). Kamu tidak perlu menulis logika verifikasi token di setiap handler. Cukup terapkan satu middleware ini ke grup route yang ingin kamu lindungi.
    - Memisahkan logika keamanan dari logika bisnis. Handler-mu (misal: getProfile) bisa fokus pada tugasnya mengambil data profil, karena ia bisa berasumsi bahwa request yang sampai padanya sudah pasti terautentikasi.
    - Tujuan Proteksi Route: Kamu akan membuat sebuah AuthMiddleware dan menerapkannya pada grup route seperti /api/v1/tasks atau /api/v1/profile, sementara route seperti /login dan /register dibiarkan publik (tidak memakai middleware ini).

3. Session Handling
- Poin Penting:
    - Stateful: Server harus menyimpan data sesi untuk setiap pengguna yang aktif. Ini bisa menjadi tantangan saat scaling karena setiap server harus punya akses ke penyimpanan sesi yang sama.
    - Klien Hanya Menyimpan ID: Klien (browser) hanya menyimpan ID sesi yang tidak berarti apa-apa (biasanya dalam bentuk cookie), bukan data pengguna. Ini sedikit lebih aman jika cookie dicuri.
    - Invalidasi Sesi: Karena sesi disimpan di server, kamu punya kontrol penuh. Kamu bisa secara paksa me-logout pengguna kapan saja dengan menghapus data sesinya dari server. Hal ini lebih sulit dilakukan dengan JWT.
    - Perbandingan: JWT lebih modern dan cocok untuk API stateless dan arsitektur microservices. Session lebih tradisional, cocok untuk aplikasi web monolitik di mana kamu butuh kontrol penuh untuk me-logout pengguna secara paksa dari sisi server.

Kesimpulan untuk Membangun Sistem Login
1.Buat endpoint /register:
- Terima name, email, password dari klien.
- Hash password-nya! Jangan pernah simpan password sebagai teks biasa. Gunakan library seperti bcrypt.
- Simpan user baru ke database.

2. Buat endpoint /login:
- Terima email dan password.
- Cari user di DB berdasarkan email.
- Bandingkan password yang masuk dengan hash yang ada di DB menggunakan fungsi bcrypt.CompareHashAndPassword.
- Jika cocok, generate JWT yang berisi user_id dan exp.
- Kirim JWT ini kembali ke klien.

3. Buat AuthMiddleware:
- Implementasikan logika untuk memeriksa dan memvalidasi JWT dari Authorization header.
- Ekstrak user_id dari token dan simpan di context Gin (c.Set("userID", claims.UserID)) agar bisa diakses oleh handler.

4. Terapkan Middleware:
- Grup semua route yang butuh login (/profile, /tasks, dll.) dan terapkan AuthMiddleware pada grup tersebut.