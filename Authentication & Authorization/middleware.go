package main

import "github.com/gin-gonic/gin"

func AuthMiddleware() gin.HandlerFunc {
	return  func(ctx *gin.Context) {
		// 1. Ambil token dari header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(401, gin.H{"error": "Request not contain an access token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func main() {
	router := gin.Default()

	// Router public
	router.POST("/register", registerHandler)
	router.POST("/login", loginHandler)

	// Group router yang dilindungi
	protected := router.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.GET("/profile", profileHandler)
		protected.GET("/tasks", getTaskHandler)
	}

	router.Run()
}