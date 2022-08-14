// 链接：https://leetcode.com/problems/task-scheduler-ii/
// 题意：给定一个正整数数组 tasks ，其中 tasks[i] 表示第 i 个需要完成的任务的类型。
//      同时给定一个正整数 space ，表示执行任何两个相同类型的任务之间所需的最小间隔天数。
//      在所有任务完成之前，每一天有两种选择：
//          1. 完成 tasks 中接下来的任务
//          2. 休息一天
//
//      返回完成所有任务最少需要的天数。


// 数据限制：
//  1 <= tasks.length <= 10 ^ 5
//  1 <= tasks[i] <= 10 ^ 9
//  1 <= space <= tasks.length


// 输入： tasks = [1,2,1,2,3,1], space = 3
// 输出： 9
// 解释： 一种合法的方案是：
//       第 1 天：完成任务 0
//       第 2 天：完成任务 1
//       第 3 天：休息一天
//       第 4 天：休息一天
//       第 5 天：完成任务 2
//       第 6 天：完成任务 3
//       第 7 天：休息一天
//       第 8 天：完成任务 4
//       第 9 天：完成任务 5


// 输入： tasks = [5,8,8,5], space = 2
// 输出： 6
// 解释： 一种合法的方案是：
//       第 1 天：完成任务 0
//       第 2 天：完成任务 1
//       第 3 天：休息一天
//       第 4 天：休息一天
//       第 5 天：完成任务 2
//       第 6 天：完成任务 3


// 思路： Map
//
//      如果第 x 天执行了类型为 y 的任务，
//      那么下一次至少要等到第 x + space + 1 天，才可再次执行类型为 y 的任务。
//
//      所以我们可以维护一个名为 min_start_day 第 map ，
//      min_start_day[y] 表示类型 y 的任务最小可以开始的天数。
//
//      然后我们维护当前天数 now ，表示处理完前 i 个任务所需的最小天数。
//
//      初始 now 为 0 ，表示前 0 天处理完了前 0 个任务。
//
//      再按顺序遍历 tasks 的第 i 个任务 task ，此时需要同时满足以下两个条件：
//          1. 至少要在第 now + 1 天处理任务 task
//          2. 至少要在第 min_start_day[task] 天处理任务 task
//
//      取两者的最大值作为真正执行的天数，并更新 now 。
//
//      那么下一次任务 task 至少要在第 now + space + 1 天进行处理，
//      更新 min_start_day[task] = now + space + 1 即可。
//
//      最后 now 就是执行完全部任务时，所需的最小天数，直接返回即可。
//
//
//      时间复杂度：O(n)
//          1. 需要遍历 tasks 中全部 O(n) 个任务一次
//      空间复杂度：O(n)
//          1. 需要维护 min_start_day 中全部不同任务的最小可以开始天数，
//              最差情况下有 O(n) 个


use std::collections::HashMap;


impl Solution {
    pub fn task_scheduler_ii(tasks: Vec<i32>, space: i32) -> i64 {
        let space = space as i64;
        // min_start_day[y] 表示类型 y 的任务最小可以开始的天数
        let mut min_start_day = HashMap::new();
        // now 表示处理完前 i 个任务所需的最小天数
        let mut now = 0;
        for task in tasks {
            // 现在需要处理第 i 个任务 task ，需要同时满足以下两个条件：
            //  1. 至少要在第 now + 1 天处理任务 task
            //  2. 至少要在第 min_start_day[task] 天处理任务 task
            // 取两者的最大值作为真正执行的天数，并更新 now
            now = (now + 1).max(*min_start_day.get(&task).unwrap_or(&0));
            // 则下一次任务 task 至少要在第 now + space + 1 天进行处理
            min_start_day.insert(task, now + space + 1);
        }

        // 此时 now 就是执行完全部任务时，所需的最小天数，直接返回即可
        now
    }
}
