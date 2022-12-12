package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	var a, b, c int

	c = len(nums1) - 1 // the last index of nums1

	// set a to last value of nums1
	a = m - 1

	// set b to last value of nums2
	b = len(nums2) - 1

	for b >= 0 && c >= 0 {
		if a < 0 {
			// nums1 is all out
			nums1[c] = nums2[b]
			b--
			c--
			continue
		}

		if nums1[a] > nums2[b] {
			nums1[c] = nums1[a]
			a--
		} else {
			nums1[c] = nums2[b]
			b--
		}
		c--
	}

}
