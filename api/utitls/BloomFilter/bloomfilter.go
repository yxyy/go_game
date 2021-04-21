package BloomFilter

import "github.com/willf/bitset"

type BloomFilter struct {
	set *bitset.BitSet
	funcs [6]SimpleHash
}

type SimpleHash struct{
	cap uint
	seed uint
}
