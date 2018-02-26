package main

import (
"fmt"
"math/rand"
"sort"
)

func main() {
	scores := make([]int, 100)

	for i := 0; i < 100; i++ {
		scores[i] = int(rand.Int31n(1000))
	}
	sort.Ints(scores)

	worst := make([]int, 15)
	copy(worst[1:5], scores[:15])
	fmt.Println(worst)
}
