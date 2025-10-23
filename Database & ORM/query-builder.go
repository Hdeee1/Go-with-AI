// Contoh di dalam handler Gin/Fiber
package main

// READ: Mencari semua task yang belum selesai
var inCompleteTask []Task
result = db.Where("is_done = ?", false).Find(&inCompleteTask)

if result.Error != nil {
	// Handle error
}

// CREATE : Membuat user baru 
newUser := User{Name: "nama", age: 0}
result = db.Create(&newUser) // ID otomatis diisi oleh gorm

// UPDATE : Menandai task sebagai selesai
var task Task
db.First(&task, "some-id-from-url") // Cari task dulu
db.Model(&task).Update("is_done", true)