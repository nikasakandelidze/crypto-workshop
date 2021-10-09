package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

const (
	START   = "START"
	NO_HASH = "NO_HASH"
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
	hasher := sha1.New()
	hasher.Write([]byte(value))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return sha
}

func getStdIn() {
	scanner := bufio.NewScanner(os.Stdin)
	head := &Node{Next: nil, Value: START, NextHash: NO_HASH}
	current := head
	for scanner.Scan() {
		fmt.Print("Please enter value for next block in blockchain: ")
		inputValue := scanner.Text()
		if isStringValid(inputValue) {
			next := &Node{Next: nil, Value: inputValue, NextHash: NO_HASH}
			current.NextHash = getHashOfValue(current.Value)
			current.Next = next
			current = next
		} else {
			break
		}
	}
	temp := head
	for temp != nil {
		fmt.Println(temp.Value)
		fmt.Println(temp.NextHash + " \n")
		temp = temp.Next
	}
}

func main() {
	getStdIn()
}
