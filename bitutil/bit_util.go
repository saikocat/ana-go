package bitutil

import "fmt"

type Int interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

// IsPowerOfTwo checks if a number is a power of two.
func IsPowerOfTwo[T Int](value T) bool {
	return value > 0 && ((value & (^value + 1)) == value)
}

// AlignU32 aligns a value to the next multiple of alignment.
// If the value is already a multiple of alignment, it returns the value unchanged.
func AlignU32(value, alignment uint32) uint32 {
	return (value + (alignment - 1)) & -alignment
}

// AlignU64 aligns a value to the next multiple of alignment.
// If the value is already a multiple of alignment, it returns the value unchanged.
func AlignU64(value, alignment uint64) uint64 {
	return (value + (alignment - 1)) & -alignment
}

// IsAligned checks if a value is aligned to a given power of 2 alignment.
func IsAligned(value, alignment int) (bool, error) {
	if !IsPowerOfTwo(alignment) {
		return false, fmt.Errorf("alignment must be a power of 2: alignment=%d", alignment)
	}
	return (value & (alignment - 1)) == 0, nil
}
