// 链接：https://leetcode.com/problems/search-in-rotated-sorted-array/
// 题意：有一个不含重复数字的递增数组，将后面一的部分（不清楚有多少数）放到前面，在 O(logn) 内找到指定的数

// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4

// 输入：nums = [4,5,6,7,0,1,2], target = 3
// 输出：-1

// 思路1：两次二分
//  一看时间复杂度就知道是二分，但中午睡了一觉后忘记了不含重复数字，导致下午无论怎么想都不觉得 O(n) 内可以搞定
//  最后一看大家的想法，明明思路一样，怎么就能过呢？再仔细一读题，才发现忘记了重要条件。。。
//
//  这道题由于数字不重复，所以很容易就能想到：先通过二分找到起点（最小值），再在相应的范围内进行二分搜索
//  1. 二分 [l, r] 内的数，每次找到中点 mid = (l + r) >> 1
//      (1) nums[l] > nums[mid] ：则起点（最小值）在 左边区间，r = mid
//      (2) nums[mid] > nums[r] ：则起点（最小值）在 右边区间，l = mid + 1
//      (3) 上面两个都不满足，则有：nums[l] <= nums[mid] <= nums[r] ，即数组一定是递增的，l 就是起点（最小值）， r = l
//  2. 找到起点（最小值），然后二分即可
//  时间复杂度：O(logn)，空间复杂度：O(1)

func search(nums []int, target int) int {
    length := len(nums)
    if length == 0 {
        return -1
    }
    
    l, r := 0, length - 1
    // 二分找到起点（前闭后闭）
    for l < r {
        mid := (l + r) >> 1
        if nums[l] > nums[mid] { // 起点在 左边区间
            r = mid
        } else if nums[mid] > nums[r] { // 起点在 右边区间
            l = mid + 1
        } else { // 此时必有：nums[l] <= nums[mid] <= nums[r] ，即该部分时递增的，l 就是起点
            r = l
        }
    }
    // 起点（最小值）下标
    start := l
    if nums[start] <= target && target <= nums[length - 1] { // 如果目标数在右边区间数值范围内，则对右边进行二分
        return binarySearch(nums, start, length, target)
    } else { // 否则，对左边进行二分
        return binarySearch(nums, 0, start, target)
    }
}

// 二分搜索，查找 target 的下标（前闭后开）
func binarySearch(nums []int, start, end int, target int) int {
    l, r := start, end
    for l < r {
        mid := (l + r) >> 1
        if nums[mid] < target { // 目标值在右边区间
            l = mid + 1
        } else if nums[mid] > target { // 目标值在左边区间
            r = mid
        } else { // 数组不重复，相等，则是答案
            return mid
        }
    }
    if r == start || l == end { // target 比 nums 内所有的数 都小 || 都大
        return -1
    }
    if nums[l] == target { // 找到的值是目标值
        return l
    }
    return -1 // 找到的值不是目标值，则不存在
}

// 思路2：一次二分
//  看到还有人说不用找到起点（最小值）即可找到答案，即只用一次二分
//  一想就是在第一次二分就让目标数参与进来即可
//  所以可以在每次确定起点（最小值）的区间后后，判断 目标数的相应区间
//  1. 起点（最小值）在左边区间（包括起点为区间开始点）：
//      (1) 目标值在右边区间数的范围内：l = mid + 1 （因为右边区间的数是连续的，所以很好比较）
//      (2) 否则认为目标数在左边区间：  r = mid - 1
//  2. 起点（最小值）在右边区间：
//      (1) 目标值在左边区间数的范围内：r = mid - 1 （因为左边区间的数是连续的，所以很好比较）
//      (2) 否则认为目标数在左边区间：  l = mid + 1
//  时间复杂度：O(logn)，空间复杂度：O(1)

func search(nums []int, target int) int {
    length := len(nums)
    if length == 0 {
        return -1
    }

    l, r := 0, length - 1
    // 二分找到目标数（前闭后闭）
    for l <= r {
        mid := (l + r) >> 1
        if nums[mid] == target {
            return mid
        }

        if nums[mid] > nums[r] { // 起点在 右边区间
            if nums[l] <= target && target < nums[mid] { // 目标数在 左边区间
                r = mid - 1
            } else { // 目标数在 右边区间
                l = mid + 1
            }
        } else { // 起点在 左边区间（包括 起点是当前区间最左边的情况）
            if nums[mid] < target && target <= nums[r] { // 目标数在 右边区间
                l = mid + 1
            } else { // 目标数在 左边区间
                r = mid - 1
            }
        }
    }
    return -1
}