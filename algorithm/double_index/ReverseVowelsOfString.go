package double_index

/**
345. 反转字符串中的元音字母
 */

func reverseVowels(s string) string {
	m := map[byte]bool {
		'a':true,
		'A':true,
		'e':true,
		'E':true,
		'i':true,
		'I':true,
		'o':true,
		'O':true,
		'u':true,
		'U':true,
	}
	sli := []byte(s)
	left, right := 0, len(s)-1
	for left < right {
		if m[sli[left]] && m[sli[right]] {
			sli[left], sli[right] = sli[right], sli[left]
			left++
			right--
		}
		if !m[sli[left]] {
			left++
		}
		if !m[sli[right]] {
			right--
		}
	}
	return string(sli)
}


