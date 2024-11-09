package collections

import "fmt"

// FastUnorderedRemoveByIndex removes element at index, but instead of copying all elements to the left, moves into the same slot the last
// element. This spoils the list order. If index is the last element it is just removed.
//
// slice to be modified. index to be removed.
func FastUnorderedRemove[T any](list []T, index int) ([]T, error) {
	if err := checkIndexBoundary(index, len(list)); err != nil {
		return list, err
	}

	lastIndex := len(list) - 1
	if index != lastIndex {
		list[index] = list[lastIndex]
	}
	return list[:lastIndex], nil
}

// FastUnorderedRemoveByIndexWithLast ...
//
// lastIndex last element index in the list to be swapped into the removed index.
func FastUnorderedRemoveWithLastIndex[T any](list []T, index int, lastIndex int) ([]T, error) {
	if err := checkIndexBoundary(index, len(list)); err != nil {
		return list, err
	}
	if err := checkIndexBoundary(lastIndex, len(list)); err != nil {
		return list, err
	}

	if index != lastIndex {
		list[index] = list[lastIndex]
	}
	return list[:lastIndex], nil
}

// FastUnorderedRemoveByValue removes the first occurrence of a value from the slice.
// return true if found and removed, false otherwise.
func FastUnorderedRemoveByValue[T comparable](list []T, e T) ([]T, bool, error) {
	for i, item := range list {
		if item != e {
			continue
		}
		result, err := FastUnorderedRemoveWithLastIndex(list, i, len(list)-1)
		if err != nil {
			return list, false, err
		}
		return result, true, nil
	}
	return list, false, nil
}

func checkIndexBoundary(index int, listLength int) error {
	if index < 0 || index >= listLength {
		return fmt.Errorf("index %d is out of bounds", index)
	}
	return nil
}
