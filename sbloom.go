package sbloom

import (
	"hash"
	"math/rand"
)

const (
	elemSize = 3
	idxShift = elemSize
	idxMask  = 1<<elemSize - 1
)

//set turns the nth bit in the slice of bytes to one.
func set(m []uint8, n uint64) {
	idx, mask := n>>idxShift, n&idxMask
	m[idx] |= 1 << mask
}

//get returns true if the nth bit in the slice of bytes is one.
func get(m []uint8, n uint64) bool {
	idx, mask := n>>idxShift, n&idxMask
	return m[idx]&(1<<mask) != 0
}

type Filter struct {
	bh hash.Hash64
	fs []*filter
	hs []sHash
}

func (f *Filter) Add(p []byte) {
	last := f.fs[len(f.fs)-1]
	last.Add(p, f.hs[:last.k])
	if last.left == 0 {
		newSize := last.size << 1
		newMask := last.mask<<1 + 1
		newK := last.k + 1
		f.addNewFilter(newSize, newMask, newK)
	}
}

func (f *Filter) Lookup(p []byte) bool {
	for _, subfil := range f.fs {
		if subfil.Lookup(p, f.hs[:subfil.k]) {
			return true
		}
	}
	return false
}

func (f *Filter) addNewFilter(size, mask uint64, k int) {
	bins := make([][]uint8, k+1) //allocate one more bin
	binSize := size >> elemSize  //size of each bin to have size bits
	for i := range bins {
		bins[i] = make([]uint8, binSize)
	}

	//make sure we have up to k hashes
	for len(f.hs) < k {
		f.hs = append(f.hs, sHash{
			ha:   f.bh,
			seed: randSeed(),
		})
	}

	//add the new bloom filter
	f.fs = append(f.fs, &filter{
		size: size,
		mask: mask,
		bins: bins,
		k:    k,
		left: size / 5 * 7,
	})
}

func randSeed() (p []byte) {
	for i := 0; i < 4; i++ {
		p = append(p, byte(rand.Intn(256)))
	}
	return
}

type filter struct {
	size uint64 // N == 1 << size
	mask uint64 // x % N == x & mask == x & (1 << size) - 1

	k    int
	bins [][]uint8 // [k][1<<size]uint8 bins
	left uint64    // number of additions left until new filter
}

func (f *filter) Add(p []byte, hs []sHash) {
	for i, h := range hs {
		val := h.Hash(p)
		set(f.bins[i], val&f.mask)
	}
	f.left--
}

func (f *filter) Lookup(p []byte, hs []sHash) bool {
	for i, h := range hs {
		val := h.Hash(p)
		if !get(f.bins[i], val&f.mask) {
			return false
		}
	}
	return true
}
