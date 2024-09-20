package popcount

// pc[i] is the population count of i.
var pc [256]byte // 1 byte = 8 bits 表示2^8=256种可能
// pc[1]=1, pc[2]=1, pc[3]=2, pc[4]=1, pc[5]=2, pc[6]=2, pc[7]=3, pc[8]=1, pc[9]=2, pc[10]=2, pc[11]=3, pc[12]=2, pc[13]=3, pc[14]=3, pc[15]=4

//p[0]=p[0/2]+0&1=0+0=0
//p[1]=p[1/2]+1&1=0+1=1
//p[2]=p[2/2]+2&1=1+0=1
//p[3]=p[3/2]+3&1=1+1=2
//p[4]=p[4/2]+4&1=1+0=1
//p[5]=p[5/2]+5&1=1+1=2

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1) //
	}
}

// 返回当前数字的二进制表示中1的个数  3 = 11  2
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

func PopCount2(x uint64) int {
	count := 0
	for x != 0 {
		if x&1 == 1 {
			count++
		}
		x >>= 1
	}
	return count
}
