package solution

type Solution interface {
	String() string
}

type Exercise interface {
	Prepare() error
	Solution1() (Solution, error)
	Solution2() (Solution, error)
}

type stringSolution struct {
	solution string
}

func (s stringSolution) String() string {
	return s.solution
}

func New(solution string) Solution {
	return stringSolution{
		solution,
	}
}