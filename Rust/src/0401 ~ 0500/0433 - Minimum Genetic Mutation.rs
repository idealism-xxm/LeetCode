// 链接：https://leetcode.com/problems/minimum-genetic-mutation/
// 题意：一个基因串是长度为 8 且仅含 'A', 'C', 'G', 'T' 四种字母的字符串。
//
//      现在给定两个基因串 start 和 end ，以及一个基因库 bank 。
//      一次操作能将基因串的一个字母替换成其他字母，产生一个基因库中的基因串。
//
//      求多少次操作能将 start 变为 end ？
//      如果无法将 start 变为 end ，则返回 -1 。


// 数据限制：
//  start.length == 8
//  end.length == 8
//  0 <= bank.length <= 10
//  bank[i].length == 8
//  start, end, bank[i] 仅由 'A', 'C', 'G', 'T' 四种字母组成


// 输入： start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
// 输出： 1
// 解释： "AACCGGTT" -> "AACCGGTA"

// 输入： start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
// 输出： 2
// 解释： "AACCGGTT" -> "AACCGGTA" -> "AAACGGTA"

// 输入： start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
// 输出： 3
// 解释： "AAAAACCC" -> "AAAACCCC" -> "AAACCCCC" -> "AACCCCCC"


// 思路：BFS
//
//      本题是单源最短路，而且边长都是 1 ，所以可以直接使用 BFS 搜索即可。
//
//      我们可以维护一个邻接表 adj ，遍历每个基因 bank[i] 的每个位置 j ，
//      将第 j 个字符替换为 '.' 形成通配字符串 source 。
//
//      然后将 i 放入 adj[source] 中，
//      那么 adj[source] 中的所有下标对应的基因都可以相互转换。
//
//      同时维护一个距离数组 distance ， 
//      distance[i] 表示转换到基因 bank[i] 时的所需操作数，
//      -1 表示无法转换至基因 bank[i] 。
//
//      BFS 每个基因出队时，遍历可替换的字符生成通配字符串 source ，
//      遍历 adj[source] 中所有能转换的基因下标 next ，
//      更新转化序列长度 distance[next] ，并将 next 放入队列中。
//
//      每次出队时，如果当前基因下标 cur 就是结束基因的下标 end_index ，
//      则直接返回 distance[cur] 。
//
//      最后如果 BFS 结束还没有返回，则直接返回 -1 ，表示无法转换到结束基因。
//
//      由于本题只用找到最短路径的长度即可，所以也可以使用 双向 BFS 。
//
//
//      设 n 为基因库长度， L 为基因长度。
//
//      时间复杂度：O(n * L ^ 2)
//          1. 需要计算全部 O(n) 个基因所属的邻接表，
//              每次计算时都需要遍历全部 O(L) 个可替换的字符，
//              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
//              总时间复杂度为 O(n * L ^ 2)
//          2. 全部 O(n) 个基因都会入队列一次
//          3. 全部 O(n) 个基因都会出队列一次，
//              每次出队列都需要遍历全部 O(L) 个可替换的字符，
//              每次遍历时都需要生成对应的长度为 O(L) 的通配符字符串。
//              总时间复杂度为 O(n * L ^ 2)
//      空间复杂度：O(nL) 
//          1. 需要维护开始基因的克隆中全部 O(L) 个字母
//          2. 需要维护邻接表中全部 O(nL) 个基因下标
//          3. 需要维护 distance 全部 O(n) 个状态
//          4. 需要维护队列 q 中全部 O(n) 个基因


use std::collections::{ HashMap, VecDeque };


impl Solution {
    pub fn min_mutation(start: String, end: String, mut bank: Vec<String>) -> i32 {
        // 先把开始基因放入基因库中，方便后续使用下标处理
        bank.push(start.clone());
        let start_index = bank.len() - 1;

        // 找到结束基因在基因库中的下标
        let end_index = bank.iter().position(|gene| gene == &end);
        // 如果结束基因不在基因库中，则无法转换，直接返回 -1
        if end_index.is_none() {
            return -1;
        }
        // 获取结束基因下标
        let end_index = end_index.unwrap();

        // 构建邻接表
        let mut adj = HashMap::new();
        for (i, gene) in bank.iter().enumerate() {
            // 枚举 gene 替换的字符
            for j in 0..gene.len() {
                // 将第 j 个字符替换为通配符 '.'
                let source: String = gene
                    .chars()
                    .enumerate()
                    .map(|(k, ch)| if j == k { '.' } else { ch })
                    .collect();
                // 所有能变为 source 的单词都能相互转换
                adj.entry(source).or_insert_with(Vec::new).push(i);
            }
        }

        // 队列 q 存储 BFS 下一次遍历的基因下标
        let mut q = VecDeque::new();
        // 初始只有开始基因的下标在其中
        q.push_back(start_index);
        // distance[i] 表示从 start_index 转换到 i 时的操作数，
        // 初始化为 -1 ，表示无法转换
        let mut distance = vec![-1; bank.len()];
        // 开始基因本身的无需任何操作就能得到
        distance[start_index] = 0;

        // 不断从 q 中获取下一个基因下标，直至 q 为空
        while let Some(cur) = q.pop_front() {
            // 如果当前基因下标就是结束基因的下标，则直接返回
            if cur == end_index {
                return distance[end_index];
            }

            // 枚举 cur 替换的字符
            for j in 0..bank[cur].len() {
                // 将第 j 个字符替换为通配符 '.'
                let source: String = bank[cur]
                    .chars()
                    .enumerate()
                    .map(|(k, ch)| if j == k { '.' } else { ch })
                    .collect();
                // 遍历邻接表
                for &next in adj.get(&source).unwrap() {
                    // 如果 next 还未遍历过，则更新 distance[next] ，
                    // 并将 next 放入队列 q 中
                    if distance[next] == -1 {
                        distance[next] = distance[cur] + 1;
                        q.push_back(next);
                    }
                }
            }
        }

        // 最后遍历完还没有找结束基因，则直接返回 -1
        -1
    }
}
