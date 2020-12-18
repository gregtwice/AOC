package main

import (
	"AOC/helper"
	"fmt"
	"regexp"
	"strings"
)

type rule struct {
	min, max int
}

type field struct {
	name   string
	rl, rh rule
}

func main() {
	line := helper.StringFromfile("./AOC2020/day16/sample.txt")
	//line := helper.StringFromfile("./AOC2020/day16/exemple.txt")

	parts := strings.Split(line, "\n\n")
	fields := makeRules(parts[0])
	fmt.Println(fields)


	ticketStr := strings.Split(parts[1],"\n")
	ticketArr := strings.Split(ticketStr[1],",")
	ticketNums := []int{}
	for _, s := range ticketArr {
		ticketNums = append(ticketNums, helper.MustParseInt(s))
	}
	fmt.Println(ticketNums)

	others := strings.Split(parts[2], "\n")
	badTickets := []int{}
	part1(others, fields, &badTickets)
	fmt.Println(badTickets)
	for _, ticket := range badTickets {
		fmt.Println(others[ticket])
		others[ticket+1] = ""
	}
	valid := map[string][]int{}
	for _, f := range fields {
		valid[f.name] = []int{}
		for i := 0; i < len(fields); i++ {
			valid[f.name] = append(valid[f.name], i)
		}
	}
	fmt.Println(valid)

	for li, s := range others[1:] {
		if s == "" {
			continue
		}
		for i, snum := range strings.Split(s, ",") {
			num := helper.MustParseInt(snum)
			for _, f := range fields {
				if num < f.rl.min || num > f.rh.max || (num > f.rl.max && num < f.rh.min) {
					if ok, pos := contains(valid[f.name], i); ok {
						fmt.Println("nok : ", li, f.name, num, s)
						valid[f.name] = removeInt(valid[f.name], pos)
					}
				}
			}
		}
	}
	fmt.Println(valid)
	for !valider(valid) {

		for s, ints := range valid {
			if len(ints) == 1 {
				// ok remove that from all the others

				for s2, _ := range valid {
					if s2 == s {
						continue
					}
					ok, pos := contains(valid[s2], ints[0])
					if ok {
						valid[s2] = removeInt(valid[s2], pos)
					}
				}

			}
		}
	}
	fmt.Println(valid)
	val := 1
	for s, ints := range valid {
		if strings.Contains(s, "departure") {
			val *= ticketNums[ints[0]]
		}
	}
	fmt.Println(val)
}

func valider(m map[string][]int) bool {
	for _, a := range m {
		if len(a) != 1 {
			return false
		}
	}
	return true
}

func contains(arr []int, e int) (bool, int) {
	for index, i := range arr {
		if i == e {
			return true, index
		}
	}
	return false, -1
}

func removeInt(s []int, i int) []int {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func removeString(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func part1(others []string, fields []field, badTickets *[]int) {
	for i, ticket := range others[1:] {
		numsAsStr := strings.Split(ticket, ",")
		for _, s := range numsAsStr {
			num := helper.MustParseInt(s)
			nOk := 0
			if num == 977 {
				fmt.Println(i)
			}
			for _, f := range fields {
				if num < f.rl.min || num > f.rh.max || (num > f.rl.max && num < f.rh.min) {
					nOk++
				}
			}
			if nOk == len(fields) {
				*badTickets = append(*badTickets, i)
			}
		}
	}
}

func makeRules(s string) []field {
	lines := strings.Split(s, "\n")
	pattern := regexp.MustCompile(`([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)`)
	fields := []field{}
	for _, line := range lines {
		groups := pattern.FindStringSubmatch(line)
		fields = append(fields, field{
			name: groups[1],
			rl: rule{
				min: helper.MustParseInt(groups[2]),
				max: helper.MustParseInt(groups[3]),
			},
			rh: rule{
				min: helper.MustParseInt(groups[4]),
				max: helper.MustParseInt(groups[5]),
			},
		})
	}
	return fields
}
