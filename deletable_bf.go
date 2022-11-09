package abloom

type DeletableBloom struct {
	bloom        *bloom
	perRegionLen int
	pallete      []byte
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
		perRegionLen: divCeil(size*8, numRegions),
		pallete:      make([]byte, divCeil(numRegions, 8)),
		bloom:        newBloom(size, hashSeeds),
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
		posReg := positions[i] / b.perRegionLen
		// if collision then set the corresponding bit in the pallet
		if getBit(b.bloom.filter, positions[i]) == 1 {
			setBit(b.pallete, posReg)
		}
	}

	return nil
}

// Check checks the existence of the element `x` in the bloom filter
func (b *DeletableBloom) Check(x []byte) (bool, error) {
	return b.bloom.check(x)
}

// Delete possibly deletes the element `x` from the bloom filter the element `x`
// returns error if any and if the element `x` was deleted or not
func (b *DeletableBloom) Delete(x []byte) (bool, error) {
	positions, err := b.bloom.positions(x)
	if err != nil {
		return false, err
	}

	// updating the pallete, get the region number
	for i := range positions {
		posReg := positions[i] / b.perRegionLen
		bit := getBit(b.pallete, posReg)
		// if pallet bit is 0 then clear the position from the filter
		if bit == 0 {
			resetBit(b.bloom.filter, positions[i])
		}
	}

	return false, nil
}
