package main

import "gorm.io/gorm"


// struct User akan menjadi table users
type User struct {
	Name	string
	Email	string	`gorm:""unique`
	Tasks	[]Task
	gorm.Model
}

// struct Task akan menajadi table tasks
type Task struct {
	Title	string
	IsDone	bool
	UserID	uint
	gorm.Model
}