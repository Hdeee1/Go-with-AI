package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func loggerMiddleware() gin.HandlerFunc {
	return  func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		latency := time.Since(start)
		fmt.Printf("[LOG] %s | %s | %d | %v\n",
				c.Request.Method,
				path,
				c.Writer.Status(),
				latency,
		)
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-KEY") 

		if apiKey != "Rahasia-negara" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API Key is not valid"})
			c.Abort()
			return 
		}

		c.Next()
	}
}

func main() {
	router := gin.New() // Kita pakai .New() agar bisa pilih middleware

	// 2. Terapkan middleware secara GLOBAL (berlaku untuk semua route)
	router.Use(loggerMiddleware())
	router.Use(gin.Recovery()) // Middleware bawaan Gin untuk pulih dari panic

	// Route publik, tidak perlu auth
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// 3. Grup route yang dilindungi middleware
	api := router.Group("/api")
	api.Use(AuthMiddleware()) // Hanya berlaku untuk /api/*
	{
		// Endpoint ini (/api/users) sekarang dilindungi
		api.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Anda berhasil masuk! Ini data user."})
		})
	}

	router.Run(":8080")

}