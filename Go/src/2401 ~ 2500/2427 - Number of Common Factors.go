// 链接：https://leetcode.com/problems/number-of-common-factors/
// 题意：给定两个正整数 a 和 b ，求 a 和 b 全部共因子的数量？


// 数据限制：
//  1 <= a, b <= 1000


// 输入： a = 12, b = 6
// 输出： 4
// 解释： 12 和 6 的公因子是： 1, 2, 3, 6

// 输入： a = 25, b = 30
// 输出： 2
// 解释： 25 和 30 的公因子是： 1, 5


// 思路： 模拟
//
//      a, b 的公因子必定是其最大公约数的因子。
//
//      所以可以先求出 a 和 b 的最大公约数 mx ，
//      然后求 mx 的因子数即可。
//
//
//      时间复杂度：O(sqrt(min(a, b)))
//          1. 需要求 a 和 b 的最大公约数，辗转相除法的时间复杂度为 O(log(max(a, b)))
//          2. 需要遍历全部 O(sqrt(mx)) 个数，最差情况下 mx 为 min(a, b)
//      空间复杂度：O(1)
//          1. 只需要使用常数个额外变量即可


func commonFactors(a int, b int) int {
    // ans 维护满足题意的公因子数
    ans := 0
    // a, b 的公因子必定是其最大公约数的因子
    mx := gcd(a, b)
    // 枚举因子 [1, sqrt(mx))
    factor := 1
    // 这里不取等号是为了最后特殊处理恰好开平方的情况
    for factor * factor < mx {
        // 如果 factor 是 mx 的因子，那么 factor / mx 也是，
        // 并且两者不想等，所以找到 2 个满足题意的公因子
        if mx % factor == 0 {
            ans += 2
        }
        factor += 1
    }
    // 如果 factor == sqrt(mx) ，那么 factor 是 1 个满足题意的公因子
    if factor * factor == mx {
        ans += 1
    }

    return ans
}

// 辗转相除法计算最大公约数
func gcd(a, b int) int {
    for b != 0 {
        a, b = b, a % b;
    }

    return a
}
