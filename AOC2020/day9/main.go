package main

import (
	"AOC/helper"
	"fmt"
)

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day9/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day9/exemple.txt")
	numbers := []int{}
	for _, line := range lines {
		numbers = append(numbers, helper.MustParseInt(line))
	}

	preambule := 25

	faulty := printPart1(preambule, numbers)
	for i := 0; i < len(numbers); i++ {
		sum := numbers[i]
		for j := i + 1; j < len(numbers); j++ {
			sum += numbers[j]
			if sum > faulty {
				break
			} else if sum == faulty {
				nums := numbers[i:j]
				min := nums[0]
				max := 0
				for _, num := range nums {
					min = helper.Min(min, num)
					max = helper.Max(max, num)
				}
				fmt.Println(min + max)
				return
			}
		}
	}

}

func printPart1(preambule int, numbers []int) int {
	for i := preambule; i < len(numbers); i++ {
		flagOK := false
		for j := i - preambule; j < i; j++ {
			for k := j + 1; k < i; k++ {
				if numbers[j]+numbers[k] == numbers[i] {
					flagOK = true
				}
			}
		}
		if flagOK {
			fmt.Println("ok")
		} else {
			fmt.Println("nok", numbers[i], i)
			return numbers[i]
		}
	}
	return 0
}
