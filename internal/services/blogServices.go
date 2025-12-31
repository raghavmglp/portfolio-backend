package services

import (
	"context"
	"time"

	"blog-backend/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostService struct {
	Collection *mongo.Collection
}

func NewPostService(db *mongo.Database) *PostService {
	return &PostService{
		Collection: db.Collection("posts"),
	}
}

func (s *PostService) ListPosts(limit int64, filterTitle, sortField string, sortOrder int) ([]models.Post, error) {
	filter := bson.M{}
	if filterTitle != "" {
		filter["title"] = bson.M{"$regex": filterTitle, "$options": "i"}
	}

	opts := options.Find()
	opts.SetLimit(limit)
	opts.SetSort(bson.D{{Key: sortField, Value: sortOrder}})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var posts []models.Post
	if err := cursor.All(ctx, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPostByID(id string) (*models.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var post models.Post
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = s.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (s *PostService) CreatePost(post *models.Post) error {
	// 1. Set default values
	post.ID = primitive.NewObjectID()
	post.DatePublished = time.Now()

	// 2. Insert into MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.Collection.InsertOne(ctx, post)
	return err
}
