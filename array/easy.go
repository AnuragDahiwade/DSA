package main

func largestElement(arr []int) int {
	if len(arr) < 1 {
		return -1
	}
	max := arr[0]
	for _, val := range arr {
		if max < val {
			max = val
		}
	}
	return max
}

func secondLargestElement(arr []int) int {
	if len(arr) < 1 {
		return -1
	}
	lar, seclar := arr[0], -1
	for i := 1; i < len(arr); i++ {
		if arr[i] > lar {
			seclar = lar
			lar = arr[i]
		}
	}
	return seclar
}

func checkIsArraySorted(nums []int) bool {

	flag := false
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if flag == true {
				return false
			}
			flag = true
		}
	}

	if flag == true && !(nums[0] >= nums[len(nums)-1]) {
		return false
	}

	return true
}

func removeDuplicates(nums []int) int {
	i, j, n := 0, 0, len(nums)

	for i < n && j < n {
		if nums[i] != nums[j] {
			i += 1
			nums[i] = nums[j]
			j += 1
		} else {
			j += 1
		}
	}
	return i + 1
}

func rotateArrayByDPlaces(nums []int, k int) {
	n := len(nums)
	rotation := k % n

	nums = reverse(nums, 0, n-rotation-1)
	nums = reverse(nums, n-rotation, n-1)
	nums = reverse(nums, 0, n-1)
}

func reverse(nums []int, l, r int) []int {

	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return nums
}

func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j := 0, 0
	for j < n {
		if nums[j] != 0 {
			nums[i] = nums[j]
			i += 1
		}
		j += 1
	}

	for i < n {
		nums[i] = 0
		i += 1
	}

}

func missingNumber(nums []int) int {
	arr := make([]int, len(nums)+1)

	for _, val := range nums {
		arr[val] = 1
	}

	for idx, val := range arr {
		if val == 0 {
			return idx
		}
	}
	return -1
}

func findMaxConsecutiveOnes(nums []int) int {
	count, max := 0, 0
	for _, val := range nums {
		if val != 1 {
			if max < count {
				max = count
			}
			count = -1
		}
		count += 1
	}
	if max < count {
		max = count
	}
	return max
}

func singleNumber(nums []int) int {
	myMap := make(map[int]int)

	for _, val := range nums {
		if _, ok := myMap[val]; ok {
			delete(myMap, val)
		} else {
			myMap[val] = 1
		}

	}

	for key, _ := range myMap {
		return key
	}
	return -1
}
