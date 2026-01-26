package main

func permutations(nums []int) [][]int {
	res := [][]int{}
	helperPermutation(nums, 0, &res)
	return res
}

func helperPermutation(nums []int, idx int, res *[][]int) {
	if idx == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, nums)
		*res = append(*res, temp)
		return
	}

	for i := idx; i < len(nums); i++ {
		nums[idx], nums[i] = nums[i], nums[idx]
		helperPermutation(nums, idx+1, res)

		nums[idx], nums[i] = nums[i], nums[idx] // Backtrack

	}
}
