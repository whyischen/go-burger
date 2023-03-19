package algorithm

// QuickSort 快速排序
func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, left, right int) {
	if right <= left {
		return
	}
	partitionIndex := partition(arr, left, right)
	quickSort(arr, left, partitionIndex-1)
	quickSort(arr, partitionIndex+1, right)

}
func partition(arr []int, left, right int) int {
	index := left + 1
	for i := index; i < right; i++ {
		if arr[left] > arr[i] {
			arr[left], arr[i] = arr[i], arr[left]
			index++
		}
	}
	return index
}
