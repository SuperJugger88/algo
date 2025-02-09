package search

func BinarySearch(value int, arr ...int) int {
	l, r := 0, len(arr)-1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == value {
			return mid
		}

		if arr[mid] > value {
			r = mid - 1
		}

		if arr[mid] < value {
			l = mid + 1
		}
	}

	return -1
}
