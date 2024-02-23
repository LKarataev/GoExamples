package SortPresents

import (
	"reflect"
	"testing"
)

func TestGetNCoolestPresentsFour(t *testing.T) {
	presents := PresentHeap{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
	}

	want := PresentHeap{
		{Value: 5, Size: 1},
		{Value: 5, Size: 2},
	}
	got, err := getNCoolestPresents(presents, 2)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}

func TestGetNCoolestPresentsFifteen(t *testing.T) {
	presents := []Present{
		{Value: 5, Size: 1},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 5, Size: 2},
		{Value: 1, Size: 1},
		{Value: 2, Size: 2},
		{Value: 5, Size: 3},
		{Value: 1, Size: 5},
		{Value: 5, Size: 4},
		{Value: 2, Size: 3},
		{Value: 3, Size: 3},
		{Value: 1, Size: 4},
		{Value: 4, Size: 4},
		{Value: 4, Size: 1},
		{Value: 2, Size: 4},
	}

	want := PresentHeap{
		{Value: 5, Size: 1},
		{Value: 5, Size: 2},
		{Value: 5, Size: 3},
		{Value: 5, Size: 4},
		{Value: 4, Size: 1},
		{Value: 4, Size: 4},
		{Value: 4, Size: 5},
		{Value: 3, Size: 1},
		{Value: 3, Size: 3},
		{Value: 2, Size: 2},
	}
	got, err := getNCoolestPresents(presents, 10)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}

func TestGetNCoolestPresentsempty(t *testing.T) {
	presents := PresentHeap{}
	_, err := getNCoolestPresents(presents, 2)
	if err == nil {
		t.Error("Expected an error but got nil")
	}
}

func TestGetNCoolestPresentsNegativeN(t *testing.T) {
	presents := PresentHeap{
		{Value: 5, Size: 1},
	}

	_, err := getNCoolestPresents(presents, -1)
	if err == nil {
		t.Error("Expected an error but got nil")
	}
}

func TestGetNCoolestPresentsTooLargeN(t *testing.T) {
	presents := PresentHeap{
		{Value: 5, Size: 1},
	}

	_, err := getNCoolestPresents(presents, 5)
	if err == nil {
		t.Error("Expected an error but got nil")
	}
}
