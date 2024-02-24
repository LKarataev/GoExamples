package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
)

type MetricsFlags struct {
	mean   bool
	median bool
	mode   bool
	sd     bool
}

func main() {
	var mf MetricsFlags
	flag.BoolVar(&mf.mean, "mean", false, "Print the mean value")
	flag.BoolVar(&mf.median, "median", false, "Print the median value")
	flag.BoolVar(&mf.mode, "mode", false, "Print the mode value")
	flag.BoolVar(&mf.sd, "sd", false, "Print the standard deviation value")
	flag.Parse()

	nums, freq, err := GetInput()
	if err != nil && err.Error() != "unexpected newline" {
		fmt.Println("Error:", err)
		if err != io.EOF {
			DeleteRemainingInput()
		}
		return
	}

	sort.Ints(nums)
	PrintMetrics(nums, freq, mf)
}

func GetInput() ([]int, map[int]int, error) {
	nums := []int{}
	freq := make(map[int]int)

	for {
		var num int
		_, err := fmt.Scanln(&num)
		if err != nil {
			return nums, freq, err
		}
		if num < -100000 || num > 100000 {
			fmt.Println("Error: ", num, " is out of range (-100000 to 100000)")
			continue
		}
		nums = append(nums, num)
		freq[num]++
	}
}

func DeleteRemainingInput() {
	//TODO refactor this
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
}

func PrintMetrics(nums []int, freq map[int]int, mf MetricsFlags) {
	if len(nums) == 0 {
		fmt.Println("Error: no numbers in input")
		return
	}
	if mf.mean {
		fmt.Printf("Mean: %.2f\n", Mean(nums))
	}
	if mf.median {
		fmt.Printf("Median: %.2f\n", Median(nums))
	}
	if mf.mode {
		fmt.Println("Mode: ", Mode(nums, freq))
	}
	if mf.sd {
		fmt.Printf("SD: %.2f\n", StandardDeviation(nums))
	}
	if !(mf.mean || mf.median || mf.mode || mf.sd) {
		fmt.Printf("Mean: %.2f\n", Mean(nums))
		fmt.Printf("Median: %.2f\n", Median(nums))
		fmt.Println("Mode: ", Mode(nums, freq))
		fmt.Printf("SD: %.2f\n", StandardDeviation(nums))
	}
}

func Mean(nums []int) float64 {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return float64(sum) / float64(len(nums))
}

func Median(nums []int) float64 {
	var median float64
	mid := len(nums) / 2
	if len(nums)%2 == 0 {
		median = float64(nums[mid-1]+nums[mid]) / 2.0
	} else {
		median = float64(nums[mid])
	}
	return median
}

func Mode(nums []int, freq map[int]int) int {
	mode := nums[0]
	max_freq := 1
	for num, f := range freq {
		if f > max_freq || (f == max_freq && num < mode) {
			mode = num
			max_freq = f
		}
	}
	return mode
}

func StandardDeviation(nums []int) float64 {
	variance := 0.0
	mean := Mean(nums)
	for _, num := range nums {
		diff := float64(num) - mean
		variance += diff * diff
	}
	variance /= float64(len(nums))

	return math.Sqrt(variance)
}
