package main

import (
	"AOC/helper"
	"strings"
)

func main() {

	s := helper.GetFile("./AOC2020/day6/sample.txt")
	//s := helper.GetFile("./AOC2020/day6/exemple.txt")

	answerMap := map[int32]int{}
	total1 := 0
	total2 := 0
	nbPers := 0
	answerStr := ""
	for	s.Scan(){
		answerStr += strings.TrimSuffix(s.Text(),"\n")
		if s.Text() =="" {
			incTotal(answerStr, answerMap, &total1)
			incTotal2(answerMap, nbPers, &total2)
			answerStr = ""
			nbPers = 0
			answerMap = map[int32]int{}
		}else {
			nbPers++
		}
	}
	incTotal(answerStr, answerMap, &total1)
	incTotal2(answerMap, nbPers, &total2)

	println(total1)
	println(total2)

}

func incTotal2(answerMap map[int32]int, nbPers int, total2* int) {
	for _, nb := range answerMap {
		if nb == nbPers {
			*total2++
		}
	}
}

func incTotal(answerStr string, answerMap map[int32]int, total *int) {
	for _, c := range answerStr {
		answerMap[c] ++
	}
	for range answerMap {
		*total++
	}
}