# 链接：https://leetcode.com/problems/customers-who-never-order/
# 题意：给定两张表，分别包含所有用户和部分用户的订单，
#       写 SQL 查询没有订单的用户？
#
#       Table: Customers
#       +----+-------+
#       | Id | Name  |
#       +----+-------+
#       | 1  | Joe   |
#       | 2  | Henry |
#       | 3  | Sam   |
#       | 4  | Max   |
#       +----+-------+
#
#       Table: Orders
#       +----+------------+
#       | Id | CustomerId |
#       +----+------------+
#       | 1  | 3          |
#       | 2  | 1          |
#       +----+------------+
#
#       查询结果
#       +-----------+
#       | Customers |
#       +-----------+
#       | Henry     |
#       | Max       |
#       +-----------+

# 思路1： 子查询 + NOT IN
#
#		最直观的反应就是自查询，先查出所有有订单的用户 Id ，
#       然后查用户表中非这些用户 Id 的用户即可

# Write your MySQL query statement below
SELECT `Customers`.`Name` AS `Customers`
FROM `Customers`
WHERE `Customers`.`Id` NOT IN (
    SELECT DISTINCT `Orders`.`CustomerId`
    FROM `Orders`
)

# 思路2： 左联
#
#		还可以想到使用左联，过滤 订单 Id 为 NULL 用户即可

# Write your MySQL query statement below
SELECT `Customers`.`Name` AS `Customers`
FROM `Customers`
         LEFT JOIN `Orders` ON `Customers`.`Id` = `Orders`.`CustomerId`
WHERE `Orders`.`Id` IS NULL
