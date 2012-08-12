package sbloom

import (
	"math/rand"
	"testing"
)

func randNum(top uint64) uint64 {
	hi, low := rand.Uint32(), rand.Uint32()
	return (uint64(hi)<<32 | uint64(low)) % top
}

const maxSize = 10

func BenchmarkSet(b *testing.B) {
	const size = 1 << maxSize
	x := make([]uint8, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := uint64(0); i < size*(1<<elemSize); i++ {
			set(x, i)
		}
		b.SetBytes(size * (1 << (elemSize - 3)))
	}
}

func BenchmarkGet(b *testing.B) {
	const size = 1 << maxSize
	x := make([]uint8, size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := uint64(0); i < size*(1<<elemSize); i++ {
			get(x, i)
		}
		b.SetBytes(size * (1 << (elemSize - 3)))
	}
}

func TestSetAll(t *testing.T) {
	sizes := []uint64{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]uint8, size)
		for i := uint64(0); i < size*(1<<elemSize); i++ {
			set(x, i)
		}
		for i, xi := range x {
			if xi != 1<<(1<<elemSize)-1 {
				t.Errorf("%d: %b", i, xi)
			}
		}
	}
}

func TestSetRandom(t *testing.T) {
	sizes := []uint64{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]uint8, size)

		//set 10 random bytes
		for i := 0; i < 10; i++ {
			idx := randNum(size * (1 << elemSize))
			set(x, idx)
			//make sure that byte is > 0
			n := idx >> elemSize
			if xi := x[n]; xi == 0 {
				t.Errorf("%d: %b", n, xi)
			}
		}
	}
}

func TestGetAll(t *testing.T) {
	sizes := []uint64{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]uint8, size)
		for i := range x {
			x[i] = 1<<(1<<elemSize) - 1
		}

		for i := uint64(0); i < size*(1<<elemSize); i++ {
			if !get(x, i) {
				t.Errorf("%d", i)
			}
		}
	}
}

func TestGetRandom(t *testing.T) {
	sizes := []uint64{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]uint8, size)

		//set 10 random bytes
		for i := 0; i < 10; i++ {
			idx := randNum(size * (1 << elemSize))
			set(x, idx)
			if !get(x, idx) {
				t.Errorf("%d", i)
			}
		}
	}
}
