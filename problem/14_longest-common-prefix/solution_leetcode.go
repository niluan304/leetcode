package main

// 方法一：横向扫描
// 用 LCP(S1…Sn)表示字符串 S1…Sn 的最长公共前缀。
//
// 可以得到以下结论：
//
// LCP(S1…Sn)=LCP(LCP(LCP(S1,S2),S3),…Sn)
// 基于该结论，可以得到一种查找字符串数组中的最长公共前缀的简单方法。
// 依次遍历字符串数组中的每个字符串，对于每个遍历到的字符串，更新最长公共前缀，
// 当遍历完所有的字符串以后，即可得到字符串数组中的最长公共前缀。
//
// 如果在尚未遍历完所有的字符串时，最长公共前缀已经是空串，则最长公共前缀一定是空串，
// 因此不需要继续遍历剩下的字符串，直接返回空串即可。
func leetcode1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

// 方法二：纵向扫描
// 方法一是横向扫描，依次遍历每个字符串，更新最长公共前缀。另一种方法是纵向扫描。
// 纵向扫描时，从前往后遍历所有字符串的每一列，比较相同列上的字符是否相同，如果相同则继续对下一列进行比较，
// 如果不相同则当前列不再属于公共前缀，当前列之前的部分为最长公共前缀。
func leetcode2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
