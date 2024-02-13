package main

import (
	"fmt"
	"testing"
)

// TODO передалать тесты

func TestInt(t *testing.T) {
	ch1 := generateChannel(1, 2, 3, 4, 5)
	ch2 := generateChannel(6, 7, 8)
	ch3 := generateChannel(9, 10)
	ch4 := generateChannel(11, 12, 13, 14, 15, 16)
	ch5 := generateChannel(17, 18, 19, 20)
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestInt output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func TestFloat(t *testing.T) {
	ch1 := generateChannel(1.0, 2.0, 3.0, 4.0, 5.0)
	ch2 := generateChannel(6.0, 7.0, 8.0)
	ch3 := generateChannel(9.0, 10.0)
	ch4 := generateChannel(11.0, 12.0, 13.0, 14.0, 15.0, 16.0)
	ch5 := generateChannel(17.0, 18.0, 19.0, 20.0)
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestFloat output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func TestString(t *testing.T) {
	ch1 := generateChannel("one", "two", "three", "four", "five")
	ch2 := generateChannel("six", "seven", "eight")
	ch3 := generateChannel("nine", "ten")
	ch4 := generateChannel("eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen")
	ch5 := generateChannel("seventeen", "eighteen", "nineteen", "twenty")
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestString output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func TestStruct(t *testing.T) {
	ch1 := generateChannel(struct {
		Name string
		Val  int
	}{Name: "one", Val: 1}, struct {
		Name string
		Val  int
	}{Name: "two", Val: 2})
	ch2 := generateChannel(struct {
		Name string
		Val  int
	}{Name: "tree", Val: 3})
	ch3 := generateChannel(struct {
		Name string
		Val  int
	}{Name: "four", Val: 4})
	ch4 := generateChannel(struct {
		Name string
		Val  int
	}{Name: "five", Val: 5}, struct {
		Name string
		Val  int
	}{Name: "six", Val: 6})
	ch5 := generateChannel(struct {
		Name string
		Val  int
	}{Name: "seven", Val: 7})
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestStruct output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func TestSlice(t *testing.T) {
	ch1 := generateChannel([]int{1, 2, 3, 4}, []int{5, 6, 7, 8})
	ch2 := generateChannel([]int{9, 10}, []int{11, 12, 13})
	ch3 := generateChannel([]int{14}, []int{15, 16, 17, 18})
	ch4 := generateChannel([]int{19, 20}, []int{21})
	ch5 := generateChannel([]int{23, 24, 25}, []int{26, 27, 28})
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestSlice output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func TestMap(t *testing.T) {
	ch1 := generateChannel(map[int]string{1: "one", 2: "two"}, map[int]string{3: "three", 4: "four"})
	ch2 := generateChannel(map[int]string{5: "five", 6: "six"})
	ch3 := generateChannel(map[int]string{7: "seven", 8: "eight"}, map[int]string{9: "nine", 10: "ten"})
	ch4 := generateChannel(map[int]string{11: "eleven", 12: "twelve"})
	ch5 := generateChannel(map[int]string{13: "thirteen", 14: "fourteen"}, map[int]string{15: "fifteen", 16: "sixteen"})
	out := multiplex(ch1, ch2, ch3, ch4, ch5)

	fmt.Println("TestMap output:")
	for val := range out {
		fmt.Printf("%v ", val)
	}
	fmt.Println()
}

func generateChannel(in ...interface{}) chan interface{} {
	out := make(chan interface{}, len(in))
	for _, val := range in {
		out <- val
	}
	close(out)

	return out
}
