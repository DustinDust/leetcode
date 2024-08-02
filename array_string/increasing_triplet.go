package arraystring

import "math"

/*
Given an integer array nums, return true if
there exists a triple of indices (i, j, k) such that
i < j < k and nums[i] < nums[j] < nums[k].
If no such indices exists, return false.
*/
func increasingTriplet(nums []int) bool {
	if len(nums) <= 0 {
		return false
	}
	a, b := math.MaxInt, math.MaxInt
	for _, v := range nums {
		if v < a {
			a = v
		}
		if v > a && v < b {
			b = v
		}
		if v > b {

			return true
		}
	}
	return false
}
