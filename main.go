package main

import (
	"blog-backend/internal/config"
	"blog-backend/internal/handlers"
	"blog-backend/internal/routes"
	"blog-backend/internal/services"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	client := config.ConnectMongo()
	db := client.Database("blog")

	postService := services.NewPostService(db)
	postHandler := handlers.NewPostHandler(postService)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "https://raghavm.dev"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	s
	routes.RegisterPostRoutes(r, postHandler)

	r.Run(":8080")
}
