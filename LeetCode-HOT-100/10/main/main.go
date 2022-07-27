package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if p[j-1] == '*' {
				if i != 0 && (s[i-1] == p[j-2] || p[j-2] == '.') {
					dp[i][j] = dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			} else {
				if i != 0 && (s[i-1] == p[j-1] || p[j-1] == '.') {
					dp[i][j] = dp[i-1][j-1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func main() {
	s := "AAABCAAB"
	p := "A*B.A*B"
	fmt.Println(isMatch(s, p))
}
