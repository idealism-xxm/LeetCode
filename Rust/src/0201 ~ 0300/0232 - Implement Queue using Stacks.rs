// 链接：https://leetcode.com/problems/implement-queue-using-stacks/
// 题意：用栈模拟队列？

// MyQueue queue = new MyQueue();
//
// queue.push(1);
// queue.push(2);
// queue.peek();  // returns 1
// queue.pop();   // returns 1
// queue.empty(); // returns false

// 思路： push 栈 + pop 栈
//
//      维护两个栈，一个用于存储 push 的元素，一个用于存储 pop 的元素，
//      当进行 push 时，往 push 栈推入，
//      当进行 pop/peek 时，如果 pop 栈为空，则将 push 栈中的元素放入 pop 栈中，
//          然后将 pop 栈中栈顶元素出栈/返回
//
//      时间复杂度： 平均 O(1) / 最差 O(n)
//      空间复杂度： O(1)

use std::vec::Vec;

// 简化创建节点
#[derive(Default)]
struct MyQueue {
    pub push_stack: Vec<i32>,
    pub pop_stack: Vec<i32>,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyQueue {

    /** Initialize your data structure here. */
    fn new() -> Self {
        MyQueue::default()
    }

    /** Push element x to the back of queue. */
    fn push(&mut self, x: i32) {
        self.push_stack.push(x);
    }

    /** Removes the element from in front of queue and returns that element. */
    fn pop(&mut self) -> i32 {
        // 如果 pop 栈为空，则需要将当前 push 栈中的数组转移到 pop 栈中
        if self.pop_stack.is_empty() {
            while !self.push_stack.is_empty() {
                self.pop_stack.push(self.push_stack.pop().unwrap());
            }
        }
        // pop 栈顶元素出栈
        return self.pop_stack.pop().unwrap()
    }

    /** Get the front element. */
    fn peek(&mut self) -> i32 {
        // 如果 pop 栈为空，则需要将当前 push 栈中的数组转移到 pop 栈中
        if self.pop_stack.is_empty() {
            while !self.push_stack.is_empty() {
                self.pop_stack.push(self.push_stack.pop().unwrap());
            }
        }
        // 返回 pop 栈顶元素
        return *self.pop_stack.last().unwrap();
    }

    /** Returns whether the queue is empty. */
    fn empty(&self) -> bool {
        // 当两个栈都为空是，队列才为空
        return self.push_stack.is_empty() && self.pop_stack.is_empty();
    }
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * let obj = MyQueue::new();
 * obj.push(x);
 * let ret_2: i32 = obj.pop();
 * let ret_3: i32 = obj.peek();
 * let ret_4: bool = obj.empty();
 */
