package SortSearch

import "sort"

//https://blog.csdn.net/weixin_47214315/article/details/124518756

type MySort struct {
	Id int
	s  int
}

func sortUsage(msarr []MySort) {
	sort.Slice(msarr, func(i, j int) bool {
		//return true：交换，return false:不交换，所以return A>B表示当A>B就交换，即降序 其中i代表后一个参数，j代表前一个
		return msarr[i].Id > msarr[j].Id
	})
}
