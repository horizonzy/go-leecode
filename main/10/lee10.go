package main

func main() {

}

func isMatch(s string, p string) bool {

	// if p[j] == '.', s[i]==p[j]
	// p[i] != *, if s[i] == p[i], dp[i][j] = dp[i-1][j-1]. if s[i] != p[i], dp[i][j] = false
	// p[i] == *, if s[i] == p[j-1], dp[i][j] = dp[i-1][j] or dp[i][j-2], if s[i] != p[j-1], dp[i][j] = dp[i][j-2]

	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	len1 := len(s)
	len2 := len(p)
	var dp = make([][]bool, len1+1)

	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}
	dp[0][0] = true

	for i := 0; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if matches(i, j-1) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			} else if matches(i, j) {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			}
		}
	}
	return dp[len1][len2]
}
