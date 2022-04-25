// 链接：https://leetcode.com/problems/peeking-iterator/
// 题意：给定一个迭代器，实现一个支持获取当前 peek 操作的迭代器。
//
//		该迭代器需要支持以下操作：
//			1. PeekingIterator(Iterator<int> iter):
//				通过给定的迭代器初始化当前迭代器实例
//			2. int next(): 返回迭代器中的下一个元素，并将指针移向下一个元素
//			3. bool hasNext(): 返回迭代器是否还有下一个元素
//			4. int peek(): 返回迭代器中的下一个元素，但不移动指针


// 数据限制：
//  1 <= nums.length <= 1000
//	1 <= nums[i] <= 1000
//	所有 next 和 peek 的调用都是合法的
//	next, hasNext, peek 的调用最多会有 1000 次


// 输入： ["PeekingIterator", "next", "peek", "next", "next", "hasNext"]
//		 [[[1, 2, 3]], [], [], [], [], []]
// 输出： [null, 1, 2, 2, 3, false]
// 解释： PeekingIterator peekingIterator = new PeekingIterator([1, 2, 3]); // [(1),2,3]
//		 peekingIterator.next();    // 返回 1, 指针移向下一个元素 [1,(2),3]
//		 peekingIterator.peek();    // 返回 2, 指针不移动 [1,(2),3].
//		 peekingIterator.next();    // 返回 2, 指针移向下一个元素 [1,2,(3)]
//		 peekingIterator.next();    // 返回 3, 指针移向下一个元素 [1,2,3]
//		 peekingIterator.hasNext(); // 返回 False


// 思路： 模拟
//
//      我们可以将迭代器 iter 保存，
//		并维护 nextVal 和 hasNextVal 两个变量，
//		其中 nextVal 表示当前迭代器中的下一个元素（存在的话），
//		hasNextVal 表示迭代器是否还有下一个元素。
//
//		然后在不同的方法中处理即可：
//			1. PeekingIterator: 通过 iter.hasNext() 初始化 nextVal 和 hasNextVal
//			2. next: 返回 nextVal ，并根据 iter.hasNext() 设置 nextVal 和 hasNextVal
//			3. peek: 返回 nextVal
//			4. hasNext: 返回 hasNextVal
//
//
//		时间复杂度： O(1)
//          1. 三个方法都只用常数次 O(1) 操作就完成
//		空间复杂度： O(1)
//          1. 只需要使用常数个额外变量


/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

 type PeekingIterator struct {
    // 原始迭代器
    iter       *Iterator
    // 下一个元素（存在的话）
    nextVal    int
    // 是否存在下一个元素
    hasNextVal bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	// 先根据 iter 创建一个存在下一个元素的迭代器实例
    peekingIterator := &PeekingIterator {
        iter:       iter,
        hasNextVal: true,
    }
	// 再调用 move() 方法，获取第一个元素
	peekingIterator.move()
    return peekingIterator
}

func (this *PeekingIterator) hasNext() bool {
	// 返回迭代器是否还有下一个元素
    return this.hasNextVal
}

func (this *PeekingIterator) next() int {
    // 先保存要返回的元素
    value := this.nextVal
	// 然后移动指针到下一个元素
	this.move()
    return value
}

func (this *PeekingIterator) peek() int {
	// 返回迭代器中下一个元素，但不移动指针
    return this.nextVal
}

// 将迭代器移动到下一个元素
func (this *PeekingIterator) move() {
    if this.iter.hasNext() {
        // 如果 iter 还有下一个元素，保存下一个元素
        this.nextVal = this.iter.next()
    } else {
		// 如果 iter 不存在下一个元素，
		// 则将 hasNextVal 设置为 false
        this.hasNextVal = false
    }
}