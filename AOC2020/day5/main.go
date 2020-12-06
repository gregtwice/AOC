package main

import (
	"AOC/helper"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	//lines := helper.StringArrayFromFile("./AOC2020/day5/exemple.txt")

	allSeats := []int{}

	lines := helper.StringArrayFromFile("./AOC2020/day5/sample.txt")
	max := 0
	for _, line := range lines {
		row := ""
		column := ""
		for i := 0; i < len(line); i++ {
			switch line[i] {
			case 'F':
				row += "0"
			case 'B':
				row += "1"
			case 'R':
				column += "1"
			case 'L':
				column += "0"
			}
		}
		intRow, err := strconv.ParseInt(row, 2, 0)
		intCol, err := strconv.ParseInt(column, 2, 0)
		if err != nil {
			panic(err)
		}

		seatId := int(intRow*8 + intCol)
		allSeats = append(allSeats, seatId)
		max = helper.Max(max, seatId)
	}
	sort.Ints(allSeats)
	last:= allSeats[0]
	for i := 1; i < len(allSeats); i++ {
		seat := allSeats[i]
		if seat-last != 1 {
			fmt.Println(seat-1)
		}
		last = seat
	}
	fmt.Println(max)
}
