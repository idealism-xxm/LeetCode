// 链接：https://leetcode.com/problems/sort-the-people/
// 题意：给定一个字符串数组 names 和一个不含重复数字的整数数组 heights ，
//      长度均为 n 。
//      其中 names[i] 和 heights[i] 分别表示第 i 个人的姓名和身高。
//      将 names 按照每个人的身高降序排序，然后返回排序后的 names 。


// 数据限制：
//  n == names.length == heights.length
//  1 <= n <= 10 ^ 3
//  1 <= names[i].length <= 20
//  1 <= heights[i] <= 10 ^ 5
//  names[i] 仅由英文字母组成
//  heights 中所有的值各不相同


// 输入： names = ["Mary","John","Emma"], heights = [180,165,170]
// 输出： ["Mary","Emma","John"]
// 解释： Mary 最高，然后是 Emma 和 John

// 输入： names = ["Alice","Bob","Bob"], heights = [155,185,150]
// 输出： ["Bob","Alice","Bob"]
// 解释： 第一个 Bob 最高，然后是 Alice 和第二个 Bob


// 思路： 排序
//
//      由于姓名和身高是一个人的，算作一个整体，很容易就能想到重新定义一个类型 People 。
//
//      将 names 和 heights 整合成 peoples ，然后对 peoples 按照身高降序排序，
//      再按顺序遍历 peoples 只取出姓名即可。
//
//
//      时间复杂度：O(nlogn)
//          1. 需要遍历 names, heights, peoples 中全部 O(n) 个元素各一次
//          2. 需要对 peoples 中全部 O(n) 个元素排序，时间复杂度为 O(nlogn)
//      空间复杂度：O(n)
//          1. 需要存储 peoples 全部 O(n) 个元素


func sortPeople(names []string, heights []int) []string {
    // 将 names 和 heights 整合成 peoples
    peoples := make([]*People, len(names))
    for i, name := range names {
        peoples[i] = &People{ name, heights[i] }
    }
    
    // peoples 按身高降序排序
    sort.Slice(peoples, func(i, j int) bool { return peoples[i].height > peoples[j].height })

    // 按顺序遍历 peoples ，只取出姓名
    for i, people := range peoples {
        names[i] = people.name
    }
    return names
}

type People struct {
    name string
    height int
}
