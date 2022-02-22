// 链接：https://leetcode.com/problems/majority-element/
// 题意：给定一个整数数组，有一个数出现的次数超过一半（向下取整），
//      找出这个数？


// 数据限制：
//  n == nums.length
//  1 <= n <= 5 * 10 ^ 4
//  -(2 ^ 31) <= nums[i] <= 2 ^ 31 - 1


// 输入：nums = [3,2,3]
// 输出：3
// 解释： 3 出现 2 次， 
//       2 出现 1 次。

// 输入：nums = [2,2,1,1,1,2,2]
// 输出：2
// 解释： 2 出现 4 次，
//       1 出现 3 次。


// 思路2：位运算
//
//		针对每一个二进制：
//		    1. 若众数的该位为 1 ，那么这位为 1 的数字个数必定超过一半
//		    2. 若众数的该位为 0 ，那么这位为 1 的数字个数必定不超过一半
//
//      所以我们维护一个长度为 32 的数组 count ，
//      count[i] 表示所有数中第 i 个二进制位为 1 的数字个数。
//
//      那么最终统计完成后，我们维护 majority 表示众数，
//      遍历所有的二进制位 i ，如果 count[i] > nums.len() / 2 ，
//      则众数的第 i 位是 1 ，执行 majority |= 1 << i
//
//
//      假设 N 为 n 的最大值，这里是 2 ^ 31 - 1
//
//		时间复杂度： O(nlogN)
//          1. 需要遍历全部 O(n) 个数字
//          2. 每个数字都要遍历全部二进制位，可以看作 O(1) ，
//              但实际严格来讲应该是 O(logN)
//		空间复杂度： O(logN)
//          1. 实际上开辟的二进制位空间与 n 有关系，严格来说应该是 O(logN)


impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        // count[i] 表示所有数中第 i 个二进制位为 1 的数字个数
        let mut count = [0; 32];
        // 遍历每个数
        for num in nums.iter() {
            // 遍历当前数的所有二进制位
            for i in 0..32 {
                // 如果第 i 位是 1 ，则 count[i] += 1
                if num & (1 << i) != 0 {
                    count[i] += 1;
                }
            }
        }

        // 维护 majority ，表示是众数
        let mut majority = 0;
        // 遍历所有二进制位
        for i in 0..32 {
            // 如果第 i 位为 1 的个数超过一半，则众数的第 i 位是 1
            if count[i] > nums.len() >> 1 {
                majority |= 1 << i;
            }
        }

        majority
    }
}


// 思路3： Boyer-Moore 投票算法
//
//		题解也提到了最优的 Boyer-Moore 投票算法
//		先指定众数 majority = random ，并且其出现的次数 count = 0
//		然后遍历整个数组：
//		    1. count == 0 时： 令 majority = num
//		    2. majority == num 时： count++
//		    3. majority != num 时： count--
//
//		算法正确性：
//		    1. 由于先判断 count == 0 时，令 majority = num ，
//              所以 count 必定是非负数
//		    2. 若 majority 就是众数，那么下一次 count 为 0 时，
//              必定抵消了相同数量的非众数，
//              剩余的数组中，众数还是占一半以上
//		    3. 若 majority 不是众数，那么下一次 count 为 0 时，
//              最多抵消了相同数量的众数，
//		    	剩余的数组中，众数还是占一半以上
//
//
//		时间复杂度： O(n)
//          1. 只需要遍历全部 O(n) 个数字一次
//		空间复杂度： O(1)
//          1. 只维护 2 个变量，所以空间复杂度为 O(1)


impl Solution {
    pub fn majority_element(nums: Vec<i32>) -> i32 {
        // 维护 majority ，表示众数
        let mut majority = 0;
        // 维护 count ，表示当前众数的个数
        let mut count = 0;
        // 遍历每个数
        for num in nums {
            // 如果当前众数的个数为 0 ，则更新当前众数为 num
            if count == 0 {
                majority = num;
            }

            if majority == num {
                // 如果当前众数是 num ，那就增加其出现次数
                count += 1
            } else {
                // 如果当前众数不是 num ，那就减少其出现次数，
                // 抵消 num 的出现次数
                count -= 1
            }
        }
        majority
    }
}
