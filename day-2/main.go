package main

import (
	"fmt"
	"strings"

	"github.com/willynogs/advent-of-code-2019/helpers"
)

func loadIntCode() []int {
	input := helpers.LoadInput()
	arr := strings.Split(input, ",")
	result := helpers.StringArrayToIntArray(arr)
	return result
}

func main() {
	intCode := loadIntCode()

	for i, v := range intCode {

	}

	fmt.Println(intCode)
}
