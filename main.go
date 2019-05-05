package main

import (
	"fmt"
	"math"
	"os"
)

func quickSort(randomNumbers []int, low int, high int) []int {
	if len(randomNumbers) > 1 {
		index := partition(randomNumbers, low, high)

		if low > index-1 {
			quickSort(randomNumbers, low, index-1)
		}

		if high > index {
			quickSort(randomNumbers, index, high)
		}
	}

	return randomNumbers
}

func partition(arr []int, low int, high int) int {
	index := int(math.Floor(float64((low + high) / 2)))
	pivot := arr[index]
	i := low
	j := high

	for i <= j {
		for arr[i] < pivot {
			i++
		}

		for arr[j] > pivot {
			j++
		}

		if i <= j {
			tempI := arr[int(i)]
			arr[int(i)] = arr[int(j)]
			arr[int(j)] = tempI
			j--
			i++
		}
	}

	return i
}

func main() {
	args := os.Args[1:]
	sortType := args[0]
	sampleArray := []int{1, 4, 0, 2}
	switch sortType {
	case "quickSort":
		quickSort(sampleArray, 0, len(sampleArray)-1)
	}
	fmt.Printf("Sorted array: %v", sampleArray)
}
