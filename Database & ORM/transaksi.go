// Membuat User baru berserta Task pertamanya dalam satu transaksi 
err := db.Transaction(func(tx *gorm.DB) error {
	// 1. Buat user baru
	newUser := User{Name: "nama", Email: "email@g.c"}

	if err:= tx.Create(&newUser).Error; err != nil {
		// Jika gagal, kembalikan error. Transaksi akan dirollback
		return err
	}

	// 2. Buat Task pertamanya, milik user yang baru dibuat
	firstTask := Task{Title: "Task", UserID: newUser.ID}
	if err := tx.Create(&firstTask).Error; err != nil {
		// Jika gagal, pembuatan user diatas juga gagal
		return err
	}
})

if err != nil {
	// Handler error transaksi disini
	//  Kita bisa yakin bahwa tidak ada "user tanpa task pertama" di database
}