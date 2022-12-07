package MyStack

//用了双端口队列，题目让用普通FIFO的队列
type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	value := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return value
}

func (this *MyStack) Top() int {
	return this.queue[len(this.queue)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}
