package service

import (
	"context"
	"github.com/sathishkumar64/funbook/masterservice/internal/bulkupload"
	"github.com/sathishkumar64/funbook/masterservice/internal/durable"
	"github.com/sathishkumar64/funbook/masterservice/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)
//RegisterCategory is used to parse CSV and insert data into mongodb.
func RegisterCategory(ctx context.Context,database *durable.Database,csvFileName string){

	var school *mongo.Collection
	var categoryObj []model.Category
	school = database.Db.Database("schoolservice").Collection("category")
	withExtFileName:=	csvFileName +".csv"
	reader:=bulkupload.ParseCSV(ctx, withExtFileName)
	categoryObj = bulkupload.Category(reader)
	var category []interface{}
	for _, p := range categoryObj {
		category = append(category, p)
	}
		dropCollection(ctx,school)
	_, err := school.InsertMany(ctx, category)
	if err != nil {
		log.Fatal(err)
	}
}

//RegisterSubCategory is used to parse CSV and insert data into mongodb.
func RegisterSubCategory(ctx context.Context,database *durable.Database,csvFileName string){
	var school *mongo.Collection
	var categoryObj []model.SubCategory
	school = database.Db.Database("schoolservice").Collection("subcategory")
	withExtFileName:=	csvFileName +".csv"
	reader:=bulkupload.ParseCSV(ctx, withExtFileName)
	categoryObj = bulkupload.SubCategory(reader)
	var category []interface{}
	for _, p := range categoryObj {
		category = append(category, p)
	}
	school.Drop(ctx)
	dropCollection(ctx,school)
	_, err := school.InsertMany(ctx, category)
	if err != nil {
		log.Fatal(err)
	}
}


func dropCollection(ctx context.Context,collection *mongo.Collection){
	error := collection.Drop(ctx)
	if error != nil {
		log.Fatal(error)
	}
}