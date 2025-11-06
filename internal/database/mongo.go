package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	client *mongo.Client
	dbName string
	coll   string
	ctx    context.Context
}

func (r *Repo) CreateNewArticle(a *NewArticle) (string, error) {
	article := CreateNewArticle(a)
	db := r.client.Database(r.dbName)
	coll := db.Collection(r.coll)
	res, err := coll.InsertOne(r.ctx, article)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *Repo) GetArticleByID(id string) (*ArticleJSON, error) {
	db := r.client.Database(r.dbName)
	coll := db.Collection(r.coll)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": objID,
	}

	res := coll.FindOne(r.ctx, filter)
	if res.Err() != nil {
		return nil, err
	}
	var article ArticleDB
	err = res.Decode(&article)
	if err != nil {
		return nil, err
	}

	return FromDBToJSON(&article), nil
}

func (r *Repo) GetArticlesList(tag string) ([]*ArticleJSON, error) {
	filter := bson.M{}
	if tag != "" {
		filter["tag"] = tag
	}

	db := r.client.Database(r.dbName)
	coll := db.Collection(r.coll)
	articles := make([]*ArticleJSON, 0)
	cur, err := coll.Find(r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(r.ctx)
	for cur.Next(r.ctx) {
		var a ArticleDB
		err = cur.Decode(&a)
		if err != nil {
			return nil, err
		}
		articles = append(articles, FromDBToJSON(&a))
	}

	return articles, nil
}

func (r *Repo) DeleteArticleByID(id string) error {
	db := r.client.Database(r.dbName)
	coll := db.Collection(r.coll)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": objID,
	}

	_, err = coll.DeleteOne(r.ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) BuildChangesBSON(a *ArticleJSON) bson.M {
	changes := bson.M{}
	if a.Title != "" {
		changes["title"] = a.Title
	}
	if a.Description != "" {
		changes["description"] = a.Description
	}
	if a.Tag != "" {
		changes["tag"] = a.Tag
	}
	changes["updated_at"] = time.Now()
	return changes
}

func (r *Repo) UpdateArticleByID(id string, a *ArticleJSON) error {
	db := r.client.Database(r.dbName)
	coll := db.Collection(r.coll)

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": r.BuildChangesBSON(a),
	}

	_, err = coll.UpdateByID(r.ctx, objID, update)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repo) Close() error {
	err := r.client.Disconnect(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(ctx context.Context, uri, dbName, coll string) (*Repo, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Repo{
		client: client,
		ctx:    ctx,
		dbName: dbName,
		coll:   coll,
	}, nil
}
