# 链接：https://leetcode.com/problems/second-highest-salary/
# 题意：给定一张表，写 SQL 查询所有人中第二高的薪水，没有第二高就返回 null ？
#
#       Table: Employee
#       +----+--------+
#       | Id | Salary |
#       +----+--------+
#       | 1  | 100    |
#       | 2  | 200    |
#       | 3  | 300    |
#       +----+--------+

# 思路： 模拟
#
#		由于需要第二高，先按照薪水降序排序，所以再 LIMIT 1, 1 ，
#		并且在 SELECT 的时候使用 DISTINCT 进行去重，
#		由于需要在不足两个不同的薪水时要返回 NULL ，所以还需要使用 IFNULL 语句进行处理

# Write your MySQL query statement below
SELECT IFNULL(
               NULL,
               (
                   SELECT DISTINCT `Employee`.`Salary`
                   FROM `Employee`
                   ORDER BY `Employee`.`Salary` DESC
                   LIMIT 1, 1
               )
           ) AS SecondHighestSalary
