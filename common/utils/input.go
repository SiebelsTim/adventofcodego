package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadInput(n int) <-chan string {
	return readInput(n, false)
}


func ReadInputTest(n int) <-chan string {
	return readInput(n, true)
}

func ReadInputArray(n int) []string {
	return chanToArray(readInput(n, false))
}

func ReadInputTestArray(n int) []string {
	return chanToArray(readInput(n, true))
}

func chanToArray(ch <-chan string) []string {
	var ret []string

	for line := range ch {
		ret = append(ret, line)
	}

	return ret
}

func readInput(n int, test bool) <-chan string {
	suffix := ""
	if test {
		suffix = "test"
	}

	filename := fmt.Sprintf("exercise%d/%sinput.txt", n, suffix)

	return readFileLines(filename)
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