// 链接：https://leetcode.com/problems/fraction-to-recurring-decimal/
// 题意：给定一个分数的分子和分母，将其转换成小数字符串（有循环节需要标记） ？

// 输入： numerator = 1, denominator = 2
// 输出： "0.5"

// 输入： numerator = 2, denominator = 1
// 输出： "2"

// 输入： numerator = 2, denominator = 3
// 输出： "0.(6)"

// 思路： 模拟
//
//		分成两部分处理，
//		整数部分 = numerator / denominator （整数除法）
//		小数部分 = (numerator % denominator) / denominator （小数除法）
//
//		整数部分很简单，直接计算即可
//		小数部分先判断余数是否为 0
//			(1) 若余数为 0 ，则没有小数部分
//			(2) 若余数不为 0 ，则模拟除法计算，并记录每次余数出现的下标，
//				直至余数为 0 （有限小数，直接返回即可）
//				或者 余数再次出现（无限循环小数，上次下标至本次下标的是循环节）
//
//		时间复杂度： O(denominator)
//		空间复杂度： O(denominator)

import (
	"bytes"
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	// 记录结果符号，后面计算时只用正数
	sign := ""
	if numerator < 0 {
		sign = "-"
		numerator = -numerator
	}
	if denominator < 0 {
		if sign == "" {
			sign = "-"
		} else {
			sign = ""
		}
		denominator = -denominator
	}
	integerPart := strconv.Itoa(numerator / denominator)
	remain := (numerator % denominator) * 10
	// 整除时直接返回
	if remain == 0 {
		// 如果是 0 ，则直接返回，避免负号干扰
		if integerPart == "0" {
			return integerPart
		}
		return fmt.Sprintf("%v%v", sign, integerPart)
	}

	// 有小数部分需要计算
	var decimalPartBuffer bytes.Buffer
	remainToIndex := make(map[int]int)
	for i := 0; ; i++ {
		if remain == 0 {
			// 有限小数，直接拼接返回即可
			return fmt.Sprintf("%v%v.%v", sign, integerPart, decimalPartBuffer.String())
		}
		if lastIndex, exists := remainToIndex[remain]; exists {
			// 无限循环小数，标记循环节后拼接返回即可
			decimalPart := decimalPartBuffer.String()
			decimalPartPre := decimalPart[:lastIndex]
			decimalPartLoop := decimalPart[lastIndex:]
			return fmt.Sprintf("%v%v.%v(%v)", sign, integerPart, decimalPartPre, decimalPartLoop)
		}
		// 先进行标记
		remainToIndex[remain] = i
		// 模拟除法计算
		decimalPartBuffer.WriteString(strconv.Itoa(remain / denominator))
		remain = (remain % denominator) * 10
	}
}
