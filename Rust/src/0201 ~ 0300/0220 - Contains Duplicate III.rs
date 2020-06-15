// 链接：https://leetcode.com/problems/contains-duplicate-iii/
// 题意：给定一个数组和一个整数 k ，判断是否存在一对数的差最多为 t ，
//      且下标差最多为 k ？

// 输入： nums = [1,2,3,1], k = 3, t = 0
// 输出： true

// 输入： nums = [1,0,1,1], k = 1, t = 2
// 输出： true

// 输入： nums = [1,5,9,1,5,9], k = 2, t = 3
// 输出： false

// 思路1： 滑动窗口
//
//		遍历数组 nums ，
//      每次判断 [i-t, i) 内是否存在与 nums[i] 的差是否小于等于 t 的数，
//      若存在，且则直接返回 true，
//
//      在循环中没有返回，则所有数都不满足题意，直接返回 false
//
//      时间复杂度： O(n * min(n, t))
//      空间复杂度： O(1)

use std::cmp;

impl Solution {
    pub fn contains_nearby_almost_duplicate(nums: Vec<i32>, k: i32, t: i32) -> bool {
        let t = t as i64;
        for (i, num) in nums.iter().enumerate() {
            let num = *num as i64;
            let start = cmp::max(0, i as i32 - k) as usize;
            // 遍历 [i-t, t) 内的数，判断其与 num 的差是否小于等于 t
            for number in nums[start..i].iter() {
                let number = *number as i64; 
                // 若存在 number 使得其与 num 的差小于等于 t ，则直接返回 true
                if (num - number).abs() <= t {
                    return true;
                }
            }
        }
        // 循环内未返回，则所有数都不满足题意，直接返回 false
        return false;
    }
}

// 思路1： 滑动窗口
//
//		遍历数组 nums ，
//      每次判断 [i-t, i) 内是否存在与 nums[i] 的差是否小于等于 t 的数，
//      若存在，且则直接返回 true，
//
//      在循环中没有返回，则所有数都不满足题意，直接返回 false
//
//      时间复杂度： O(n * min(n, t))
//      空间复杂度： O(1)

use std::cmp;

impl Solution {
    pub fn contains_nearby_almost_duplicate(nums: Vec<i32>, k: i32, t: i32) -> bool {
        let t = t as i64;
        for (i, num) in nums.iter().enumerate() {
            let num = *num as i64;
            let start = cmp::max(0, i as i32 - k) as usize;
            // 遍历 [i-t, t) 内的数，判断其与 num 的差是否小于等于 t
            for number in nums[start..i].iter() {
                let number = *number as i64;
                // 若存在 number 使得其与 num 的差小于等于 t ，则直接返回 true
                if (num - number).abs() <= t {
                    return true;
                }
            }
        }
        // 循环内未返回，则所有数都不满足题意，直接返回 false
        return false;
    }
}

// 思路2： 桶 + 鸽笼原理
//
//		我们将所有的数都分配一个桶，桶的大小为 t + 1 ，
//      这样我们就可以根据 nums[i] 计算出其所在的桶 bucket ，
//      同时，我们再维护一个 map ，存放最近 k 个数的桶及其数，
//      这样我们可以遍历 nums 数组，
//      1. 进行如下判断：
//          (1). 若 bucket 在 map 中，则两个数的差必定在 t 以内，直接返回 true
//          (2). 若 bucket - 1 在 map 中，
//              且其对应的数 number 与 num[i] 的差在 t 以内，直接返回 true
//          (3). 若 bucket + 1 在 map 中，
//              且其对应的数 number 与 num[i] 的差在 t 以内，直接返回 true
//      2. 将 bucket 及 nums[i] 放入 map 中
//      3. 若 i >= k ，则将 nums[i - k] 对应的桶从 map 中移除
//
//      若循环内没有返回，则必定不满足题意，返回 false
//
//      时间复杂度： O(n)
//      空间复杂度： O(min(n, t))

use std::collections::HashMap;

impl Solution {
    pub fn contains_nearby_almost_duplicate(nums: Vec<i32>, k: i32, t: i32) -> bool {
        // 绝对值不可能小于 0
        if t < 0 {
            return false;
        }
        let k = k as usize;
        let t = t as i64;
        // 下标相差不超过 t ，则每个桶可以放 t + 1 个数
        let bucket_size = t + 1;
        // 记录每个桶及桶内但数
        let mut map: HashMap<i64, i64> = HashMap::new();
        for (i, num) in nums.iter().enumerate() {
            let num = *num as i64;
            // 获取 num 在桶尺寸为 t 时的桶
            let bucket = Solution::get_bucket(num, bucket_size);
            // 若该桶已存在一个数，则直接返回 true
            if map.get(&bucket).is_some() {
                return true;
            }
            // 获取该桶前一个桶中的数
            if let Some(number) = map.get(&(bucket - 1)) {
                // 若两个数相差不超过 t ，则直接返回 true
                if (num - number).abs() <= t {
                    return true;
                }
            }
            // 获取该桶后一个桶中的数
            if let Some(number) = map.get(&(bucket + 1)) {
                // 若两个数相差不超过 t ，则直接返回 true
                if (num - number).abs() <= t {
                    return true;
                }
            }

            // 将当前桶及数放入 map
            map.insert(bucket, num);

            if i >= k {
                // 将 nums[i - k] 的桶移出 map
                let bucket = Solution::get_bucket(nums[i - k] as i64, bucket_size);
                map.remove(&bucket);
            }
        }
        // 循环内未返回，则所有数都不满足题意，直接返回 false
        return false;
    }

    pub fn get_bucket(num: i64, bucket_size: i64) -> i64 {
        // 除法是向 0 取整，即： -1 / 2 = 0 ，
        // 但我们希望是向负无穷取整数，即： -1 / 2 = -1
        if num < 0 {
            (num + 1) / bucket_size - 1
        } else {
            num / bucket_size
        }
    }
}
