package main

import (
	"AOC/helper"
	"fmt"
	"log"
)

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day12/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day12/exemple.txt")
	//depart := helper.Point{0,0}
	currPosP1 := helper.Point{0, 0}
	currPosP2 := helper.Point{0, 0}
	waypoint := helper.Point{10, -1}
	directions := []byte{'W', 'N', 'E', 'S'}
	currDir := 2
	for _, line := range lines {
		action := line[0]
		amount := helper.MustParseInt(line[1:])
		log.Println(currPosP2, string(action), amount, waypoint)
		switch action {
		case 'F':
			currPosP2.X += waypoint.X * amount
			currPosP2.Y += waypoint.Y * amount
			switch directions[currDir] {
			case 'E':
				currPosP1.X += amount
			case 'N':
				currPosP1.Y -= amount
			case 'S':
				currPosP1.Y += amount
			case 'W':
				currPosP1.X -= amount
			}
		case 'N':
			currPosP1.Y -= amount
			waypoint.Y -=amount
		case 'S':
			currPosP1.Y += amount
			waypoint.Y += amount
		case 'E':
			currPosP1.X += amount
			waypoint.X += amount
		case 'W':
			currPosP1.X -= amount
			waypoint.X -= amount
		case 'L':
			currDir += (-amount / 90)
			if currDir < 0 {
				currDir = 4 + currDir
			}
			waypoint.Rotate(amount)

		case 'R':
			currDir = (currDir + (amount / 90)) % 4
			waypoint.Rotate(-amount)
		}
	}
	fmt.Println(currPosP1.ManDist(helper.Point{0, 0}))
	fmt.Println(currPosP2.ManDist(helper.Point{0,0}))

}

func rotateShip(amount int, curDir int) {

}
