package double_index

import "fmt"

/**
443. 压缩字符串
给你一个字符数组 chars ，请使用下述算法压缩：

从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：

如果这一组长度为 1 ，则将字符追加到 s 中。
否则，需要向 s 追加字符，后跟这一组的长度。
压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中。需要注意的是，如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符。

请在 修改完输入数组后 ，返回该数组的新长度。

 */

func compress(chars []byte) int {
	length := len(chars)
	if length <= 1 {
		return length
	}
	idx := -1
	var count  = 1
	pre := chars[0]
	for i:=0; i < length; i++ {
		if pre == chars[i] {
			if idx+1 == i {
				continue
			} else {
				count++
			}
		} else {
			idx++
			chars[idx] = pre
			if count > 1 {
				s := fmt.Sprintf("%d", count)
				for j := 0; j < len(s); j++ {
					idx++
					chars[idx] = s[j]
				}
			}
			pre = chars[i]
			count = 1
			continue
		}
	}
	idx++
	chars[idx] = pre
	if count > 1 {
		s := fmt.Sprintf("%d", count)
		for j := 0; j < len(s); j++ {
			idx++
			chars[idx] = s[j]
		}
	}
	return idx+1
}
