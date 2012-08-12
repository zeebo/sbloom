package sbloom

import (
	"math/rand"
	"testing"
)

const maxSize = 10

func TestSetAll(t *testing.T) {
	sizes := []int{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]byte, size)
		for i := 0; i < size*8; i++ {
			set(x, uint32(i))
		}
		for i, xi := range x {
			if xi != 255 {
				t.Errorf("%d: %b", i, xi)
			}
		}
	}
}

func TestSetRandom(t *testing.T) {
	sizes := []int{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]byte, size)

		//set 10 random bytes
		for i := 0; i < 10; i++ {
			idx := rand.Uint32() % uint32(size*8)
			set(x, idx)
			//make sure that byte is > 0
			n := idx >> 3
			if xi := x[n]; xi == 0 {
				t.Errorf("%d: %b", n, xi)
			}
		}
	}
}

func TestGetAll(t *testing.T) {
	sizes := []int{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]byte, size)
		for i := range x {
			x[i] = 255
		}

		for i := 0; i < size*8; i++ {
			if !get(x, uint32(i)) {
				t.Errorf("%d", i)
			}
		}
	}
}

func TestGetRandom(t *testing.T) {
	sizes := []int{}
	for i := uint(0); i < maxSize; i++ {
		sizes = append(sizes, 1<<i)
	}

	for _, size := range sizes {
		x := make([]byte, size)

		//set 10 random bytes
		for i := 0; i < 10; i++ {
			idx := rand.Uint32() % uint32(size*8)
			set(x, idx)
			if !get(x, idx) {
				t.Errorf("%d", i)
			}
		}
	}
}
