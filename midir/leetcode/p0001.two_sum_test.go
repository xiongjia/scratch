package leetcode

import (
	"fmt"
	"testing"
)

func Test_p0001(t *testing.T) {
	m := make(map[string]int)
	m["abc"] = 1
	fmt.Printf("Debug %v\n", m["abc1"])

	fmt.Printf("%v \n", TwoSum([]int{3, 2, 4}, 6))
}
