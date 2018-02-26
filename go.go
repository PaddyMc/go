package main

import (
	"fmt"
)

func main() {
	scores := make([]int, 0, 5)
	c := cap(scores)
	fmt.Println(c)
	for i := 0; i < 25; i++ {
		scores = append(scores, i)

		if cap(scores) != c {
			c = cap(scores)
			fmt.Println(c)
		}
	}
	scores1 := make([]int, 5)
	scores1 = append(scores, 9332)
	fmt.Println(scores1)
}
