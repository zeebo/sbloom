package sbloom

import (
	"hash/fnv"
	"reflect"
	"testing"
)

func TestGobDeepEqual(t *testing.T) {
	f := NewFilter(fnv.New64(), 8)

	m, err := f.GobEncode()
	if err != nil {
		t.Fatal(err)
	}

	g := new(Filter)
	if err := g.GobDecode(m); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(f, g) {
		t.Fatal("Not equal after gob")
	}
}
