# 链接：https://leetcode.com/problems/rank-scores/
# 题意：给定一张表，写 SQL 查询所有分数，按高到低排序，
#       并返回每个分数对应的排名（分数相同，排名相同，排名必须连续）？
#
#       Table: Scores
#       +----+-------+
#       | Id | Score |
#       +----+-------+
#       | 1  | 3.50  |
#       | 2  | 3.65  |
#       | 3  | 4.00  |
#       | 4  | 3.85  |
#       | 5  | 4.00  |
#       | 6  | 3.65  |
#       +----+-------+
#
#       查询结果
#       +-------+------+
#       | Score | Rank |
#       +-------+------+
#       | 4.00  | 1    |
#       | 4.00  | 1    |
#       | 3.85  | 2    |
#       | 3.65  | 3    |
#       | 3.65  | 3    |
#       | 3.50  | 4    |
#       +-------+------+

# 思路1： 子查询
#
#		很自然地就能想到我们按照分数降序排序查询每个分数后，
#		再对每一个分数进行一个子查询，查询其对应对排名即可

# Write your MySQL query statement below
SELECT `a`.`Score`,
       (
           SELECT COUNT(DISTINCT `b`.Score)
           FROM `Scores` AS `b`
           WHERE `a`.`Score` <= `b`.`Score`
       ) AS `RANK`
FROM `Scores` AS `a`
ORDER BY `a`.`Score` DESC

# 思路2： 内联 + GROUP BY
#
#		后来可以通过内联去除子查询，
#		先按照子查询中 `a`.`Score` <= `b`.`Score` 进行内敛，
#       然后按照 `a`.`Id` 进行 GROUP BY ，并按照 `a`.`Score` 降序排序，
#       最后查询 `a`.`Score` 和大于等于 `a`.`Score` 对个数即可

# Write your MySQL query statement below
SELECT `a`.`Score`, COUNT(DISTINCT `b`.Score) AS `Rank`
FROM `Scores` as `a`
         INNER JOIN `Scores` as `b`
                    ON `a`.`Score` <= `b`.`Score`
GROUP BY `a`.`Id`
ORDER BY `a`.`Score` DESC
