package main

import "fmt"

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}

func Split(s, sep string) []string {
	var res []string
	word := ""
	i := 0
	for i < len(s) {
		if s[i] == sep[0] {
			frag := s[i : i+len(sep)]
			if frag == sep {
				res = append(res, word)
				word = ""
				i += len(sep)
				continue
			}
		}
		word += string(s[i])
		i++
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}

/*
package piscine
func Split(s, sep string) []string {
	var res []string
	str := ""
	i := 0
	for len(s) > i {
		if s[i] == sep[0] {
			flag := s[i : len(sep)+i]
			if flag == sep {
				res = append(res, str)
				str = ""
				i += len(sep)
				continue
			}
		}
		str += string(s[i])
		i++
	}
	res = append(res, str)
	return res
}
*/
