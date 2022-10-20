package sortalgorithm

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1
		for j >= 0 && cur < arr[j] { //如果插入元素小于插入对象元素，将插入对象元素后移一位，继续往前比较
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = cur //将插入元素放置在最后位移的插入对象前，假如此轮未移动，那么插入对象元素为本位
	}
	return arr
}
