// 链接：https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// 题意：给定一个整型数组 nums 和一个数组 queries ，其中 queries[i] = [val_i, index_i] 。
//      对于第 i 个查询，我们执行 nums[index_i] += nums[index_i] + val_i;
//
//      返回一个整型数组 answer ，其中 answer[i] 就是执行第 i 个查询后， nums 中所有偶数的和。


// 数据限制：
//  1 <= nums.length <= 10 ^ 4
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4
//  1 <= queries.length <= 10 ^ 4
//  -(10 ^ 4) <= val_i <= 10 ^ 4
//  0 <= index_i < nums.length


// 输入： nums = [1,2,3,4], queries = [[1,0],[-3,1],[-4,0],[2,3]]
// 输出： [8,6,2,4]
// 解释： 最开始 nums 为 [1,2,3,4]
//       - 第 0 次查询执行 nums[0] += 1  ， nums 变为 [2,2,3,4] ，
//          所有偶数和为 2 + 2 + 4 = 8
//       - 第 1 次查询执行 nums[1] += -3 ， nums 变为 [2,-1,3,4] ，
//          所有偶数和为 2 + 4 = 6
//       - 第 2 次查询执行 nums[0] += -4 ， nums 变为 [-2,-1,3,4] ，
//          所有偶数和为 (-2) + 4 = 2
//       - 第 3 次查询执行 nums[3] += 2  ， nums 变为 [-2,-1,3,6] ，
//          所有偶数和为 (-2) + 6 = 4

// 输入： nums = [1], queries = [[4,0]]
// 输出： [0]
// 解释： 最开始 nums 为 [1]
//       - 第 0 次查询执行 nums[0] += 5  ， nums 变为 [5] ，
//          所有偶数和为 0


// 思路： 模拟
//
//      由于每次执行查询后， nums 中最多只有一个数会改变，
//      所以我们可以提前维护 nums 中所有偶数的和 evenSum 。
//
//      对于第 i 个查询按顺序进行如下处理：
//          1. 如果 nums[index_i] 原本是偶数，则先从 evenSum 减去
//          2. 执行加法操作： evenSum += val_i;
//          3. 如果 nums[index_i] 现在是偶数，则再加到 evenSum 中
//          4. 此时， evenSum 就是执行完前 i 次查询后， nums 中所有偶数的和
//      
//
//      时间复杂度： O(m + n)
//          1. 需要遍历 nums 中全部 O(m) 个数字一次
//          2. 需要遍历 queries 中全部 O(n) 个查询一次
//      空间复杂度： O(1)
//          1. 只需要维护常数个额外遍历即可


func sumEvenAfterQueries(nums []int, queries [][]int) []int {
    // evenSum 维护 nums 中所有偶数的和，初始化最开始数组中的偶数和
    evenSum := 0
    for _, num := range nums {
        if num & 1 == 0 {
            evenSum += num
        }
    }
    // ans 维护每一次查询的结果
    ans := make([]int, len(queries))
    // 遍历第 i 次查询
    for i := range queries {
        val, idx := queries[i][0], queries[i][1]
        // 1. 如果 nums[idx] 原本是偶数，则先从 evenSum 减去
        if nums[idx] & 1 == 0 {
            evenSum -= nums[idx]
        }
        // 2. 执行加法操作
        nums[idx] += val
        // 3. 如果 nums[index_i] 现在是偶数，则再加到 evenSum 中
        if nums[idx] & 1 == 0 {
            evenSum += nums[idx]
        }
        // 4. 此时， evenSum 就是执行完前 i 次查询后， nums 中所有偶数的和
        ans[i] = evenSum
    }

    return ans
}
