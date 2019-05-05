package algorithms

import "math"

// QuickSort Quick sorting algorithm
func QuickSort(randomNumbers []int, low int, high int) []int {
	if len(randomNumbers) > 1 {
		index := partition(randomNumbers, low, high)

		if low > index-1 {
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
