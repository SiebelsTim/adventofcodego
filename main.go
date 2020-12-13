package main

import (
	"exercise1/exercise1"
	"exercise1/exercise2"
	. "exercise1/utils"
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strconv"
	"time"
)

var exercises = []Exercise{
	&exercise1.Exercise1{},
	&exercise2.Exercise2{},
}

func main() {
	if len(os.Args) < 2 {
		color.Red("Must give an exericse number as first argument.")
		os.Exit(1)
	}

	exerciseNumber, err := strconv.Atoi(os.Args[1])
	exerciseNumber -= 1
	if err != nil {
		color.Red("%s is not a number.", os.Args[1])
		os.Exit(1)
	}
	if len(exercises) <= exerciseNumber || exerciseNumber < 0 {
		color.Red("Could not find exercise%d.", exerciseNumber)
		os.Exit(1)
	}

	green := color.New(color.FgGreen)
	boldGreen := color.New(color.FgGreen).Add(color.Bold)

	color.Blue("Starting exercise %d", exerciseNumber+1)
	color.Blue("-------------------------------------")

	startTime := time.Now()
	exercise := exercises[exerciseNumber]
	if exercise.ReadInput() != nil {
		color.Red("ERROR: %s\n", err)
	} else {
		chSol1 := make(chan string)
		chSol2 := make(chan string)
		go exercise.Solution1(chSol1)
		go exercise.Solution2(chSol2)

		if sol1 := <-chSol1; sol1 == "" {
			color.Red("ERROR\n")
		} else {
			fmt.Printf("Solution1: %s\n", boldGreen.Sprint(sol1))
		}

		if sol2 := <-chSol2; sol2 == "" {
			color.Red("ERROR\n")
		} else {
			fmt.Printf("Solution2: %s\n", boldGreen.Sprint(sol2))
		}
	}
	endTime := time.Since(startTime) // nanoseconds

	color.Green("\n-------------------------------------")
	color.Green("Finished exercise %d", exerciseNumber+1)

	fmt.Printf("%s %s\n", green.Sprint("Took"), boldGreen.Sprint(endTime))

	memory := getUsedMemory()
	fmt.Printf("%s %s\n", green.Sprint("RAM:"), boldGreen.Sprintf("%d KB", memory.TotalAlloc/1000))
}

func getUsedMemory() runtime.MemStats {
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	return memstats
}
