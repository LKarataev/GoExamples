package Wc

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"unicode/utf8"
)

type CountFlags struct {
	lines bool
	words bool
	chars bool
}

func main() {
	var cli_flags CountFlags
	flag.BoolVar(&cli_flags.lines, "l", false, "Count lines")
	flag.BoolVar(&cli_flags.words, "w", false, "Count words")
	flag.BoolVar(&cli_flags.chars, "m", false, "Count characters")
	flag.Parse()

	if !cli_flags.lines && !cli_flags.chars {
		cli_flags.words = true
	}

	var wg sync.WaitGroup
	for _, file := range flag.Args() {
		wg.Add(1)
		go CountInFile(file, &wg, cli_flags)
	}

	wg.Wait()
}

func CountInFile(filename string, wg *sync.WaitGroup, flags CountFlags) {
	defer wg.Done()

	file, err := os.Open(filename)
	CheckError(err)
	defer file.Close()

	var CountFunc func(io.Reader) (int, error)
	switch {
	case flags.lines:
		CountFunc = CountLines
	case flags.chars:
		CountFunc = CountChars
	default:
		CountFunc = CountWords
	}

	count, err := CountFunc(file)
	CheckError(err)

	fmt.Printf("%d\t%s\n", count, filename)
}

func CountLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lines, nil
}

func CountWords(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan() {
		words++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return words, nil
}

func CountChars(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}
		if n == 0 {
			break
		}
		count += utf8.RuneCount(buf[:n])
	}
	return count, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
