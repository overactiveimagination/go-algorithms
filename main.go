package main

import (
	"fmt"
	"os"

	"github.com/overactiveimagination/go-algorithms/algorithms"
)

func main() {
	args := os.Args[1:]
	sortType := args[0]
	sampleArray := []int{1, 4, 2, 0}
	var result []int
	switch sortType {
	case "quickSort":
		result = algorithms.QuickSort(sampleArray, 0, len(sampleArray)-1)
	case "mergeSort":
		result = algorithms.MergeSort(sampleArray)
	}
	fmt.Printf("Sorted array: %v", result)
}
