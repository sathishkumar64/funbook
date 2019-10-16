package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

//Category is used to define model.
type Category struct {
	CategoryID  string
	Name        string
	Alias       string
	Description string
	Active      string
}

//SubCategory is used to define model.
type SubCategory struct {
	SubCategoryID string
	CategoryID    string
	Name          string
	Alias         string
	Description   string
	Active        string
}

//InsertCollection is used to insert new collections
func InsertCollection(ctx context.Context,collection *mongo.Collection,category []interface{}){
	_, err := collection.InsertMany(ctx, category)
	if err != nil {
		log.Fatal(err)
	}
}

//DropCollection is used to drop collection completely
func DropCollection(ctx context.Context,collection *mongo.Collection){
	error := collection.Drop(ctx)
	if error != nil {
		log.Fatal(error)
	}
}