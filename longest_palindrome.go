package main

func longestPalindrome(s string) string {
	start := 0
	pos := -1
	maxLength := 0
	strLength := len(s)
	if 0 == strLength {
		return ""
	}
	for start < strLength && start+maxLength < strLength {

		start++
	}

	// --  @# 没有找到回文串
	if -1 == pos {
		return ""
	}

	return s[pos : pos+maxLength]
}
