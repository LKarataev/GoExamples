package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func crawlWeb(ctx context.Context, urls <-chan string, response chan<- string) {
	wg := sync.WaitGroup{}
	sem := make(chan struct{}, 8)
	defer close(response)

	for url := range urls {
		sem <- struct{}{}
		wg.Add(1)

		go func(url string) {
			defer func() {
				<-sem
				wg.Done()
			}()

			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error creating request for URL %s: %s\n", url, err)
				return
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error requesting URL %s: %s\n", url, err)
				return
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading response body for URL %s: %s\n", url, err)
				return
			}

			select {
			case response <- string(body):
			case <-ctx.Done():
				return
			}
		}(url)

		select {
		case <-ctx.Done():
			os.Exit(1)
		default:
		}
	}

	wg.Wait()
}

func main() {
	urls := make(chan string)
	results := make(chan string)
	done := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for result := range results {
			fmt.Printf("Received result: %s\n", result)
		}
		done <- "Done receiving results"
	}()

	go crawlWeb(ctx, urls, results)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Received signal, shutting down gracefully...")
		cancel()
	}()

	for i := 6; i < 30; i++ {
		time.Sleep(1 * time.Second)
		urls <- fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", i)
	}

	close(urls)

	<-done

	fmt.Println("All done!")
}

//TODO реализовать противодавление с помощью следующих функций
/*
type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func main() {
	pg := New(10)
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := pg.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})
		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
*/
