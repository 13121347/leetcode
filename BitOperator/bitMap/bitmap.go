package bitMap

type bitMap struct {
	arr      []byte
	capacity int
}

type bitOperator interface {
	add(n int)
	delete(n int)
	contain(n int) bool
}

func (bm *bitMap) add(n int) {
	index := n >> 3                // n >> 3 ,n右移3位，相当于n / 8
	position := n & 0x07           // n & 0x07, 相当于n % 8，因为这里等于n 和111进行与操作，会保留后三位
	bm.arr[index] |= 1 << position //把 1 移动到position位，变成 0..100..,然后和数组这位进行或操作，一定会把这位变成1
}

func (bm *bitMap) delete(n int) {
	index := n >> 3
	position := n & 0x07
	//bm.arr[index] &= 0 << position //这样不可以，因为这样会把positionn后面的数都变成0，相当于都删除了
	bm.arr[index] &= ^(1 << position) //把 1 移动到position位，变成 0..100..,再异或操作，1^0 = 0,然后和数组这位进行与操作,相当于对nXXXX & 01111,不会改变其他位的值
	//异或操作，1^0 = 0,1^1 = 0,0^0 = 1
}

func (bm *bitMap) contain(n int) bool {
	index := n >> 3
	position := n & 0x07
	return (bm.arr[index] & (1 << position)) != 0 //看看这位是不是0，0的话就是不包括
}

func InitBitMap(n int) *bitMap {
	bm := &bitMap{
		arr:      make([]byte, n/8+1),
		capacity: n,
	}
	return bm
}
