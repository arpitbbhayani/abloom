package abloom

import (
	"hash"

	"github.com/twmb/murmur3"
)

type bloom struct {
	filter []byte
}

var fns [2]hash.Hash32

func init() {
	fns[0] = murmur3.SeedNew32(3719237192)
	fns[1] = murmur3.SeedNew32(8981437)
}

// NewBloom creates a bloom filter of length `len` bytes
// The number of hash functions are fixed at 2
func NewBloom(len int) *bloom {
	return &bloom{
		filter: make([]byte, len),
	}
}

// Put puts the element `x` in the bloom filter
func (b *bloom) Put(x []byte) error {
	var err error
	for i := range fns {
		fns[i].Reset()
		if _, err = fns[i].Write(x); err != nil {
			return err
		}
		pos := int(fns[i].Sum32()) % (len(b.filter) * 8)
		idx, offset := pos/8, pos%8
		b.filter[idx] = b.filter[idx] | 1<<offset
	}
	return nil
}

// Check checks the existence of the element `x` in the bloom filter
func (b *bloom) Check(x []byte) (bool, error) {
	var err error
	var keyExists bool = true
	for i := range fns {
		fns[i].Reset()
		if _, err = fns[i].Write(x); err != nil {
			return false, err
		}
		pos := int(fns[i].Sum32()) % (len(b.filter) * 8)
		idx, offset := pos/8, pos%8
		if b.filter[idx]&(1<<offset) == 0 {
			keyExists = false
			break
		}
	}
	return keyExists, nil
}
