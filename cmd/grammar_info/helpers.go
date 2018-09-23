package main

import "fmt"

func printCommaList(indices []int, lookup []string, qc string) {
	for n, index := range indices {
		if n != 0 {
			if n == len(indices)-1 {
				fmt.Printf(" and ")
			} else {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("%s%s%s", qc, lookup[index], qc)
	}
}

func intRange(n int) []int {
	list := []int{}
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	return list
}

func plural(n int, sing, plur string) string {
	if n == 1 {
		return fmt.Sprintf("%d %s", n, sing)
	}
	return fmt.Sprintf("%d %s", n, plur)
}
