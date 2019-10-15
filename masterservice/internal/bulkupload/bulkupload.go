package bulkupload

import (
	"bufio"
	"context"
	"encoding/csv"
	"github.com/sathishkumar64/funbook/masterservice/internal/model"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

)
//ParseCSV will parse csv and send it to DB.
func ParseCSV(ctx context.Context, csvFilename string) {
	var categoryObj []model.Category
	var subCategoryObj []model.SubCategory
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	if strings.TrimSuffix(csvFilename, filepath.Ext(csvFilename)) == "Category" {
		categoryObj= category(reader)
		log.Println(categoryObj)
	}else {
		subCategoryObj=subCategory(reader)
		log.Println(subCategoryObj)
	}
}

func subCategory(reader *csv.Reader) []model.SubCategory{

	var subCategory []model.SubCategory
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i := model.SubCategory{
			SubCategoryID: record[0],
			CategoryID:    record[1],
			Name:          record[2],
			Alias:         record[3],
			Description:   record[4],
			Active:        record[5],
		}
		subCategory = append(subCategory, i)
	}
	return subCategory
}

func category(reader *csv.Reader)[]model.Category{
	var category []model.Category
	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i := model.Category{
			CategoryID:   record[0],
			Name:         record[1],
			Alias:        record[2],
			Description:  record[3],
			Active: 	  record[4],
		}
		category = append(category, i)
	}
	return category
}