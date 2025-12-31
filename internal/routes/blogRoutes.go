package routes

import (
	"blog-backend/internal/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// 1. Define the Guard Middleware (Checks for your password)
// internal/routes/routes.go

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		expected := "Bearer " + os.Getenv("ADMIN_SECRET")
		fmt.Println("Received token:", token)
		fmt.Println("Expected token:", expected)
		if token != expected {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid or missing API Key"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func RegisterPostRoutes(r *gin.Engine, postHandler *handlers.PostHandler) {
	// === PUBLIC ROUTES (Everyone can read) ===
	r.GET("/posts", postHandler.ListPosts)
	r.GET("/post/:id", postHandler.GetPostByID)

	// === PRIVATE ROUTES (Only you can write) ===
	// We create a protected group using the middleware above
	protected := r.Group("/")
	protected.Use(authMiddleware())
	{
		protected.POST("/posts", postHandler.CreatePost)
	}
}
