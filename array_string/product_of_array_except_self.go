package arraystring

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.
*/

func productExceptSelf(nums []int) []int {
	totalProduct := 1
	nonZeroTotalProduct := 1
	indexesOfZeroes := make([]int, 0)
	for i, v := range nums {
		totalProduct *= v
		if v == 0 {
			indexesOfZeroes = append(indexesOfZeroes, i)
		} else {
			nonZeroTotalProduct *= v
		}
	}
	result := make([]int, len(nums))
	if len(indexesOfZeroes) >= 2 {
		return result
	}
	if len(indexesOfZeroes) == 1 {
		result[indexesOfZeroes[0]] = nonZeroTotalProduct
		return result
	}
	for i, v := range nums {
		result[i] = totalProduct / v
	}
	return result
}
