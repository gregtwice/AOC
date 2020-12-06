package helper

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

func MustParseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(str)
		debug.PrintStack()
		log.Fatal(err)
	}
	return i
}
