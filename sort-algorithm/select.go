package sortalgorithm

func SelectSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i <= len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j <= len(arr)-1; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if i != minIndex { //如果本轮第一个数字不是最小值，则交换
			swap(arr, minIndex, i)
		}
	}

	return arr
}
