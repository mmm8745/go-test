package sortalgorithm

func BubbleSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	for i := 0; i < len(arr)-1; i++ {
		isSwaped := false //本轮有无交换的标志
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
				isSwaped = true
			}
		}
		if isSwaped == false { //若无交换，说明数组已经有序
			break
		}
	}
	return arr
}

func swap(arr []int, left, right int) {
	tmp := arr[left]
	arr[left] = arr[right]
	arr[right] = tmp
}
