package basicBitOperator

//按位与 & ：两位全为1，结果才是1，否则是0
//按位或 | ：两位有一个为1，结果为1，否则是0
//按位异或^ ：两位一个为0，一个为1，结果为1，否则为0——把二进制位的一个值变成另外一个值
//https://www.php.cn/be/go/496132.html

func testBasicBit() {
	a := 0xCEFF
	a ^= 0xFF00
}
