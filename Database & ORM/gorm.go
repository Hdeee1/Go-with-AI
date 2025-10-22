package main

import "gorm.io/gorm"


// struct User akan menjadi table users
type User struct {
	Name	string
	Email	string	`gorm:""unique`
	Tasks	[]Task	// Relasi one to many: satu user bisa memiliki banyak task
	gorm.Model		// Otomatis menambahkan ID, CreatedAt, UpdateAt, DeleteAt
}

// struct Task akan menajadi table tasks
type Task struct {
	Title	string
	IsDone	bool
	UserID	uint	// Ini adalah foreign key yang menghubungkan ke User
	gorm.Model
}