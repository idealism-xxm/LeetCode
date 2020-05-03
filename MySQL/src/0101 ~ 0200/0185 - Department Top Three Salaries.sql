# 链接：https://leetcode.com/problems/department-top-three-salaries/
# 题意：给定两张表，分别包含员工相关信息和部门信息，
#       写 SQL 查询工资是所在部门前 3 高的所有人？
#
#       Table: Employee
#       +----+-------+--------+--------------+
#       | Id | Name  | Salary | DepartmentId |
#       +----+-------+--------+--------------+
#       | 1  | Joe   | 85000  | 1            |
#       | 2  | Henry | 80000  | 2            |
#       | 3  | Sam   | 60000  | 2            |
#       | 4  | Max   | 90000  | 1            |
#       | 5  | Janet | 69000  | 1            |
#       | 6  | Randy | 85000  | 1            |
#       | 7  | Will  | 70000  | 1            |
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
#       | IT         | Randy    | 85000  |
#       | IT         | Joe      | 85000  |
#       | IT         | Will     | 70000  |
#       | Sales      | Henry    | 80000  |
#       | Sales      | Sam      | 60000  |
#       +------------+----------+--------+

# 思路1： 内联 + IN + LIMIT
#
#       0184 的加强版，但很容易就能将解法扩展为符合本题的情况
#
#		前半部分查询需要先内联查询所有字段，
#		然后条件就是当前员工的工资 IN 其所在部门的前 3 高工资，
#       而 其所在部门的前 3 高工资 可以通过子查询获得
#
#       但是直接用 IN + LIMIT 提示当前版本无法这种方式查询，
#       可以再 SELECT 一次绕过这个限制

# Write your MySQL query statement below
SELECT `Department`.`Name` AS `Department`, `Employee`.`Name` AS `Employee`, `Employee`.`Salary`
FROM `Employee`
         INNER JOIN `Department` ON `Employee`.`DepartmentId` = `Department`.`Id`
WHERE `Employee`.`Salary` IN (
    SELECT `b`.`Salary`
    FROM (
             SELECT DISTINCT `a`.`Salary`
             FROM `Employee` AS `a`
             WHERE `a`.`DepartmentId` = `Employee`.`DepartmentId`
             ORDER BY `a`.`Salary` DESC
             LIMIT 3
         ) AS `b`
)

# 思路2： 内联 + COUNT
#
#       看了题解发现其实不用正面解决，
#       可以通过子查询中查其所在部门比其工资高的不同的工资数，不超过 3 即可
#
#		前半部分查询需要先内联查询所有字段，
#		然后条件就是比 当前员工的工资 高的不同的工资数不超过 3 即可

# Write your MySQL query statement below
SELECT `Department`.`Name` AS `Department`, `Employee`.`Name` AS `Employee`, `Employee`.`Salary`
FROM `Employee`
         INNER JOIN `Department` ON `Employee`.`DepartmentId` = `Department`.`Id`
WHERE 3 > (
    SELECT COUNT(DISTINCT `a`.`Salary`)
    FROM `Employee` AS `a`
    WHERE `a`.`DepartmentId` = `Employee`.`DepartmentId`
    AND `a`.`Salary` > `Employee`.`Salary`
)
