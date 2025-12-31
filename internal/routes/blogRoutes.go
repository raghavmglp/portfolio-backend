package routes

import (
	"blog-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

// 1. Define the Guard Middleware (Checks for your password)
// internal/routes/routes.go

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// OLD: adminSecret := os.Getenv("ADMIN_SECRET")

		// NEW: Hardcode it for testing
		adminSecret := "mySuperSecretPassword123"

		clientToken := c.GetHeader("X-Admin-Token")

		if adminSecret == "" || clientToken != adminSecret {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized: Invalid or missing API Key"})
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
