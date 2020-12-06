package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	min int
	max int
	letter string
	password string
}

func main() {

	//file ,err := os.Open("./day2/exemple.txt")
	file ,err := os.Open("./day2/sample.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	rules := []Rule{}
	re := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	for scanner.Scan() {
		match := re.FindStringSubmatch(scanner.Text())
		min,_ := strconv.Atoi(match[1])
		max,_ := strconv.Atoi(match[2])
		rules = append(rules, Rule{
			min:      min,
			max:      max,
			letter:   match[3],
			password: match[4],
		})
	}

	cmpt := 0
	for _, rule := range rules {
		if (rule.password[rule.min-1] == []byte(rule.letter)[0]  ) != ( rule.password[rule.max-1] == []byte(rule.letter)[0]){
			cmpt++
		}
	}
	fmt.Println(cmpt)

}

func printPart1(rules []Rule) {
	cmpt := 0

	for _, rule := range rules {
		nb := strings.Count(rule.password, rule.letter)
		if nb >= rule.min && nb <= rule.max {
			cmpt++
		}
	}

	fmt.Println(cmpt)
}
