package main

import (
	"fmt"
	"math/rand"
	"sync"
	// "time"
)

func multiplex(chs ...chan interface{}) chan interface{} {
	out := make(chan interface{})
	var wg sync.WaitGroup

	for _, ch := range chs {
		wg.Add(1)

		go func(ch chan interface{}) {
			defer wg.Done()
			for msg := range ch {
				out <- msg
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- rand.Intn(100)
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch2 <- rand.Float64()
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch2)
	}()

	go func() {
		for i := 0; i < 10; i++ {
			ch3 <- fmt.Sprintf("Hello %d", i)
			// time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(ch3)
	}()

	out := multiplex(ch1, ch2, ch3)

	for msg := range out {
		fmt.Println(msg)
	}
}
