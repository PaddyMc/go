package main

import (
	"fmt"
)

func main() {
	//power := getPower()

	//name, power := "Goku", 9000
	//fmt.Printf("%s's power is over %d\n", name, power)

	a, exists := power("pat")
	//fmt.Printf(a)
	//fmt.Printf(exists)
	fmt.Printf("over %d\n",a)
	if(exists == false){

	}

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
}