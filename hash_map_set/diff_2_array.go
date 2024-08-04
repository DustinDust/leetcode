package hashmapset

/*
Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:

    answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
    answer[1] is a list of all distinct integers in nums2 which are not present in nums1.

    Note that the integers in the lists may be returned in any order.
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
    set1 := make(map[int]interface{}, 0)
    set2 := make(map[int]interface{}, 0)
    answer := make([][]int, 2)
    // convert 2 array into set
    for _, v := range nums1 { 
        set1[v] = struct{}{}
    }

    for _, v := range nums2 { 
        set2[v] = struct{}{}
    }

    for k := range set1 {
        if _, ok := set2[k]; !ok {
            answer[0] = append(answer[0], k)
        }
    }

    for k := range set2 {
        if _, ok := set1[k]; !ok {
            answer[1] = append(answer[1], k)
        }
    }

    return answer
}
