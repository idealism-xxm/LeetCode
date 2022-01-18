// 链接：https://leetcode.com/problems/find-the-town-judge/
// 题意：一个小镇有 n 个人，现在有一个信任关系数组 trust ，
//      trust[i] = [a_i, b_i] 表示第 a_i 个人信任第 b_i 个人。
//      现在找到一个人满足以下条件：
//          1. 这个人不信任任何人
//          2. 这个人被其他所有人信任
//          3. 只有一个人会同时满足 1 和 2 两个条件
//      不存在这样的人时，返回 -1

// 数据限制：
//  1 <= n <= 1000
//  0 <= trust.length <= 10 ^ 4
//  trust[i].length == 2
//  trust 所有的数对都不同
//  a_i != b_i
//  1 <= a_i, b_i <= n

// 输入： n = 2, trust = [[1,2]]
// 输出： 2
// 解释： 第 1 个人信任 1 个人，被 0 个人信任，不满足题意
//       第 2 个人信任 0 个人，被 1 个人信任，满足题意

// 输入： n = 3, trust = [[1,3],[2,3]]
// 输出： 3
// 解释： 第 1 个人信任 1 个人，被 0 个人信任，不满足题意
//       第 2 个人信任 1 个人，被 0 个人信任，不满足题意
//       第 3 个人信任 0 个人，被 2 个人信任，满足题意

// 输入： n = 3, trust = [[1,3],[2,3],[3,1]]
// 输出： -1
// 解释： 第 1 个人信任 1 个人，被 1 个人信任，不满足题意
//       第 2 个人信任 1 个人，被 0 个人信任，不满足题意
//       第 3 个人信任 1 个人，被 2 个人信任，不满足题意


// 思路：模拟
//
//      按照题意维护两个长度为 n + 1 的数组 in_degree 和 out_degree ，
//      （因为人的下标是从 1 开始，所以要数组长度为 n + 1 ）
//          in_degree[i] 表示 i 被多少个人信任，初始均为 0
//          out_degree[i] 表示 i 信任多少个人，初始均为 0
//
//      然后我们遍历 trust 数组的每个信任关系 trust[i] = [a_i, b_i] ，
//      即： a_i 信任 b_i ，
//      所以就可以更新 in_degree[i] 和 out_degree[i] ：
//          out_degree[a_i] += 1
//          in_degree[b_i] += 1
//
//      此时我们已经统计了每个人的信任与被信任关系，我们找到第一个满足题意的人即可，
//          1. 不信任任何人 -> out_degree[a_i] == 0
//          2. 被其他人信任 -> in_degree[b_i] == n - 1
//          3. 最多只有一人满足 1 和 2 ：假设存在第二个人，
//               那么他们都不会信任对方，也就都不会被其他所有人信任，就不满足 2 了
//
//      最后还没有返回时，说明不存在这样的人，直接返回 -1 即可
//
//      注意一个人不能信任自己，所以 in_degree 和 out_degree 可以合成一个数组 degree 统计，
//      关系如下： degree[i] = in_degree[i] - out_degree[i]
//      即对于每一个信任关系 trust[i] = [a_i, b_i] 有：
//          degree[a_i] -= 1
//          degree[b_i] += 1
//      最后只有 degree[i] == n - 1 的那个人满足题意
//      
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)


impl Solution {
    pub fn find_judge(n: i32, trust: Vec<Vec<i32>>) -> i32 {
        // 如果只有 1 个人，那这个人必定就满足题意
        if n == 1 {
            return 1;
        }
        // 初始化长度为 n + 1 的 in_degree 和 out_degree ，
        //  in_degree[i] 表示 i 被多少个人信任，初始均为 0
        //  out_degree[i] 表示 i 信任多少个人，初始均为 0
        let (mut in_degree, mut out_degree) = (vec![0; (n + 1) as usize], vec![0; (n + 1) as usize]);
        // 遍历信任关系列表
        for edge in trust {
            // 信任关系如下： edge[0] 信任 edge[1] ，
            // 所以 out_degree[edge[0]] += 1
            //      in_degree[edge[1]] += 1
            out_degree[edge[0] as usize] += 1;
            in_degree[edge[1] as usize] += 1;
        }
        // 使用迭代器遍历 in_degree
        in_degree.iter()
            // 同时遍历 out_degree
            .zip(out_degree.iter())
            // 找到第一个被其他所有人信任，但不信任其他任何人的那个人的下标
            // （这样的人最多只存在一个，假设存在第二个人，
            //  那么他们都不会信任对方，也就都不会被其他所有人信任）
            .position(|(&ind, &outd)| ind == n - 1 && outd == 0)
            // 如果这个人存在就直接返回，如果不存在则默认返回 -1
            .map_or(-1, |i| i as i32)
    }
}
