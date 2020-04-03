// 链接：https://leetcode.com/problems/lru-cache/
// 题意：实现一个 LRU cache ，保证 get 和 set 操作的时间复杂度为 O(1) ？

// LRUCache cache = new LRUCache( 2 /* capacity */ );
//
// cache.put(1, 1);
// cache.put(2, 2);
// cache.get(1);       // returns 1
// cache.put(3, 3);    // evicts key 2
// cache.get(2);       // returns -1 (not found)
// cache.put(4, 4);    // evicts key 1
// cache.get(1);       // returns -1 (not found)
// cache.get(3);       // returns 3
// cache.get(4);       // returns 4

// 思路： map + 双向链表
//
//		因为要保证 get 和 set 时间复杂度均为 O(1) ，
//		又需要保证有序，所以要用链表存储 value
//
//		为了找到每个 key 对应的结点，需要使用 map 进行存储对应关系
//
//		时间复杂度： O(1)
//		空间复杂度： O(n)

type node struct {
	Key int
	Val int
	Pre *node
	Next *node
}

type LRUCache struct {
	keyToNode map[int]*node
	head *node
	tail *node
	capacity int
}


func Constructor(capacity int) LRUCache {
	head, tail := new(node), new(node)
	head.Next, tail.Pre = tail, head
	return LRUCache{
		keyToNode: make(map[int]*node, capacity),
		head:      head,
		tail:      tail,
		capacity:  capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	if valueNode, exists := this.keyToNode[key]; exists {
		// 把该结点从链表中移除
		valueNode.Pre.Next, valueNode.Next.Pre = valueNode.Next, valueNode.Pre
		// 把该结点从插入链表头
		valueNode.Pre, valueNode.Next = this.head, this.head.Next
		this.head.Next, this.head.Next.Pre = valueNode, valueNode

		return valueNode.Val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	var valueNode *node
	if cur, exists := this.keyToNode[key]; exists {
		// 如果存在，则先把该结点从链表中移除
		cur.Pre.Next, cur.Next.Pre = cur.Next, cur.Pre

		valueNode = cur
	} else {
		// 如果已经达到容量，则删除最后一个结点
		if len(this.keyToNode) == this.capacity {
			last := this.tail.Pre
			last.Pre.Next, last.Next.Pre = last.Next, last.Pre
			delete(this.keyToNode, last.Key)
		}

		// 结点不存在，则直接创建新的
		valueNode = new(node)
		this.keyToNode[key] = valueNode
	}
	// 赋值
	valueNode.Key, valueNode.Val = key, value
	// 插入到链表头
	valueNode.Pre, valueNode.Next = this.head, this.head.Next
	this.head.Next, this.head.Next.Pre = valueNode, valueNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
