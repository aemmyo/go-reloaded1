package main

import (
	"strconv"
	"strings"
)

func HexBin(text string) string {
	s := strings.Fields(text)
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			num, err := strconv.ParseInt(s[i-1], 16, 64)
			if err == nil {
				s[i-1] = strconv.FormatInt(num, 10)
				s = append(s[:i], s[i+1:]...)
				i--
			}
		}
			if s[i] == "(bin)" {
			num, err := strconv.ParseInt(s[i-1], 2, 64)
			if err == nil {
				s[i-1] = strconv.FormatInt(num, 10)
				s = append(s[:i], s[i+1:]...)
				i--
			}
		}
				
			
		}
	
	return strings.Join(s, " ")
}
