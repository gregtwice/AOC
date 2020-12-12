package main

import (
	"AOC/helper"
	"fmt"
	"sort"
)

func main() {
	//lines := helper.StringArrayFromFile("./AOC2020/day10/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day10/exemple.txt")
	lines := helper.StringArrayFromFile("./AOC2020/day10/ex2.txt")
	nums := []int{}
	for _, line := range lines {
		nums = append(nums, helper.MustParseInt(line))
	}
	sort.Ints(nums)
	diff3 := 1
	diff1 := 1
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]

		if diff == 3 {
			diff3++
		} else if diff == 1 {
			diff1++
		}
	}
	fmt.Println(diff1, diff3, diff3*diff1)

	nums = append(nums, 0)
	sort.Ints(nums)

	slices := []permuts{}

	baseI := 0
	for i := 1; i < len(nums); i++ {
		fmt.Print(nums[i-1], " ")
		if nums[i]-nums[i-1] == 3 {
			ints := nums[baseI:i]
			fmt.Println(ints)
			baseI = i
			slices = append(slices, permuts{
				nums:  ints,
				nbPos: 0,
			})
		}
	}
	slices = append(slices, permuts{
		nums:  nums[baseI:len(nums)],
		nbPos: 0,
	})
	fmt.Println(slices)

	for i := 0; i < len(slices); i++ {
		part2(&slices[i], 0)
	}
	fmt.Println("\n\n\n")
	fmt.Println(nums)
	fmt.Println(slices)
	acc := 1
	for _, s := range slices {
		acc *= s.nbPos
	}
	fmt.Println(acc)
	dp = map[int]int{}
	fmt.Println(part2DP(nums,0))

}

type permuts struct {
	nums  []int
	nbPos int
}

func part2(perm *permuts, curI int) {
	if perm.nums[curI] == perm.nums[len(perm.nums)-1] /*49*/ /*19*/ { //167 {
		perm.nbPos++
		return
	}
	for i := curI + 1; i < curI+4 && i < len(perm.nums); i++ {
		if perm.nums[i]-perm.nums[curI] < 4 {
			part2(perm, i)
		}
	}
}

var dp map[int]int

func part2DP(nums []int, i int) int {

	if i == len(nums)-1 {
		return 1
	}
	if val, ok := dp[i]; ok {
		return val
	}
	ans := 0
	for j := i + 1; j < len(nums); j++ {
		if nums[j]-nums[i] <= 3 {
			ans += part2DP(nums, j)
		}
	}
	dp[i] = ans
	return ans
}
