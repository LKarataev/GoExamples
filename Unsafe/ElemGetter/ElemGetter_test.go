package ElemGetter

import (
	"testing"
)

func TestGetElementEmpty(t *testing.T) {
	arr := []int{}
	idx := 3

	want := 4
	got, err := getElement(arr, idx)
	if err == nil {
		t.Errorf("No error: %v", err)
	}

	if got != 0 {
		t.Errorf("Expected %d but got %d", want, got)
	}
}

func TestGetElementOne(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	idx := 3

	want := 4
	got, err := getElement(arr, idx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}
}

func TestGetElementTwo(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	idx := 9

	want := 10
	got, err := getElement(arr, idx)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if got != want {
		t.Errorf("Expected %d but got %d", want, got)
	}
}
