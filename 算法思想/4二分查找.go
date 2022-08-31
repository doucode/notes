package 算法思想

//二分查找 1/6

//69. x 的平方根
//Input: 4
//Output: 2
//
//Input: 8
//Output: 2
func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		m := l + (r-l)/2
		if m*m > x {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
