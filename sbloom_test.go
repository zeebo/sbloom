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

var sizes []uint64

func init() {
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}
}

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

func TestOneItem(t *testing.T) {
	x := make([]uint8, 1)
	for j := uint64(0); j < 1<<elemSize; j++ {
		if get(x, j) {
			t.Errorf("0[%d] %08b", j, x[0])
		}
	}
	for i := uint64(0); i < 1<<elemSize; i++ {
		set(x, i)
		t.Logf("%08b", x[0])
		for j := uint64(0); j < 1<<elemSize; j++ {
			if get(x, j) != (j <= i) { //get is true when j <= i
				t.Errorf("%d[%d] %08b", i, j, x[0])
			}
		}
	}
}

func TestPanicGetLarge(t *testing.T) {
	recov := func() {
		if recover() == nil {
			t.Fatal("no panic")
		}
	}

	for _, size := range sizes {
		func() {
			defer recov()
			x := make([]uint8, size)
			get(x, size*(1<<elemSize))
		}()

		func() {
			defer recov()
			x := make([]uint8, size)
			set(x, size*(1<<elemSize))
		}()
	}
}

func TestSetAll(t *testing.T) {
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

func TestSetAndGetRandom(t *testing.T) {
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
