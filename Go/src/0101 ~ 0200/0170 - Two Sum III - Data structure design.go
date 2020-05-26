// 链接：https://leetcode.com/problems/two-sum-iii-data-structure-design/
// 题意：设计一个 两数之和 类，支持两个操作 add 和 find ，
//		add: 向类添加一个数
//		find: 寻找内部是否存在一个数，使两数只和等于给定的数

// add(1); add(3); add(5);
// find(4) -> true
// find(7) -> false

// add(3); add(1); add(2);
// find(3) -> true
// find(6) -> false

// 思路： map
//
//		和最简单的两数只和一样，最开始用 map 统计每个数出现的次数即可，
//		然后 add 的时候就对出现次数 + 1
//		find 的时候就遍历这个 map ，查看是否有满足条件的两个数
//
//		时间复杂度： add: O(1)    find: O(n)
//		空间复杂度： O(n)
type TwoSum struct {
	numCount map[int]int
}


/** Initialize your data structure here. */
func Constructor() TwoSum {
	return TwoSum {
		numCount: make(map[int]int),
	}
}


/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int)  {
	// 标记 number 出现次数 + 1
	this.numCount[number]++
}


/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	// 遍历 map 中的每个数
	for num := range this.numCount {
		// 求得需要的另一个数的大小
		anotherNum := value - num
		// 如果两个数相等，则出现次数必须大于等于 2
		if num == anotherNum && this.numCount[anotherNum] >= 2 {
			return true
		}
		// 如果两个数相等，则出现次数必须大于等于 1
		if num != anotherNum && this.numCount[anotherNum] >= 1 {
			return true
		}
	}
	return false
}


/**
 * Your TwoSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * param_2 := obj.Find(value);
 */
