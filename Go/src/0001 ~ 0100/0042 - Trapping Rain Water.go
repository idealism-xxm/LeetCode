// 链接：https://leetcode.com/problems/trapping-rain-water/
// 题意：给定长度为 n 的非负整数数组，表示每一处的海拔高度，
//      并且每一处的宽度都为 1 ，计算下雨后能接住多少雨水？

// 输入：[0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6

// 思路1：枚举
//      很容易就可以看出若一处能够接住雨水，则左右两边最高处必定比其海拔高
//      所以，我们只需要先遍历一次数组，记录当前一处 i 左边最高的海拔 leftMaxHeight[i] 和右边最高的海拔 rightMaxHeight[i]
//      再遍历一次，当前一处 i 能接住的雨水 = min(leftMaxHeight[i], rightMaxHeight[i]) - height[i] （最小值必须比当前海拔高）
//      时间复杂度： O(n) ，空间复杂度： O(n)

func trap(height []int) int {
    length := len(height)
    leftMaxHeight, rightMaxHeight := make([]int, length, length), make([]int, length, length)
    for i, j := 1, length - 2; i < length; i, j = i + 1, j - 1 {
        leftMaxHeight[i] = max(leftMaxHeight[i - 1], height[i - 1])  // 记录当前一处 i 左边最高的海拔
        rightMaxHeight[j] = max(rightMaxHeight[j + 1], height[j + 1])  // 记录当前一处 j 右边最高的海拔
    }
    result := 0
    for i := 1; i < length; i++ {
        minMaxHeight := min(leftMaxHeight[i], rightMaxHeight[i])  // 两侧最高海拔的较小值
        if minMaxHeight > height[i] {  // 若两侧最高海拔的较小值比当前海拔高，则可以接住雨水
            result += minMaxHeight - height[i]
        }
    }
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// 思路2：双指针
//      看了题解发现还可以用双指针解答，就又继续用这样的思路思考
//      从思路 1 中可以发现我们只关心当前一处处两侧最高海拔的最小值
//      所以我们可以从区间 [0, len(height) - 1] 开始慢慢收缩区间，并不断更新能够接住的雨水
//      leftMaxHeight, rightMaxHeight 分别表示区间 [l, r] 左右两侧最高的海拔
//      1. 若 leftMaxHeight <= rightMaxHeight
//          则我们应该收缩左侧，因为 l 处能够接住的雨水取决于 leftMaxHeight ，这两个都是已知的
//          （若我们收缩右侧，则 r 处左侧可能存在更高的海拔，使得 r 处能够接住更多的雨水）
//      2. 同理：leftMaxHeight > rightMaxHeight
//          则我们应该收缩右侧，因为 r 处能够接住的雨水取决于 rightMaxHeight ，这两个都是已知的
//          （若我们收缩左侧，则 l 处右侧可能存在更高的海拔，使得 l 处能够接住更多的雨水）
//      时间复杂度： O(n) ，空间复杂度： O(1)

func trap(height []int) int {
    length := len(height)
    leftMaxHeight, rightMaxHeight := 0, 0  // 分别表示区间 [l, r] 左右两侧最高的海拔
    result := 0
    for l, r := 0, length - 1; l <= r; {
        if leftMaxHeight <= rightMaxHeight {  // 若左侧最高海拔不大于右侧最高海拔，则统计 l 处接住的雨水，更新 leftMaxHeight 并移动 l
            if height[l] < leftMaxHeight {  // height[l] < leftMaxHeight ，则可接住雨水
                result += leftMaxHeight - height[l]
            } else {  // height[l] >= leftMaxHeight ，则可更新 leftMaxHeight
                leftMaxHeight = height[l]
            }
            l ++  // 移动 l
        } else {
            if height[r] < rightMaxHeight {  // height[r] < rightMaxHeight ，则可接住雨水
                result += rightMaxHeight - height[r]
            } else {  // height[r] >= rightMaxHeight ，则可更新 rightMaxHeight
                rightMaxHeight = height[r]
            }
            r --
        }
    }
    return result
}
