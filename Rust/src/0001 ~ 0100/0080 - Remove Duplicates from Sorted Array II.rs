// 链接：https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// 题意：给你一个升序数组 nums ，原地 删除重复出现的数字，
//      使每个数字最多出现 2 次 ，将剩余的数字都放在原数组前面，
//		返回删除后数组的新长度。
//      
//      要求：不使用额外数组，空间复杂度为 O(1)

// 数据限制：
//  1 <= nums.length <= 3 * 10 ^ 4
//  -(10 ^ 4) <= nums[i] <= 10 ^ 4
//  nums 已按升序排序

// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 
//      并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。 
//      不需要考虑数组中超出新长度后面的元素。

// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 解释：函数应返回新长度 length = 7, 
//      并且原数组的前五个元素被修改为 0, 0, 1, 1, 2, 3, 3 。 
//      不需要考虑数组中超出新长度后面的元素。


// 思路：双指针
//
//      使用双指针 l 和 r 来处理，
//      左指针 l 指向下一个可放入的位置，同时也表示结果数组的长度，
//      右指针 r 指向下一个可放入的数字
//
//      初始化 l = r = 2 ，表示前 2 个数字必定满足题意，无需特殊考虑，
//      也方便后续处理。
//      不过同时也要在最开始判断 s.len() <= 2 时，
//      直接返回 s.len() ，防止结果数组长度超过原有数组长度。
//
//      不断右移右指针 r ，
//      判断当前数字 nums[r] 与结果数组中倒数第 2 个数字 nums[l - 2] 的关系：
//          1. nums[r] == nums[l - 2]: 则说明结果数组中后 2 个数字都是 nums[r] ，
//              如果此时再放入 nums[r] ，则不满足题意，所以不处理
//          2. nums[r] != nums[l - 2]: 则说明结果数组中 nums[r] 出现的次数，
//              还不到 2 次，放入 nums[r] 后仍满足题意，所以放入后右移指针 l
//
//      这个方法非常通用，所以我们可以定义一个常量 MAX_APPEAR_TIMES ，
//      表示一个数字最大出现次数，这样只需要改动这个常量，就可以适用于其他题目了
//
//
//      时间复杂度：O(n)
//      空间复杂度：O(n)


// 一个数字最大出现次数
const MAX_APPEAR_TIMES: usize = 2;


impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        // 如果 Vec 长度小于等于数字最大出现次数，则不用处理，
        // 整个 Vec 都满足题意
        if nums.len() <= MAX_APPEAR_TIMES {
            return nums.len() as i32;
        }

        // nums[..MAX_APPEAR_TIMES] 满足题意，直接放入结果中即可
        let mut l = MAX_APPEAR_TIMES;
        // 右指针从 l 开始遍历
        for r in l..nums.len() {
            // 如果当前数字 nums[r] 不等于结果数组中倒数第 MAX_APPEAR_TIMES 个数，
            // 则说明 nums[r] 在结果数组中出现次数还未到 MAX_APPEAR_TIMES 次 ，
            // 可以继续把 nums[r] 放入到结果数组中
            if nums[r] != nums[l - MAX_APPEAR_TIMES] {
                nums[l] = nums[r];
                l += 1;
            }
        }

        // 现在左指针 l 就是结果数组的长度
        l as i32
    }
}
