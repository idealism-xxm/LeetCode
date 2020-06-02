// 链接：https://leetcode.com/problems/isomorphic-strings/
// 题意：给定两个字符串，判断它们是否是同构的（两个字符不能映射为同一个字符） ？

// 输入： s = "egg", t = "add"
// 输出： true

// 输入： s = "foo", t = "bar"
// 输出： false

// 输入： s = "paper", t = "title"
// 输出： true

// 思路：map + set
//
//		用一个 map 记录 s 中每个字符映射的字符，
//		再用一个 set 记录被映射为的每个字符，
//		同时遍历 s 和 t ，
//		若 s[i] 已被映射为 ch
//			1. ch 不为 t[i] 时，直接返回 false
//		若 s[i] 未被映射为任何字符
//			1. t[i] 已被其他字符映射时，直接返回 false
//			2. t[i] 为被其他字符映射，将 s[i] 映射为 t[i] ，并标记
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

use std::collections::HashMap;
use std::collections::HashSet;

impl Solution {
	pub fn is_isomorphic(s: String, t: String) -> bool {
		// 标记 s 中每个字符映射的字符
		let mut ch_map: HashMap<char, char> = HashMap::new();
		// 存储已经被用于映射的字符
		let mut used: HashSet<char> = HashSet::new();
		// 同时遍历 s 和 t
		for (s_ch, t_ch) in s.chars().zip(t.chars()) {
			if let Some(ch) = ch_map.get(&s_ch) {
				// 此时 s_ch 已经映射为 ch
				// 如果 ch 不为 t_ch ，则直接返回 false
				if *ch != t_ch {
					return false
				}
			} else {
				// 此时 s_ch 没有映射过，则准备映射为 t_ch
				// 如果 t_ch 已经被其他字符映射过，则直接返回 false
				if used.get(&t_ch) != None {
					return false
				}
				// 将 s_ch 映射为 t_ch
				ch_map.insert(s_ch, t_ch);
				// 标记为已经被映射过
				used.insert(t_ch);
			}
		}
		true
	}
}
