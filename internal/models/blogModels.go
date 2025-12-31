package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title         string             `bson:"title" json:"title"`
	ContentMD     string             `bson:"contentMD" json:"contentMD"`
	DatePublished time.Time          `bson:"datePublished" json:"datePublished"`
	Tags          []string           `bson:"tags" json:"tags"`
	ImageLink     string             `bson:"imageLink" json:"imageLink"`
}
