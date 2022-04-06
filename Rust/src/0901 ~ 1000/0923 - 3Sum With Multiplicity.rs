// 链接：https://leetcode.com/problems/3sum-with-multiplicity/
// 题意：给定一个非负整数数组 arr ，求满足以下两个条件的三元组个数？
//          1. i < j < k
//          2. arr[i] + arr[j] + arr[k] == target


// 数据限制：
//  3 <= arr.length <= 3000
//  0 <= arr[i] <= 100
//  0 <= target <= 300


// 输入： arr = [1,1,2,2,3,3,4,4,5,5], target = 8
// 输出： 20
// 解释： (1, 2, 5) 出现 8 次
//       (1, 3, 4) 出现 8 次
//       (2, 2, 4) 出现 2 次
//       (2, 3, 3) 出现 2 次

// 输入： arr = [1,1,2,2,2,2], target = 5
// 输出： 12
// 解释： (1, 2, 2) 出现 12 次


// 思路： Map + 分类讨论
//
//      注意题目中限制 i < j < k ，
//      本质目的是限制 arr 中的每个数在三元组中只能出现一次，
//      所以我们可以通过分类讨论的方法来保证这个条件必定成立，
//      这样就只需要考虑 arr[i] + arr[j] + arr[k] == target 的条件。
//
//      我们先统计 arr 中每个数出现的次数到 cnt 中，
//      其中 cnt[a] 表示 arr 中 a 出现的次数。
//
//      考虑所有三元组 (a, b, c) 的情况，
//      保证 a <= b <= c ，使得方案数不会重复：
//          1. a == b == c: 则需要从 cnt[a] 个 a 中选择 3 个数，
//              共有 C(cnt[a], 3) 种选择方案
//          2. a == b < c: 则需要从 cnt[a] 个 a 中选择 2 个数，
//              再从 cnt[c] 个 c 中选择 1 个数，
//              共有 C(cnt[a], 2) * C(cnt[c], 1) 种选择方案
//          3. a < b == c: 则需要从 cnt[b] 个 b 中选择 2 个数，
//              再从 cnt[a] 个 a 中选择 1 个数，
//              共有 C(cnt[b], 2) * C(cnt[a], 1) 种选择方案
//          4. a < b < c: 则需要选择不同的 a, b, c 各一个数，
//              共有 C(cnt[a], 1) * C(cnt[b], 1) * C(cnt[c], 1) 种选择方案
//
//      那么我们只需要枚举所有的二元组 (a, b) 即可，其中 a <= b ，
//      然后计算出对应的 c ，再按照上面的分类讨论情况计算对答案的贡献。
//
//
//      设 n 为 arr 中的数的个数， w 为 arr 中的最大数。
//
//		时间复杂度： O(n + w ^ 2)
//          1. 需要遍历全部 O(n) 个数
//          2. 需要枚举 O(w ^ 2) 个二元组 (a, b) 的方案数
//		空间复杂度： O(w)
//          1. 需要维护全部 O(w) 个不同数字的出现次数


// arr[i] 的最大值
const MAX_NUM: usize = 100;
// ans 需要取的模
const MOD: i64 = 1_000_000_007;


impl Solution {
    pub fn three_sum_multi(arr: Vec<i32>, target: i32) -> i32 {
        // 统计 arr 中每个数字出现的次数
        let mut cnt = vec![0; MAX_NUM + 1];
        for &num in &arr {
            cnt[num as usize] += 1;
        }

        // 统计满足题意的三元组的个数
        let mut ans = 0;

        // 枚举 a 的值
        for a in 0..=MAX_NUM {
            // 枚举 b 的值，保证 a <= b
            for b in a..=MAX_NUM {
                // 计算 c 的值
                let c = target - a as i32 - b as i32;
                // 保证 a <= b <= c <= MAX_NUM
                if c < b as i32 || c > MAX_NUM as i32 {
                    continue;
                }

                let c = c as usize;
                if a == b && b == c {
                    // 如果 a == b == c ，
                    // 则需要从 cnt[a] 个 a 中选择 3 个数，
                    // 共有 C(cnt[a], 3) 选择方案
                    ans += cnt[a] * (cnt[a] - 1) * (cnt[a] - 2) / 6;
                } else if a == b {
                    // 此时是 a == b < c ，
                    // 则需要从 cnt[a] 个 a 中选择 2 个数，
                    // 再从 cnt[c] 个 c 中选择 1 个数，
                    // 共有 C(cnt[a], 2) * C(cnt[c], 1) 种选择方案
                    ans += cnt[a] * (cnt[a] - 1) / 2 * cnt[c];
                } else if b == c {
                    // 此时是 a < b == c ，
                    // 则需要从 cnt[b] 个 b 中选择 2 个数，
                    // 再从 cnt[a] 个 a 中选择 1 个数，
                    // 共有 C(cnt[b], 2) * C(cnt[a], 1) 种选择方案
                    ans += cnt[b] * (cnt[b] - 1) / 2 * cnt[a];
                } else {
                    // 此时必定是 a < b < arr[k] ，
                    // 则需要选择不同的 a, b, c 各一个数，
                    // 共有 cnt[a] * cnt[b] * cnt[c] 种选择方案
                    ans += cnt[a] * cnt[b] * cnt[c];
                }
                // 对 ans 取模
                ans %= MOD;
            }
        }

        ans as i32
    }
}
