package exercise8

type VM interface {
	Run()
	Step()
	Ip() uint
	Accumulator() int
	Stopped() bool
}

type basicVM struct{
	program Program
	ip uint
	acc int
}

func NewVM(program Program) VM {
	return &basicVM{program: program}
}

func (b *basicVM) Run() {
	for !b.Stopped() {
		b.Step()
	}
}

func (b *basicVM) Step() {
	if b.Stopped() {
		return
	}

	inst := b.program[b.ip]

	switch inst.Kind() {
	case INST_ACC:
		b.acc += inst.Argument()
		break
	case INST_JMP:
		b.ip = uint(int(b.ip) + inst.Argument())
		return
	case INST_NOP:
		break
	}

	b.ip++
}

func (b *basicVM) Ip() uint {
	return b.ip
}

func (b *basicVM) Accumulator() int {
	return b.acc
}

func (b *basicVM) Stopped() bool {
	return int(b.ip) >= len(b.program)
}


