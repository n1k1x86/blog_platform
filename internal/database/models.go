package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleDB struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Tag         string             `bson:"tage"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type ArticleJSON struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tag         string    `json:"tage"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromJSONToDB(article *ArticleJSON) (*ArticleDB, error) {
	id, err := primitive.ObjectIDFromHex(article.ID)
	if err != nil {
		return nil, err
	}
	return &ArticleDB{
		ID:          id,
		Title:       article.Title,
		Description: article.Description,
		Tag:         article.Tag,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}, nil
}

func FromDBToJSON(article *ArticleDB) *ArticleJSON {
	return &ArticleJSON{
		ID:          article.ID.Hex(),
		Title:       article.Title,
		Description: article.Description,
		Tag:         article.Tag,
		CreatedAt:   article.CreatedAt,
		UpdatedAt:   article.UpdatedAt,
	}
}
