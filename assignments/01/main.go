package main

import "fmt"

func main() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, v := range s {
		oe := "even"
		if v % 2 == 1 {
			oe = "odd"
		}
		fmt.Printf("%d is %s\n", v, oe)
	}
}
