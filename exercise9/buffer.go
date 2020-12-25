package exercise9

type RingBuffer interface {
	Values() []int
	Add(int)
	Get(int) int
	Length() int
}

// This RingBuffer is always filled up to len(values)
// startIdx is the index of the first element. The last is right before this one
type sliceRingerBuffer struct {
	values   []int
	startIdx int
}

func (b *sliceRingerBuffer) Values() []int {
	var ret = make([]int, len(b.values))

	for i := 0; i < len(b.values); i++ {
		ret[i] = b.values[(b.startIdx+i)%len(b.values)]
	}

	return ret
}

func (b *sliceRingerBuffer) Add(value int) {
	b.startIdx = (b.startIdx + 1) % len(b.values)

	b.values[(b.startIdx + len(b.values) - 1) % len(b.values)] = value
}

func (b *sliceRingerBuffer) Get(idx int) int {
	return b.values[(b.startIdx + idx) % len(b.values)]
}

func (b *sliceRingerBuffer) Length() int {
	return len(b.values)
}

func NewBuffer(values []int) RingBuffer {
	return &sliceRingerBuffer{
		values:   values,
		startIdx: 0,
	}
}
