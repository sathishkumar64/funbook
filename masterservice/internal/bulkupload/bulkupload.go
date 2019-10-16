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

// ParseFileName method is used to get type of category.
func ParseFileName(ctx context.Context, csvFilename string)string{
	return  strings.TrimSuffix(csvFilename, filepath.Ext(csvFilename))
}




//ParseCSV will parse csv and send it to DB.
func ParseCSV(ctx context.Context, csvFilename string)*csv.Reader{
	csvFile, err := os.Open(csvFilename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	return reader
}
//SubCategory method is used to read all SubCategory data.
func SubCategory(reader *csv.Reader) []model.SubCategory{

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
			SubCategoryID: strings.TrimSpace(record[0]),
			CategoryID:    strings.TrimSpace(record[1]),
			Name:          strings.TrimSpace(record[2]),
			Alias:         strings.TrimSpace(record[3]),
			Description:   strings.TrimSpace(record[4]),
			Active:      	strings.TrimSpace(record[5]),
		}
		subCategory = append(subCategory, i)
	}
	return subCategory
}
//Category method is used to read all Category data.
func Category(reader *csv.Reader)[]model.Category{
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
			CategoryID:   strings.TrimSpace(record[0]),
			Name:         strings.TrimSpace(record[1]),
			Alias:       strings.TrimSpace( record[2]),
			Description: strings.TrimSpace( record[3]),
			Active: 	  strings.TrimSpace(record[4]),
		}
		category = append(category, i)
	}
	return category
}