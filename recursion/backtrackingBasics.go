package main

import "sort"

func printAllSubsets(nums []int) [][]int {
	res := [][]int{}
	helperPrintAllSubsets(nums, 0, []int{}, &res)
	return res
}

func helperPrintAllSubsets(nums []int, idx int, curr []int, res *[][]int) {
	if idx == len(nums) {
		// Make a copy of curr before appending
		tmp := make([]int, len(curr))
		copy(tmp, curr)
		*res = append(*res, tmp)
		return
	}

	// Exclude current element
	helperPrintAllSubsets(nums, idx+1, curr, res)

	// Include current element
	curr = append(curr, nums[idx])
	helperPrintAllSubsets(nums, idx+1, curr, res)
}

// Subsets-II
func printSubsets(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	helper(nums, 0, []int{}, &res)
	return res
}

func helper(nums []int, idx int, curr []int, res *[][]int) {
	if idx == len(nums) {
		temp := make([]int, len(curr))
		copy(temp, curr)
		*res = append(*res, temp)
		return
	}

	curr = append(curr, nums[idx])
	helper(nums, idx+1, curr, res)

	curr = curr[:len(curr)-1]

	for idx < len(nums)-1 && nums[idx] == nums[idx+1] {
		idx += 1
	}

	helper(nums, idx+1, curr, res)
}

// func helper(nums []int, idx int, curr []int, res *[][]int) {

// 	temp := make([]int, len(curr))
// 	copy(temp, curr)
// 	*res = append(*res, temp)

// 	for i := idx; i < len(nums); i++ {
// 		if i > idx && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		curr = append(curr, nums[i])
// 		helper(nums, i+1, curr, res)
// 		curr = curr[:len(curr)-1]
// 	}
// }
