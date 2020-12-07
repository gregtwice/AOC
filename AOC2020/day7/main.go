package main

import (
	"AOC/helper"
	"fmt"
	"regexp"
	"strings"
)

type sac struct {
	name string
	sacs map[string]int
}

var flagFound = false
var cmpt1 = 0

func main() {

	//strArr := helper.StringArrayFromFile("./AOC2020/day7/exemple.txt")
	strArr := helper.StringArrayFromFile("./AOC2020/day7/sample.txt")
	sacArr := []sac{}
	for _, s := range strArr {
		//fmt.Println(i)
		arr := strings.Split(s, " bags contain ")
		sacs := strings.Split(arr[1], ",")
		currSac := sac{name: arr[0]}
		currSac.sacs = map[string]int{}

		if !(sacs[0] == "no other bags.") {

			pattern := regexp.MustCompile(`(\d) (\w+ \w+)`)
			for _, str := range sacs {
				fields := pattern.FindStringSubmatch(str)
				currSac.sacs[fields[2]] = helper.MustParseInt(fields[1])
			}
		}
		sacArr = append(sacArr, currSac)

	}
	part1(sacArr)

	for i := 0; i < len(sacArr); i++ {
		if sacArr[i].name == "shiny gold" {
			fmt.Println(countAll(sacArr[i], sacArr))
		}
	}
}

func countAll(s sac, arr []sac) int {
	if len(s.sacs) == 0 {
		return 1
	}
	cmpt := 0
	for sacName, nb := range s.sacs {
		cmpt += countAll(arr[findSac(sacName, arr)], arr) * nb
		if len(arr[findSac(sacName, arr)].sacs) > 0 {
			cmpt+= nb
		}
	}
	return cmpt
}

func part1(sacArr []sac) {

	for _, s := range sacArr {
		canContain(s, sacArr)
		flagFound = false
	}

	fmt.Println(cmpt1)
}

func canContain(s sac, arr []sac) {
	if len(s.sacs) == 0 {
		return
	}
	for sacName := range s.sacs {
		if flagFound {
			return
		}
		if sacName == "shiny gold" {
			cmpt1++
			flagFound = true
			return
		} else {
			canContain(arr[findSac(sacName, arr)], arr)
		}
	}
}

func findSac(s string, arr []sac) int {
	for i, sac := range arr {
		if sac.name == s {
			return i
		}
	}
	return -1
}
