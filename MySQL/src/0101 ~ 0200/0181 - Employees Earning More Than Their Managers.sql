# 链接：https://leetcode.com/problems/employees-earning-more-than-their-managers/
# 题意：给定一张表，包含所有的员工及其管理者的 id ，
#       写 SQL 查询所有工资高于其管理者的员工？
#
#       Table: Employee
#       +----+-------+--------+-----------+
#       | Id | Name  | Salary | ManagerId |
#       +----+-------+--------+-----------+
#       | 1  | Joe   | 70000  | 3         |
#       | 2  | Henry | 80000  | 4         |
#       | 3  | Sam   | 60000  | NULL      |
#       | 4  | Max   | 90000  | NULL      |
#       +----+-------+--------+-----------+
#
#
#       查询结果
#       +----------+
#       | Employee |
#       +----------+
#       | Joe      |
#       +----------+

# 思路： 内联
#
#		直接内联即可，然后过滤工资大于其管理者的员工即可

# Write your MySQL query statement below
SELECT `a`.`Name` AS Employee
FROM `Employee` AS `a`
         INNER JOIN `Employee` AS `b` ON `a`.`ManagerId` = `b`.`Id`
WHERE `a`.`Salary` > `b`.`Salary`
