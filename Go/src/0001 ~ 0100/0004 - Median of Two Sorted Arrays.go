// 链接：https://leetcode.com/problems/median-of-two-sorted-arrays/
// 题意：给定两个有序数组 a 和 b，长度分别为 m 和 n，使用时间复杂度不超过 O(log(m + n)) 的算法求出两个数组合并后的中位数
// 思路：二分

// 一看时间复杂度就知道要二分，但是想了半天，一直认为要对两个数组都二分，导致很复杂，思路也比较混乱
// 看了题解一半就明白如何去做了，还是要会分析，利用已知的信息降低复杂度

// 假设两个数组的分割点分别是 i, j，即分为 a[0]...a[i - 1] | a[i]...a[m - 1] 和 b[0]...b[j - 1] | b[j]...b[n - 1]
// 且中点值从 a[i - 1]、a[i]、b[j - 1]、b[j] 中产生
// 设 i + j = lenLeft，(m - i) + (n - j) = lenRight，m <= n
// 1. 若 m + n 是奇数
// 		则必有：
// 		(1) lenLeft == lenRight - 1
// 		(2) a[i - 1] <= b[j] && b[j - 1] <= a[i]
//		且中点值为：min(a[i], b[j])
// 2. 若 m + n 是偶数
// 		则必有：
// 		(1) lenLeft == lenRight
// 		(2) a[i - 1] <= b[j] && b[j - 1] <= a[i]
//		且中点值为：float64(max(a[i - 1], b[j - 1]) + min(a[i], b[j])) / 2
// 综上，我们可以统一找到分割点，然后根据 m + n 的奇偶性返回中点值
// 条件可以统一为：
// (1) lenLeft = (lenLeft + lenRight) >> 1 = (m + n) >> 1
// (2) a[i - 1] <= b[j] && b[j - 1] <= a[i]
// 二分 数组a，得到 i，从而为了满足 条件(1) 可以计算出 j = half - i，然后判断 条件(2) 是否满足即可
// 若 条件(2) 满足，则二分结束，返回相应的中点值即可
// 若 条件(2) 不满足：
//		若 a[i - 1] > b[j]，则 a 数组的分割点必定可以左移（r = i - 1），即 i = 左边区间的中点
// 		若 b[j - 1] > a[i]，则 a 数组的分割点必定可以右移（l = i + 1），即 i = 右边区间的中点
//		【注意】不可能出现 a[i - 1] > b[j] && b[j - 1] > a[i]
// 同时，可以适当，将上面三种情况融合成两种，即调整左移的端点 r = i，即可满足 条件(2) 也并入，最终 l 即 i 的分割点
// 【注意】同时还要边界情况，防止越界
// 实际上由于 m < n，所以
// 		j == 0 时必定有 i >= m
// 		j == n 时必定有 i == 0 && m == n
func findMedianSortedArrays(a []int, b []int) float64 {
	m, n := len(a), len(b)
	// 让 a 的长度更小，使算法复杂度为：O(log(min(m, n)))
	if m > n {
		a, b = b, a
		m, n = n, m
	}
	half := (m + n) >> 1 // 由于 m <= n，则必定有：m <= half && lenLeft ==  half

	l, r := 0, m
	// 二分
	for ; l < r;  {
		i := (l + r) >> 1
		j := half - i
		// 【注意】不可能出现 a[i - 1] > b[j] && b[j - 1] > a[i]
		if j > 0 && b[j - 1] > a[i] {
			// 此处要取 i + 1，因为本次判定 i 必定不能为分界线，所以要从右区间要从 i + 1 开始
			l = i + 1 //若 b[j - 1] > a[i]，则 a 数组的分隔点必定可以右移，即 i = 右边区间的中点
		} else {
			r = i // 若 a[i - 1] > b[j]，则 a 数组的分隔点必定可以左移，即 i = 左边区间的中点
		}
	}

	i := l
	j := half - i
	var rightMin int
	// 处理好边界情况，防止越界
	if i == m {
		rightMin = b[j]
	} else if j == n {
		rightMin = a[i]
	} else {
		rightMin = min(a[i], b[j])
	}
	if ((m + n) & 1) == 1 {
		// 长度为奇数
		return float64(rightMin)
	}

	var leftMax int
	if i == 0 {
		leftMax = b[j - 1]
	} else if j == 0 {
		leftMax = a[i - 1]
	} else {
		leftMax = max(a[i - 1], b[j - 1])
	}
	// 长度为偶数
	return float64(leftMax + rightMin) / 2.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}