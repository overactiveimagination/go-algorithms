package algorithms

import "math"

// QuickSort Quick sorting algorithm
func QuickSort(randomNumbers []int, low int, high int) []int {
	if len(randomNumbers) > 1 {
		index := partition(randomNumbers, low, high)

		if low < index-1 {
			QuickSort(randomNumbers, low, index-1)
		}

		if high > index {
			QuickSort(randomNumbers, index, high)
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
			j--
		}

		if i <= j {
			arr[int(i)], arr[int(j)] = arr[int(j)], arr[int(i)]
			j--
			i++
		}
	}

	return i
}

func merge(left []int, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}

// MergeSort sort main
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2

	return merge(MergeSort(arr[:mid]), MergeSort(arr[mid:]))
}
