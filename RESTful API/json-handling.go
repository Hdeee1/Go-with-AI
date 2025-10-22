package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 1. Definisikan struct untuk data kita. Ini adalah blue print
//Konsep struct dan error handling sangat dipakai disini
type User struct {
	ID		string	`json:"id"`
	Name	string	`json:"name" binding:"required"` //binding:"required" berarti field ini wajib ada
	Age		int		`json:"age"`
}

// Anggap saja ini database
var dbUsers = make(map[string]User)

// 2. Membaca JSON (Binding) untuk membuat user baru
func createUser(c *gin.Context) {
	var newUser User

	// Gin akan otomatis membaca body request, memvalidasi, dan mengisi struct 'newUser'
	if err := c.ShouldBindJSON(&newUser); err != nil {
		// Jika JSON nya tidak valid (misal: name kosong), kirim  balasan error. Ini adalah error handling di API
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// (Logika bisnis, misal simpan ke database)
	newUser.ID = "user-0045"	//Anggap dapat ID dari database
	dbUsers[newUser.ID]  = newUser

	// 3. Menulis JSON sebagai balasan sukses
	c.JSON(http.StatusCreated, newUser) //Gin otomatis konversi 'newUser' jadi JSON
}

func main() {
	router := gin.Default()
	router.POST("/users", createUser)
	router.Run(":8080")
}