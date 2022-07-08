package search

func BinarySearch(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		mid := (left + right) / 2

		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchRecursive(numbers []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if numbers[mid] == target {
		return mid
	} else if numbers[mid] < target {
		return BinarySearchRecursive(numbers, target, mid+1, right)
	} else {
		return BinarySearchRecursive(numbers, target, left, mid-1)
	}
}
