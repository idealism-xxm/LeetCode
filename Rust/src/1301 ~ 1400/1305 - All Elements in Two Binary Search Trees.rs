// 链接：https://leetcode.com/problems/all-elements-in-two-binary-search-trees/
// 题意：给定两个二叉搜索树 root1 和 root2 ，收集它们中的所有数字，并返回升序排序后的数组。

// 数据限制：
//  每棵树的节点数在 [0, 5000] 内
//  -(10 ^ 5) <= Node.val <= 10 ^ 5

// 输入： root1 = [2,1,4], root2 = [1,0,3]
// 输出： [0,1,1,2,3,4]

// 输入： root1 = [1,null,8], root2 = [8,1]
// 输出： [1,1,8,8]


// 思路：递归 + 中序遍历
//
//      可以先用递归中序遍历按升序分别收集 root1 和 root2 的数字到 list1 和 list2 中，
//      时间复杂度为 O(n1 + n2) ，空间复杂度为 O(n1 + n2)
//
//      然后用归并操作将升序的 list1 和 list2 中合并成一个升序的 Vec ，
//      时间复杂度为 O(n1 + n2) ，空间复杂度为 O(n1 + n2)
//
//      综上： 总时间复杂度为 O(n1 + n2) ，总空间复杂度为 O(n1 + n2)
//      
//
//      时间复杂度： O(n)
//      空间复杂度： O(n)


// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//   pub val: i32,
//   pub left: Option<Rc<RefCell<TreeNode>>>,
//   pub right: Option<Rc<RefCell<TreeNode>>>,
// }

// impl TreeNode {
//   #[inline]
//   pub fn new(val: i32) -> Self {
//     TreeNode {
//       val,
//       left: None,
//       right: None
//     }
//   }
// }
use std::rc::Rc;
use std::cell::RefCell;
impl Solution {
    pub fn get_all_elements(root1: Option<Rc<RefCell<TreeNode>>>, root2: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        // 定义两个 Vec 用于收集 root1 和 root2 中的数字
        let (mut list1, mut list2) = (vec![], vec![]);
        // 中序遍历收集 BST root1 中的数字到 list1 中，保证 list1 中的数字是升序的
        Self::collect_elements_inorder(root1.as_ref(), &mut list1);
        // 中序遍历收集 BST root2 中的数字到 list2 中，保证 list2 中的数字是升序的
        Self::collect_elements_inorder(root2.as_ref(), &mut list2);
        // 合并两个有序 Vec 成一个有序的 Vec ，并返回
        Self::merge_sorted_vecs(list1, list2)
    }

    //  合并两个有序 Vec 成一个有序的 Vec
    fn merge_sorted_vecs(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
        // 建立一个新的 Vec 用于收集结果
        let mut result = vec![];
        // 初始 a 和 b 都还未遍历
        let (mut i, mut j) = (0, 0);
        // 如果 a 和 b 中都还有未遍历的数字，则继续循环处理
        while i < a.len() && j < b.len() {
            if a[i] <= b[j] {
                // 如果当前 a 中最小的数字 小于等于 b 中最小的数字，
                // 则优先收集 a[i]
                result.push(a[i]);
                i += 1;
            } else {
                // 如果当前 a 中最小的数字 大于 b 中最小的数字，
                // 则优先收集 b[j]
                result.push(b[j]);
                j += 1;
            }
        }
        // 此时 a 和 b 中至少一个已经收集完全部数字，
        // 直接把剩余的数字放入到 result 末尾即可
        // （这里为了简洁都是用了 extend ，但实际最多只有一个会继续收集数字）
        result.extend(&a[i..]);
        result.extend(&b[j..]);

        result
    }

    // 递归中序遍历 root ，保证数字按升序收集到 list 中
    fn collect_elements_inorder(root: Option<&Rc<RefCell<TreeNode>>>, list: &mut Vec<i32>) {
        // 如果当前子树还有节点，则继续递归中序遍历
        if let Some(node) = root {
            let node = node.borrow();
            // 先收集左子树的数字
            Self::collect_elements_inorder(node.left.as_ref(), list);
            // 再收集当前节点的数字
            list.push(node.val);
            // 最后收集右子树的数字
            Self::collect_elements_inorder(node.right.as_ref(), list);
        }
    }
}
