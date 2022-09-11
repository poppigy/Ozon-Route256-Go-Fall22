package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type ver struct {
	right string
	left  string
}

func main() {
	var t, n int
	poly := map[string]ver{}
	in := bufio.NewReader(os.Stdin)
	fmt.Fscanln(in, &t)
	for ; t > 0; t-- {
		fmt.Fscanln(in, &n)
		primaryKey := ""
		for j := 0; j < n; j++ {
			s, _ := in.ReadString('\n')
			vertices := strings.Fields(s)
			if _, ok := poly[vertices[0]]; !ok {
				poly[vertices[0]] = ver{vertices[1], vertices[2]}
			}
			primaryKey = vertices[0]
		}
		var sorted []string
		value := poly[primaryKey].right
		for ; n > 0; n -= 2 {
			sorted = append(sorted, primaryKey)
			sorted = append(sorted, value)
			if poly[value].right == primaryKey {
				primaryKey = poly[value].left
			} else {
				primaryKey = poly[value].right
			}
			if poly[primaryKey].right == value {
				value = poly[primaryKey].left
			} else {
				value = poly[primaryKey].right
			}
		}

		offset := len(sorted) / 2
		for s := 0; s < len(sorted)/2; s++ {
			fmt.Println("sorted: ", sorted[s], sorted[s+offset])
		}
		for k, _ := range poly {
			delete(poly, k)
		}
	}
}
