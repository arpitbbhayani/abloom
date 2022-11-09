package abloom

type SimpleBF struct {
	bloom *bloom
}

// NewSimpleBF creates a simple bloom filter of length `size` bytes and
// `hashSeeds` are the seed values for themurmur hash functions
// bloom will be initialized with len(hashSeeds) hash functions
// with provided seeds.
// if no hashSeeds are provided then 2 hash functions will be used
// with random seeds.
func NewSimpleBF(size int, hashSeeds []int) *SimpleBF {
	return &SimpleBF{
		bloom: newBloom(size, hashSeeds),
	}
}

// Put puts the element `x` in the bloom filter
func (b *SimpleBF) Put(x []byte) error {
	_, err := b.bloom.put(x)
	return err
}

// Check checks the existence of the element `x` in the bloom filter
func (b *SimpleBF) Check(x []byte) (bool, error) {
	return b.bloom.check(x)
}
