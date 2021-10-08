package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	Next     *Node
	Value    string
	NextHash string
}

func main() {
	a := 1
	fmt.Println("vim-go " + strconv.Itoa(a))
}
