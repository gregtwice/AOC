package main

import (
	"AOC2016/helper"
	"fmt"
)

func main() {
	instructions := helper.StringArrayFromFile("C:\\Users\\grego\\Documents\\go\\src\\AOC2016\\day2\\sample.txt")

	keyPad := [][]byte{
		{'*', '*', '1', '*', '*'},
		{'*', '2', '3', '4', '*'},
		{'5', '6', '7', '8', '9'},
		{'*', 'A', 'B', 'C', '*'},
		{'*', '*', 'D', '*', '*'}}
	c, l := 0, 2
	for _, instruction := range instructions {
		for i := 0; i < len(instruction); i++ {
			switch instruction[i] {
			case 'U':
				l--
				if l < 0 || keyPad[l][c] == '*' {
					l++
				}
			case 'D':
				l++
				if l > 4 || keyPad[l][c] == '*' {
					l--
				}
			case 'L':
				c--
				if c < 0 || keyPad[l][c] == '*' {
					c++
				}
			case 'R':
				c++
				if c > 4 || keyPad[l][c] == '*' {
					c--
				}
			}
		}
		fmt.Printf("%c",keyPad[l][c])
	}
}

func printPart1(instructions []string) {
	num := 5
	for _, instruction := range instructions {
		for i := 0; i < len(instruction); i++ {

			switch instruction[i] {
			case 'U':
				if num >= 4 {
					num -= 3
				}
			case 'D':
				if num <= 6 {
					num += 3
				}
			case 'L':
				if num != 1 && num != 4 && num != 7 {
					num -= 1
				}
			case 'R':
				if num != 3 && num != 6 && num != 9 {
					num += 1
				}
			}
		}
		fmt.Print(num)
	}
}
