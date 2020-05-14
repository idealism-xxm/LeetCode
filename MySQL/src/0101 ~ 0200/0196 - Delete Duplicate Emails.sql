# 链接：https://leetcode.com/problems/delete-duplicate-emails/
# 题意：给定一张表，包含了一些邮箱，写 SQL 删除重复的邮箱，重复的只保留 ID 最小的？
#
#       Table: Person
#       +----+------------------+
#       | Id | Email            |
#       +----+------------------+
#       | 1  | john@example.com |
#       | 2  | bob@example.com  |
#       | 3  | john@example.com |
#       +----+------------------+
#       Id is the primary key column for this table.
#
#
#       查询结果
#       +----+------------------+
#       | Id | Email            |
#       +----+------------------+
#       | 1  | john@example.com |
#       | 2  | bob@example.com  |
#       +----+------------------+

# 思路1： GROUP BY + MIN
#
#       既然是相同邮箱只保留最小的，那么就可以先按照 Email 进行 GROUP BY ，
#       然后取其中 ID 最小的，最后删除 ID 不在这其中的即可

# Write your MySQL query statement below
DELETE
FROM `Person`
WHERE `Person`.`Id` NOT IN
      (
          SELECT *
          FROM (
                   SELECT MIN(`p`.`Id`)
                   FROM `Person` AS `p`
                   GROUP BY `p`.`Email`
               ) AS `a`
      )

# 思路2： 内联
#
#       刚开始也想通过按照 Email 内联搞，
#       但是没反应过来可以删除 Id 大的行即可

# Write your MySQL query statement below
DELETE `p1`
FROM `Person` AS `p1`
         INNER JOIN `Person` AS `p2` ON `p1`.`Email` = `p2`.`Email`
WHERE `p1`.`Id` > `p2`.`Id`
