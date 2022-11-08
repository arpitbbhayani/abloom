package abloom

import (
	"fmt"
	"math"
)

func setBit(buf []byte, b int) {
	idx, offset := b/8, b%8
	buf[idx] = buf[idx] | 1<<offset
}

func getBit(buf []byte, b int) byte {
	idx, offset := b/8, b%8
	return buf[idx] & (1 << offset)
}

func printFilter(buf []byte) {
	for i := range buf {
		fmt.Printf("%08b ", buf[i])
	}
	fmt.Println()
}

func divCeil(a, b int) int {
	return int(math.Ceil(float64(a) / float64(b)))
}
