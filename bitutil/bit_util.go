package bitutil

type Int interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

// IsPowerOfTwo checks if a number is a power of two.
func IsPowerOfTwo[T Int](value T) bool {
	return value > 0 && ((value & (^value + 1)) == value)
}

// Align aligns a value to the next multiple of alignment.
// If the value is already a multiple of alignment, it returns the value unchanged.
func Align(value, alignment int) int {
	return (value + (alignment - 1)) & -alignment
}
