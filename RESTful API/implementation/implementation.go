package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// model data
type Task struct {
	ID			string		`json:"id"`
	Title		string		`json:"title" binding:"required"`
	IsDone		bool		`json:"is_done"`
	CreatedAt	time.Time	`json:"created_at"`
}

// "Database" in-memory
var tasks = []Task{}

// Handle untuk mmembuat task bare (CREATE)
func createTask(c *gin.Context) {
	var newTask Task

	// Bind JSON dari request ke struct newTask
	// Jika ada error (misal: title kosong), kirim bad request
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set nilai default untuk task baru
	newTask.ID = uuid.New().String()
	newTask.IsDone = false
	newTask.CreatedAt = time.Now()

	// Tambahkan task baru ke database
	tasks = append(tasks, newTask)

	// Kirim response success dengan data task yang baru dibuat
	c.JSON(http.StatusCreated, newTask)
}

// Handler untuk mendapatkan semua task (READ ALL)
func getAllTasks(c *gin.Context) {
	// kirim seluruh slice 'tasks' sebagai json
	c.JSON(http.StatusOK, tasks)
}

// Handler untuk mendapatkan satu task berdasarkan ID (READ ONE)
func getTaskByID(c *gin.Context) {
	// Ambil ID dari parameter
	taskID := c.Param("id")

	// Cari task di dalam slice
	for _, task := range tasks {
		if task.ID == taskID {
			c.JSON(http.StatusOK, task)
			return // hentikan fungsi jika ditemukan
		}
	}

	// Jika loop selesai dan task tidak ditemukan, kirim error
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Handler untuk updata task
func updateTask(c *gin.Context) {
	taskID := c.Param("id")

	var updateTask Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari index dari task yang akan diupdate (UPDATE)
	for i, task := range tasks {
		if task.ID == taskID {
			// update field yang relefan
			tasks[i].Title = updateTask.Title
			tasks[i].IsDone =updateTask.IsDone
			c.JSON(http.StatusOK, tasks[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Handler untuk delete task (DELETE)
func deleteTask(c *gin.Context) {
	// cari id
	taskId := c.Param("id")

	// cari index dari task yang akan dihapus
	for i, task := range tasks {
		if task.ID == taskId {
			// hapus element dari slice dengan menggabungkan task
			// sebelum dan sesudah element dihapus
			tasks = append(tasks[:i], tasks[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// Auth middleware
func AuthMiddleware() gin.HandlerFunc {
	return  func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY")
		if apiKey != "rahasia dunia" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API Key not valid"})
			c.Abort()
			return 
		}

		c.Next()
	}
}

func main() {
	// Inisialisasi beberapa data awal
	tasks = append(tasks, Task{ID: uuid.New().String(), Title: "Learning Golang", IsDone: true, CreatedAt: time.Now()})
	tasks = append(tasks, Task{ID: uuid.New().String(), Title: "CREATE API CRUD", IsDone: true, CreatedAt: time.Now()})

	router := gin.Default()

	api := router.Group("/api/v1")
	api.Use(AuthMiddleware())
	{
		api.POST("/tasks", createTask)
		api.GET("/tasks", getAllTasks)
		api.GET("/tasks/:id", getTaskByID)
		api.PUT("/tasks/:id", updateTask)
		api.DELETE("/tasks/:id", deleteTask)
	}

	fmt.Println("server running at http://localhost:8080")
	router.Run(":8080")
}
