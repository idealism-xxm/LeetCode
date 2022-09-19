# 链接：https://leetcode.com/problems/find-duplicate-file-in-system/
# 题意：给定一个字符串数组 paths ，表示一些文件夹的信息，
#      包含文件夹路径以及其下所有文件和内容，
#      返回所有内容重复的文件的路径（按文件内容分组）。
#
#      paths[i] 格式如下： 
#      "root/d1/d2/.../dm f1.txt(f1_content) f2.txt(f2_content) ... fn.txt(fn_content)"


# 数据限制：
#  1 <= paths.length <= 2 * 10 ^ 4
#  1 <= paths[i].length <= 3000
#  1 <= sum(paths[i].length) <= 5 * 10 ^ 5
#  paths[i] 仅由以下字符组成：英文小写字母，数字， '/', '.', '(', ')', ' '
#  相同目录下没有文件名相同的文件
#  所有的文件夹路径都不相同，文件信息和文件夹路径之间通过一个空格分隔


# 输入： paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)","root 4.txt(efgh)"]
# 输出： [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]

# 输入： paths = ["root/a 1.txt(abcd) 2.txt(efgh)","root/c 3.txt(abcd)","root/c/d 4.txt(efgh)"]
# 输出： [["root/a/2.txt","root/c/d/4.txt"],["root/a/1.txt","root/c/3.txt"]]


# 思路： Map
#
#      根据题意，我们只需要按照给定的格式，解析出：文件夹路径、文件名和文件内容，
#      然后用文件夹路径和文件名拼接出文件路径，再按照文件内容分组即可。
#
#      对于第 i 个路径 paths[i] 来说，我们先按照空格分隔出 parts 数组，
#      则 parts[0] 就是文件夹路径， parts[1:] 就是文件信息列表。
#
#      再遍历 parts[1:] 中的每个文件信息 part ，对其按照左小括号分隔一次，
#      则必定会分隔出两个字符串，前者就是文件名，
#      后者就是文件内容（最后一个字符是右小括号，但不影响结果，可不处理）。
#
#
#      设字符串数组 paths 的字符串平均长度为 x 。
#
#      时间复杂度：O(nx)
#          1. 需要遍历 paths 全部 O(n) 个字符串，每次需要遍历全部 O(x) 个字符
#      空间复杂度：O(nx)
#          1. 需要维护总大小为 O(nx) 的所有文件路径，包括分组用的 map 和结果数组


class Solution:
    def findDuplicate(self, paths: List[str]) -> List[List[str]]:
        # content_to_paths 维护文件内容对应的所有文件路径列表
        content_to_paths: defaultdict = defaultdict(list)
        # 遍历所有路径
        for path in paths:
            # 按照空格分隔，第一个是文件夹路径
            parts: List[str] = path.split(' ')
            directory: str = parts[0]
            # 遍历后续的文件信息
            for i in range(1, len(parts)):
                file: str = parts[i]
                # 按照前小括号分隔，前者是文件名，
                # 后者是文件内容（最后一个字符是右小括号，但不影响结果，可不处理）
                filename, content = file.split('(', 1)
                # 拼接成该文件的文件路径，然后根据文件内容放入到对应的列表中
                content_to_paths[content].append(f"{directory}/{filename}")

        # 遍历分组后的文件列表，仅收集内容重复的（文件路径数量大于 1 ）
        return [
            paths
            for paths in content_to_paths.values()
            if len(paths) > 1
        ]
