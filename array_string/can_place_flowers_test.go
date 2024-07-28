package arraystring

import "testing"

/*
605. Can Place Flowers
You have a long flowerbed in which some of the plots are planted, and some are not.
However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's,
where 0 means empty and 1 means not empty, and an integer n,
return true if n new flowers can be planted in the flowerbed
without violating the no-adjacent-flowers rule and false otherwise.
*/

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	result := canPlaceFlowers(flowerbed, n)
	expected := true

	if result != expected {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}

func TestCanPlaceFlowers2(t *testing.T) {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 2
	result := canPlaceFlowers(flowerbed, n)
	expected := false

	if result != expected {
		t.Errorf("expected: %v, got: %v", expected, result)
	}
}
