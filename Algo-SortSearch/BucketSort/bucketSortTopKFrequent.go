package BucketSort

//前K个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/
func topKFrequentBucket(nums []int, k int) (res []int) {
	//[元素值]：[元素频率]
	numFrequencyMap := make(map[int]int)
	//初始化map
	for _, v := range nums {
		numFrequencyMap[v]++
	}
	//初始化桶
	buckets := make([][]int, len(nums)+1)
	for numValue, numFrequency := range numFrequencyMap {
		if len(buckets[numFrequency]) == 0 {
			buckets[numFrequency] = make([]int, 0)
		}
		buckets[numFrequency] = append(buckets[numFrequency], numValue)
	}
	//结果列表,利用了桶——数组的有序性，nums --> map numvalue:frequency --> buckets[frequency]:[num1,num2...]
	for i := len(buckets) - 1; k > 0; i-- {
		if buckets[i] != nil {
			for _, num := range buckets[i] {
				res = append(res, num)
				k-- //只把k个元素放进result中了
			}
		}
	}

	return res
}
