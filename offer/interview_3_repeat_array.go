package offer

/*
题目一: 找出数组中重复的数字
	在一个长度为n的数组里的所有数字都在0~n-1的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，
	也不知道每个数字重复几次，请找出数组中任意一个重复的数字

*/

func repeat_array(dst []int) int {
	if dst == nil || len(dst) == 0 {
		return -1
	}

	n := len(dst)
	for i := 0; i < n; i++ {
		for dst[i] != i {
			if dst[dst[i]] == dst[i] {
				return dst[i]
			}
			tmp := dst[i]
			dst[i] = dst[dst[i]]
			dst[dst[i]] = tmp
		}
	}
	return -1
}
