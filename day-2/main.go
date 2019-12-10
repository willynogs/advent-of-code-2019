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

func runIntCode(noun int, verb int) int {
	intCode := loadIntCode()

	intCode[1] = noun
	intCode[2] = verb

	var operation int
	var newValue int
	for i, v := range intCode {
		if i%4 == 0 {
			if v == 99 {
				break
			} else {
				operation = v
			}
		} else if i%4 == 1 {
			newValue = intCode[v]
		} else if i%4 == 2 {
			if operation == 1 {
				newValue += intCode[v]
			} else if operation == 2 {
				newValue *= intCode[v]
			}
		} else if i%4 == 3 {
			intCode[v] = newValue
		}
	}

	return intCode[0]
}

func main() {
	result := runIntCode(12, 2)

	fmt.Println(result)
}
