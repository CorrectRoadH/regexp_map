package benchmark

import "testing"

func benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkMap(b *testing.B) {
	benchmark(b)
}

func BenchmarkRegexpMap(b *testing.B) {
	benchmark(b)
}
