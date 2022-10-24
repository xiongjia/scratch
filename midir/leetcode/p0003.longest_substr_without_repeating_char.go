package leetcode

func LengthOfLongestSubstringSolution1(s string) int {
	srcLen := len(s)
	if srcLen == 0 {
		return 0
	}
	var bitSet [256]bool
	r, left, right := 0, 0, 0
	for left < srcLen {
		if bitSet[s[right]] {
			bitSet[s[left]] = false
			left++
		} else {
			bitSet[s[right]] = true
			right++
		}
		if r < right-left {
			r = right - left
		}
		if right >= srcLen || left+r >= srcLen {
			break
		}
	}
	return r
}

func LengthOfLongestSubstringSolution2(s string) int {
	srcLen := len(s)
	if srcLen == 0 {
		return 0
	}
	bitMap := make(map[byte]bool, 256)
	r, left, right := 0, 0, 0
	for left < srcLen {
		if val, ok := bitMap[s[right]]; ok && val {
			bitMap[s[left]] = false
			left++
		} else {
			bitMap[s[right]] = true
			right++
		}
		if r < right-left {
			r = right - left
		}
		if right >= srcLen || left+r >= srcLen {
			break
		}
	}
	return r
}
