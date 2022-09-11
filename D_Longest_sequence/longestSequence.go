package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t, n int
	var a, b string
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &t)
	for ; t > 0; t-- {
		fmt.Fscanln(in, &n)
		s, _ := in.ReadString('\n')
		clients := strings.Fields(s)
		max := 0
		if len(clients) < 3 {
			fmt.Println(len(clients))
		} else {
			a = clients[0]
			for _, v := range clients {
				if v == a {
					max += 1
					continue
				}
				b = v
				max += 1
				break
			}
			curr := max
			for i := max; i < len(clients); i++ {
				if clients[i] == a || clients[i] == b {
					curr += 1
					if curr > max {
						max = curr
					}
					continue
				}
				if curr > max {
					max = curr
				}
				a = clients[i-1]
				curr = 0
				for j := i - 1; j >= 0; j-- {
					if clients[j] == a {
						curr += 1
						continue
					}
					break
				}
				b = clients[i]
				curr += 1
			}
			fmt.Println(max)
		}
	}
}
