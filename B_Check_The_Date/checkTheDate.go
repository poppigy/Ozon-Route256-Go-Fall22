package main

import "fmt"

var leapYear bool

func isLeap(year int) {
	if year%400 == 0 {
		leapYear = true
	} else if year%4 == 0 {
		if year%100 == 0 {
			leapYear = false
		} else {
			leapYear = true
		}
	} else {
		leapYear = false
	}
}

func daysInMonth(month int) int {
	if month == 2 {
		if leapYear {
			return 29
		} else {
			return 28
		}
	}
	if month == 4 || month == 6 || month == 9 || month == 11 {
		return 30
	}
	return 31
}

func main() {
	var a int
	fmt.Scanf("%d\n", &a)
	var dd, mm, yyyy []int
	for i := 0; i < a; i++ {
		var d, m, y int
		fmt.Scanf("%d %d %d\n", &d, &m, &y)
		dd = append(dd, d)
		mm = append(mm, m)
		yyyy = append(yyyy, y)
	}
	for i := 0; i < a; i++ {
		isLeap(yyyy[i])
		if yyyy[i] > 2300 || yyyy[i] < 1950 {
			fmt.Println("NO")
			continue
		}
		if mm[i] < 1 || mm[i] > 12 {
			fmt.Println("NO")
			continue
		}
		if dd[i] < 1 || dd[i] > daysInMonth(mm[i]) {
			fmt.Println("NO")
			continue
		}
		fmt.Println("YES")
	}
}
