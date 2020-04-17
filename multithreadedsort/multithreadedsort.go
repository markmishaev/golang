package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	in := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a sequence of integers: ")
	in.Scan()

	var numbersSlice []int

	integersSequence := in.Text()
	ints := strings.Fields(integersSequence)

	for _, numStr := range ints {
		numInt, err := strconv.Atoi(numStr)
		if err == nil {
			numbersSlice = append(numbersSlice, numInt)
		}
	}

	fmt.Println("Numbers before sorting: ")
	fmt.Println(numbersSlice)

	MergeSort(numbersSlice)

	fmt.Println("Numbers after sorting: ")
	fmt.Println(numbersSlice)
}

// MergeSort performs the merge sort algorithm with goroutines
func MergeSort(src []int) {

	// We subtract 1 goroutine which is the one we are already running in.
	extraGoroutines := 4
	semChan := make(chan struct{}, extraGoroutines)
	defer close(semChan)
	mergesort(src, semChan)
}

func mergesort(src []int, semChan chan struct{}) {
	if len(src) <= 1 {
		return
	}

	mid := len(src) / 2

	left := make([]int, mid)
	right := make([]int, len(src)-mid)
	copy(left, src[:mid])
	copy(right, src[mid:])

	wg := sync.WaitGroup{}

	select {
	case semChan <- struct{}{}:
		wg.Add(1)
		go func() {
			mergesort(left, semChan)
			<-semChan
			wg.Done()
		}()
	default:
		// Can't create a new goroutine, let's do the job ourselves.
		mergesort(left, semChan)
	}

	mergesort(right, semChan)

	wg.Wait()

	merge(src, left, right)
}

func merge(result, left, right []int) {
	var l, r, i int

	for l < len(left) || r < len(right) {
		if l < len(left) && r < len(right) {
			if left[l] <= right[r] {
				result[i] = left[l]
				l++
			} else {
				result[i] = right[r]
				r++
			}
		} else if l < len(left) {
			result[i] = left[l]
			l++
		} else if r < len(right) {
			result[i] = right[r]
			r++
		}
		i++
	}
}
