// 链接：https://leetcode.com/problems/course-schedule-ii/
// 题意：给定 n 个课程 (0 ~ n-1) ，给定一些课程的前置课程，
//		给定一个合法的课程学习顺序，不存在则会返回空数组？

// 输入： 2, [[1,0]]
// 输出： [0, 1]
// 解释： 总共有两门课，学习课程 1 之前，要先学习课程 0 ，
//		 因此学习顺序为 [0, 1]

// 输入： 4, [[1,0],[2,0],[3,1],[3,2]]
// 输出： [0,1,2,3] 或 [0,2,1,3]
// 解释： 总共有四门课，学习课程 3 之前，要先学习课程 1 和 2 ；
//		 而学习课程 1 和 2 之前，又要学习课程 0 ，
//		 因此学习顺序为 [0,1,2,3] 或 [0,2,1,3]

// 思路：拓扑排序
//
//		0207 的加强版，思路一致，改一下就行
//
//		根据 prerequisites 建立邻接表，然后使用拓扑排序，
//		将所有入度为 0 度点放入结果列表中，
//		若最后 num_courses 点都在结果列表，则直接返回，否则返回空列表
//
//		当然也可以不用统计入度，用 dfs 的方法
//
//		时间复杂度： O(V + E)
//		空间复杂度： O(V + E)

impl Solution {
	pub fn find_order(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> Vec<i32> {
		// 邻接表
		let mut edges: Vec<Vec<i32>> = Vec::new();
		for _ in 0..num_courses {
			edges.push(Vec::new())
		}
		// 每个课程的入度
		let mut indegree = vec![0; edges.len()];
		for prerequisite in prerequisites.iter() {
			// 添加有向边
			edges[prerequisite[1] as usize].push(prerequisite[0]);
			// 入度 + 1
			indegree[prerequisite[0] as usize] += 1;
		}

		// 结果列表
		let mut result: Vec<i32> = Vec::new();
		for i in 0..num_courses {
			// 如果当前课程入度为 0 ，则可以进行学习，放入结果列表中
			if indegree[i as usize] == 0 {
				result.push(i);
			}
		}

		// 入度为 0 的课程的个数
		let mut count = 0;
		// 如果还存在未学过的课程
		let mut i: usize = 0;
		while i < result.len() {
			// 学习课程 i
			let cur = result[i];
			i += 1;
			for next in edges[cur as usize].iter() {
				let next = *next;
				// 依赖课程 i 的课程入度 - 1
				indegree[next as usize] -= 1;
				// 如果 next 的入度变为 0 ，则可以进行学习，放入结果列表中
				if indegree[next as usize] == 0 {
					result.push(next);
				}
			}
		}

		// 若所有点都变成了入度为 0 度点，则返回 result ，否则返回 空列表
		if result.len() == num_courses as usize {
			result
		} else {
			Vec::new()
		}
	}
}
