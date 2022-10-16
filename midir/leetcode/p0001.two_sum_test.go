package leetcode

import (
	"fmt"
	"testing"
)

func updateTest(a []int) {
	a[1] = 100
	fmt.Printf("Debug a = %v\n", a)
}

func Test_p0001(t *testing.T) {
	u := []int{1, 2, 3, 4, 5}
	fmt.Printf("Debug u = %v\n", u)

	updateTest(u)
	for i, v := range u {
		fmt.Println(i, v)
	}

	_ = append(u[:3], u[4:]...)
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
