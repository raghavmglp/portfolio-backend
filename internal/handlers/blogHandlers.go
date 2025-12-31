package handlers

import (
	"blog-backend/internal/models"
	"net/http"
	"strconv"

	"blog-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	Service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{Service: service}
}

func (h *PostHandler) ListPosts(c *gin.Context) {
	limit := int64(10)
	if l := c.Query("limit"); l != "" {
		if val, err := strconv.ParseInt(l, 10, 64); err == nil {
			limit = val
		}
	}

	filter := c.Query("filter")
	sortField := c.DefaultQuery("sort", "datePublished")
	sortOrder := 1
	if c.Query("order") == "desc" {
		sortOrder = -1
	}

	posts, err := h.Service.ListPosts(limit, filter, sortField, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := h.Service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post models.Post

	// 1. Parse the incoming JSON
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 2. Call the service to save it
	if err := h.Service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// 3. Return success
	c.JSON(http.StatusCreated, post)
}
