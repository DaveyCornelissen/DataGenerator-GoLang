package main

import (
	. "dataGenerator/intern/models"
	"fmt"
)
import FileService "dataGenerator/intern/services"
import MockService "dataGenerator/intern/services"

var path = "config.json"

func main() {
	var config Configuration
	//var columnList []Key

	config = FileService.LoadConfiguration(path)
	serialize(config.Objects)

}

func serialize(rootObjs []Object) interface{} {

	var _rootList []RootColumn
	var _childList []ChildColumn

	for i, rootObj := range rootObjs {
		if rootObj.Key != "" || rootObj.ChildObjects == nil {
			for _, v := range rootObj.Values {
				_rootColumn := RootColumn{
					Id:    i,
					Name:  rootObj.Key,
					Value: v,
				}
				_rootList = append(_rootList, _rootColumn)
			}

			for _, childObj := range rootObj.ChildObjects {
				_childColumn := ChildColumn{
					RootId: i,
					Name:   childObj.Key,
				}
				_childColumn.Value = MockService.Generate(childObj)
				_childList = append(_childList, _childColumn)
			}

		} else {
			fmt.Println("Error no Root Key or ChildObject provided!")
		}

	}
}
