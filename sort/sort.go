package sort

func BubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		for j := i + 1; j < lenNumbers; j++ {
			if numbers[i] > numbers[j] {
				numbers[i], numbers[j] = numbers[j], numbers[i]
			}
		}
	}
	return numbers
}

func SelectionSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		var minimumIndex = i

		for j := i + 1; j < lenNumbers; j++ {
			if numbers[minimumIndex] > numbers[j] {
				minimumIndex = j
			}
		}
		numbers[i], numbers[minimumIndex] = numbers[minimumIndex], numbers[i]
	}
	return numbers
}

func InsertionSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		temp := numbers[i]

		j := i - 1

		for j >= 0 && temp < numbers[j] {
			numbers[j+1] = numbers[j]
			j -= 1
		}
		numbers[j+1] = temp
	}
	return numbers
}

func QuickSort(numbers []int) []int {
	quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

func quickSort(number []int, low int, high int) {
	if low < high {
		partitionIndex := partition(number, low, high)
		quickSort(number, low, partitionIndex-1)
		quickSort(number, partitionIndex+1, high)
	}
}

func partition(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i += 1
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	i += 1
	numbers[high], numbers[i] = numbers[i], numbers[high]
	return i
}

func MergeSort(numbers []int) []int {
	lenNumbers := len(numbers)
	if lenNumbers == 1 {
		return numbers
	}

	mid := lenNumbers / 2
	left := numbers[:mid]
	right := numbers[mid:]

	MergeSort(left)
	MergeSort(right)

	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i += 1
		} else {
			numbers[k] = right[j]
			j += 1
		}
		k += 1
	}

	for i < len(left) {
		numbers[k] = left[i]
		i += 1
		k += 1
	}

	for j < len(right) {
		numbers[k] = right[j]
		j += 1
		k += 1
	}

	return numbers

}
