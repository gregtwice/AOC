package main

import (
	"AOC/helper"
	"fmt"
)

const (
	EMPTY    = 'L'
	OCCUPIED = '#'
	FLOOR    = '.'
)

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day11/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day11/exemple.txt")
	newState := make([]string, len(lines))
	for {
		hasChanged := false
		for r := 0; r < len(lines); r++ {
			currStateStr := ""
			for c := 0; c < len(lines[r]); c++ {
				if r == 1 && c == 5 {
					fmt.Println("LA")
				}
				switch lines[r][c] {
				case FLOOR:
					currStateStr += string(FLOOR)
				case EMPTY:
					if canOcc2(lines, r, c) {
						currStateStr += string(OCCUPIED)
					} else {
						currStateStr += string(lines[r][c])
					}
				case OCCUPIED:
					if canLeave2(lines, r, c) {
						currStateStr += string(EMPTY)
					} else {
						currStateStr += string(lines[r][c])
					}
				}

			}
			newState[r] = currStateStr
			if currStateStr != lines[r] {
				hasChanged = true
			}
		}
		if !hasChanged {
			c := 0
			for _, line := range lines {
				for _, s := range line {
					if s == OCCUPIED {
						c++
					}
				}
			}
			fmt.Println(c)
			break
		}
		for _, s := range newState {
			fmt.Println(s)
		}
		fmt.Println("\n")
		copy(lines, newState)
	}

}

func canOcc(lines []string, r, c int) bool {
	g, d, h, b := false, false, false, false
	hg, hd, bg, bd := false, false, false, false
	if r-1 >= 0 {
		h = true
	}
	if r+1 < len(lines) {
		b = true
	}
	if c-1 >= 0 {
		g = true
	}
	if c+1 < len(lines[r]) {
		d = true
	}
	hg = r-1 >= 0 && c-1 >= 0
	hd = r-1 >= 0 && c+1 < len(lines[r])
	bg = r+1 < len(lines) && c-1 >= 0
	bd = r+1 < len(lines) && c+1 < len(lines[r])

	toRet := true
	if h && lines[r-1][c] == OCCUPIED {
		toRet = false
	}
	if b && lines[r+1][c] == OCCUPIED {
		toRet = false
	}
	if g && lines[r][c-1] == OCCUPIED {
		toRet = false
	}
	if d && lines[r][c+1] == OCCUPIED {
		toRet = false
	}
	if hg && lines[r-1][c-1] == OCCUPIED {
		toRet = false
	}
	if hd && lines[r-1][c+1] == OCCUPIED {
		toRet = false
	}
	if bg && lines[r+1][c-1] == OCCUPIED {
		toRet = false
	}
	if bd && lines[r+1][c+1] == OCCUPIED {
		toRet = false
	}
	return toRet
}

func findFirstSeat(lines []string, r, c, dirr, dirc int) (int, int) {
	i := r + dirr
	j := c +dirc
	for {
		if i < len(lines) && i >= 0 &&j < len(lines[r]) && j >=0 {
			if  lines[i][j] != FLOOR {
				return i, j
			}
			i+= dirr
			j+=dirc
		}else {
			break
		}

	}

	//for i := r; i < len(lines) && i >= 0; i += dirr {
	//	for j := c; j < len(lines[r]) && j >=0; j += dirc {
	//		if i != r && j != c && lines[i][j] != FLOOR {
	//			return i, j
	//		}
	//	}
	//}
	return r, c

}

func canLeave(lines []string, r, c int) bool {
	g, d, h, b := false, false, false, false
	hg, hd, bg, bd := false, false, false, false
	if r-1 >= 0 {
		h = true
	}
	if r+1 < len(lines) {
		b = true
	}
	if c-1 >= 0 {
		g = true
	}
	if c+1 < len(lines[r]) {
		d = true
	}
	hg = (r-1 >= 0) && (c-1 >= 0)
	hd = (r-1 >= 0) && (c+1 < len(lines[r]))
	bg = (r+1 < len(lines)) && c-1 >= 0
	bd = (r+1 < len(lines)) && (c+1 < len(lines[r]))

	reasons := 0
	if h && lines[r-1][c] == OCCUPIED {
		reasons++
	}
	if b && lines[r+1][c] == OCCUPIED {
		reasons++
	}
	if g && lines[r][c-1] == OCCUPIED {
		reasons++
	}
	if d && lines[r][c+1] == OCCUPIED {
		reasons++
	}
	if hg && lines[r-1][c-1] == OCCUPIED {
		reasons++
	}
	if hd && lines[r-1][c+1] == OCCUPIED {
		reasons++
	}
	if bg && lines[r+1][c-1] == OCCUPIED {
		reasons++
	}
	if bd && lines[r+1][c+1] == OCCUPIED {
		reasons++
	}
	return reasons > 4

}

func canOcc2(lines []string, r, c int) bool {
	g, d, h, b := false, false, false, false
	hg, hd, bg, bd := false, false, false, false
	if r-1 >= 0 {
		h = true
	}
	if r+1 < len(lines) {
		b = true
	}
	if c-1 >= 0 {
		g = true
	}
	if c+1 < len(lines[r]) {
		d = true
	}
	hg = r-1 >= 0 && c-1 >= 0
	hd = r-1 >= 0 && c+1 < len(lines[r])
	bg = r+1 < len(lines) && c-1 >= 0
	bd = r+1 < len(lines) && c+1 < len(lines[r])

	toRet := true
	x, y := findFirstSeat(lines, r, c, -1, 0)
	if h && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, +1, 0)
	if b && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, 0, -1)
	if g && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, 0, +1)
	if d && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, -1, -1)
	if hg && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, -1, +1)
	if hd && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, +1, -1)
	if bg && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}
	x, y = findFirstSeat(lines, r, c, +1, +1)
	if bd && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		toRet = false
	}

	return toRet
}
func canLeave2(lines []string,r,c int)  bool{
	g, d, h, b := false, false, false, false
	hg, hd, bg, bd := false, false, false, false
	if r-1 >= 0 {
		h = true
	}
	if r+1 < len(lines) {
		b = true
	}
	if c-1 >= 0 {
		g = true
	}
	if c+1 < len(lines[r]) {
		d = true
	}
	hg = r-1 >= 0 && c-1 >= 0
	hd = r-1 >= 0 && c+1 < len(lines[r])
	bg = r+1 < len(lines) && c-1 >= 0
	bd = r+1 < len(lines) && c+1 < len(lines[r])

	reasons := 0
	x, y := findFirstSeat(lines, r, c, -1, 0)
	if h &&(x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, +1, 0)
	if b && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, 0, -1)
	if g && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, 0, +1)
	if d && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, -1, -1)
	if hg && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, -1, +1)
	if hd && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, +1, -1)
	if bg &&  (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}
	x, y = findFirstSeat(lines, r, c, +1, +1)
	if bd && (x != r || y !=c) && lines[x][y] == OCCUPIED {
		reasons++
	}

	return reasons >= 5
}
