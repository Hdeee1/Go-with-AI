package main

import (
	"fmt"
	"log"
)

type User struct {}
type Task struct {}

func main() {
	// Contoh kode, bisa berada di fungsi main atau fungsi setup database
	
	// (Kode untuk koneksi ke database, misal: Postgres/MySQL)
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// Menjalankan Auto Migration
	// GORM akan membuat tabel 'users' dan 'tasks' jika belum ada.
	// GORM juga akan menambahkan foreign key constraint secara otomatis.
	
	err = db.AutoMigrate(&User{}, &Task{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}

	fmt.Println("Successfully migrate the database")
}