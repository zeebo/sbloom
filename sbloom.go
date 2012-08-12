package sbloom

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
