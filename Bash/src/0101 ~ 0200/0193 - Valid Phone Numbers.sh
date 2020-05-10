# 链接：https://leetcode.com/problems/valid-phone-numbers/
# 题意：给定一些手机号（在文件中，每个一行），写一行 bash 输出所有的合法手机号？
#		合法手机号格式： (xxx) xxx-xxxx 或 xxx-xxx-xxxx （x 表示一个数字）

# 输入：
#		987-123-4567
#		123 456 7890
#		(123) 456-7890
# 输出：
#		987-123-4567
#		(123) 456-7890

# 思路： grep -E
#
#		首先通过 cat 获取文件的内容，然后可以通过管道传递给下一个命令，
#		而 grep 正好支持正则匹配，所以可以直接使用 grep 即可
#   由于每个手机号一行，且不含首位空格，所以要使用 ^ 和 $ 使得每一行必须精确匹配
#
#   当然，可以使用其他命令，如： sed 或和 grep -E 等价的 egrep

# Read from the file file.txt and output all valid phone numbers to stdout.
cat file.txt | grep -E '(^[0-9]{3}-[0-9]{3}-[0-9]{4}$)|(^\([0-9]{3}\) [0-9]{3}-[0-9]{4}$)'