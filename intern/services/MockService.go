package services

import (
	. "dataGenerator/intern/models"
	"math/rand"
	"strconv"
)
import "strings"
import . "dataGenerator/intern/Utils/Handlers"

const (
	random = "*"
	length = "-"
)

func GenerateMockData(childObject ChildObject) string {

	//Random attr
	if strings.Contains(childObject.Value, random) {
		switch childObject.Type {
		case "bool":
			return strconv.FormatBool(randBoolean())
		case "float":
			return floatToString(rand.Float64())
		}
	}

	//length between certain value
	if strings.Contains(childObject.Value, length) {
		switch childObject.Type {
		case "float":
			numbers := strings.Split(childObject.Value, length)
			return floatToString(randFloat(numbers))
		case "int":
			numbers := strings.Split(childObject.Value, length)
			return strconv.Itoa(randInt(numbers))
		case "string":
			return randString(stringToInt(childObject.Value))
		}
	}
	return childObject.Value
}

func randBoolean() bool {
	return rand.Float32() < 0.5
}

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func randFloat(numbers []string) float64 {
	min := stringToFloat(numbers[0])
	max := stringToFloat(numbers[1])

	return min + rand.Float64()*(max-min)
}

func randInt(numbers []string) int {
	min := stringToInt(numbers[0])
	max := stringToInt(numbers[1])

	return min + rand.Intn(max-min)
}

func stringToInt(input_num string) int {
	i, err := strconv.Atoi(input_num)
	if err != nil {
		CheckError("could not convert type:", err)
	}
	return i
}

func stringToFloat(input_num string) float64 {
	f, err := strconv.ParseFloat(input_num, 64)
	if err != nil {
		CheckError("could not convert type:", err)
	}

	return f
}

func floatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}
