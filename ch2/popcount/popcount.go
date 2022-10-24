// 定义PopCount函数, 用于返回一个数字中含二进制1bit的个数
package popcount

// pc[i] is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount1(x uint64) int {
	var r byte
	for i := 0; i < 8; i++ {
		r += pc[byte(x>>(i*8))]
	}
	return int(r)
}

func PopCount2(x uint64) int {
	var r byte
	for i := 0; i < 64; i++ {
		r += byte((x >> i) & 1)
	}
	return int(r)
}

func PopCount3(x uint64) int {
	var cnt int
	for x != 0 {
		cnt++
		x = x & (x - 1)
	}
	return cnt
}
