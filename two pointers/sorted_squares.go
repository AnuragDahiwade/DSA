package main

func sortedSquares(nums []int) []int {
	l := 0
	for l < len(nums) {
		if nums[l] >= 0 {
			break
		}
		l++
	}

	l = l - 1
	r := l + 1
	res := []int{}
	for l >= 0 && r < len(nums) {
		if -nums[l] < nums[r] {
			res = append(res, nums[l]*nums[l])
			l--
		} else {
			res = append(res, nums[r]*nums[r])
			r++
		}
	}

	for l >= 0 {
		res = append(res, nums[l]*nums[l])
		l--
	}
	for r < len(nums) {
		res = append(res, nums[r]*nums[r])
		r++
	}
	return res
}
