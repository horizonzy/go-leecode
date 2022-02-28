package main

func main() {
	c := longestPalindrome("cbbd")
	println(c)
}

func longestPalindrome(s string) string {
	//dp[i,j] = dp[i+1,j-1]&& (s[i]==s[j])
	len := len(s)
	var dp = make([][]bool, len)

	var maxLen = 0
	var start int

	for i := 0; i < len; i++ {
		dp[i] = make([]bool, len)
		dp[i][i] = true

	}

	for L := 1; L < len; L++ {

		for i := 0; i < len; i++ {
			j := i + L
			if j > len-1 {
				break
			}
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
			if dp[i][j] && j-i > maxLen {
				maxLen = j - i
				start = i
			}
		}
	}
	return s[start : start+maxLen+1]
}
