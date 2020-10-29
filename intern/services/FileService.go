package services

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"time"
)

import . "dataGenerator/intern/models"
import . "dataGenerator/intern/Utils/Handlers"

func GenerateFile(extensionType string, filename string) {
	switch extensionType {
		case "TSV":
			generateTSV(filename)
	}
}

func generateTSV(filename string) {

	cTime := time.Now().String()

	file, err := os.Create(filename + cTime + ".tsv")
	CheckError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = '\t'
	defer writer.Flush()

	//for _, value := range data {
	//	err := writer.Write(string(value))
	//	checkError("Cannot write to file", err)
	//}
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


