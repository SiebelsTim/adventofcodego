package exercise9

import (
	"reflect"
	"testing"
)

func TestSliceRingerBuffer(t *testing.T) {
	values := []int{
		1, 2, 3, 4, 5, 6,
	}
	buf := NewBuffer(values)

	if !reflect.DeepEqual(values, buf.Values()) {
		t.Errorf("Values don't match. %v", buf.Values())
	}

	buf.Add(7)

	expect := []int{
		2, 3, 4, 5, 6, 7,
	}
	if !reflect.DeepEqual(expect, buf.Values()) {
		t.Errorf("Values don't match. %v", buf.Values())
	}

	buf.Add(10)
	buf.Add(11)
	buf.Add(12)
	buf.Add(13)
	buf.Add(14)
	buf.Add(15)

	expect = []int{
		10, 11, 12, 13, 14, 15,
	}
	if !reflect.DeepEqual(expect, buf.Values()) {
		t.Errorf("Values don't match. %v", buf.Values())
	}
}
