// 链接：https://leetcode.com/problems/max-points-on-a-line/
// 题意：对链表进行排序，要求时间复杂度为 O(nlogn) ，空间复杂度为 O(1) ？

// 输入： 4->2->1->3
// 输出： 1->2->3->4

// 输入： -1->5->3->4->0
// 输出： -1->0->3->4->5

// 思路： map
//
//		其实很容易就想到 y = kx + b 解析式，
//		所以第一反应就是将 k, b 求出来作为键，
//		统计每一对 k, b 的数量，进而求的对应的点数
//
//		但有特殊情况需要处理，就是 k 不存在的时候，就会很快想到另一种方法，
//		使用 "{dy}/{dx}" 作为 k ，其中 dy 和 dx 是除以最大公约数的结果，
//		这样这个“斜率”字符串既可以避免斜率为不存在的情况，也可以避免浮点数精度的误差
//
//		但还有另一个问题就是 b 还是需要求出来，并有可能是浮点数，
//		这点还是忘了以前的集合知识造成的，陷入了盲区
//
//		后来看了题解才记起来，直线的解析式还有点斜式可以使用：
//		只需要知道直线的斜率 k 和直线上一点，就可以唯一确一条直线，
//		而我们已经知道这两个了，所以可以直接统计即可
//		对于每个 points[i] 的每个斜率维护一个斜率统计数，
//		计算其和 points[j] (i < j) 所在直线的斜率并计数，
//		最后计算这些统计数的最大值即可
//		当然还有一种特殊情况需要考虑，那就是 points[i] 和 points[j] 是同一个点，
//		那么只需对这种情况维护一个统计值即可，最后取最大值的时候合并计算即可
//
//		其实中间还蹦出来另外一种思路，不过空间复杂度是 O(n ^ 2) 的，
//		就是计算所有边的斜率，并处理成上述的斜率字符串形式，
//		然后根据三点共线的条件，用并查集合并斜率相同且有公共点的边集即可，
//		最后根据： 边的数量 = 点数 * (点数 - 1) / 2
//		可以求得点数
//
//		时间复杂度： O(n ^ 2)
//		空间复杂度： O(n)

import "fmt"

func maxPoints(points [][]int) int {
	length := len(points)
	// 如果不够三个，则直接返回长度即可
	if length < 3 {
		return length
	}

	result := 0
	for i := 0; i < length; i++ {
		// 记录斜率为 k 时的点点个数
		kToNum := make(map[string]int)
		// 记录与 points[i] 时同一个点的数量（包括 points[i] 自身）
		duplicate := 1
		for j := i + 1; j < length; j++ {
			dy := points[j][1] - points[i][1]
			dx := points[j][0] - points[i][0]
			if dy == 0 && dx == 0 {
				// 同一个点，则直接对计数 +1
				duplicate++
				continue
			}

			_gcd := gcd(dy, dx)
			// 都除以最大公约数后就不用直接计算 k ，可以避免浮点数的精度误差
			dy, dx = dy / _gcd, dx / _gcd
			k := fmt.Sprintf("%d/%d", dy, dx)
			kToNum[k]++
		}

		// 遍历所有斜率 k ，更新最大值
		for k := range kToNum {
			result = max(result, kToNum[k] + duplicate)
		}
		// 所有点都是同一个点
		result = max(result, duplicate)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func gcd(a, b int) int {
	for ; b != 0; {
		a, b = b, a % b
	}
	return a
}
