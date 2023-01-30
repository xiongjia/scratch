package leetcode

func FindMedianSortedArraysSolution1(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return FindMedianSortedArraysSolution1(nums2, nums1)
	}

	// TODO
	return 0.0
}
