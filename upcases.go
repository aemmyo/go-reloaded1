package main

import (
	"strconv"
	"strings"
)

func HandleCases(text string) string {
	s := strings.Fields(text)
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}
		if s[i] == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}
		if s[i] == "(cap)" {
			s[i-1] = strings.Title(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--
		}
		if strings.HasPrefix(s[i], "(") && i+1 < len(s) {
			numStr := strings.TrimRight(s[i+1], ")")
			num, err := strconv.Atoi(numStr)
			if err == nil {
				for j := 1; j <= num && i-j >= 0; j++ {
					if strings.Contains(s[i], "up") {
						s[i-j] = strings.ToUpper(s[i-j])
					}
					if strings.Contains(s[i], "low") {
						s[i-j] = strings.ToLower(s[i-j])
					}
					if strings.Contains(s[i], "cap") {
						s[i-j] = strings.Title(s[i-j])
					}
				}
				s = append(s[:i], s[i+2:]...)
				i--
			}
			
		}
	}
	return strings.Join(s, " ")

}