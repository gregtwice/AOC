package main

import (
	"AOC2016/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	N = iota
	E
	S
	O
)

type Point struct {
	x, y int
}

func main() {
	fmt.Println(os.Getwd())
	file, err := os.Open("./day1/exemple.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	direction := N
	scanner.Scan()
	instructions := strings.Split(scanner.Text(), ", ")

	visited := map[Point]bool{}
	visited[Point{0, 0}] = true
	x, y := 0, 0
	for _, instruction := range instructions {

		//amount, _ := strconv.Atoi(string(instruction[1:]))
		amount := helper.MustParseInt(instruction[1:])

		if instruction[0] == 'R' {
			direction = (direction + 1) % 4
		} else {
			direction = (direction + 3) % 4
		}

		switch direction {
		case N:
			x += amount
			for i := 0; i < amount; i++ {
				if _, ok := visited[Point{x - i, y}]; ok {
					fmt.Println(Point{x - i, y})
					os.Exit(0)
				}
				visited[Point{x - i, y}] = true
			}
		case E:
			y += amount
			for i := 0; i < amount; i++ {
				if _, ok := visited[Point{x, y - i}]; ok {
					fmt.Println(Point{x, y - i})
					os.Exit(0)
				}
				visited[Point{x, y - i}] = true
			}
		case S:
			x -= amount
			for i := 0; i < amount; i++ {
				if _, ok := visited[Point{x + i, y}]; ok {
					fmt.Println(Point{x + i, y})
					os.Exit(0)
				}
				visited[Point{x + i, y}] = true
			}
		case O:
			y -= amount
			for i := 0; i < amount; i++ {
				if _, ok := visited[Point{x, y + i}]; ok {
					fmt.Println(Point{x, y + i})
					os.Exit(0)
				}
				visited[Point{x, y + i}] = true
			}
		}
	}

	println(helper.Abs(x) + helper.Abs(y))

}
