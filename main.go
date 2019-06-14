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
	switch sortType {
	case "quickSort":
		algorithms.QuickSort(sampleArray, 0, len(sampleArray)-1)
	}
	fmt.Printf("Sorted array: %v", sampleArray)
}
