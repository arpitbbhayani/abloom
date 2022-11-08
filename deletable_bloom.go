package abloom

import "fmt"

type DeletableBloom struct {
	Bloom
	numRegions  int
	regionCover int
	pallete     []byte
}

// NewDeletableBloom creates a deletable bloom filter of length `size` bytes
// and `hashSeeds` are the seed values for the murmur hash functions
// bloom will be initialized with len(hashSeeds) hash functions
// with provided seeds.
// if no hashSeeds are provided then 2 hash functions will be used
// with random seeds.
// numRegions are the number of regions in which the bloom filter needs to be split
// so as to make it deletable. For each region, one extra bit of memory will be allocated.
func NewDeletableBloom(size int, hashSeeds []int, numRegions int) *DeletableBloom {
	bf := &DeletableBloom{
		numRegions:  numRegions,
		regionCover: divCeil(size*8, numRegions),
		pallete:     make([]byte, divCeil(numRegions, 8)),
		Bloom:       *NewBloom(size, hashSeeds),
	}
	return bf
}

// Put puts the element `x` in the deletable bloom filter
func (b *DeletableBloom) Put(x []byte) error {
	var err error
	for i := range b.fns {
		b.fns[i].Reset()
		if _, err = b.fns[i].Write(x); err != nil {
			return err
		}
		pos := int(b.fns[i].Sum32()) % (len(b.filter) * 8)
		setBit(b.filter, pos)
		fmt.Println(pos)

		// updating the pallete, get the region number
		posReg := pos / b.regionCover
		setBit(b.pallete, posReg)
	}
	printFilter(b.pallete)
	printFilter(b.filter)
	return nil
}

// Check checks the existence of the element `x` in the bloom filter
func (b *DeletableBloom) Check(x []byte) (bool, error) {
	return b.Bloom.Check(x)
}
