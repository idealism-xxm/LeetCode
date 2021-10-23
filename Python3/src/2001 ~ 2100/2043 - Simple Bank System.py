# 链接：https://leetcode.com/problems/simple-bank-system/
# 题意：实现一个银行系统，初始该银行系统有 n 个人，
#       账户余额在 balance 数组中，然后有一系列的事务操作，
#       该银行系统只会成功执行合法的事务，总共需要实现以下三种事务操作：
#       1. transfer(account1, account2, money) ：
#           账户 account1 向账户 account2 转账 money 元，
#           如果 account1 余额不足 money ，则不执行该操作，并返回 false ，
#           否则，执行转账操作，返回 true 
#       2. deposit(account, money) ：
#           向账户 account 存入 money 元，并返回 true
#       3. withdraw(account, money) ：
#           从账户 account 中取出 money 元，
#           如果 account 余额不足 money ，则不执行该操作，并返回 false ，
#           否则，执行取出操作，返回 true 

# 数据限制：
#   n == balance.length
#   1 <= n, account, account1, account2 <= 10 ^ 5
#   0 <= balance[i], money <= 10 ^ 12
#   transfer, deposit, withdraw 最多各有 10 ^ 4 次调用

# 输入： s = "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
# 输出： true
# 解释： s 中的数字 1, 3, 4, 6, 12 ，是严格单调递增的。

# 输入： ["Bank", "withdraw", "transfer", "deposit", "transfer", "withdraw"]
#       [[[10, 100, 20, 50, 30]], [3, 10], [5, 1, 20], [5, 20], [3, 4, 15], [10, 50]]
# 输出： [null, true, true, true, false, false]
# 解释： 
#   Bank bank = new Bank([10, 100, 20, 50, 30]);
#   bank.withdraw(3, 10);    # 返回 true ， 账号 3 有余额 20 ，所以取出 10 是合法的，
#                            # 账号 3 现在有 20 - 10 = 10
#   bank.transfer(5, 1, 20); # 返回 true ， 账号 5 有余额 30 ，所以转账 20 是合法的，
#                            # 账号 5 现在有 30 - 20 = 10 ，账号 1 现在有 10 + 20 = 30
#   bank.deposit(5, 20);     # 返回 true ， 账号 5 存入 20 是合法的
#                            # 账号 5 现在有 10 + 20 = 30
#   bank.transfer(3, 4, 15); # 返回 false ，账号 3 现在余额是 10 ，无法转账 15
#   bank.withdraw(10, 50);   # 返回 false ，账号 10 不存在


# 思路： 模拟
#
#       按照题意模拟即可，先判断每个账号是否合法，再判断转账和取出是否合法，
#       如果都合法，则可以执行操作，否则不执行操作
#
#       时间复杂度： O(q) ，其中 q 表示所有的操作次数
#       空间复杂度： O(n)


class Bank:

    def __init__(self, balance: List[int]):
        self.length = len(balance)
        self.balance = balance

    def transfer(self, account1: int, account2: int, money: int) -> bool:
        # 如果账号不存在，则直接返回 False
        if account1 > self.length or account2 > self.length:
            return False
        
        # 如果账号余额不足，则直接返回 False
        if self.balance[account1 - 1] < money:
            return False

        # 执行转账操作
        self.balance[account1 - 1] -= money
        self.balance[account2 - 1] += money
        return True

    def deposit(self, account: int, money: int) -> bool:
        # 如果账号不存在，则直接返回 False
        if account > self.length:
            return False

        # 执行存入操作
        self.balance[account - 1] += money
        return True
        

    def withdraw(self, account: int, money: int) -> bool:
        # 如果账号不存在，则直接返回 False
        if account > self.length:
            return False

        # 如果账号余额不足，则直接返回 False
        if self.balance[account - 1] < money:
            return False
        
        # 执行取出操作
        self.balance[account - 1] -= money
        return True
        


# Your Bank object will be instantiated and called as such:
# obj = Bank(balance)
# param_1 = obj.transfer(account1,account2,money)
# param_2 = obj.deposit(account,money)
# param_3 = obj.withdraw(account,money)
