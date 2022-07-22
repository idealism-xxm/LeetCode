// 链接：https://leetcode.com/problems/smallest-number-in-infinite-set/
// 题意：有一个含有全部正整数的无限集合，实现一个数据结构，支持以下操作：
//          1. SmallestInfiniteSet(): 初始化一个含有全部正整数的无限集合
//          2. int popSmallest(): 移除并返回无限集合中最小的数
//          3. void addBack(int num): 如果 num 不在无限集合中，就添加到无限集合中


// 数据限制：
//  1 <= num <= 1000
//  最多会调用 popSmallest 和 addBack 共 1000 次


// 输入： ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
//       [[], [2], [], [], [], [1], [], [], []]
// 输出： [null, null, 1, 2, 3, null, 1, 4, 5]
// 解释： SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
//       smallestInfiniteSet.addBack(2);    // 2 已在集合中，无事发生
//       smallestInfiniteSet.popSmallest(); // 返回 1, 1 是集合中最小的数，移除并返回
//       smallestInfiniteSet.popSmallest(); // 返回 2, 2 是集合中最小的数，移除并返回
//       smallestInfiniteSet.popSmallest(); // 返回 3, 3 是集合中最小的数，移除并返回
//       smallestInfiniteSet.addBack(1);    // 1 被加回到集合中
//       smallestInfiniteSet.popSmallest(); // 返回 1, 因为 1 被加回到集合中，
//                                             现在 1 是集合中最小的数，移除并返回
//       smallestInfiniteSet.popSmallest(); // 返回 4, 4 是集合中最小的数，移除并返回
//       smallestInfiniteSet.popSmallest(); // 返回 5, 5 是集合中最小的数，移除并返回


// 思路： 模拟
//
//      可以注意到最多只会移除 1000 个数字，这个范围很小。
//
//      那我们可以逆向思维，只维护被删除数字的集合 deleted_nums 即可。
//
//      这样我们就能在 O(n) 内移除并返回最小的数字，并在 O(1) 内将数字加回到无限集合中。
//
//      同时我们可以再稍微做一点优化：记录无限集合中最小的数字 min_num 。
//
//      在进行移除操作时，我们先保存 min_num 到 ans 中，
//      然后不断对 min_num 加 1 ，直至 min_num 不在 deleted_nums 中，
//      最后返回 ans 即可。
//
//      在加回数字 num 操作时，我们将 num 从 deleted_nums 中移除，
//      如果 num < min_num ，则更新 min_num 为 num 即可。
//
//
//      设两个函数共调用 n 次。​
//
//      时间复杂度： popSmallest - O(n) | addBack - O(1)
//          1. popSmallest: 需要不断对 min_num 执行加 1 ，最差情况在有 O(n) 次
//          2. addBack: 在 O(1) 内就能从 deleted_nums 中移除 num
//      空间复杂度：O(n)
//          1. 需要存储全部 O(n) 个被删除的数字


use std::collections::HashSet;


struct SmallestInfiniteSet {
    // deleted_nums 维护从无限集合中删除的数字
    deleted_nums: HashSet<i32>,
    // min_num 维护无限集合中最小的数字
    min_num: i32,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl SmallestInfiniteSet {

    fn new() -> Self {
        SmallestInfiniteSet {
            deleted_nums: HashSet::new(),
            // 初始未删除任何数字，最小的数字就是 1
            min_num: 1,
        }
    }
    
    fn pop_smallest(&mut self) -> i32 {
        let ans = self.min_num;
        // 将 min_num 删除，放入到 deleted_nums 中
        self.deleted_nums.insert(self.min_num);
        // 不断将 min_num 加 1 ，直至 min_num 不在 deleted_nums 中
        while self.deleted_nums.contains(&self.min_num) {
            self.min_num += 1;
        }

        ans
    }
    
    fn add_back(&mut self, num: i32) {
        // 将 num 从 deleted_nums 中移除
        self.deleted_nums.remove(&num);
        // 如果 num < min_num ，则 num 就是无限集合中最小的数
        if num < self.min_num {
            self.min_num = num;
        }
    }
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * let obj = SmallestInfiniteSet::new();
 * let ret_1: i32 = obj.pop_smallest();
 * obj.add_back(num);
 */
