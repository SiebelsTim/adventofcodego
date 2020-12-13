package utils

import "testing"

func TestStringSolution_String(t *testing.T) {
	solution := New("my message")

	if solution.String() != "my message" {
		t.Error("Message mismatch")
	}
}
