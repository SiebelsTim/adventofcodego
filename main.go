package main

import (
	"adventofcode/common/format"
	"adventofcode/common/solution"
	"adventofcode/exercise1"
	"adventofcode/exercise2"
	"adventofcode/exercise3"
	"adventofcode/exercise4"
	"adventofcode/exercise5"
	"adventofcode/exercise6"
	"adventofcode/exercise7"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"time"
)

var exercises = []solution.Exercise{
	&exercise1.Exercise1{},
	&exercise2.Exercise2{},
	&exercise3.Exercise3{},
	&exercise4.Exercise4{},
	&exercise5.Exericse5{},
	&exercise6.Exericse6{},
	&exercise7.Exericse7{},
}


var green = color.New(color.FgGreen)
var boldGreen = color.New(color.FgGreen).Add(color.Bold)

func main() {
	isTest := flag.Bool("test", false, "Specify -test to run with test data")
	isAll := flag.Bool("all", false, "Use to run all exercise afert each other")
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}


	if *isTest {
		color.Yellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		color.Yellow("!!  You are using test data          !!")
		color.Yellow("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
	}

	if *isAll {
		for idx, _ := range exercises {
			runProgram(idx, isTest)
		}
	} else {
		exerciseNumber, err := strconv.Atoi(flag.Arg(0))
		exerciseNumber -= 1
		if err != nil {
			color.Red("%s is not a number.", os.Args[1])
			os.Exit(1)
		}

		count := 1
		if *cpuprofile != "" {
			count = 2000
		}
		for i:=0; i < count; i++ {
			runProgram(exerciseNumber, isTest)
		}
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

func runProgram(exerciseNumber int, isTest *bool) {
	if len(exercises) <= exerciseNumber || exerciseNumber < 0 {
		color.Red("Could not find exercise%d.", exerciseNumber+1)
		os.Exit(1)
	}

	color.Blue("Starting exercise %d", exerciseNumber+1)
	color.Blue("-------------------------------------")

	endTime := runExercise(exerciseNumber, *isTest)

	color.Green("\n-------------------------------------")
	color.Green("Finished exercise %d", exerciseNumber+1)

	fmt.Printf("%s %s\n", green.Sprint("Took"), boldGreen.Sprint(endTime))

	memory, m := format.GetUsedMemory()
	fmt.Printf("%s %s\n", green.Sprint("RAM:"), boldGreen.Sprintf("%s", m.String()))
	fmt.Printf("%s %s\n", green.Sprint("Total Number of GCs:"), boldGreen.Sprintf("%d", memory.NumGC))
	fmt.Printf("%s %s\n", green.Sprint("Total Time stopped:"), boldGreen.Sprintf("%d ns", memory.PauseTotalNs))
	fmt.Printf("%s %s\n", green.Sprint("Total Time stopped:"), boldGreen.Sprintf("%d ns", memory.PauseTotalNs))
	fmt.Printf("Goroutine count: %s\n", boldGreen.Sprintf("%d", runtime.NumGoroutine()))
}

func runExercise(n int, isTest bool) time.Duration {

	startTime := time.Now()
	exercise := exercises[n]
	if err := exercise.Prepare(isTest); err != nil {
		color.Red("ERROR: %s\n", err)
	} else {
		fmt.Printf("Preparing took %s\n\n", boldGreen.Sprint(time.Since(startTime)))
		chSol1 := make(chan string, 1)
		chTime1 := make(chan time.Duration)
		chSol2 := make(chan string, 1)
		chTime2 := make(chan time.Duration)
		go func() {
			startTime := time.Now()
			solution, err := exercise.Solution1()

			if err != nil {
				color.Red("Error %s", err.Error())
			} else {
				chSol1 <- solution.String()
			}
			chTime1 <- time.Since(startTime)
		}()

		go func() {
			startTime := time.Now()
			solution, err := exercise.Solution2()

			if err != nil {
				color.Red("Error %s", err.Error())
				close(chSol2)
			} else {
				chSol2 <- solution.String()
			}
			chTime2 <- time.Since(startTime)
		}()

		if sol1 := <-chSol1; sol1 == "" {
			color.Red("ERROR\n")
		} else {
			fmt.Printf("Solution1: %s\n", boldGreen.Sprint(sol1))
			fmt.Printf("%s %s\n\n", green.Sprint("Took"), boldGreen.Sprint(<-chTime1))
		}

		if sol2 := <-chSol2; sol2 == "" {
			color.Red("ERROR\n")
		} else {
			fmt.Printf("Solution2: %s\n", boldGreen.Sprint(sol2))
			fmt.Printf("%s %s\n\n", green.Sprint("Took"), boldGreen.Sprint(<-chTime2))
		}
	}

	return time.Since(startTime) // nanoseconds
}