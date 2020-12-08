package main

import (
	"AOC/helper"
	"fmt"
	"strings"
)

func main() {

	lines := helper.StringArrayFromFile("./AOC2020/day8/sample.txt")
	//lines := helper.StringArrayFromFile("./AOC2020/day8/exemple.txt")

	for i := 0; i < len(lines); i++ {
		orig := lines[i]
		inst := strings.Split(lines[i], " ")
		if inst[0] == "nop" {
			inst[0] = "jmp"
		} else if inst[0] == "jmp" {
			inst[0] = "nop"
		}
		lines[i] = inst[0] + " " + inst[1]
		tryRun(lines)
		lines[i] = orig
	}

}

func tryRun(lines []string) {
	visited := make([]bool, len(lines))
	i := 0
	acc := 0
	for {
		if i >= len(lines) {
			fmt.Println("FINI !!!", acc)
			break
		}
		inst := strings.Split(lines[i], " ")

		if visited[i] {
			//println(acc)
			break
		}
		//
		//if inst[0] == "nop" {
		//	inst[0] = "jmp"
		//} else if inst[0] == "jmp" {
		//	inst[0] = "nop"
		//}

		visited[i] = true
		switch inst[0] {
		case "nop":
			i++
		case "acc":
			acc += helper.MustParseInt(inst[1])
			i++
		case "jmp":
			i += helper.MustParseInt(inst[1])

		}

	}
}
