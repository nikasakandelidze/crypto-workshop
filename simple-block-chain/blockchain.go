package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	START = 1
)

type Node struct {
	Next     *Node
	Value    string
	NextHash string
}

func isStringValid(str string) bool {
	return len(str) > 0
}

func getHashOfValue(value string) string {
	return "1"
}

func getStdIn() {
	scanner := bufio.NewScanner(os.Stdin)
	head := &Node{Next: nil, Value: START, NextHash: nil}
	current := &head
	for scanner.Scan() {
		fmt.Println("Please enter value for next block in blockchain")
		inputValue := scanner.Text()
		if isStringValid(inputValue) {
			next := &Node{Next: nil, Value: inputValue, NextHash: nil}
			current.Hash = getHashOfValue(current.Value)
			current.Next = &next
			current = &next
		}
	}
}

func main() {
	getStdIn()
}
