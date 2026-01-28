package main

import "slices"

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				ans := []int{nums[i], nums[j], nums[k]}
				res = append(res, ans)

				j++
				k--

				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}

	return res
}
