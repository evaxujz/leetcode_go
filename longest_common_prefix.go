package main

func main() {

}

// 最长公共前缀

// 编写一个函数来查找字符串数组中的最长公共前缀。
//
// 如果不存在公共前缀，返回空字符串 ""。
func longestCommonPrefix(strs []string) string {
	if strs == nil || len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for _, str := range strs {
		length := 0
		if len(prefix) < len(str) {
			length = len(prefix)
		} else {
			length = len(str)
		}
		prefix = prefix[0:length]
		for i := 0; i < length; i++ {
			if prefix[i] != str[i] {
				prefix = prefix[0:i]
				break
			}
		}
	}
	return prefix
}
