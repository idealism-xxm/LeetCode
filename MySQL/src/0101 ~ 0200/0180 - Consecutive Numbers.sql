# 链接：https://leetcode.com/problems/consecutive-numbers/
# 题意：给定一张表，写 SQL 查询所有数字连续出现 3 次的数字？
#
#       Table: Logs
#       +----+-----+
#       | Id | Num |
#       +----+-----+
#       | 1  |  1  |
#       | 2  |  1  |
#       | 3  |  1  |
#       | 4  |  2  |
#       | 5  |  1  |
#       | 6  |  2  |
#       | 7  |  2  |
#       +----+-----+
#
#       查询结果
#       +-----------------+
#       | ConsecutiveNums |
#       +-----------------+
#       | 1               |
#       +-----------------+

# 思路1： 内联
#
#		官方题解认为连续出现就意味着 id 必定连续，
#		所以可以内联 Logs 表三次即可，
#       这样可以解决当前题目的数据，
#       但是总觉得有点奇怪，题目并没有用明确说明 id 连续，
#       并且实际场景中肯定存在跳跃的情况，所以不具备通用性

# Write your MySQL query statement below
SELECT DISTINCT `a`.`Num` AS ConsecutiveNums
FROM `Logs` AS `a`
         INNER JOIN `Logs` AS `b` ON `a`.`Id` = `b`.`Id` - 1
         INNER JOIN `Logs` AS `c` ON `b`.`Id` = `c`.`Id` - 1
WHERE `a`.`Num` = `b`.`Num`
  AND `b`.`Num` = `c`.Num


# 思路2： 变量 + IF
#
#		看了题解后，发现还可以使用变量分别记录每一个数字连续出现的次数 和 上一次数字，
#		然后在根据 IF 每次更新两个变量，其实就和平时其他语言写代码一样，
#       最后求数字连续出现次数大于 2 的即可

# Write your MySQL query statement below
SELECT DISTINCT `a`.`Num` AS ConsecutiveNums
FROM (
         SELECT `b`.`Num`, @cnt := IF(@pre = `b`.`Num`, @cnt + 1, 1) AS cnt, @pre := `b`.`Num`
         FROM `Logs` AS `b`, (SELECT @pre := NULL, @cnt := 0) AS `c`
         ORDER BY `b`.`Id`
     ) AS `a`
WHERE `a`.`cnt` > 2