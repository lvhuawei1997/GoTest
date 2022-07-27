package main

import "fmt"

func minWindow(s string, t string) string {
	len_s, len_t := len(s), len(t)
	if len_t > len_s {
		return ""
	}
	m := make(map[byte]int)
	for i := 0; i < len_t; i++ {
		m[t[i]]++
	}
	var str string
	minLen := len_s
	for i := 0; i < len_s; i++ {
		mp := m
		for j := i; j < len_s; j++ {
			if _, ok := mp[s[i]]; ok {
				mp[s[i]]--
				if mp[s[i]] == 0 {
					delete(mp, s[i])
				}
				if len(mp) == 0 {
					if j-i+1 < minLen {
						minLen = j - i + 1
						str = s[i : j+1]
					}
				}
			}
		}
	}
	return str
}

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}
