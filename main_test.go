package main__test

import (
	"math"
	"testing"
)

// BenchmarkIntMod benchmarks integer modulus operation.
func BenchmarkIntMod(b *testing.B) {
	var result int
	x := 123456789
	y := 12345

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = x % y
		result += x % y
		result += x % y
	}
	_ = result // Use result to avoid compiler optimizations
}

// BenchmarkFloatMod benchmarks floating-point modulus operation using math.Mod.
func BenchmarkFloatMod(b *testing.B) {
	var result float64

	x := 123456789.0
	y := 12345.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = math.Mod(x, y)
		result += math.Mod(x, y)
		result += math.Mod(x, y)
	}
	_ = result // Use result to avoid compiler optimizations
}

func BenchmarkFloatModTrunc(b *testing.B) {
	var result float64

	x := 123456789.0
	y := 12345.0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result = x - math.Trunc(x/y)*y
		result += x - math.Trunc(x/y)*y
		result += x - math.Trunc(x/y)*y
	}
	_ = result // Use result to avoid compiler optimizations
}
