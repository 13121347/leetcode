package InsertSort

// 讲解 https://blog.csdn.net/hcz666/article/details/126488359

//不断后退
func insertSort(arr []int) {
	//外循环，从0开始遍历每个值
	for outIndex := 0; outIndex < len(arr)-1; outIndex++ {
		//内循环，比较有序区后一个数和有序区内所有数
		toSort := arr[outIndex+1]
		for inIndex := outIndex; inIndex >= 0; inIndex-- {
			if arr[inIndex] > toSort {
				arr[inIndex+1] = arr[inIndex]
			}
			if arr[inIndex] <= toSort {
				arr[inIndex+1] = toSort
				break
			}
			if inIndex == 0 {
				arr[inIndex] = toSort
			}
		}
	}
}

//不断交换
func insertSort2(arr []int) {
	count := len(arr)
	for outIndex := 1; outIndex < count; outIndex++ {
		for inIndex := outIndex; inIndex > 0; inIndex-- {
			if arr[inIndex-1] > arr[inIndex] {
				arr[inIndex-1], arr[inIndex] = arr[inIndex], arr[inIndex-1]
			} else {
				break
			}
		}
	}

}
