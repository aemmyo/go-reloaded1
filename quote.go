package main

import (
	"regexp"
)

func ItQuote(s string) string {
	 a := regexp.MustCompile(`\s+`)
	 s = a.ReplaceAllString(s, " ")

	b := regexp.MustCompile(`'\s*([^']*?)(\s*)'`)
	s = b.ReplaceAllString(s, "'$1'")
	return s
}

