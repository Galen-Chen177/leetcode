package leetcode

import "math/big"

/*
 * @lc app=leetcode.cn id=67 lang=golang
 *
 * [67] 二进制求和
 *
 * https://leetcode-cn.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (54.14%)
 * Likes:    710
 * Dislikes: 0
 * Total Accepted:    208.3K
 * Total Submissions: 384.8K
 * Testcase Example:  '"11"\n"1"'
 *
 * 给你两个二进制字符串，返回它们的和（用二进制表示）。
 *
 * 输入为 非空 字符串且只包含数字 1 和 0。
 *
 *
 *
 * 示例 1:
 *
 * 输入: a = "11", b = "1"
 * 输出: "100"
 *
 * 示例 2:
 *
 * 输入: a = "1010", b = "1011"
 * 输出: "10101"
 *
 *
 *
 * 提示：
 *
 *
 * 每个字符串仅由字符 '0' 或 '1' 组成。
 * 1 <= a.length, b.length <= 10^4
 * 字符串如果不是 "0" ，就都不含前导零。
 *
 *
 */

// @lc code=start
func addBinary(a string, b string) string {
	ai, ok := new(big.Int).SetString(a, 2)
	if !ok {
		return ""
	}
	bi, ok := new(big.Int).SetString(b, 2)
	if !ok {
		return ""
	}

	ai.Add(ai, bi)

	return ai.Text(2)
}

// @lc code=end
