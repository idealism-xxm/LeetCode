// 链接：https://leetcode.com/problems/course-schedule/
// 题意：给定 n 个课程 (0 ~ n-1) ，给定一些课程的前置课程，
//		判断是否能学完所有课程？

// 输入： numCourses = 2, prerequisites = [[1,0]]
// 输出： true
// 解释： 总共有两门课，学习课程 1 之前，要先学习课程 0

// 输入： numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出： true
// 解释： 总共有两门课，学习课程 1 之前，要先学习课程 0 ；
//		 而学习课程 0 之前，又要学习课程 1 ，所以无法学完所有课程

// 思路1：拓扑排序
//
//		根据 prerequisites 建立邻接表，然后使用拓扑排序，记录入度为 0 度点的个数，
//		若最后 num_courses 点的入度都为 0 ，则返回 true ，否则返回 false
//
//		时间复杂度： O(V + E)
//		空间复杂度： O(V + E)

use std::collections::VecDeque;

impl Solution {
	pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
		// 邻接表
		let mut edges: Vec<Vec<i32>> = Vec::new();
		for _ in 0..num_courses {
			edges.push(Vec::new())
		}
		// 每个点的入度
		let mut indegree = vec![0; edges.len()];
		for prerequisite in prerequisites.iter() {
			// 添加有向边
			edges[prerequisite[1] as usize].push(prerequisite[0]);
			// 入度 + 1
			indegree[prerequisite[0] as usize] += 1;
		}

		// 入度为 0 的点的队列
		let mut queue: VecDeque<i32> = VecDeque::new();
		for i in 0..num_courses {
			// 如果当前点入度为 0 ，则放入队列中
			if indegree[i as usize] == 0 {
				queue.push_back(i);
			}
		}

		// 入度为 0 的点的个数
		let mut count = 0;
		// 如果还存在入度为 0 的点
		while let Some(cur) = queue.pop_front() {
			// 入度为 0 的点的个数 + 1
			count += 1;
			for next in edges[cur as usize].iter() {
				let next = *next;
				// 入度 - 1
				indegree[next as usize] -= 1;
				// 如果 next 的入度变为 0 ，则放入队列中
				if indegree[next as usize] == 0 {
					queue.push_back(next);
				}
			}
		}

		// 若所有点都变成了入度为 0 度点，则返回 true ，否则返回 false
		count == num_courses
	}
}


// 思路2：tarjan
//
//		题目本质是求是否存在环，所以可以使用 tarjan 算法判断有向图中是否存在环
//
//		设 vis 为每个节点的三种状态：
//			0: 未访问过
//			1: 访问中，该点在当前路径上
//			2: 已访问过，该点不在当前路径上
//
//		可以发现：
//			如果当前状态是 1 ，则存在环，直接返回 false
//			如果当前状态是 2 ，则不存在环，直接返回 true
//			如果当前状态是 0 ，则先标记为 1 ，然后递归遍历所有相邻的节点，
//				任意一个返回 false 就返回 false ，否则继续处理，
//				最后在标记为 2 ， 返回 true
//
//		时间复杂度： O(V + E)
//		空间复杂度： O(V + E)

impl Solution {
	pub fn can_finish(num_courses: i32, prerequisites: Vec<Vec<i32>>) -> bool {
		// 邻接表
		let mut edges: Vec<Vec<i32>> = Vec::new();
		for _ in 0..num_courses {
			edges.push(Vec::new())
		}
		for prerequisite in prerequisites.iter() {
			// 添加有向边
			edges[prerequisite[0] as usize].push(prerequisite[1]);
		}

		// 标记每个点被访问的状态（ 0: 未访问过 1: 当前路径上的点 2: 非当前路径上的点）
		let mut vis = vec![0; num_courses as usize];
		for i in 0..num_courses {
			// 若当前点开始点路径存在环，则直接返回 false
			if !Solution::dfs(i, &edges, &mut vis) {
				return false;
			}
		}
		// 所有节点都访问过，且不存在环
		true
	}

	fn dfs(cur: i32, edges: &Vec<Vec<i32>>, vis: &mut Vec<i32>) -> bool {
		let cur = cur as usize;
		// 如果当前点在当前路径上，则存在环
		if vis[cur] == 1 {
			return false;
		}
		// 如果当前点不在当前路径上，则不存在环
		if vis[cur] == 2 {
			return true;
		}
		// 当前点未访问过，标记其为当前路径上点点
		vis[cur] = 1;
		// 递归访问领接表中的点
		for next in edges[cur].iter() {
			// 如果接下来存在环，则直接返回 false
			if !Solution::dfs(*next, edges, vis) {
				return false;
			}
		}
		// 标记 cur 不在后面的路径上
		vis[cur] = 2;
		// 当前路径上不存在环
		true
	}
}
