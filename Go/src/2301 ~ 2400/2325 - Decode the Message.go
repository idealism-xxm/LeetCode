// 链接：https://leetcode.com/problems/decode-the-message/
// 题意：给定两个字符串 key 和 message ，其中 key 表示密钥， message 表示加密消息。
//      解密 message 的步骤如下：
//          1. 使用 key 中 26 个小写字母第一次出现的位置，作为替换表中的字母顺序
//          2. 将替换表与普通字母表对齐，形成对照表
//          3. 按照对照表替换 message 的每个字母
//          4. 空格 ' ' 保持不变
//
//      返回解密后的消息。


// 数据限制：
//  26 <= key.length <= 2000
//  key 只含有英文小写字母和 ' '
//  key 含有全部英文小写字母 ('a'-'z') 至少一次
//  1 <= message.length <= 2000
//  message 只含有英文小写字母和 ' '


// 输入： key = "the quick brown fox jumps over the lazy dog", message = "vkbs bs t suepuv"
// 输出： "this is a secret"
// 解释： 对照表如下：
//          thequickbrownfxjmpsvlazydg
//          abcdefghijklmnopqrstuvwxyz

// 输入： key = "eljuxhpwnyrdgtqkviszcfmabo", message = "zwx hnfx lqantp mnoeius ycgk vcnjrdb"
// 输出： "the five boxing wizards jump quickly"
// 解释： 对照表如下：
//          eljuxhpwnyrdgtqkviszcfmabo
//          abcdefghijklmnopqrstuvwxyz


// 思路： 模拟
//
//      按照题意构建对照表名，用一个名为 chs 的字节数组维护。
//      其中 chs[i] 表示字母 i （这里 i 为字母表中的下标）对应的解密后的字母。
//
//      初始化均为 0 ，表示还未找到解密后的字母，
//      同时维护下一个加密字母对应的解密后的字母 origin ，初始化为 'a'。
//
//      遍历 key 中的每个字母 ch （这里 i 为字母表中的下标），
//      如果 chs[ch] 为 ' ' ，则将 chs[ch] 设置为 origin ，
//      并将 origin 移向字母表中的下一个字母。
//
//      最后遍历 message 中的每个字母 ch ，将其按照​对照表转换成原字母即可。
//
//
//      设字符集大小为 C ， key 的长度为 m ， message 的长度为 n 。
//
//      时间复杂度： O(m + n)
//          1. 需要遍历 key 中全部 O(m) 个字母
//          2. 需要遍历 message 全部 O(n) 个字母
//      空间复杂度：O(C + n)
//          1. 需要存储全部 O(C) 个字母的解密后的字母
//          2. 需要存储 message 解密后消息的全部 O(n) 个字母


func decodeMessage(key string, message string) string {
	// 初始化对照表， 0 表示还未找到解密后的字母
	chs := make([]byte, 26)
	// 初始化下一个加密字母对应的解密后的字母
	origin := byte('a')
	// 遍历密钥中的每个字母
	for _, ch := range key {
		// 如果 ch 不是空格，且是第一次出现，则设置对照表中的字母为 origin
		if ch != ' ' && chs[ch-'a'] == 0 {
			chs[ch-'a'] = origin
			origin += 1
		}
	}

	// 解密 message
	ans := make([]byte, len(message))
	for i, ch := range message {
		// 如果 ch 不是空格，则解密
		if ch == ' ' {
			// 如果 ch 是空格，则依旧为空格
			ans[i] = ' '
		} else {
			// // 如果 ch 是字母，则设置为对照表中的字母
			ans[i] = chs[ch-'a']
		}
	}

	return string(ans)
}
