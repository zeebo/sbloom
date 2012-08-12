package sbloom

import "hash"

//sHash is a hash with an initial seed.
type sHash struct {
	ha   hash.Hash64
	seed []byte
}

func (s sHash) Hash(p []byte) uint64 {
	s.ha.Reset()
	s.ha.Write(s.seed)
	s.ha.Write(p)
	return s.ha.Sum64()
}
