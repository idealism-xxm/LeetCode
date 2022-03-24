// 链接：https://leetcode.com/problems/boats-to-save-people/
// 题意：给定一个整型数组 people ， people[i] 表示第 i 个人的体重。
//      现在有不限量的船，每艘船可以承载的最大重量为 limit ，
//      并且最多承载 2 人，求最少要多少船才能将所有人都带上？


// 数据限制：
//  1 <= people.length <= 5 * 10 ^ 4
//  1 <= people[i] <= limit <= 3 * 10 ^ 4


// 输入： people = [1,2], limit = 3
// 输出： 1
// 解释： 第 1 艘船承载体重为 (1, 2) 的 2 个人

// 输入： people = [3,2,2,1], limit = 3
// 输出： 3
// 解释： 第 1 艘船承载体重为 (1, 2) 的 2 个人
//       第 2 艘船承载体重为 (2) 的 1 个人
//       第 3 艘船承载体重为 (3) 的 1 个人

// 输入： people = [3,5,3,4], limit = 5
// 输出： 4
// 解释： 第 1 艘船承载体重为 (3) 的 1 个人
//       第 2 艘船承载体重为 (5) 的 1 个人
//       第 3 艘船承载体重为 (3) 的 1 个人
//       第 4 艘船承载体重为 (4) 的 1 个人


// 思路： 贪心 + 双指针
//
//      我们可以先对 people 按照体重升序排序，
//      然后用双指针贪心地求解。
//
//      我们维护左指针 l 表示下一个需要承载的体重最小的人的下标，
//      并维护右指针 r 表示下一个需要承载的体重最大的人的下标。
//
//      那么当还有人需要承载时，我们优先让体重最大的 r 这个人上船，
//      如果此时还能承载体重最小的 l 这个人，那么就能承载这两个人，
//      否则就只能承载 r 这个人。
//
//      
//      时间复杂度：O(nlogn)
//          1. 需要对 people 进行排序，时间复杂度为 O(nlogn)
//          2. 需要遍历 people 中全部 O(n) 个元素
//      空间复杂度：O(n)
//          1. 需要存储 people 排序后的全部 O(n) 个元素


impl Solution {
    pub fn num_rescue_boats(mut people: Vec<i32>, limit: i32) -> i32 {
        // 先对 people 按照体重升序排序
        people.sort();
        // ans 表示当前所需要的船数
        let mut ans = 0;
        // 左指针 l 表示下一个需要承载的体重最小的人的下标
        let mut l = 0;
        // 右指针 r 表示下一个需要承载的体重最大的人的下标
        let mut r = (people.len() - 1) as i32;
        // 如果还有人需要承载，则继续循环处理
        while l <= r {
            // 此时我们需要一条船承载体重最重的那个人
            ans += 1;
            // 如果能同时把最轻的那个人也带上，
            // 那么我们就贪心地把他也带上
            if people[l as usize] + people[r as usize] <= limit {
                // 承载第 l 个人，后续要继续承载第 l + 1 个人
                l += 1;
            }
            // 必定会承载第 r 个人
            r -= 1;
        }

        // 最后 ans 就是总共需要的船数
        ans
    }
}
