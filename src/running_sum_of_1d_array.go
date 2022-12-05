package src

/*
Difficult: Easy
Link: https://leetcode.com/problems/running-sum-of-1d-array
*/
func runningSum(nums []int) []int {
	prefix := 0

	for i, n := range nums {
		prefix += n
		nums[i] = prefix
	}
	return nums
}