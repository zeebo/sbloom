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
--- PASS: TestFillSize (2.43 seconds)
	fill_test.go:39: 2: 1.00 (1 1.00) std dev: 0.00
	fill_test.go:39: 4: 2.32 (2 1.16) std dev: 0.66
	fill_test.go:39: 8: 5.08 (4 1.27) std dev: 1.21
	fill_test.go:39: 16: 10.58 (8 1.32) std dev: 1.96
	fill_test.go:39: 32: 21.67 (16 1.35) std dev: 3.09
	fill_test.go:39: 64: 43.77 (32 1.37) std dev: 4.32
	fill_test.go:39: 128: 88.33 (64 1.38) std dev: 6.39
	fill_test.go:39: 256: 176.93 (128 1.38) std dev: 9.03
	fill_test.go:39: 512: 354.78 (256 1.39) std dev: 12.94
	fill_test.go:39: 1024: 710.29 (512 1.39) std dev: 17.79
	fill_test.go:39: 2048: 1419.28 (1024 1.39) std dev: 26.61
	fill_test.go:39: 4096: 2837.60 (2048 1.39) std dev: 35.16
	fill_test.go:39: 8192: 5677.99 (4096 1.39) std dev: 51.15
	fill_test.go:39: 16384: 11355.13 (8192 1.39) std dev: 71.27
PASS
*/
