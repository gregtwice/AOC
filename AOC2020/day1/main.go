package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {



	//file ,err := os.Open("exemple.txt")
	file ,err := os.Open("sample.txt")
	if err != nil {
		panic(err)
	}

	nums := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n,err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			for k := j+1; k < len(nums); k++ {
				if  nums[i] + nums[j] + nums[k] == 2020{
					fmt.Println(nums[i] * nums[j] * nums[k])
				}
			}
		}
	}

}

func PrintPart1(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				fmt.Println(i, j, nums[i], "+", nums[j], nums[i]*nums[j])
			}
		}
	}
}
