package main

import (
	"AOC/helper"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	//ecl pid eyr hcl byr iyr hgt

	required := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}

	//file, err := os.Open("./day4/exemple.txt")
	file, err := os.Open("./day4/sample.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	passports := []string{}
	currPassport := ""
	for scanner.Scan() {
		currPassport += scanner.Text() + " "
		if scanner.Text() == "" {
			passports = append(passports, currPassport)
			currPassport = ""
		}
	}
	passports = append(passports, currPassport)

	part1 := 0
	part2 := 0
	for _, passport := range passports {
		flagOKP1 := true
		flagOKP2 := true
		infos := strings.Split(strings.TrimSpace(passport), " ")
		validMap := map[string]bool{}
		for _, info := range infos {

			keyval := strings.Split(info, ":")
			key := keyval[0]
			val := keyval[1]

			switch key {
			case "byr":
				date := helper.MustParseInt(val)
				if date < 1920 || date > 2002 {
					flagOKP2 = false
				}
			case "iyr":
				date := helper.MustParseInt(val)
				if date < 2010 || date > 2020 {
					flagOKP2 = false
				}
			case "eyr":
				date := helper.MustParseInt(val)
				if date < 2020 || date > 2030 {
					flagOKP2 = false
				}

			case "hgt":
				if strings.Contains(val, "cm") {
					h := helper.MustParseInt(val[0 : len(val)-2])
					if h < 150 || h > 193 {
						flagOKP2 = false
					}
				} else if strings.Contains(val, "in") {
					h := helper.MustParseInt(val[0 : len(val)-2])
					if h < 59 || h > 76 {
						flagOKP2 = false
					}
				} else {
					flagOKP2 = false
				}
			case "hcl":
				ok, err := regexp.MatchString(`^#[0-9a-f]{6}$`, val)
				if !ok {
					flagOKP2 = false
				}
				if err != nil {
					panic(err)
				}
				//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
			case "ecl":
				if val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth" {
					// oups
				} else {
					flagOKP2 = false
				}
				//pid (Passport ID) - a nine-digit number, including leading zeroes.
			case "pid":
				ok, err := regexp.MatchString(`^[0-9]{9}$`, val)
				if err != nil {
					panic(err)
				}
				if !ok {
					flagOKP2 = false
				}
			}

			validMap[key] = true
		}

		for _, s := range required {
			if _, ok := validMap[s]; !ok {
				flagOKP1 = false
				flagOKP2 = false
			}
		}
		if flagOKP1 {
			part1++
		}
		if flagOKP2 {
			fmt.Println(passport)
			part2++
		}
	}
	fmt.Println(part1)
	fmt.Println(part2)

}
