package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 1. Setup dan connect database
// Menyimpan koneksi DB di variable global agar bisa diakses semua handler
var DB *gorm.DB

func connectToDatabase() {
	var err error
	//GORM akan otomatis membuat file 'todolist.db' jika belum ada
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Success to connect the database")
}

// 2. Definisikan Model dan jalankan migrasi
// Model user
type User struct {
	Name	string
	Email	string
	Tasks	[]Task
	gorm.Model
}

type Task struct {
	Title	string
	IsDone	bool
	UserID	uint
	gorm.Model
}

func runMigrate() {
	// AutoMigrate akan menambahkan table sesuai struct User dan Task
	err := DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}
	fmt.Println("Success to migrate")
}

// 3. Buat handler CRUD dengan GORM
// Handler untuk membuat task baru CREATE
func createTask(c *gin.Context) {
	var input struct {
		Title	string	`json:"title" binding:"required"`
		UserID	uint	`json:"user_id" binding:"required"`	
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	// Buat task baru 
	task := Task{Title: input.Title,  UserID: input.UserID, IsDone: false}
		

	// simpan ke database menggunakan Gorm
	result := DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error to save task"})
		return
	}

	c.JSON(http.StatusCreated, task)
}

// handler untuk mendapatkan semua task (READ)
func getAllTasks(c *gin.Context) {
	var tasks []Task

	// Ambil semua data dari table tasks
	DB.Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

//  Handler untuk mendapatkan task berdasarkan id
func getTaskByID(c *gin.Context) {
	var task Task

	taskID := c.Param("id")

	// Cari task dengan ID tersebut, jika tidak ada, first akan mengembalikan error
	if err := DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Handler untuk mengubah task (UPDATE)
func updateTask(c *gin.Context) {
	var task Task
	taskID := c.Param("id")
	
	if err := DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input struct {
		Title  string `json:"title"`
		IsDone *bool  `json:"is_done"` // Gunakan pointer agar bisa membedakan 'false' dan 'tidak diisi'
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update task menggunakan Model().Updates()
	DB.Model(&task).Updates(Task{Title: input.Title, IsDone: *input.IsDone})
	c.JSON(http.StatusOK, task)
}

// Handler untuk menghapus task (DELETE)
func deleteTask(c *gin.Context) {
	var task Task
	taskID := c.Param("id")
	
	// Menggunakan 'Unscoped' untuk benar-benar menghapus dari DB, bukan soft delete.
	if err := DB.Unscoped().Delete(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func registerUser(c *gin.Context) {
	var input struct {
		Name 	string	`json:"name" binding:"required"`
		Email 	string	`json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := User{Name: input.Name, Email: input.Email}

	// Mulai transaksi
	err := DB.Transaction(func(tx *gorm.DB) error {
		// 1. Buat User di dalam transaksi
		if err := tx.Create(&newUser).Error; err != nil {
			return err
		}

		// 2. Buat task pertama di dalam transaksi 
		firstTask := Task{Title: "Complete your profile!", UserID: newUser.ID, IsDone: false, }
		if err := tx.Create(&firstTask).Error; err != nil {
			// Jika gagal daalm pembuatan user diatas akan otomatis di rollback
			return err
		}

		// Jika semua berhasil, return nil untuk commit transaksi 
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully created", "user": newUser})
}

func getUserWithTasks(c *gin.Context) {
	var user User
	userID := c.Param("id")

	if err := DB.Preload("Tasks").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func main() {
	connectToDatabase()

	runMigrate()

	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/users", registerUser)
		api.GET("/users/:id", getUserWithTasks)
		
		api.POST("/tasks", createTask)
		api.GET("/tasks", getAllTasks)
		api.GET("/tasks/:id", getTaskByID)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}

	fmt.Println("server run at http://localhost:8080")
	router.Run(":8080")
}