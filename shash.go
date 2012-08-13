package sbloom

import (
	"hash"
	"math/rand"
)

//sHash is a hash with an initial seed.
type sHash struct {
	ha   hash.Hash64
	seed []byte
}

func newsHash(ha hash.Hash64) (s sHash) {
	s.ha = ha
	s.seed = randSeed()
	return
}

func randSeed() (p []byte) {
	for i := 0; i < 10; i++ {
		p = append(p, byte(rand.Intn(256)))
	}
	return
}

func (s sHash) Hash(p []byte) uint64 {
	s.ha.Reset()
	s.ha.Write(s.seed)
	s.ha.Write(p)
	return s.ha.Sum64()
}
