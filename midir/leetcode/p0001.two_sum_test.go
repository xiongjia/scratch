package leetcode

import (
	"fmt"
	"testing"
)

func Test_p0001(t *testing.T) {
	u := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Debug u = %v\n", u)
	u[1] = 10
	fmt.Printf("Debug u = %v\n", u)

	s2 := u[1:3]
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	s2[0] = 12
	fmt.Printf("Debug u = %v\n", u)
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))

	s2 = append(s2, 9)
	fmt.Printf("s2 = %v, len = %d, cap = %d\n", s2, len(s2), cap(s2))
	fmt.Printf("Debug u = %v\n", u)

	s1 := make([]int, 5, 10)
	s1 = append(s1, 1)
	fmt.Printf("Debug s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))

	m := make(map[string]int)
	m["abc"] = 1
	fmt.Printf("Debug %v\n", m["abc1"])

	fmt.Printf("%v \n", TwoSum([]int{3, 2, 4}, 6))
}
