package main

import (
	"fmt"
	"sync"
	"time"
)

func SleepSort(nums []int) chan int {
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, i := range nums {
		wg.Add(1)

		go func(num int) {
			defer wg.Done()
			time.Sleep(time.Duration(num) * time.Second)
			ch <- num
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return ch
}

func main() {
	nums := []int{7, 5, 6, 3, 9, 6, 2, 1, 7, 3, 0, 8}
	sorted := SleepSort(nums)
	for n := range sorted {
		fmt.Println(n)
	}
	fmt.Println("Done")
}
