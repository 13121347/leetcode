package ThreeSumNumber

import (
	"math"
	"sort"
)

//排序+双指针
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	var (
		n    = len(nums)
		best = math.MaxInt32
	)

	//根据差值的绝对值来更新答案
	update := func(cur int) {
		if abs(cur-target) < abs(best-target) {
			best = cur
		}
	}

	//枚举a
	for a := 0; a < n; a++ {
		//保证和上一次枚举的元素不相等
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		//使用双指针枚举b和c
		b, c := a+1, n-1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				//如果和大于target,移动C的指针
				c0 := c - 1
				//移动到下一个不相等的元素
				for c < c0 && nums[c0] == nums[c] {
					c0--
				}
				c = c0
			} else {
				// 如果和小于 target，移动 b 对应的指针
				b0 := b + 1
				// 移动到下一个不相等的元素
				for b0 < b && nums[b0] == nums[b] {
					b0++
				}
				b = b0
			}
		}
	}
	return best
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
