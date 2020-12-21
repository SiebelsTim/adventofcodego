package utils

import (
	"bufio"
	"fmt"
	"os"
)


func ReadInput(n int, test bool) <-chan string {
	suffix := ""
	if test {
		suffix = "test"
	}

	filename := fmt.Sprintf("exercise%d/%sinput.txt", n, suffix)

	return readFileLines(filename)
}

func ReadInputArray(n int, isTest bool) []string {
	return chanToArray(ReadInput(n, isTest))
}

func chanToArray(ch <-chan string) []string {
	var ret []string

	for line := range ch {
		ret = append(ret, line)
	}

	return ret
}

func readFileLines(path string) <-chan string {
	fp, err := os.Open(path)
	if err != nil {
		panic("Cannot open file: " + err.Error())
	}

	ch := make(chan string)
	scanner := bufio.NewScanner(fp)
	scanner.Split(bufio.ScanLines)

	go func () {
		defer fp.Close()
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}