package helpers

import (
	"io/ioutil"
	"log"
	"strconv"
)

// LoadInput loads the input file data
func LoadInput() string {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return string(data)
}

// StringArrayToIntArray turns an array of strings into an array of ints
func StringArrayToIntArray(input []string) []int {
	var result []int
	for _, v := range input {
		conv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		result = append(result, conv)
	}
	return result
}
