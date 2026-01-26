package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, val := range nums {
		complement := target - val

		if idx, isExist := numMap[complement]; isExist {
			return []int{i, idx}
		}
		numMap[val] = i
	}

	return nil
}

func sortColors(nums []int) {
	countMap := make(map[int]int)
	for _, val := range nums {
		countMap[val] += 1
	}

	idx := 0
	for i := 0; i < countMap[0]; i++ {
		nums[idx] = 0
		idx += 1
	}
	for i := 0; i < countMap[1]; i++ {
		nums[idx] = 1
		idx += 1
	}
	for i := 0; i < countMap[2]; i++ {
		nums[idx] = 2
		idx += 1
	}
}

func majorityElement(nums []int) int {

	countMap := make(map[int]int)

	maxId, maxVal := 0, 0
	for _, val := range nums {
		countMap[val] += 1

		if maxVal < countMap[val] {
			maxId = val
			maxVal = countMap[val]
		}
	}

	return maxId
}

func maxSubArray(nums []int) int {
	max, sum := nums[0], 0

	for _, val := range nums {
		sum += val

		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}

// Cloud := The cloud may be defined as the on-demand self-service access to computing resources that are pooled and broadly available for network access providing rapid elasticity and measured use.
