// 链接：https://leetcode.com/problems/merge-sorted-array/
// 题意：给定两个升序的整型数组 nums1 和 nums2 ，以及他们有效的数字个数 m 和 n ，
//		nums1 长度不小于 m + n ，将 nums2 的数字都放入 nums1 中，且保持升序。

// 输入：
// nums1 = [1,2,3,0,0,0], m = 3
// nums2 = [2,5,6],       n = 3
// 输出：[1,2,2,3,5,6]

// 思路1：模拟即可
//		先将 nums1[:m] 的元素都移动到 nums1[n:m+n]
//		然后按照升序，每次选择两个数组中第一个数字中较小者放入即可
//		时间复杂度： O(m + n)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	maxNum := 0
	if m > 0 && maxNum <= nums1[m - 1] {
		maxNum = nums1[m - 1] + 1
	}
	if n > 0 && maxNum <= nums2[n - 1] {
		maxNum = nums2[n - 1] + 1
	}
	for i := m - 1; i >= 0; i-- {
		nums1[n + i] = nums1[i]
	}
	for i, j, k := n, 0, 0; k < m + n; k++ {
		num, di, dj := maxNum, 0, 0
		if i < m + n {
			if nums1[i] < num {
				num, di, dj = nums1[i], 1, 0
			}
		}
		if j < n {
			if nums2[j] < num {
				num, di, dj = nums2[j], 0, 1
			}
		}
		nums1[k] = num
		i, j = i + di, j + dj
	}
}

// 思路2：模拟即可
//		看了题解才发现自己思维顽固，没有充分利用数组的特点，
//		既然 nums1 中后半部分为空，那么从大到小放入即可，这样和 思路1 想法一样而不用循环两次
//		时间复杂度： O(m + n)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	for i, j, k := m - 1, n - 1, m + n - 1; j >= 0; k-- {
		// 由于从后往前遍历，所以只要 nums2 全部放入 nums1 了， nums1 中剩余的数字就是有序的
		if i >= 0 && nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}
}
