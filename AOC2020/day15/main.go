package main

import (
	"AOC/helper"
	"fmt"
	"strings"
)

func main() {
	//lines := helper.StringArrayFromFile("./AOC2020/day8/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day8/exemple.txt")

	line := "17,1,3,16,19,0"
	//line = "0,3,6"
	//line = "2,1,3"
	//line = "1,2,3"
	//line = "2,3,1"
	//line = "3,2,1"
	numbers := processInput(line)

	max := 30000000
	turns := make([]int, max+1)
	spoken := map[int][]int{}

	compute(max, numbers, turns, spoken)
	fmt.Println(turns[max])
}

func compute(max int, numbers []int, turns []int, spoken map[int][]int) {
	for i := 1; i <= max; i++ {
		if i < len(numbers)+1 {
			turns[i] = numbers[i-1]
			spoken[numbers[i-1]] = append(spoken[numbers[i-1]], i)
		} else {
			lastSpoken := turns[i-1]
			if len(spoken[lastSpoken]) == 1 { // firstTimespoken
				turns[i] = 0
				//if len(spoken[0]) == 0 {
				spoken[0] = append(spoken[0], i)
				//} else {
				//	spoken[0] = append(spoken[0], i)
				//}
			} else {
				num := spoken[lastSpoken]
				newNum := num[len(num)-1] - num[len(num)-2]
				turns[i] = newNum
				//if len(spoken[newNum]) == 0 {
				//	spoken[newNum][0] = i
				//} else {
				//	spoken[newNum][1], spoken[newNum][0] = i, spoken[newNum][1]
				//}
				spoken[newNum] = append(spoken[newNum], i)
			}
		}
	}
}

func processInput(str string) []int {
	arr := strings.Split(str, ",")

	nums := []int{}
	for _, s := range arr {
		nums = append(nums, helper.MustParseInt(s))
	}
	return nums
}
