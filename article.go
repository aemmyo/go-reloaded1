package main

import(
	"strings"
)
func aOrAn(text string)string{
	s:= strings.Fields(text)
	for i := 0; i < len(s); i++{
		k := "aeiouhAEIOUH"
		if s[i] == "a" && strings.ContainsAny(k,s[i+1][:1]){
			s[i]= "an"
		}
		if s[i] == "A" && strings.ContainsAny(k,s[i+1][:1]){
			s[i]= "An"
		}
	}
	return strings.Join(s, " ")
}
