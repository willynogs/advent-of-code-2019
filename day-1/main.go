package main

import (
	"fmt"
	"strings"

	"github.com/willynogs/advent-of-code-2019/helpers"
)

func loadModuleMasses() []int {
	input := helpers.LoadInput()
	arr := strings.Split(string(input), "\n")
	result := helpers.StringArrayToIntArray(arr)
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
