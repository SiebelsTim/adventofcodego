package exercise7

import (
	"testing"
)

func BenchmarkExericse7_Prepare(b *testing.B) {
	exercise := Exericse7{}
	for i := 0; i < b.N; i++ {
		exercise.Prepare(false)
	}
}

