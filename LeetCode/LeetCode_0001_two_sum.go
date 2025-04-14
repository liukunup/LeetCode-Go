// sss

package main

import (
	"testing"

	c "github.com/smartystreets/goconvey/convey"
)

func twoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TestT(t *testing.T) {

	c.Convey("TestCase_1", t, func() {
		var (
			numns  = []int{2, 7, 11, 15}
			target = 9
			expect = []int{0, 1}
		)
		result := twoSum(numns, target)
		c.So(result, c.ShouldResemble, expect)
	})

	c.Convey("TestCase_2", t, func() {
		var (
			numns  = []int{3, 2, 4}
			target = 6
			expect = []int{1, 2}
		)
		result := twoSum(numns, target)
		c.So(result, c.ShouldResemble, expect)
	})

	c.Convey("TestCase_3", t, func() {
		var (
			numns  = []int{3, 3}
			target = 9
			expect = []int{0, 1}
		)
		result := twoSum(numns, target)
		c.So(result, c.ShouldResemble, expect)
	})
}
