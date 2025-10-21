package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ini adalah fungsi handler yang akan dipanggil oleh router
func getAllUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "This is a list of all users"})
}

func createUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "New user successfully created"})
}

func getUserByID(c *gin.Context) {
	// menangkap parameter id
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "These are the details for the user" + id})
}

func main() {
	// 1. Buat router
	router := gin.Default()

	// 2. Definisikan semua router 
	router.GET("/users", getAllUsers)			//Read all
	router.POST("/users", createUser)			//Create
	router.GET("/users/:id", getUserByID)		//Read one
	// router.PUT("/users/:id", updateUser)		//Update
	// router.DELETE("/users/:id", deleteUser)	//Delete

	// 3. Jalankan server
	router.Run(":8080")	//server berjalan di port 8080
}