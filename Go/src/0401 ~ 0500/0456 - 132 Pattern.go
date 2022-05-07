// 链接：https://leetcode.com/problems/132-pattern/
// 题意：给定一个数组 nums ，判断是否存在满足以下条件的 i, j, k ？
//          1. i < j < k
//          2. nums[i] < nums[k] < nums[j]


// 数据限制：
//  n == nums.length
//  1 <= n <= 2 * 10 ^ 5
//  -(10 ^ 9) <= nums[i] <= 10 ^ 9


// 输入： nums = [1,2,3,4]
// 输出： false
// 解释： 不存在满足题意的 i, j, k

// 输入： nums = [3,1,4,2]
// 输出： true
// 解释： 子序列 [1,4,2] 满足题意

// 输入： nums = [-1,3,2,0]
// 输出： true
// 解释： 子序列 [-1,3,2], [-1,3,0] 和 [-1,2,0] 均满足题意


// 思路： 单调栈
//
//      我们维护 nums[:k] 中的最小值 leftMin ，
//      以及一个单调递减栈 stack ，
//      其中 stack 存放二元组 (min, max) ：
//          min 是 nums[:k] 中的最小值 nums[i] ， 
//          max 是 nums[i:k] 中的最大值 nums[j] 。
//
//      令所有元素的 min 单调递减，且每个元素的 min < max ，
//      这样就保证满足题意的 i < j && nums[i] < nums[j] 的条件。
//
//      那么后续只需要维护这个关系，找到满足题意的 nums[k] 即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历全部 O(n) 个数字
//      空间复杂度：O(n)
//          1. 需要维护一个栈 stack ，最差情况下需要保存全部 O(n) 个数字


func find132pattern(nums []int) bool {
    // stack 存放 (min, max) 二元组，其中：
    //      min 是 nums[:k] 中的最小值 nums[i] ， 
    //      max 是 nums[i:k] 中的最大值 nums[j] 。
    //
    // 令所有元素的 min 单调递减，且每个元素的 min < max ，
    // 这样就保证满足题意的 i < j && nums[i] < nums[j] 的条件。
    //
    // 那么后续只需要维护这个关系，找到满足题意的 nums[k] 即可。
    stack := make([]*Pair, 0)
    // 维护 nums[:k] 中的最小值，初始化为第一个数字
    leftMin := nums[0]
    // 遍历 nums 中的每个数字
    for _, num := range nums {
        // 当栈不为空 且 stack.top().max <= num 时，不断弹出栈顶元素。
        //
        // 因为 leftMin 只减不增，必有 leftMin <= stack.top().min ，
        // 那么当 num >= stack.top().max 时，
        // (leftMin, num) 代表的范围不会更小，且包含栈顶元素的范围，
        // 可以直接替换栈顶元素，能保证答案不会更差。
        for len(stack) != 0 && stack[len(stack) - 1].max <= num {
            stack = stack[:len(stack) - 1]
        }
        // 此时有 num < stack.top().max ，
        // 如果 num > stack.top().min ，则说明满足题意，直接返回 true
        if len(stack) != 0 && stack[len(stack) - 1].min < num {
            return true
        }

        if leftMin < num {
            // 如果 leftMin < num ，则将 (leftMin, num) 压入栈中。
            //
            // 因为 leftMin 只减不增，必有 leftMin <= stack.top().min ，
            // 那么此时 (leftMin, num) 入栈后，
            // 满足所有元素的 min 是单调递减的，只要能保证 min < max 就可以入栈
            stack = append(stack, &Pair{ min: leftMin, max: num })
        } else {
            // 如果 leftMin >= max ，则更新 leftMin 为 num
            leftMin = num
        }
    }

    // 遍历完没找到，则没有满足题意的子序列，返回 false
    return false
}

type Pair struct {
    min, max int
}
