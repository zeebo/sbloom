package sbloom

import (
	"bytes"
	"encoding/gob"
	"errors"
	"hash"
	"hash/fnv"
)

func init() {
	//register the fnv type as it is commonly used
	gob.Register(fnv.New64())
}

//gobFilter is an internal type that gob will use to represent a Filter.
type gobFilter struct {
	Hash    hash.Hash64
	Hashes  []sHash
	Filters []*filter
}

//GobEncode returns the gob marshalled value of the filter.
func (f *Filter) GobEncode() (p []byte, err error) {
	gf := gobFilter{
		Hash:    f.bh,
		Hashes:  f.hs,
		Filters: f.fs,
	}
	var buf bytes.Buffer
	err = gob.NewEncoder(&buf).Encode(gf)
	if err == nil {
		p = buf.Bytes()
	}
	return
}

//GobDecode sets the filters state to the gob marshalled value in the buffer.
func (f *Filter) GobDecode(p []byte) (err error) {
	var gf gobFilter
	buf := bytes.NewReader(p)
	err = gob.NewDecoder(buf).Decode(&gf)
	if err != nil {
		return
	}

	if len(gf.Hashes) == 0 {
		err = errors.New("no hash function specified")
		return
	}

	f.bh = gf.Hash
	f.hs = gf.Hashes
	f.fs = gf.Filters
	return
}
