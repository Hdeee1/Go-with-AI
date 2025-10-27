package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var JWT_SECRET = []byte("my_secret_key_ku")
var DB *gorm.DB

func connectToDatabase() {
	var err error
	
	DB, err = gorm.Open(sqlite.Open("todolist.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	fmt.Println("Success to connect the database")
}

type User struct {
	Name		string
	Email		string	`json:"email" gorm:"unique"`	
	Password	string	`json:"-"`
	Tasks		[]Task
	gorm.Model
}

type Task struct {
	Title	string
	IsDone	bool
	UserID	uint
	gorm.Model
}

func runMigrate() {
	
	err := DB.AutoMigrate(&User{}, &Task{})
	if err != nil {
		log.Fatal("Failed to migrate:", err)
	}
	fmt.Println("Success to migrate")
}

func createTask(c *gin.Context) {
	var input struct {
		Title	string	`json:"title" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")

	task := Task{Title: input.Title, UserID: userID.(uint), IsDone: false}
	DB.Create(&task)

	c.JSON(http.StatusOK, task)
}

func getAllTasks(c *gin.Context) {
	var tasks []Task
	userID, _ := c.Get("userID") 

	DB.Where("user_id = ?", userID.(uint)).Find(&tasks)
	c.JSON(http.StatusOK, tasks)
}

func getTaskByID(c *gin.Context) {
	var task Task

	taskID := c.Param("id")
	userID, _ := c.Get("userID")

	
	if err := DB.Where(" id = ? AND user_id = ?", taskID, userID.(uint)).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func updateTask(c *gin.Context) {
	var task Task
	userID, _ := c.Get("userID")
	taskID := c.Param("id")
	
	if err := DB.Where("id = ? AND user_id = ?", taskID, userID.(uint)).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input struct {
		Title  string `json:"title"`
		IsDone *bool  `json:"is_done"` 
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	DB.Model(&task).Updates(Task{Title: input.Title, IsDone: *input.IsDone})
	c.JSON(http.StatusOK, task)
}

func deleteTask(c *gin.Context) {
	var task Task
	userID, _ := c.Get("userID")
	taskID := c.Param("id")
	
	result := DB.Where("id = ? AND user_id = ?", taskID, userID.(uint)).Delete(&task)
	
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found or you don't have an access"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func registerUser(c *gin.Context) {
	var input struct {
		Name 		string	`json:"name" binding:"required"`
		Email 		string	`json:"email" binding:"required"`
		Password	string	`json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hashing password"})
		return
	}

	user := User{Name: input.Name, Email: input.Email, Password: string(hashedPassword)}
	result := DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"email": "Email is already exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully created"})
}

func loginUser(c *gin.Context) {
	var input struct {
		Email		string `json:"email" binding:"required"`
		Password	string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			return 
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header)
			}
			return JWT_SECRET, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Ekstarak user_id dari token dan simpan di context gin agar bisa digunakan oleh handler selanjutnya.
			userID := uint(claims["user_id"].(float64))
			c.Set("userID", userID)
		}  else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid", "detail": err.Error()})
			c.Abort()
			return 
		}
		c.Next()
	}
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

	// Grup untuk route publik (tidak butuh token)
	public := router.Group("/auth")
	{
		public.POST("/register", registerUser)
		public.POST("/login", loginUser)
	}

	// Grup untuk route yang dilindungi (butuh token)
	protected := router.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.POST("/tasks", createTask)
		protected.GET("/tasks", getAllTasks)
		// Tambahkan route lain yang dilindungi di sini...
        protected.GET("/tasks/:id", getTaskByID)
        protected.PUT("/tasks/:id", updateTask)
        protected.DELETE("/tasks/:id", deleteTask)
	}

	fmt.Println("Server berjalan di http://localhost:8080")
	router.Run(":8080")
}