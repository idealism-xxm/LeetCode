# 链接：https://leetcode.com/problems/rising-temperature/
# 题意：给定一张表，包含了一些日期及其温度，
#       写 SQL 查询某一天温度比前一天温度高的所有 Id ？
#
#       Table: Weather
#       +---------+------------------+------------------+
#       | Id(INT) | RecordDate(DATE) | Temperature(INT) |
#       +---------+------------------+------------------+
#       |       1 |       2015-01-01 |               10 |
#       |       2 |       2015-01-02 |               25 |
#       |       3 |       2015-01-03 |               20 |
#       |       4 |       2015-01-04 |               30 |
#       +---------+------------------+------------------+
#
#
#       查询结果
#       +----+
#       | Id |
#       +----+
#       |  2 |
#       |  4 |
#       +----+

# 思路： 内联
#
#       按照日期内联，内联条件为两个日期相差一天，
#       然后查询后一天温度比前一天温度大的 Id 即可

# Write your MySQL query statement below
SELECT `w1`.`Id`
FROM `Weather` AS `w1`
         INNER JOIN `Weather` AS `w2` ON `w1`.`RecordDate` = DATE_ADD(`w2`.`RecordDate`, INTERVAL 1 DAY)
WHERE `w1`.`Temperature` > `w2`.`Temperature`
