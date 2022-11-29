// 链接：https://leetcode.com/problems/insert-delete-getrandom-o1/
// 题意：实现一个数据结构 RandomizedSet ，支持在一个不重复集合中插入、删除数，
//      以及随机返回集合中的一个数，需要实现以下操作:
//          1. RandomizedSet(): 初始化一个 RandomizedSet 对象
//          2. bool insert(int val): 往集合中插入 val 。
//              如果 val 原本不在集合中，则返回 true ；否则返回 false 。
//          3. bool remove(int val): 从集合中删除 val 。
//              如果 val 原本在集合中，则返回 true ；否则返回 false 。
//          4. int getRandom(): 返回集合中随机一个数。（调用时保证集合中至少存在一个数）
//              需要保证每个数字被返回的可能性都相等。
//
//      进阶：集合的三个方法的平均时间复杂度都为 O(1) 。


// 数据限制：
//  -(2 ^ 31) <= val <= 2 ^ 31 - 1
//  最多调用 insert, remove, getRandom 共 2 * 10 ^ 5 次
//  调用 getRandom 时保证集合中至少存在一个数


// 输入： ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
//       [[], [1], [2], [2], [], [1], [2], []]
// 输出： [null, true, false, true, 2, true, false, 2]
// 解释： RandomizedSet randomizedSet = new RandomizedSet();
//       randomizedSet.insert(1);   // 往集合中插入 1 。 1 原本不在集合中，返回 true 。集合变为 {1}
//       randomizedSet.remove(2);   // 从集合中移除 2 。 2 原本不在集合中，返回 false 。集合变为 {1}
//       randomizedSet.insert(2);   // 往集合中插入 2 。 2 原本不在集合中，返回 true 。集合变为 {1,2}
//       randomizedSet.getRandom(); // 集合为 {1,2} ，应该随机返回 1 或 2
//       randomizedSet.remove(1);   // 从集合中移除 1 。 1 原本在集合中，返回 true 。集合变为 {2}
//       randomizedSet.insert(2);   // 往集合中插入 2 。 2 原本在集合中，返回 false 。集合变为 {2}
//       randomizedSet.getRandom(); // 集合为 {2} ，应该随机返回 2


// 思路： Map
//
//      如果数据结构题需要用到 O(1) 的查询操作，一般都要使用 map 来辅助。
//
//      我们可以用数组 nums 维护集合中的数。
//
//      注意到使用类似 random 的库函数，
//      能在 O(1) 内随机获得范围 [0, nums.len()) 中的一个数 i ，
//      那么也就能在 O(1) 内随机返回集合中的数 nums[i] 。
//
//      同时我们也需要在 O(1) 内从 nums 中删除一个数 val ，
//      那么就需要知道 val 在 nums 中的下标，这就需要使用 map 来维护这个关系。
//
//      删除 val 时，如果 val 在 map 中，则获取对应下标 index ，
//      将最后一个数移动至 index （注意更新下标）。
//      并从 map 中删除 val ，再删除 nums 中最后一个数。
//
//      插入 val 时，如果 val 不在 map 中，
//      则将 val 放入 nums 末尾，并将其下标插入 map 中。
// 
//
//      时间复杂度： O(1)
//          1. 集合的三个方法都只需要执行常数次操作
//      空间复杂度： O(n)
//          1. 需要维护全部不同的数字，最差情况下全部 O(n) 个数字都不同


type RandomizedSet struct {
    // nums 维护集合中的全部数字，用于 O(1) 随机返回一个数字
    nums []int
    // numToIndex[num] 表示 num 在 nums 中的下标，用于 O(1) 插入/删除一个数字
    numToIndex map[int]int
}


func Constructor() RandomizedSet {
    return RandomizedSet { numToIndex: make(map[int]int) }
}


func (this *RandomizedSet) Insert(val int) bool {
    // 如果 val 原本在集合中，则直接返回 false
    if _, exists := this.numToIndex[val]; exists {
        return false
    }
    
    // 否则将 val 放入 nums 中，并记录其在 nums 中的下标
    this.numToIndex[val] = len(this.numToIndex)
    this.nums = append(this.nums, val)
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    // 如果 val 原本不在集合中，则直接返回 false
    if _, exists := this.numToIndex[val]; !exists {
        return false
    }

    index := this.numToIndex[val]
    lastIndex := len(this.nums) - 1
    // 将最后一个数移动至 index 处，并更新其下标
    this.numToIndex[this.nums[lastIndex]] = index
    this.nums[index] = this.nums[lastIndex]

    // O(1) 移除最后一个数字（即 val ）
    delete(this.numToIndex, val)
    this.nums = this.nums[:len(this.nums)-1]

    return true
}


func (this *RandomizedSet) GetRandom() int {
    // 随机返回 nums 中的一个数
    return this.nums[rand.Int31n(int32(len(this.nums)))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
