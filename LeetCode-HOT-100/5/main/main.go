package main

import "fmt"

func longestPalindrome(s string) (ans string) {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s); i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
				} else if dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] && j-i+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}
	return
}

func main() {
	s := "aba"
	res := longestPalindrome(s)
	fmt.Println(res)
}
