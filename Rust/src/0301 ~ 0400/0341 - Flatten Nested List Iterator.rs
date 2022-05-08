// 链接：https://leetcode.com/problems/flatten-nested-list-iterator/
// 题意：给定整数的嵌套列表 nestedList ，其中每个元素的类型是 NestedInteger ，
//      NestedInteger 要么是一个整数，要么是一个 NestedInteger 的列表。
//
//      实现一个迭代器，支持以下操作：
//          1. NestedIterator(List<NestedInteger> nestedList): 
//              用 nestedList 初始化一个迭代器对象
//          2. int next(): 返回列表中的下一个数字
//          3. boolean hasNext(): 判断是否还有数字可以返回


// 数据限制：
//  1 <= nestedList.length <= 500
//  nestedList 中的整数范围在 [-(10 ^ 6), 10 ^ 6] 内


// 输入： nestedList = [[1,1],2,[1,1]]
// 输出： [1,1,2,1,1]
// 解释： 通过迭代器不断执行 next 获取数字，直至 hasNext 返回 false ，
//       获得的列表应该是 [1,1,2,1,1]

// 输入： nestedList = [1,[4,[6]]]
// 输出： [1,4,6]
// 解释： 通过迭代器不断执行 next 获取数字，直至 hasNext 返回 false ，
//       获得的列表应该是 [1,4,6]


// 思路： 栈
//
//      我们先不考虑实现迭代器，而是直接收集所有的整数到一个整数列表中，
//      那么其实只要使用 DFS 递归遍历即可。
//
//      但这样是直接使用递归，所以我们只能一次性收集全部的整数，
//      无法按需获取下一个数字，很可能浪费大量空间存储本不会访问到的数字。
//
//      所以我们可以使用栈将递归转换成迭代，这样存储了上下文信息，
//      我们就可以仅在需要时才获取下一个数字。
//
//      在迭代器中维护两个变量 stack 和 next_num ，
//      其中 stack 是存放 NestedInteger 的栈， next_num 存放下一个数字。
//
//      初始化时，我们将 nestedList 反转作为栈 stack ，并令 next_num 为空。
//      然后使用 advance_next 函数获取下一个数字。
//
//      1. 当调用 hasNext 时，只要 next_num 不为空，就返回 true
//      2. 当调用 next 时，我们先暂存 next_num ，
//          然后调用 advance_next 函数获取下一个数字，
//          最后再返回刚刚暂存的 next_num
//
//      在 advance_next 函数中，我们就是不断将栈顶的元素 top 弹出，
//      不断处理，直至栈为空：
//          1. top 是整数，则将 top 赋值给 next_num ，并返回
//          2. top 是 NestedInteger 列表，
//              则将 top 中的 NestedInteger 倒序压入栈中，继续处理
//      最后栈为空时，则不存在下一个数字，将 next_num 设置为空
//
//
//      时间复杂度：平均 O(1)
//          1. 遍历全部 O(n) 个数字，需要 O(n) 次 next 操作，
//              平均下来每次操作的时间复杂度为 O(1)
//      空间复杂度：O(n)
//          1. 需要维护一个栈 stack ，
//              最差情况下需要保存全部 O(n) 个 NestedInteger


// #[derive(Debug, PartialEq, Eq)]
// pub enum NestedInteger {
//   Int(i32),
//   List(Vec<NestedInteger>)
// }
struct NestedIterator {
    // 存放 NestedInteger 的栈
    stack: Vec<NestedInteger>,
    // 存放下一个要返回的数
    next_num: Option<i32>,
}


/** 
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl NestedIterator {

    fn new(mut nestedList: Vec<NestedInteger>) -> Self {
        // 先将 nestedList 反转，方便后续可以从栈顶开始遍历
        nestedList.reverse();
        // 定义并初始化 NestedIterator
        let mut iter = NestedIterator {
            stack: nestedList,
            next_num: None,
        };
        // 找到下一个要返回的数
        iter.advande_next();

        iter
    }
    
    fn next(&mut self) -> i32 {
        // 暂存当前要返回的数
        let next_num = self.next_num.unwrap();
        // 找到下一个要返回的数
        self.advande_next();

        next_num
    }

    fn has_next(&self) -> bool {
        // 如果有下一个要返回的数，则返回 true
        self.next_num.is_some()
    }

    fn advande_next(&mut self) {
        // 当栈不为空时，继续寻找下一个要返回的数
        while !self.stack.is_empty() {
            // 取出栈顶元素
            match self.stack.pop().unwrap() {
                NestedInteger::Int(num) => {
                    // 如果栈顶元素是整数，
                    // 则将其作为下一个要返回的数，并返回
                    self.next_num = Some(num);
                    return;
                }
                NestedInteger::List(mut list) => {
                    // 如果栈顶元素是列表，
                    // 则先将其反转，方便后续可以从栈顶开始遍历
                    list.reverse();
                    // 将 list 压入栈中
                    self.stack.extend(list);
                }
            }
        }
        // 此时栈中没有元素，则不存在下一个数，
        // 将 next_num 设置为 None
        self.next_num = None
    }
}

/**
 * Your NestedIterator object will be instantiated and called as such:
 * let obj = NestedIterator::new(nestedList);
 * let ret_1: i32 = obj.next();
 * let ret_2: bool = obj.has_next();
 */
