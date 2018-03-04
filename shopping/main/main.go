package main

import (
	"shopping"
	"fmt"
	//"github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(shopping.PriceCheck(4343))
}

/*type Logger interface {
	Log(message string)
}

type SqlLogger struct {}
type ConsoleLogger struct {}
type FileLogger struct {}

type Server struct {
	logger Logger
}

func process(logger Logger) {
	logger.Log("hello!")
}

type ConsoleLogger struct {}
	func (l ConsoleLogger) Log(message string) {
			fmt.Println(message)
}
*/
