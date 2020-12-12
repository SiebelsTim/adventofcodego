package utils

type Exercise interface {
	ReadInput() error
	Solution1(solution chan string)
	Solution2(solution chan string)
}

type SingleSolution struct {
	Solution string
	Err error
}

type Solution struct {
	Solution1 string
	Solution2 string
	Err error
}

func Err(err error) Solution {
	return Solution { "", "", err }
}

func Success(solution1 string, solution2 string) Solution {
	return Solution { solution1, solution2, nil}
}