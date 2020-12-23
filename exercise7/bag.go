package exercise7

type bag string

type containedBag struct {
	bag
	count int
}

func (b bag) String() string {
	return string(b)
}

func NewBag(variation string, color string) bag {
	return bag(variation + " " + color)
}