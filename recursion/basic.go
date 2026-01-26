package main

import "fmt"

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}

func sum(num int) int {
	if num == 1 {
		return 1
	}
	return num + sum(num-1)
}

// 0, 1, 1, 2, 3, 5, 8, 13, 21...
func fiboSeq(num int) {
	i, j := 0, 1
	k := 0
	for k < num {
		fmt.Println(i)
		temp := i
		i = j
		j = j + temp
		k += 1
	}

}

func fiboNum(num int) int {
	if num == 0 || num == 1 {
		return num
	}

	i, j := 0, 1
	k := 0
	for k < num-1 {
		temp := i
		i = j
		j = j + temp
		k += 1
	}
	return j
}

func fibo(num int) int {
	if num == 0 || num == 1 {
		return num
	}
	return fibo(num-1) + fibo(num-2)
}

func isSorted(nums []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}

	if nums[n-1] < nums[n-2] {
		return false
	}

	return isSorted(nums, n-1)
}

func binarySearch(nums []int, target, s, e int) int {
	if e < s {
		return -1
	}

	mid := s + (e-s)/2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		e = mid - 1
	} else {
		s = mid + 1
	}

	return binarySearch(nums, target, s, e)
}
