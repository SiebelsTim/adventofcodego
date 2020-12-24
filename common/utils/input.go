package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var basepath = getBasePath()

func ReadInput(n int, test bool) <-chan string {
	suffix := ""
	if test {
		suffix = "test"
	}

	filename := fmt.Sprintf("%s/exercise%d/%sinput.txt", basepath, n, suffix)

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

	go func() {
		defer fp.Close()
		defer close(ch)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	return ch
}

func getBasePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd[:strings.Index(cwd, "adventofcodego")+len("adventofcodego")]
}
