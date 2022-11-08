package abloom

import (
	"hash"
	"math/rand"

	"github.com/twmb/murmur3"
)

const DefaultHashFns int = 2

type Bloom struct {
	fns    []hash.Hash32
	filter []byte
}

// NewBloom creates a bloom filter of length `size` bytes and
// `hashSeeds` are the seed values for themurmur hash functions
// bloom will be initialized with len(hashSeeds) hash functions
// with provided seeds.
// if no hashSeeds are provided then 2 hash functions will be used
// with random seeds.
func NewBloom(size int, hashSeeds []int) *Bloom {
	bf := &Bloom{filter: make([]byte, size)}

	if hashSeeds == nil || len(hashSeeds) == 0 {
		hashSeeds = make([]int, DefaultHashFns)
		for i := 0; i < DefaultHashFns; i++ {
			hashSeeds[i] = rand.Int()
		}
		bf.fns = make([]hash.Hash32, DefaultHashFns)
	} else {
		bf.fns = make([]hash.Hash32, len(hashSeeds))
	}

	for i := 0; i < len(bf.fns); i++ {
		bf.fns[i] = murmur3.SeedNew32(uint32(hashSeeds[i]))
	}

	return bf
}

// Put puts the element `x` in the bloom filter
func (b *Bloom) Put(x []byte) error {
	var err error
	for i := range b.fns {
		b.fns[i].Reset()
		if _, err = b.fns[i].Write(x); err != nil {
			return err
		}
		pos := int(b.fns[i].Sum32()) % (len(b.filter) * 8)
		idx, offset := pos/8, pos%8
		b.filter[idx] = b.filter[idx] | 1<<offset
	}
	return nil
}

// Check checks the existence of the element `x` in the bloom filter
func (b *Bloom) Check(x []byte) (bool, error) {
	var err error
	var keyExists bool = true
	for i := range b.fns {
		b.fns[i].Reset()
		if _, err = b.fns[i].Write(x); err != nil {
			return false, err
		}
		pos := int(b.fns[i].Sum32()) % (len(b.filter) * 8)
		idx, offset := pos/8, pos%8
		if b.filter[idx]&(1<<offset) == 0 {
			keyExists = false
			break
		}
	}
	return keyExists, nil
}
