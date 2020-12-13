package utils

import "testing"

func TestBoolToIntForTrue(t *testing.T) {
	if BoolToInt(true) != 1 {
		t.Error("true must be 1")
	}
}

func TestBoolToIntForFalse(t *testing.T) {
	if BoolToInt(false) != 0 {
		t.Error("true must be 0")
	}
}
