// 链接：https://leetcode.com/problems/implement-trie-prefix-tree/
// 题意：实现 Trie ？

// Trie trie = new Trie();
//
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true

// 思路：模拟
//
//		创建一个 reverse(pre, cur) 函数，用于递归处理
//		1. cur 是空节点， pre 就是新的头节点，直接返回 pre
//		2. cur 不是头节点，获取 next = cur.next ，
//			将 cur.next 置为 pre ，然后返回 reverse(cur, next)
//
//		时间复杂度： O(n)
//		空间复杂度： O(n)

use std::boxed::Box;
use std::option::Option;

// 简化创建节点
#[derive(Default)]
struct TrieNode {
	next: [Option<Box<TrieNode>>; 26],
	is_end: bool,
}

// 简化创建 Trie
#[derive(Default)]
struct Trie {
	root: TrieNode
}


/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl Trie {
	/** Initialize your data structure here. */
	fn new() -> Self {
		Self::default()
	}

	/** Inserts a word into the trie. */
	fn insert(&mut self, word: String) {
		let mut cur = &mut self.root;
		// 遍历 word 每一个字符
		for ch in word.bytes() {
			// 计算 next 数组中的下标
			let i = (ch - b'a') as usize;
			// 找到该字符对应的节点
			let mut next = &mut cur.next[i];
			// 若 next 不存在，则创建新的节点。
			// 返回目前 next 包含的节点的可变引用
			cur = next.get_or_insert_with(Box::<TrieNode>::default);
		}
		// 标记最后一个字符对应的节点为单词的结束节点
		cur.is_end = true;
	}

	/** Returns if the word is in the trie. */
	fn search(&self, word: String) -> bool {
		return self.get_trie_node(word).map_or(false, |cur| cur.is_end)
	}

	/** Returns if there is any word in the trie that starts with the given prefix. */
	fn starts_with(&self, prefix: String) -> bool {
		return self.get_trie_node(prefix).is_some()
	}

	fn get_trie_node(&self, prefix: String) -> Option<&TrieNode> {
		let mut cur = &self.root;
		// 遍历 word 每一个字符
		for ch in prefix.bytes() {
			// 计算 next 数组中的下标
			let i = (ch - b'a') as usize;
			match cur.next[i].as_ref() {
				// 存在则继续遍历下一个字符
				Some(next) => cur = next,
				// 如果不存在，则直接返回 false
				None => return None,
			}
		}
		return Some(cur)
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * let obj = Trie::new();
 * obj.insert(word);
 * let ret_2: bool = obj.search(word);
 * let ret_3: bool = obj.starts_with(prefix);
 */