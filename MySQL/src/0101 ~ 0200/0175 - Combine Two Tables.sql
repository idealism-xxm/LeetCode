# 链接：https://leetcode.com/problems/combine-two-tables/
# 题意：给定两张表，写 SQL 查询所有人的以下四列数据： FirstName, LastName, City, State
#      并且无论每个人是否有对应的地址，都要返回？
#
#       Table: Person
#       +-------------+---------+
#       | Column Name | Type    |
#       +-------------+---------+
#       | PersonId    | int     |
#       | FirstName   | varchar |
#       | LastName    | varchar |
#       +-------------+---------+
#       PersonId is the primary key column for this table.
#
#       Table: Address
#       +-------------+---------+
#       | Column Name | Type    |
#       +-------------+---------+
#       | AddressId   | int     |
#       | PersonId    | int     |
#       | City        | varchar |
#       | State       | varchar |
#       +-------------+---------+
#       AddressId is the primary key column for this table.

# 思路： 左联
#
#		由于左表的数据不能缺失，所以当 Person 没有地址时需要后两列为 null ，
#       这样时左联查询，直接写左联 SQL 即可

# Write your MySQL query statement below
SELECT `Person`.`FirstName`, `Person`.`LastName`, `Address`.`City`, `Address`.`State`
FROM `Person`
         LEFT JOIN `Address` ON `Person`.`PersonId` = `Address`.`PersonId`
