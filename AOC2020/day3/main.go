package main

import (
	"bufio"
	"fmt"
	"os"
)

type increment struct {
	x, y int
}

func main() {

	tobMap := [323][31]byte{}
	file, err := os.Open("./day3/sample.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		for i, b := range scanner.Text() {

			tobMap[line][i] = byte(b)
		}
		line++
	}
	fmt.Println(tobMap)
	all := 1
	incs := []increment{{x: 1, y: 1,}, {x: 3, y: 1,}, {x: 5, y: 1,}, {x: 7, y: 1,}, {x: 1, y: 2,}}
	for i := 0; i < len(incs); i++ {
		trees := 0
		x, y := 0, 0

		for y < 323 {
			if tobMap[y][x] == '#' {
				trees++
			}
			x += incs[i].x
			y += incs[i].y
			x %= 31
		}
		all *= trees
	}
	fmt.Println(all)

}
