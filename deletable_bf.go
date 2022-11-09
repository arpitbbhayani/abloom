package abloom

type DeletableBloom struct {
	bloom       *bloom
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
		bloom:       newBloom(size, hashSeeds),
	}
	return bf
}

// Put puts the element `x` in the deletable bloom filter
func (b *DeletableBloom) Put(x []byte) error {
	positions, err := b.bloom.put(x)
	if err != nil {
		return err
	}

	// updating the pallete, get the region number
	for i := range positions {
		posReg := positions[i] / b.regionCover
		setBit(b.pallete, posReg)
	}

	printFilter(b.pallete)
	printFilter(b.bloom.filter)
	return nil
}

// Check checks the existence of the element `x` in the bloom filter
func (b *DeletableBloom) Check(x []byte) (bool, error) {
	return b.bloom.check(x)
}
