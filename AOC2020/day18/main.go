package main

import (
	"AOC/helper"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day18/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day8/exemple.txt")

	line := "2 * 3 + (4 * 5)"
	//line := "1 + 2 * 3 + 4 * 5 + 6"

	//line := "6 * (5 + (6 * 7 * 3 * 3 + 2 + 3) + 6 * (4 + 7 * 3 + 5 + 5 + 2) * 4 + 4) * 9"

	fmt.Println(parseExp(line))

	sum := 0
	for i, line := range lines {
		fmt.Println(i, line)
		sum += parseExp(line)
	}
	fmt.Println(sum)
}

func parseExp(s string) int {

	pattern := regexp.MustCompile(`(\([+ \d*]+\))`)
	loops := 0
	//fmt.Println(pars)
	for strings.Contains(s, "(") {
		results := []int{}
		pars := pattern.FindAllString(s, -1)
		//indices := pattern.FindAllStringIndex(s, -1)
		fmt.Println(s)
		for _, par := range pars {
			expr := par[1 : len(par)-1]
			r := parseExp(expr)
			results = append(results, r)
		}

		//n := 0
		for i, res := range results {
			s = strings.Replace(s, pars[i], strconv.Itoa(res), 1)
			/*if indices[i][1]+1 < len(s) {

				s = s[0:indices[i][0]-n] + strconv.Itoa(res) + s[indices[i][1]-n:]
			} else {
				s = s[0:indices[i][0]-n] + strconv.Itoa(res)
			}
			n += len(pars[i]) - len(strconv.Itoa(res))*/
		}
		loops++
		if loops == 10 {
			os.Exit(0)
		}
	}

	elems := strings.Split(s, " ")

	for i, elem := range elems {
		if elem == "+" {
			a := helper.MustParseInt(elems[i-1])
			b := helper.MustParseInt(elems[i+1])
			elems[i+1] = strconv.Itoa(a + b)
			elems[i-1] = ""
			elems[i] = ""
		}
	}
	e := strings.Join(elems, " ")
	elems = strings.FieldsFunc(e, func(r rune) bool {
		return r == ' '
	})
	for i, elem := range elems {
		if elem == "*" {
			a := helper.MustParseInt(elems[i-1])
			b := helper.MustParseInt(elems[i+1])
			elems[i+1] = strconv.Itoa(a * b)
		}
	}

	return helper.MustParseInt(elems[len(elems)-1])
}
