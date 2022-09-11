package main

import "fmt"

var a = "00"
var b = "100"
var c = "101"
var d = "11"

func decode(s string) string {
	res := ""
	for i := 0; i < len(s); {
		if s[i] == '1' {
			if s[i+1] == '1' { // 11
				res += "d"
				i += 2
			} else {
				if s[i+2] == '1' { //101
					res += "c"
				} else { //100
					res += "b"
				}
				i += 3
			}
		} else {
			res += "a"
			i += 2
		}
	}
	return res
}

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	var s string
	var res []string
	for i := 0; i < t; i++ {
		s = ""
		fmt.Scanf("%s\n", &s)
		res = append(res, s)
	}
	for _, r := range res {
		fmt.Println(decode(r))
	}
}
