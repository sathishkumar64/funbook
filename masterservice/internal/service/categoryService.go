package service

import (
	"context"
	"github.com/sathishkumar64/funbook/masterservice/internal/bulkupload"
	"github.com/sathishkumar64/funbook/masterservice/internal/durable"
	"github.com/sathishkumar64/funbook/masterservice/internal/model"
	"go.mongodb.org/mongo-driver/mongo"
)
//RegisterCategory is used to parse CSV and insert data into mongodb.
func RegisterCategory(ctx context.Context,database *durable.Database,csvFileName string){

	var category *mongo.Collection
	var categoryObj []model.Category
	category = database.Db.Database("schoolservice").Collection("category")
	withExtFileName:=	csvFileName +".csv"
	reader:=bulkupload.ParseCSV(ctx, withExtFileName)
	categoryObj = bulkupload.Category(reader)
	var categoryInterface []interface{}
	for _, p := range categoryObj {
		categoryInterface = append(categoryInterface, p)
	}
	model.DropCollection(ctx,category)
	model.InsertCollection(ctx,category,categoryInterface)
}

//RegisterSubCategory is used to parse CSV and insert data into mongodb.
func RegisterSubCategory(ctx context.Context,database *durable.Database,csvFileName string){
	var subCategory *mongo.Collection
	var categoryObj []model.SubCategory
	subCategory = database.Db.Database("schoolservice").Collection("subcategory")
	withExtFileName:=	csvFileName +".csv"
	reader:=bulkupload.ParseCSV(ctx, withExtFileName)
	categoryObj = bulkupload.SubCategory(reader)
	var subCategoryInterface []interface{}
	for _, p := range categoryObj {
		subCategoryInterface = append(subCategoryInterface, p)
	}
	model.DropCollection(ctx,subCategory)
	model.InsertCollection(ctx,subCategory,subCategoryInterface)
}

