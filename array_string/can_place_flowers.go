package arraystring

/*
605. Can Place Flowers
Easy
Topics
Companies
You have a long flowerbed in which some of the plots are planted, and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n, return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
*/

func canPlaceFlowers(flowerbed []int, n int) bool {
	planted := 0
	for i, v := range flowerbed {
		var prev, next = 0, 0
		if i < len(flowerbed)-1 {
			next = flowerbed[i+1]
		}
		if i > 0 {
			prev = flowerbed[i-1]
		}
		if prev == 0 && next == 0 && v == 0 {
			flowerbed[i] = 1
			planted += 1
		}
	}
	return planted >= n
}
