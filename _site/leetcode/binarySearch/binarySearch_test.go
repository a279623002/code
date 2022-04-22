package binarySearch

import "testing"

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 4, 5, 7, 8}
	if index := BinarySearch(arr, 3); index != 1 {
		t.Errorf("search 3 expected be 1, but %d got", index)
	}
}
