// 链接：https://leetcode.com/problems/different-ways-to-add-parentheses/
// 题意：给定一个只包含加减乘的算数表达式，求添加括号后所有可能的结果？

// 输入： "2-1-1"
// 输出： [0, 2]
// 解释：
// ((2-1)-1) = 0
// (2-(1-1)) = 2

// 输入： "2*3-4*5"
// 输出： [-34, -14, -10, -10, 10]
// 解释：
// (2*(3-(4*5))) = -34
// ((2*3)-(4*5)) = -14
// ((2*(3-4))*5) = -10
// (2*((3-4)*5)) = -10
// (((2*3)-4)*5) = 10

// 思路： 递归
//
//      我们用一个 map 统计每一层中输入对应的结果，
//      若以前计算过，则直接返回，否则计算一遍并记忆化，
//
//      递归每一层遍历当前输入，当遇到运算符时，递归处理左右部分，
//      然后分别遍历这两部分，计算他们组合起来可能产生的结果并收集为当前输入对应的结果

use std::collections::HashMap;
use std::vec::Vec;

impl Solution {
    pub fn diff_ways_to_compute(input: String) -> Vec<i32> {
        let map: HashMap<String, Vec<i32>> = HashMap::new();
        // 递归处理
        let (_, result) = Solution::dfs(input, map).clone();
        result
    }

    fn dfs(input: String, map: HashMap<String, Vec<i32>>) -> (HashMap<String, Vec<i32>>, Vec<i32>) {
        // 记忆化，如果以前计算过，则直接返回
        if let Some(result) = map.get(&input) {
            let result = result.to_vec();
            return (map, result);
        }

        let mut map = map;
        // 当前输入对应的结果
        let mut result: Vec<i32> = Vec::new();
        // 遍历输入，找到所有运算符，然后递归处理
        for (i, ch) in input.chars().enumerate() {
            // 如果不是运算符，则继续处理下一个
            if ch != '+' &&  ch != '-' &&  ch != '*' {
                continue;
            }

            // 递归计算左右部分可能产生的结果
            let (returned_map, left_result) = Solution::dfs(input[..i].to_string(), map);
            let (returned_map, right_result) = Solution::dfs(input[i+1..].to_string(), returned_map);
            // 赋值给原有的 map ，供下次循环使用
            map = returned_map;
            // 组装当前输入对应的结果
            for left_num in left_result.iter() {
                for right_num in right_result.iter() {
                    result.push(match ch {
                        '+' => left_num + right_num,
                        '-' => left_num - right_num,
                        '*' => left_num * right_num,
                        _ => unreachable!(),
                    });
                }
            }
        }
        // 如果没有任何结果，则当前输入是一个数字
        if result.is_empty() {
            result.push(input.parse().unwrap())
        }

        // 放入 map 记忆化
        let result_copy = result.to_vec();
        map.insert(input, result_copy);
        (map, result)
    }
}
