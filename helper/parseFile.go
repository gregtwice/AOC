package helper

import (
	"bufio"
	"os"
	"strings"
)

func GetFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	return scanner
}

func StringArrayFromFile(filename string) []string {
	scanner := GetFile(filename)
	arr := []string{}
	for scanner.Scan() {
		arr = append(arr, scanner.Text())
	}
	return arr
}

func StringArrayFromSingleLine(filename, sep string) []string {
	scanner := GetFile(filename)
	scanner.Scan()
	return strings.Split(scanner.Text(), sep)
}
