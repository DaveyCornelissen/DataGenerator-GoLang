package main

import (
	. "dataGenerator/intern/models"
)
import FileService "dataGenerator/intern/services"

var path = "config.json"

func main() {
	var config Configuration
	config = FileService.LoadConfiguration(path)
	FileService.GenerateFile(config.Objects, config.Format, config.FileName)
}
