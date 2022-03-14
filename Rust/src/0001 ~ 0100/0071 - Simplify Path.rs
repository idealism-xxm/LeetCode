// 链接：https://leetcode.com/problems/simplify-path/
// 题意：给定一个文件的绝对路径，将其简化为规范路径。
//
//      规范路径遵循下述格式：
//          1. 始终以斜杠 '/' 开头
//          2. 两个目录名之间必须只有一个斜杠 '/'
//          3. 最后一个目录名（如果存在）不能 以 '/' 结尾
//          4. 路径仅包含从根目录到目标文件或目录的路径上的目录
//              （即不含 '.' 或 '..' ）


// 数据限制：
//  1 <= path.length <= 3000
//  路径仅有英文字母、数字、 '.' 、 '/' 和 '-' 组成
//  路径是一个合法的 Unix 绝对路径


// 输入： path = "/home/"
// 输出： "/home"
// 解释： 注意，最后一个目录名后面没有 '/'

// 输入： path = "/../"
// 输出： "/"
// 解释： 从根目录向上一级是不可行的，因为根目录是可以到达的最高级

// 输入： path = "/home//foo/"
// 输出： "/home/foo"
// 解释： 在规范路径中，多个连续 '/' 需要用一个 '/' 替换

// 输入： path = "/a/./b/../../c/"
// 输出： "/c"


// 思路： 模拟
//
//      维护规范路径中的目录列表，初始放入空串 "" ，
//      方便最后拼接时在最前面有 "/" 。
//
//	    用 '/' 将 path 划分成目录列表 parts ，并遍历每一个目录 part ，
//      根据 part 的值进行处理：
//          1. part 是 ".": 则表明当前目录，无需处理
//          2. part 是 "": 则表明有连续多个 '/' ，无需处理
//          3. part 是 "..": 则表明父目录，需要往上一层，
//              移除 result 的最后一个元素。
//              （注意只有非根目录才能往上一层）
//          4. part 是其他字符串：则 part 是一个合法的目录名，
//              直接放入 result
//
//      最后返回时需要进行判断：
//          1. 如果 result 只有一个元素，那么这个元素必定是 "" ，
//              则说明是根目录 "/" ，直接返回即可
//          2. 如果 result 不止一个元素，那么需要用 "/" 拼接各级目录
//
//
//		时间复杂度： O(n)
//          1. 需要遍历 path 中的全部 O(n) 个字符
//          2. 需要遍历 path 中的目录，最差情况下有 O(n) 级目录
//		空间复杂度： O(n)
//          1. 需要记录 path 中的目录，最差情况下有 O(n) 级目录


impl Solution {
    pub fn simplify_path(path: String) -> String {
        // 记录规范路径的目录列表，
        // 初始放入空串 "" ，方便最后拼接时在最前面有 "/"
        let mut result = vec![""];
        // 将 path 用 '/' 划分成多个部分，并遍历
        for part in path.split('/') {
            // 根据 part 的值决定处理逻辑
            match part {
                // 如果是 "." ，则表明当前目录，无需处理
                // 如果是 "" ，则表明有连续多个 '/' ，无需处理
                "." | "" => {},
                // 如果是 "..": 则表明父目录，需要往上一层
                ".." => {
                    // 注意只有非根目录才能往上一层
                    if result.len() > 1 {
                        result.pop();
                    }
                },
                // 其他情况是一个合法的目录名，直接放入即可
                _ => result.push(part),
            }
        }

        if result.len() == 1 {
            // 如果 result 只有一个元素，那么必定是 "" ，
            // 则说明是根目录 "/" ，直接返回即可
            "/".to_string()
        } else {
            // 如果 result 不止一个元素，那么需要用 "/" 拼接各级目录
            result.join("/")
        }
    }
}
