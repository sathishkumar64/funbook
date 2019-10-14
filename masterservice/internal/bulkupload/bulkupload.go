package bulkupload

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func parseCSV() {

	csvfile, err := os.Open("test.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(bufio.NewReader(csvfile))

	for {
		// Read each record from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
}
