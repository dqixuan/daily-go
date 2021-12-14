package offer

/**
	判断一个数组是否是二叉搜索树的后序遍历
 */


func VerifySquenceOfBST( sequence []int ) bool {
	// write code here
	if len(sequence) == 0 {
		return false
	}
	return helpVerifySequenceOfBST(sequence)
}

func helpVerifySequenceOfBST(sequence []int) bool {
	if len(sequence) < 2 {
		return true
	}
	rootValue := sequence[len(sequence)-1]
	idx := -1
	for i:=0; i < len(sequence); i++ {
		if sequence[i] > rootValue {
			idx = i
			break
		}
	}
	if idx >= 0 {
		for i:=idx; i < len(sequence)-1; i++ {
			if sequence[i] < rootValue {
				return false
			}
		}
		return helpVerifySequenceOfBST(sequence[:idx]) && helpVerifySequenceOfBST(sequence[idx:len(sequence)-1])
	}
	return helpVerifySequenceOfBST(sequence[:len(sequence)-1])
}