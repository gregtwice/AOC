package main

import (
	"AOC2016/helper"
	"fmt"
	"regexp"
)

func main() {

	triangles := helper.StringArrayFromFile("C:\\Users\\grego\\Documents\\go\\src\\AOC2016\\day3\\sample.txt")
	nb := 0

	for _, triangleSpec := range triangles {
		var a, b, c int

		pattern := regexp.MustCompile(`\s+(\d+)\s+(\d+)\s+(\d+)`)
		matchs := pattern.FindStringSubmatch(triangleSpec)

		a = helper.MustParseInt(matchs[1])
		b = helper.MustParseInt(matchs[2])
		c = helper.MustParseInt(matchs[3])

		triangle := []int{a, b, c}
		maxI := 0
		maxCote := 0
		for i, cote := range triangle {
			if cote > maxCote {
				maxCote = cote
				maxI = i
			}
		}


		sum := triangle[0] + triangle[1] + triangle[2] - triangle[maxI]
		fmt.Println(a,b,c,maxCote,maxI,sum)
		if sum > triangle[maxI] {
			nb++
		}
	}

	fmt.Println(nb)

}
