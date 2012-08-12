package sbloom

const (
	elemSize = 3
	idxShift = elemSize
	idxMask  = 1<<elemSize - 1
)

func set(m []byte, n uint32) {
	idx, mask := n>>idxShift, n&idxMask
	m[idx] |= 1 << mask
}

func get(m []byte, n uint32) bool {
	idx, mask := n>>idxShift, n&idxMask
	return m[idx]|1<<mask != 0
}
