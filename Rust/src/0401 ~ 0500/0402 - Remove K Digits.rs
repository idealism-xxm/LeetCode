// 链接：https://leetcode.com/problems/remove-k-digits/
// 题意：给你一个以字符串表示的非负整数 num 和一个整数 k ，
//      移除这个数中的 k 个数位，使得剩下的数字最小。
//      以字符串形式返回这个最小的数字。


// 数据限制：
//  1 <= k <= num.length <= 10 ^ 5
//  num 仅由 '0' - '9' 组成
//  除了 '0' 本身之外，num 不含任何前导零


// 输入：num = "1432219", k = 3
// 输出："1219"
// 解释：移除 4, 3, 2 后，最小的数是 1219

// 输入：num = "10200", k = 1
// 输出："200"
// 解释：移除 1 后，最小的数是 200 ，
//      注意结果字符串不能含有前导零

// 输入：num = "10", k = 2
// 输出：移除 1, 0 后，最小的数是 0


// 思路：贪心 + 单调栈
//
//      所有位数相同的数比较大小，肯定是从最左侧开始比较，
//      在第一处不同的数位处就能区分出大小。
//
//      那么我们可以采取贪心的思想，每次只移除一个数，
//      从最左侧开始遍历，不断对比 num[i] 和 num[i - 1] ：
//          1. num[i] > num[i - 1]: 移除 num[i - 1] 会使的结果变大，
//              所以不进行移除操作，直接处理下一个
//          2. nus[i] == num[i - 1]: 移除 num[i - 1] 在目前看来不会使结果变得更优，
//              所以暂时不进行移除操作，先处理下一个
//          3. num[i] < num[i - 1]: 移除 num[i - 1] 会使得结果变小，
//              所以必定要移除 num[i - 1] ，然后进入下一轮移除循环
//
//      如果 num 中的数位都是非递减的，那么移除最后一个数字即可。
//
//      但目前这个解法是 O(nk) 的时间复杂度，当 num 是非递减时为最差情况。
//      由于 n 和 k 都可以取到 10 ^ 5 ，所以必须使用 O(n) 或 O(nlogn) 的时间复杂度才行。
//
//
//      那我们思考一下这个贪心过程，其实存在大量重复无意义的遍历。
//
//      假设上次删除了 num[i - 1] ，那么说明 num[..i - 1] 是非递减的。
//
//      如果本次从最左侧开始遍历，那么前 num[..i - 2] 一定是还是非递减的，
//      这些遍历就是无效遍历。
//
//      实际上我们可以利用这些信息，本次直接从 num[i] 继续开始遍历，
//      因为只有从 num[i - 2] 和 num[i] 开始的相邻的数字关系不确定。
//
//      这种其实就是单调栈的行为，所以可以直接使用单调栈维护合法的数位序列即可。
//
//      需要注意的有两点：
//          1. 我们最多只能删除 k 个数位，所以单调栈出栈操作最多有 k 次
//          2. 结果数字不含前导 0 ，所以只有栈非空 或者 当前数位不是 0 时，才能入栈
//      
//      单调栈处理完成后，如果发现还需要继续删除数位，则把栈顶的 k 个数位弹出。
//
//      最后根据不同情况返回：
//          1. 如果栈为空，则所有数字都可以被删除 或者 只剩前导零，那么直接返回 "0"
//          2. 如果栈不为空，则将栈中的数位收集成字符串返回
//
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)


impl Solution {
    pub fn remove_kdigits(num: String, mut k: i32) -> String {
        let mut k = k as usize;
        // 定义一个单调递增栈，最多存储 num.len() 个数位
        let mut stack = Vec::with_capacity(num.len());
        // 遍历 num 中的每个字符
        for ch in num.chars() {
            // 为了维持单调递增栈，只要满足以下三个条件，就需要不断移除栈顶元素
            //  1. k > 0: 还能移除数位
            //  2. !stack.is_empty(): 栈中还有数位
            //  3. stack.last().unwrap() > &ch: 栈顶元素比 ch 大
            while k > 0 && !stack.is_empty() && stack.last().unwrap() > &ch {
                // 移除栈顶元素
                stack.pop();
                // 数位可以出次数 -1
                k -= 1;
            }
            // 当栈不为空 或 ch 不是 '0' 时，可以将 ch 入栈，防止产生前导零
            if !stack.is_empty() || ch != '0' {
                stack.push(ch);
            }
        }

        if stack.len() <= k {
            // 如果栈为空，则直接返回 "0"
            "0".to_string()
        } else {
            // 如果栈不为空，则组装成字符串
            stack
                // 转成迭代器
                .iter()
                // 只取前 stack.len() - k 个字符，避免移除的数字不够
                .take(stack.len() - k)
                // 收集成字符串
                .collect()
        }
    }
}
