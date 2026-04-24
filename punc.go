package main

import (
	"regexp"
)

func ItPunc(s string) string {

	 a := regexp.MustCompile(`\s*([.,;:!?])`)
	 s = a.ReplaceAllString(s, "$1")

	 b := regexp.MustCompile(`([.,;:!?])(\w)`)
	 s = b.ReplaceAllString(s, "$1 $2")
	 return s
}
