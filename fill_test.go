package sbloom

import (
	"math"
	"math/rand"
	"testing"
)

func TestFillSize(t *testing.T) {
	sizes := []int{}
	for i := uint(1); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	const iters = 100
	for _, size := range sizes {
		var sum, sumsq float64
		for i := 0; i < iters; i++ {
			var n int
			x := make([]int, size)
			count := size / 2
			for n = 0; count > 0; n++ {
				idx := rand.Intn(size)
				if x[idx] == 0 {
					count--
				}
				x[idx]++
			}
			sum += float64(n)
			sumsq += float64(n) * float64(n)
		}

		t.Logf("%d: %.2f (%d %.2f) std dev: %.2f",
			size,
			sum/iters,
			size/2,
			(sum/iters)/float64(size/2),
			math.Sqrt(sumsq*iters-sum*sum)/iters,
		)
	}
}

/*
=== RUN TestFillSize
--- PASS: TestFillSize (0.01 seconds)
	fill_test.go:39: 2: 1.00 (1 1.00) std dev: 0.00
	fill_test.go:39: 4: 2.32 (2 1.16) std dev: 0.58
	fill_test.go:39: 8: 5.10 (4 1.27) std dev: 1.28
	fill_test.go:39: 16: 10.49 (8 1.31) std dev: 2.10
	fill_test.go:39: 32: 21.53 (16 1.35) std dev: 2.96
	fill_test.go:39: 64: 45.41 (32 1.42) std dev: 4.96
	fill_test.go:39: 128: 88.83 (64 1.39) std dev: 6.13
	fill_test.go:39: 256: 178.09 (128 1.39) std dev: 8.72
	fill_test.go:39: 512: 354.53 (256 1.38) std dev: 13.39
PASS
*/
