package services

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"time"
)

import . "dataGenerator/intern/models"

func GenerateFile(extensionType string, filename string) {
	switch extensionType {
		case "TSV":
			generateTSV(filename)
	}
}

func generateTSV(filename string) {

	cTime := time.Now().String()

	file, err := os.Create(filename + cTime + ".tsv")
	checkError("Cannot create file", err)
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
		checkError("cannont read file", err)
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&con)
	return con
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
