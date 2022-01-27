// 链接：https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 题意：给定一个长度为 n 的正整数数组 nums ，
//      求 nums[i] XOR nums[j] 的最大值？ (0 <= i <= j < n)

// 数据限制：
//  1 <= nums.length <= 2 * 10 ^ 5
//  0 <= nums[i] <= 2 ^ 31 - 1
 

// 输入： nums = [3,10,5,25,2,8]
// 输出： 28
// 解释： 最大的结果是 5 XOR 25 = 28

// 输入： nums = [14,70,53,83,49,91,36,80,92,51,66,70]
// 输出： 127


// 思路：Trie
//
//      按照题意模拟就是枚举两个数 nums[i] 和 nums[j] (i < j) ，
//      然后找到这些数的异或最大值 max(nums[i] ^ nums[j]) 即可。
//
//      但这种方法的时间复杂度是 O(n ^ 2) ，在给定的数据范围内会超时，
//      所以需要使用其他方法优化，求异或最值的问题很多时候可以用 trie 来进行优化。
//
//      我们枚举其中一个数 nums[j] ，先将 nums[j] 加入到 trie 中，
//      然后我们可以在 nums[0..=j] 中，贪心地去找与 nums[j] 异或后最大的值 ans ，
//      方式就是从高到低枚举 nums[j] 的第 i 位，
//      利用 trie 就可以每次找到符合要求的那一部分数，直至枚举完所有的位。
//
//      设 bit = (nums[j] >> i) & 1 ，则先贪心判断当前是否存在第 i 位为 1 - bit 的数：
//          1. 如果存在，则答案 ans 的第 i 位可以为 1 ，然后走 trie 的 1 - bit 的子节点
//          2. 如果不存在，则答案 ans 的第 i 位可以为 0 ，然后走 trie 的 bit 的子节点 
//       
//      时间复杂度： O(n)
//      空间复杂度： O(n)

// 简化创建节点
#[derive(Default)]
pub struct TrieNode {
    chilren: [Option<Box<TrieNode>>; 2]
}

// 简化创建 trie
#[derive(Default)]
pub struct Trie {
    root: TrieNode
}

impl Trie {
    fn new() -> Self {
		Self::default()
	}

    // 将 num 插入到 trie 中
    pub fn insert(&mut self, num: i32) {
        // 先获取 trie 的根结点
        let mut node = &mut self.root;
        // 从高位开始插入每一位的情况
        for i in (0..=31).rev() {
            // 获取第 i 位的值
            let bit = ((num >> i) & 1) as usize;
            // 获取对应的子节点，如果没有则创建一个返回
            node = node.chilren[bit].get_or_insert_with(Default::default);
        }
    }

    pub fn max_xor(&mut self, num: i32) -> i32 {
        // 先获取 trie 的根节点
        let mut node = &mut self.root;
        // ans 维护 num 与当前 trie 内某个数的异或最大值
        let mut ans = 0;
        // 从高位开始判断
        for i in (0..=31).rev() {
            // 获取第 i 位的值
            let bit = ((num >> i) & 1) as usize;
            if let Some(ref mut child) = node.chilren[1 - bit] {
                // 如果 tire 存在与 num 第 i 位的值相反的数，
                // 则 ans 第 i 位可以设置为 1
                ans |= 1 << i;
                // 下一位就可以从 node.chilren[1 - bit] 中继续判断处理
                node = child;
            } else {
                // 如果 tire 只存在与 num 第 i 位的值相同的数，
                // 下一位就只能从 node.chilren[bit] 中继续判断处理
                node = node.chilren[bit].as_mut().unwrap();
            }
        }

        // 返回异或最大值
        ans
    }
}

impl Solution {
    pub fn find_maximum_xor(nums: Vec<i32>) -> i32 {
        // 使用迭代器遍历
        nums.iter()
            // 使用 flod 积累最大值和最新的 trie
            .fold(
                // 初始 ans = 0 ，trie 为空
                (0, Trie::new()), 
                // 接收当前积累的 ans 和 trie ，并通过 num 转换成新的值返回
                |(ans, mut trie), &num| 
                {
                    // 先将 num 插入到 num 中
                    trie.insert(num);
                    // 然后返回最新的异或最大值和最新的 trie
                    (ans.max(trie.max_xor(num)), trie)
                }
            )
            // 二元组第一个值就是所求的异或最大值，直接返回
            .0
    }
}
