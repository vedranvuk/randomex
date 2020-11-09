package randomex

import (
	"testing"
)

func TestString(t *testing.T) {
	s := ""
	for i := 0; i < 100; i++ {
		if ns := String(true, true, true, true, 16); s == ns {
			t.Fatal("well, it failed.")
		} else {
			s = ns
		}
	}
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		String(true, true, true, true, 16)
	}
}
