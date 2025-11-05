package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	client *mongo.Client
	dbName string
	ctx    context.Context
}

func (r *Repo) Close() error {
	err := r.client.Disconnect(r.ctx)
	if err != nil {
		return err
	}
	return nil
}

func NewRepo(ctx context.Context, uri, dbName string) (*Repo, error) {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &Repo{
		client: client,
		ctx:    ctx,
		dbName: dbName,
	}, nil
}
