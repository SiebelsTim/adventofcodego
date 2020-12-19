package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadInput(n int) []string {
	return readInput(n, false)
}

func ReadInputTest(n int) []string {
	return readInput(n, true)
}

func readInput(n int, test bool) []string {
	suffix := ""
	if test {
		suffix = "test"
	}

	filename := fmt.Sprintf("exercise%d/%sinput.txt", n, suffix)
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	content := string(data)

	return strings.Split(content, "\n")
}
