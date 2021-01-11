/**
*   @Author:qixuan ding
*   @Date: 2021/1/11 21:40
**/

package Traversal_Multiply

// 递归乘法。 写一个递归函数，不使用 * 运算符， 实现两个正整数的相乘。可以使用加号、减号、位移，但要吝啬一些。

// 方法1：a * b = a + a * (b - 1)
func multiply(a, b int) int {
	if b == 0 {
		return 0
	}
	return a + multiply(a, b-1)
}

// 方法2：寻找偶数项  a*b=(a*2)*(b/2) = (a << 1) (b / 2)  需要判断b是否是偶数
func multiply1(a, b int) int {
	if b == 0 {
		return 0
	}
	if b % 2 == 0 {
		return multiply1(a << 1, b / 2)
	} else {
		return a + multiply1(a, b -1)
	}
}

// 方法3：在方法2的基础上进行优化
/** a * b
	4 * 2 = 8 * 1 = 8 + 0 = 8
	2 * 4 = 4 * 2 = 8 * 1 = 8 + 0 = 8
	说明，对较大数进行位运算，较小数进行凑偶数，能减少运算次数
 */
func multiply2(a, b int) int {
	if b == 0 {
		return 0
	}
	if a < b {
		a, b = b, a
	}
	if b % 2 == 0 {
		return multiply2(a << 1, b / 2)
	} else {
		return a + multiply2(a, b -1)
	}
}