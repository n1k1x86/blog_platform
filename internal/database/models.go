package database

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleDB struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Tag         string             `bson:"tag"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

type ArticleJSON struct {
	ID          string
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Tag         string    `json:"tag"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NewArticle struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
}

func CreateNewArticle(a *NewArticle) *ArticleDB {
	return &ArticleDB{
		ID:          primitive.NewObjectID(),
		Title:       a.Title,
		Description: a.Description,
		Tag:         a.Tag,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func FromJSONToDB(id string, article *ArticleJSON) (*ArticleDB, error) {
	return &ArticleDB{
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
