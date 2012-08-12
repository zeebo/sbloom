package sbloom

import "hash"

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
}

func (f Filter) Add(p []byte) {
	last := f.fs[len(f.fs)-1]
	last.Add(p)
	if last.left == 0 {
		//this one filled, add a new one
	}
}

func (f Filter) Lookup(p []byte) bool {
	for _, subfil := range f.fs {
		if subfil.Lookup(p) {
			return true
		}
	}
	return false
}

type filter struct {
	size uint   // N == 1 << size
	mask uint64 // x % N == x & mask == x & (1 << size) - 1

	bins [][]uint8 // [k][1<<size]uint8 bins
	hs   []sHash   // [k]sHash hashes
	left uint64    // number of additions left until new filter
}

func (f *filter) Add(p []byte) {
	for i, h := range f.hs {
		val := h.Hash(p)
		set(f.bins[i], val&f.mask)
	}
	f.left--
}

func (f *filter) Lookup(p []byte) bool {
	for i, h := range f.hs {
		val := h.Hash(p)
		if !get(f.bins[i], val&f.mask) {
			return false
		}
	}
	return true
}
