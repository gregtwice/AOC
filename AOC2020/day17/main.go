package main

import (
	"AOC/helper"
	"fmt"
)

type Point struct {
	x, y, z,w int
}

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day17/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day17/exemple.txt")

	xMin := 0
	xMax := 2
	yMin := 0
	yMax := len(lines)
	zMin := 0
	zMax := 0
	wMin := 0
	wMax := 0
	m := map[Point]bool{}
	for y, line := range lines {
		xMax = len(line)
		for x, c := range line {
			m[Point{x, y, 0,0}] = c == '#'
		}
	}
	fmt.Println(xMax,yMax)
	//println(m[Point{2, 2, 0}])

	for g := 1; g <= 6; g++ {
		toAct := []Point{}
		toDeact := []Point{}
		//PrintMap(m, xMin, xMax, yMin, yMax, zMin, zMax)
		xMin--
		xMax++
		yMin--
		yMax++
		zMin--
		zMax++
		wMin--
		wMax++
		//	pour chaque point
		for x := xMin; x <= xMax; x++ {
			for y := yMin; y <= yMax; y++ {
				for z := zMin; z <= zMax; z++ {
					for w := wMin; w <= wMax; w++ {

						p := Point{x, y, z, w}
						nbVoisins := 0

						for i := x - 1; i <= x+1; i++ {
							for j := y - 1; j <= y+1; j++ {
								for k := z - 1; k <= z+1; k++ {
									for l := w - 1; l <= w+1; l++ {
										if i == x && j == y && k == z && l == w {
											continue
										}
										nbVoisins += b2i(m[Point{i, j, k, l}])
									}
								}
							}
						}
						if ok, b := m[Point{x, y, z, w}]; ok && b {
							if nbVoisins == 2 || nbVoisins == 3 {
								toAct = append(toAct, p)
							} else {
								toDeact = append(toDeact, p)
							}
						} else {
							if nbVoisins == 3 {
								toAct = append(toAct, p)
							}
						}
					}
				}
			}
		}
		for _, point := range toAct {
			m[point] = true
		}
		for _, point := range toDeact {
			m[point] = false
		}
		nb := 0
		for _, b := range m {
			nb += b2i(b)
		}
		fmt.Println(nb)
	}
	//PrintMap(m, xMin, xMax, yMin, yMax, zMin, zMax)
}

func b2i(bool2 bool) int {
	if bool2 {
		return 1
	} else {
		return 0
	}

}


func PrintMap(m map[Point]bool, xmn, xmx, ymn, ymx, zmn, zmx int) {

	//for z := zmn; z <= zmx; z++ {
	//	fmt.Println("z = ", z)
	//	for y := ymn; y <= ymx; y++ {
	//		for x := xmn; x <= xmx; x++ {
	//			if m[Point{x, y, z}] {
	//				fmt.Print("@")
	//			} else {
	//				fmt.Print(".")
	//			}
	//		}
	//		fmt.Println("\t", y)
	//	}
	//}
	//fmt.Println("----------------------")
}
