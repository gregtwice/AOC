package main

import (
	"AOC/helper"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day14/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day14/exemple.txt")
	mem := map[int]int{}
	mask := ""
	pattern := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)


	for _, line := range lines {
		if strings.Contains(line, "mask") {
			mask = strings.Split(line, " = ")[1]
			fmt.Println(mask)
		} else {
			groups := pattern.FindStringSubmatch(line)
			address := helper.MustParseInt(groups[1])
			value := helper.MustParseInt(groups[2])
			addresses := getAllAddresses(mask, address)
			for _, addr := range addresses {
				mem[addr] = value
			}
		}
	}
	s:= 0
	for _, i := range mem {
		s+=i
	}
	fmt.Println(s)
}

func getAllAddresses(mask string, addrValue int) []int {
	numX := strings.Count(mask, "X")
	poss := int(math.Pow(2, float64(numX)))
	posX := []int{}

	addresses := []int{}

	for i, c := range mask {
		if c == 'X' {
			posX = append(posX, i)
		}
	}
	fmt.Println(mask)
	intMask, _ := strconv.ParseInt(strings.Replace(mask, "X", "0", -1), 2, 64)
	for i := 0; i < poss; i++ {
		addresses = append(addresses, int(intMask)|addrValue)
		bitString := fmt.Sprintf("%0"+strconv.Itoa(numX)+"b", i)
		//fmt.Println(bitString)
		newaddr := []rune(fmt.Sprintf("%036b", addresses[i]))
		for j, b := range bitString {
			newaddr[posX[j]] = b
		}
		newMaskInt, err := strconv.ParseInt(string(newaddr), 2, 64)
		if err != nil {
			panic(err)
		}
		//addresses[i] &= int(newMaskInt)
		addresses[i] = int(newMaskInt)

	}
	return addresses
}

func printPart1(lines []string) int {
	mem := map[int]int{}
	mask := ""
	pattern := regexp.MustCompile(`mem\[(\d+)\] = (\d+)`)

	for _, line := range lines {
		if strings.Contains(line, "mask") {
			mask = strings.Split(line, " = ")[1]
			fmt.Println(mask)
		} else {
			groups := pattern.FindStringSubmatch(line)

			address := helper.MustParseInt(groups[1])
			value := helper.MustParseInt(groups[2])
			mem[address] = value

			mask1, err := strconv.ParseInt(strings.Replace(mask, "X", "0", -1), 2, 64)
			if err != nil {
				panic(err)
			}
			mask0, err := strconv.ParseInt(strings.Replace(mask, "X", "1", -1), 2, 64)
			if err != nil {
				panic(err)
			}

			fmt.Println(mem[address])
			mem[address] |= int(mask1)
			mem[address] &= int(mask0)
			fmt.Println(mem[address])
		}
	}
	s := 0
	for _, i := range mem {
		s += i
	}
	return s
}
