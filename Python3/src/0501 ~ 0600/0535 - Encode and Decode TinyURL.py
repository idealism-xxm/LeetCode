# 链接：https://leetcode.com/problems/encode-and-decode-tinyurl/
# 题意：实现一个支持短链接编码和解码的数据结构。
#      该数据结构需要支持以下操作：
#          1. String encode(String longUrl): 返回原始链接 longUrl 对应的短链接
#          2. String decode(String shortUrl): 返回短链接 shortUrl 对应的原始链接


# 数据限制：
#  1 <= url.length <= 10 ^ 4
#  url 确保是一个合法的链接


# 输入： url = "https://leetcode.com/problems/design-tinyurl"
# 输出： "https://leetcode.com/problems/design-tinyurl"
# 解释： Solution obj = new Solution();
#       string tiny = obj.encode(url); # 返回编码后的短链接
#       string ans = obj.decode(tiny); # 返回解码后的原始链接


# 思路： Map
#
#      可以使用随机生成 6 位 uid 的方式来标识别一个原始链接，
#      并且使用两个 map 来存储原始链接和短链接的相互映射关系。
#
#      注意随机生成的 uid 可能会重复冲突，
#      所以需要使用循环的方式生成一个全新的 uid 。
#
#      uid 使用 a-zA-Z0-9 的字符集，
#      那么不同 uid 的个数为 62 ^ 6 ≈ 5.68 * 10 ^ 10 个。
#
#      uid 重复冲突的概率极低，在本题数据量下基本可以忽略不计。
#
#
#      设 l 为原始链接的平均长度， n 为原始链接的个数
#
#      时间复杂度：O(l)
#          1. 每次编码都需要遍历原始链接中的全部 O(l) 个字符
#      空间复杂度：O(n)
#          1. 需要存储全部 O(n) 个原始链接和短链接的相互映射关系


from random import choices
import string


class Codec:

    def __init__(self):
        # 存储链接对应的 uid
        self.url_to_uid = {}
        # 存储 uid 对应的链接
        self.uid_to_url = {}
        # 字符集
        self.charset = string.ascii_letters + string.digits

    def encode(self, longUrl: str) -> str:
        """Encodes a URL to a shortened URL.
        """
        # 如果不存在该链接，则继续循环生成，防止生成的 uid 重复
        while longUrl not in self.url_to_uid:
            # 生成长度为 6 的 uid
            uid: str = "".join(choices(self.charset, k=6))
            # 如果 uid 未存在于哈希表中，则可以使用
            if uid not in self.uid_to_url:
                self.url_to_uid[longUrl] = uid
                self.uid_to_url[uid] = longUrl

        # 生成并返回对应短链接
        return "http://tinyurl.com/" + self.url_to_uid[longUrl]
        

    def decode(self, shortUrl: str) -> str:
        """Decodes a shortened URL to its original URL.
        """
        # 返回对应的原始链接
        return self.uid_to_url.get(shortUrl[-6:])
