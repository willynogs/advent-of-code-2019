package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func loadModuleMasses() []int {
	data, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	arr := strings.Split(string(data), "\n")
	var result []int
	for _, v := range arr {
		conv, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		result = append(result, conv)
	}

	return result
}

func calcFuelRequirement(mass int) int {
	fuelRequirement := (mass / 3) - 2
	if fuelRequirement > 0 {
		fuelRequirement += calcFuelRequirement(fuelRequirement)
	} else {
		fuelRequirement = 0
	}

	return fuelRequirement
}

func main() {
	totalFuelRequirement := 0
	moduleMasses := loadModuleMasses()

	for _, v := range moduleMasses {
		totalFuelRequirement += calcFuelRequirement(v)
	}

	fmt.Println(totalFuelRequirement)
}
