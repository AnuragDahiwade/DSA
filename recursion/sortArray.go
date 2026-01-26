package main

func sortArray(nums []int) []int {
	helperSort(&nums, len(nums)-1)
	return nums
}

func helperSort(nums *[]int, n int) {
	if n < 0 {
		return
	}

	maxIdx := 0
	for i := 0; i <= n; i++ {
		if (*nums)[i] > (*nums)[maxIdx] {
			maxIdx = i
		}
	}

	(*nums)[maxIdx], (*nums)[n] = (*nums)[n], (*nums)[maxIdx]

	helperSort(nums, n-1)

}
