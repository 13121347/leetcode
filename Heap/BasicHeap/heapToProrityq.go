package BasicHeap

import (
	"container/heap"
	"fmt"
	"sort"
)

// https://blog.csdn.net/C_L99/article/details/124595316
/*
heap库之前，首先要实现其定义的五个函数
type Interface interface {
    // 堆的大小，通常返回的是切片的长度
    Len() int
    // 比较器，函数返回值为true时表示切片中i索引的元素比较“小”
	Less(i, j int) bool
    // 交换i和j两个坐标的元素
	Swap(i, j int)
    // 将元素x加入到堆中
	Push(x any)
    // 弹出堆顶元素
	Pop() any
}
*/

//大根堆实现如下
type bHp struct {
	//sort.IntSlice是golang标准库sort下定义的一个类型
	//它实现了上述提到的接口函数的前三个，分别是Len()、Less()、Swap()
	//结构体bHp直接使用sort.IntSlice类型作为值，可以将它默认实现的三个函数添加到bHp中（一般情况下，不需要再自己再去实现一次了）
	sort.IntSlice
}

//heap库实现的是小根堆，但是可以由我们自己去定义何为“小”。
////这里要实现的是一个大根堆，所以我们需要重写Less()函数
func (hp bHp) Less(i, j int) bool {
	return hp.IntSlice[i] > hp.IntSlice[j]
}

func (hp *bHp) Push(x interface{}) {
	hp.IntSlice = append(hp.IntSlice, x.(int))
}

func (hp *bHp) Pop() interface{} {
	a := hp.IntSlice
	x := a[len(a)-1]
	//这里不弹出hp.IntSilice[0]是因为heap框架处理的时候，会把[0]和[len-1]元素对调，再从头执行下沉操作
	/**
	func Pop(h Interface) any {
		n := h.Len() - 1
		h.Swap(0, n)
		down(h, 0, n)
		return h.Pop()
	}
	*/
	//调用我们自定义pop函数的时候，已经进行交换了
	hp.IntSlice = a[:len(a)-1]
	return x
}

func (hp bHp) BigHeap() interface{} {
	a := hp.IntSlice
	return a[0]
}

//使用方法
func useBigHeap() {
	// 实例化自定义堆
	pq := &bHp{}
	// 插入值
	heap.Push(pq, 8)
	heap.Push(pq, 9)
	// 获取堆顶元素但不删除
	fmt.Printf("当前堆顶元素为：%d\n", pq.BigHeap())
	fmt.Printf("当前堆大小为：%d\n", pq.Len())
	// 将所有的元素弹出
	for pq.Len() > 0 {
		//这里用的是heap.Pop,而不是我们自定义的bhp.pop,因为heap帮我们做了down操作，这里相当于是一个模板方法
		fmt.Println(heap.Pop(pq))
	}

}
