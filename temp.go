package main

import (
	"fmt"
)

func main() {
	player := &Player{Name: "a"}
	player.Power = 9000

	//Improve(player);
	player.Improve()

	fmt.Println(player.Power)

	player2 := NewPlayer("s",100)
	fmt.Println(player2.Power)

	//scores := [4]int{9001, 9333, 212, 33}
	/*for index, value := range scores {
		fmt.Println(index)
		fmt.Println(value)
	}*/

	//scores := []int{1,4,293,4,9}
 

	//scores := make([]int, 10)
	scores := make([]int, 0, 10)
	scores = scores[0:8]
	scores[7] = 9033
	fmt.Println(scores)
}



func getPower() int {
	return 9001
}

func log(message string) {

}
func add(a int, b int) int {
	return a+b
}
func power(name string) (int, bool) {
	return 1,true
}

type Player struct {
	Name string
	Power int
	Father *Player
}

func (p *Player) Improve (){
	p.Power += 10000
}

func NewPlayer(name string, power int) *Player {
	return &Player{
		Name: name,
		Power: power,
	}
}

/*func NewPlayer(name string, power int) Player {
	return Player{
		Name: name,
		Power: power,
	}
}*/

