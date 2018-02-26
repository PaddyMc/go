package main

import (
	"fmt"
)

func main() {
	playerList := make(map[string]int)
	playerList["p"] = 900
	playerList["a"] = 90000
	power, exists := playerList["p"]

	total := len(playerList)
	fmt.Println(total)
	fmt.Println(power, exists)

	delete(playerList, "p")

	power2, exists2 := playerList["p"]
	fmt.Println(power2, exists2)

	for key, value := range playerList {
		
	}

}

type Player struct {
	Name string
	Power int
	Friends map[string]*Player
}


