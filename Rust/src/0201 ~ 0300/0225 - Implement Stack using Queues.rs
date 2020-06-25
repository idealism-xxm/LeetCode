// 链接：https://leetcode.com/problems/implement-stack-using-queues/
// 题意：用队列模拟栈？

// MyStack stack = new MyStack();
//
// stack.push(1);
// stack.push(2);
// stack.top();   // returns 2
// stack.pop();   // returns 2
// stack.empty(); // returns false

// 思路： 模拟
//
//      由于队列是先进先出，无论怎么都无法换方向
//      因此我们有两种选择：
//          1. push 时候为 O(n) ， pop/top 时候为 O(1)
//          2. push 时候为 O(1) ， pop/top 时候为 O(n)
//      第一种方法可以只使用一个队列，也更容易处理
//      我们将队列维护成栈的顺序，即后进的元素在前面，先进的元素在后面，
//      所以我们在 push 的时候，将元素放入队首即可，
//      这样在 pop/top 的时候只用返回队首元素即可
//
//      时间复杂度： O(n)
//      空间复杂度： O(1)

use std::collections::VecDeque;

// 简化创建节点
#[derive(Default)]
struct MyStack {
    pub queue: VecDeque<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyStack {

    /** Initialize your data structure here. */
    fn new() -> Self {
        MyStack::default()
    }

    /** Push element x onto stack. */
    fn push(&mut self, x: i32) {
        let len = self.queue.len();
        // 先将元素推入队列
        self.queue.push_back(x);
        for i in 0..len {
            // 每次将队首元素，推入队尾部，保证队列后进元素在前
            let front = self.queue.pop_front().unwrap();
            self.queue.push_back(front);
        }
    }

    /** Removes the element on top of the stack and returns that element. */
    fn pop(&mut self) -> i32 {
        // 直接 pop 队首元素
        self.queue.pop_front().unwrap()
    }

    /** Get the top element. */
    fn top(&self) -> i32 {
        // 直接返回队首元素
        *self.queue.front().unwrap()
    }

    /** Returns whether the stack is empty. */
    fn empty(&self) -> bool {
        // 队列不为空，则栈不为空
        self.queue.is_empty()
    }
}

/**
 * Your MyStack object will be instantiated and called as such:
 * let obj = MyStack::new();
 * obj.push(x);
 * let ret_2: i32 = obj.pop();
 * let ret_3: i32 = obj.top();
 * let ret_4: bool = obj.empty();
 */
