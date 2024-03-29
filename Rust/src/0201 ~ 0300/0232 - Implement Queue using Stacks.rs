// 链接：https://leetcode.com/problems/implement-queue-using-stacks/
// 题意：用两个栈模拟队列？
//
//      进阶：使用平均时间复杂度为 O(1) 的算法。


// 数据限制：
//  1 <= x <= 9
//  最多回调用 100 次 push, pop, peek, empty
//  pop 和 peek 的所有调用都是合法的


// 输入： ["MyQueue", "push", "push", "peek", "pop", "empty"]
//       [[], [1], [2], [], [], []]
// 输出： [null, null, null, 1, 1, false]
// 解释： MyQueue myQueue = new MyQueue();
//       myQueue.push(1); // 队列变为 [1]
//       myQueue.push(2); // 队列变为 [1, 2] （左侧为队首）
//       myQueue.peek();  // 返回 1
//       myQueue.pop();   // 返回 1 ， 队列变为 [2]
//       myQueue.empty(); // 返回 false


// 思路： 栈
//
//      维护两个栈，一个用于存储 push 的元素，一个用于存储 pop 的元素。
//
//      当调用 push 时，将元素往 push 栈推入即可。
//
//      当调用 pop/peek 时，如果 pop 栈为空，
//      则将 push 栈中的元素放入 pop 栈中，这样就将先进后出转换为了先进先出。
//      然后将 pop 栈中栈顶元素出栈/返回。
//
//      当调用 empty 时， push 栈和 pop 栈都为空时，队列为空。
//
//
//      时间复杂度： push/empty - O(1) | pop/peek - 平均 O(1)
//          1. push: 直接往 push 栈中推入，时间复杂度为 O(1)
//          2. empty: 直接判断 push 栈和 pop 栈是否为空，时间复杂度为 O(1)
//          3. pop/peek: pop 栈为空时，需要将 push 栈中全部 O(n) 个元素放入 pop 栈中。
//              调用 O(n) 次，只有一次需要这样处理，
//              所以最差时间复杂度为 O(n) ，平均时间复杂度为 O(1)
//      空间复杂度： O(n)
//          1. 需要维护 push 栈和 pop 栈中共 O(n) 个元素


use std::vec::Vec;


// 简化创建节点
#[derive(Default)]
struct MyQueue {
    // push 栈维护已放入的元素
    push_stack: Vec<i32>,
    // pop 栈维护待移除的元素。
    // 将 push 栈中的元素放入 pop 栈时，就将先进后出转换为了先进先出
    pop_stack: Vec<i32>,
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
        // 如果 pop 栈为空，则需要将当前 push 栈中的元素转移到 pop 栈中，
        // 这样就将先进后出转换为了先进先出
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
        // 如果 pop 栈为空，则需要将当前 push 栈中的元素转移到 pop 栈中，
        // 这样就将先进后出转换为了先进先出
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
