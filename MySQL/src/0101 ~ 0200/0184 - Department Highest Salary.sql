# 链接：https://leetcode.com/problems/department-highest-salary/
# 题意：给定两张表，分别包含员工相关信息和部门信息，
#       写 SQL 查询工资等于所在部门最高工资的所有人？
#
#       Table: Employee
#       +----+-------+--------+--------------+
#       | Id | Name  | Salary | DepartmentId |
#       +----+-------+--------+--------------+
#       | 1  | Joe   | 70000  | 1            |
#       | 2  | Jim   | 90000  | 1            |
#       | 3  | Henry | 80000  | 2            |
#       | 4  | Sam   | 60000  | 2            |
#       | 5  | Max   | 90000  | 1            |
#       +----+-------+--------+--------------+
#
#       Table: Department
#       +----+----------+
#       | Id | Name     |
#       +----+----------+
#       | 1  | IT       |
#       | 2  | Sales    |
#       +----+----------+
#
#       查询结果
#       +------------+----------+--------+
#       | Department | Employee | Salary |
#       +------------+----------+--------+
#       | IT         | Max      | 90000  |
#       | IT         | Jim      | 90000  |
#       | Sales      | Henry    | 80000  |
#       +------------+----------+--------+

# 思路1： 内联 + 子查询
#
#		很容易就能想到前半部分查询需要先内联查询所有字段，
#		然后条件就是当前员工的工资等于其所在部门的最高工资，
#       而 其所在部门的最高工资 可以通过子查询获得

# Write your MySQL query statement below
SELECT `Department`.`Name` AS `Department`, `Employee`.`Name` AS `Employee`, `Employee`.`Salary`
FROM `Employee`
         INNER JOIN `Department` ON `Employee`.`DepartmentId` = `Department`.`Id`
WHERE `Employee`.`Salary` = (
    SELECT MAX(`a`.`Salary`)
    FROM `Employee` AS `a`
    WHERE `a`.`DepartmentId` = `Employee`.`DepartmentId`
)

# 思路2： 内联 + 多字段 IN
#
#		看了题解后发现 MySQL 支持多列一起 IN ，
#		所以我们可以在子查询里选出每个部门及其最高工资，然后 多列 IN 即可
#
#       不过看起来 多列 IN 效率不太好，可以使用等价的 EXISTS 替代
#       （平常业务中都是分别 IN ，然后程序中手动处理）

# Write your MySQL query statement below
SELECT `Department`.`Name` AS `Department`, `Employee`.`Name` AS `Employee`, `Employee`.`Salary`
FROM `Employee`
         INNER JOIN `Department` ON `Employee`.`DepartmentId` = `Department`.`Id`
WHERE (`Department`.`Id`, `Employee`.`Salary`) IN (
    SELECT `a`.`DepartmentId`, MAX(`a`.`Salary`)
    FROM `Employee` AS `a`
    GROUP BY `a`.`DepartmentId`
)
