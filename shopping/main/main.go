package main

import (
	"shopping"
	"fmt"
	"github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println(shopping.PriceCheck(4343))
}
