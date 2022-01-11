package mymath

import (
	"testing"
)

func benchmarkAverage(b *testing.B, input ...float64){
	for i:=0 ; i< b.N; i++ {
		Average(input...)
	}
}

func BenchmarkAverageNeg(b *testing.B) {
	benchmarkAverage(b, -2, -3)	// Case 1
}

func BenchmarkAveragePos(b *testing.B) {
	benchmarkAverage(b, 2, 4) 	// Case 2
}

func BenchmarkAveragePosNeg(b *testing.B) {
	benchmarkAverage(b, -2, 3) 	// Case 3
}

func BenchmarkAverageZero(b *testing.B) {
	benchmarkAverage(b, 0) 		// Case 4
}
