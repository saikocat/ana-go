package bitutil

type Int interface {
	int | uint | int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64
}

// IsPowerOfTwo checks if a number is a power of two.
func IsPowerOfTwo[T Int](value T) bool {
	return value > 0 && ((value & (^value + 1)) == value)
}
