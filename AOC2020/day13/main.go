package main

import (
	"AOC/helper"
	"fmt"
	"strings"
)

type Bus struct {
	id     int
	offset int
}

func main() {
	lines := helper.StringArrayFromFile("./AOC2020/day13/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day13/exemple.txt")
	busesStr := strings.Split(lines[1], ",")
	bArr := []Bus{}
	for i, bus := range busesStr {
		offset := i
		if bus == "x" {
			continue
		}
		//if bus == "7" {
		//	offset = 2
		//}
		bArr = append(bArr, Bus{
			id:     helper.MustParseInt(bus),
			offset: offset,
		})
	}
	merde(bArr)

	//printPart1(lines)

}

func merde(bArr []Bus) {
	n := 1
	for _, b := range bArr {
		n *= b.id
	}
	somme := 0
	for _, b := range bArr {
		ni := b.id
		nip := n / ni
		_, u, v := euclide(ni, nip)
		fmt.Println("euclide := ", u*ni+v*nip)
		ei := v * nip
		fmt.Println("ei := ", ei)
		fmt.Println("ei % ni = " ,ei%ni)
		diff := b.id - b.offset
		//for diff < 0 {
		//	diff += b.id
		//}
		somme += ei * diff
	}
	fmt.Println(somme%n)
	for somme < 0 {
		somme += n
	}
	fmt.Println(somme)
	fmt.Println(bArr)
	for _, bus := range bArr[0:] {
		fmt.Println("somme = ", somme, " % id (", bus.id, ") = ", somme%bus.id, )
	}

}

func euclide(a, b int) (r, u, v int) {
	/*
	   	r, u, v, r', u', v') := (a, 1, 0, b, 0, 1)
	                        q  quotient entier

	      les égalités r = a*u+b*v et r' = a*u'+b*v' sont des invariants de boucle

	      tant que (r' ≠ 0) faire
	          q := r÷r'
	          (r, u, v, r', u', v') := (r', u', v', r - q *r', u - q*u', v - q*v')
	          fait
	      renvoyer (r, u, v)
	*/

	r, u, v, r1, u1, v1 := a, 1, 0, b, 0, 1
	q := 0
	for r1 != 0 {
		q = r / r1
		r, u, v, r1, u1, v1 = r1, u1, v1, r-q*r1, u-q*u1, v-q*v1
	}
	return r, u, v

}

func part2long(bArr []Bus) {
	//btime := 0 //1201000
	//btime := 1202100000
	btime := 100000000000000
	//btime := 0
	busTimes := make([]int, len(bArr))
	busTimes[0] = btime
	for !isOk(bArr, busTimes, len(bArr)-1) {
		//fmt.Println(busTimes)
		for i := 1; i < len(bArr); i++ {
			a, b := getMultiples2(bArr[0].id, bArr[i].id, btime, bArr[i].offset)

			//fmt.Println(a,b)
			//if a > 1261476{
			//	return
			//}
			busTimes[i] = b
			busTimes[0] = helper.Max(busTimes[0], a)
		}
		btime = busTimes[0] - bArr[0].id
	}
	fmt.Println(busTimes)

}

func printPart1(lines []string) {
	timestamp := helper.MustParseInt(lines[0])

	buses := strings.Split(strings.Replace(lines[1], ",x", "", -1), ",")

	bestBus := 0
	bestTime := timestamp + 100000
	for _, bus := range buses {
		j := 0
		for ; j < timestamp; j += helper.MustParseInt(bus) {
		}
		if j < bestTime {
			bestBus = helper.MustParseInt(bus)
			bestTime = j
		}
	}
	fmt.Println((bestTime - timestamp) * bestBus)
}

func isOk(buses []Bus, btimes []int, max int) bool {
	for i := 1; i <= max; i++ {
		if btimes[i]-btimes[0] != buses[i].offset {
			return false
		}
	}
	return true
}

func findNextMultiple(baseNum, num int) int {
	multiple := baseNum + num + 1
	multiple -= (multiple % num)
	return multiple
}

func getMultiples2(a, b, base, offset int) (int, int) {
	mulA := findNextMultiple(base, a)
	for (mulA+offset)%b != 0 {
		mulA += a
	}
	return mulA, mulA + offset
}

func getMultiples(a, b, base, offset int) (int, int) {

	mulA := findNextMultiple(base, a)
	mulB := findNextMultiple(base, b)
	for mulB-mulA != offset {
		if mulA < mulB {
			diff := mulB - mulA - a
			toMul := diff / a
			mulA += a * toMul
		} else if mulB < mulA {
			diff := mulA - mulB - b
			toMul := diff / b
			mulB += b * toMul
		}
	}
	return mulA, mulB
}
