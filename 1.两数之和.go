package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if key, ok := m[v]; ok {
			return []int{key, i}
		}
		m[target-v] = i
	}
	return nil
}

// @lc code=end
