package Knapsack

import (
	"reflect"
	"testing"
)

func TestGrabPresentsEmpty(t *testing.T) {
	presents := []Present{}
	got := grabPresents(presents, 10)
	if len(got) != 0 {
		t.Errorf("Expected empty list of selected presents, but got %v", got)
	}
}

func TestGrabPresentsOne(t *testing.T) {
	presents := []Present{
		{Value: 10, Size: 5},
		{Value: 20, Size: 10},
		{Value: 30, Size: 15},
	}

	got := grabPresents(presents, 9)
	want := []Present{
		{Value: 10, Size: 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}

func TestGrabPresentsTwo(t *testing.T) {
	presents := []Present{
		{Value: 10, Size: 5},
		{Value: 15, Size: 5},
		{Value: 20, Size: 5},
		{Value: 20, Size: 10},
		{Value: 70, Size: 15},
		{Value: 10, Size: 15},
		{Value: 60, Size: 20},
	}

	got := grabPresents(presents, 45)
	want := []Present{
		{Value: 60, Size: 20},
		{Value: 70, Size: 15},
		{Value: 20, Size: 5},
		{Value: 15, Size: 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}

func TestGrabPresentsThree(t *testing.T) {
	presents := []Present{
		{Value: 10, Size: 5},
		{Value: 15, Size: 5},
		{Value: 20, Size: 5},
		{Value: 20, Size: 10},
		{Value: 70, Size: 15},
		{Value: 10, Size: 15},
		{Value: 60, Size: 20},
	}

	got := grabPresents(presents, 0)
	want := []Present{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}

func TestGrabPresentsFour(t *testing.T) {
	presents := []Present{
		{Value: 10, Size: 5},
		{Value: 15, Size: 5},
		{Value: 20, Size: 5},
		{Value: 20, Size: 10},
		{Value: 70, Size: 15},
		{Value: 10, Size: 15},
		{Value: 60, Size: 20},
	}

	got := grabPresents(presents, 600)
	want := []Present{
		{Value: 60, Size: 20},
		{Value: 10, Size: 15},
		{Value: 70, Size: 15},
		{Value: 20, Size: 10},
		{Value: 20, Size: 5},
		{Value: 15, Size: 5},
		{Value: 10, Size: 5},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected result. Expected %v, but got %v", want, got)
	}
}
