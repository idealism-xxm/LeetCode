# 链接：https://leetcode.com/problems/nth-highest-salary/
# 题意：给定一张表，写 SQL 查询所有人中第 n 高的薪水，没有第 n 高就返回 null ？
#
#       Table: Employee
#       +----+--------+
#       | Id | Salary |
#       +----+--------+
#       | 1  | 100    |
#       | 2  | 200    |
#       | 3  | 300    |
#       +----+--------+

# 思路： FUNCTION
#
#		由于需要第 n 高，所以先计算需要跳过的人数 skp = n - 1 ，
#		再按照薪水降序排序，并 LIMIT skp, 1 ，
#		并且在 SELECT 的时候使用 DISTINCT 进行去重，
#		由于本次我们只用写内部的方法，所以不需要再使用 IFNULL 处理

CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    DECLARE skp INT;
    SET skp = N - 1;
    RETURN (
        # Write your MySQL query statement below.
        SELECT DISTINCT `Employee`.`Salary`
        FROM `Employee`
        ORDER BY `Employee`.`Salary` DESC
        LIMIT skp, 1
    );
END

