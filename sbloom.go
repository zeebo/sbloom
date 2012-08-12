package sbloom

const (
	elemSize = 3
	idxShift = elemSize
	idxMask  = 1<<elemSize - 1
)

func set(m []uint8, n uint64) {
	idx, mask := n>>idxShift, n&idxMask
	m[idx] |= 1 << mask
}

func get(m []uint8, n uint64) bool {
	idx, mask := n>>idxShift, n&idxMask
	return m[idx]|1<<mask != 0
}
