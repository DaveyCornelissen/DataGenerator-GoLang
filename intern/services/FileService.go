package services

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

import . "dataGenerator/intern/models"
import . "dataGenerator/intern/Utils/Handlers"

func GenerateFile(data []Object, extensionType string, filename string) {

	switch extensionType {
	case "TSV":
		t := serializeTsv(data)
		for _, tableData := range t {
			generateTsv(filename, tableData)
		}
	}
}

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

func generateTsv(filename string, table *Table) {
	tableName := table.Headers[0]
	cTime := time.Now()
	newFileName := filename + "-" + tableName + "-" + cTime.Format("2006-01-02") + ".tsv"

	file, err := os.Create(newFileName)
	if err != nil {
		CheckError("Cannot create file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	var rows [][]string
	rows = append(rows, table.Headers)
	rows = append(rows, table.Columns...)

	for _, v := range rows {
		err := writer.Write(v)
		if err != nil {
			CheckError("Cannot write to file", err)
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
