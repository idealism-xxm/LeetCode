# 链接：https://leetcode.com/problems/duplicate-emails/
# 题意：给定一张表，包含一些邮箱，
#       写 SQL 查询所有重复的邮箱？
#
#       Table: Person
#       +----+---------+
#       | Id | Email   |
#       +----+---------+
#       | 1  | a@b.com |
#       | 2  | c@d.com |
#       | 3  | a@b.com |
#       +----+---------+
#
#       查询结果
#       +---------+
#       | Email   |
#       +---------+
#       | a@b.com |
#       +---------+

# 思路： GROUP BY + HAVING
#
#		直接按照 email GROUP BY ，然后使用 HAVING 过滤 COUNT 数大于 1 的即可

# Write your MySQL query statement below
SELECT `Person`.`Email`
FROM `Person`
GROUP BY `Person`.`Email`
HAVING COUNT(`Person`.`Id`) > 1
