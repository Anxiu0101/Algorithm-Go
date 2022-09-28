package hashtable

// Question Addr: https://leetcode.cn/problems/valid-anagram/
// 给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
// 注意：若 `s` 和 `t` 中每个字符出现的次数都相同，则称 `s` 和 `t` 互为字母异位词。

// 在提示了要使用哈希表的情况下，打算将第一个字符串先
func isAnagram(s string, t string) bool {

	// 官方答案中两种方法，
	// 第一个是排序法，t 是 s 的异位词等价于「两个字符串排序后相等」
	//s1, s2 := []byte(s), []byte(t)
	//sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	//sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
	//return string(s1) == string(s2)

	// 第二个是哈希表，t 是 s 的异位词等价于「两个字符串中字符出现的种类和次数均相等」。
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
