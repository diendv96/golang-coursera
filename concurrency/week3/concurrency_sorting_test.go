package week3

import (
	"testing"
)

func BenchmarkConcurrentSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
