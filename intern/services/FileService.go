package services

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

import . "dataGenerator/intern/models"
import . "dataGenerator/intern/Utils/Handlers"

//can be expanded with for instance protocol buffer only add new case and method
func GenerateFile(data []Object, extensionType string, filename string) {
	switch extensionType {
	case "TSV":
		t := serializeTsv(data)
		generateTsv(filename, t)
	}
}

//serialize the raw objects to usable tsv format
func serializeTsv(rootObjs []Object) []*Table {

	var _rootList []*Table
	var _totalChildrenCount int

	for _, rootObj := range rootObjs {
		if rootObj.Key != "" || rootObj.ChildObjects == nil {
			var table Table

			table.Headers = []string{
				rootObj.Key,
			}
			for _, h := range rootObj.ChildObjects {
				table.Headers = append(table.Headers, h.Key)
			}
			_rootList = append(_rootList, &table)

			for _, v := range rootObj.Values {
				for r := 1; r <= rootObj.TotalObjects; r++ {
					_column := []string{
						v,
					}
					for _, childObj := range rootObj.ChildObjects {
						_column = append(_column, GenerateMockData(childObj))
						_totalChildrenCount++
					}
					table.Columns = append(table.Columns, _column)
					fmt.Println("Created child row: ", r)
				}
			}
		} else {
			CheckError("Error no Root Key or ChildObject provided!")
		}
	}
	fmt.Printf("Done! Total RootColumns: %d  Total ChildColumns: %d \n", len(_rootList), _totalChildrenCount)
	return _rootList
}
//Generate the actual files and compress them into zip
func generateTsv(filename string, tables []*Table) {
	cTime := time.Now()
	zipFileName := filename + "-" + cTime.Format("2006-01-02") + ".zip"
	file, err := os.Create(zipFileName)
	if err != nil {
		CheckError("Cannot create file", err)
	}
	zipArchive := zip.NewWriter(file)
	defer zipArchive.Flush()
	defer zipArchive.Close()
	defer file.Close()

	for _, t := range tables {
		tableName := t.Headers[0]
		tsvFileName := filename + "-" + tableName + "-" + cTime.Format("2006-01-02") + ".tsv"

		zipWriter, _ := zipArchive.Create(tsvFileName)
		tsvWriter := csv.NewWriter(zipWriter)
		tsvWriter.Comma = '\t'
		defer tsvWriter.Flush()

		var rows [][]string
		rows = append(rows, t.Headers)
		rows = append(rows, t.Columns...)

		for _, v := range rows {
			err := tsvWriter.Write(v)
			if err != nil {
				CheckError("Cannot write to file", err)
			}
		}
	}
}

func LoadConfiguration(file string) Configuration {
	var con Configuration

	configFile, err := os.Open(file)

	if err != nil {
		CheckError("cannont read file", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&con)
	return con
}
