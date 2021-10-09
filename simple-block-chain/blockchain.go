package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
)

const (
	START            = "START"
	NO_HASH          = "NO_HASH"
	VALIDATION_OK    = 1
	VALIDATION_ERROR = 2
)

type Node struct {
	Next     *Node
	Value    string
	NextHash string
}

type ValidationResult struct {
	Status  int
	Message string
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

func initBlockChain(head *Node) {
	scanner := bufio.NewScanner(os.Stdin)
	current := head
	for scanner.Scan() {
		fmt.Print("Please enter value for next block in blockchain: ")
		inputValue := scanner.Text()
		if isStringValid(inputValue) {
			next := &Node{Next: nil, Value: inputValue, NextHash: NO_HASH}
			current.NextHash = getHashOfValue(inputValue)
			current.Next = next
			current = next
		} else {
			break
		}
	}

}

func validateBlockChain(head *Node) ValidationResult {
	temp := head
	result := ValidationResult{Status: VALIDATION_OK, Message: "blockchain is valid"}
	hash := head.NextHash
	if temp.Next != nil {
		temp = temp.Next
	}
	for temp != nil {
		currentNodeHash := getHashOfValue(temp.Value)
		if currentNodeHash != hash {
			fmt.Println(currentNodeHash + "  " + hash)
			result.Status = VALIDATION_ERROR
			result.Message = "Validation error blockchain was tempered"
			return result
		}
		hash = temp.NextHash
		temp = temp.Next
	}
	return result
}

func printResults(validationResult ValidationResult) {
	fmt.Println("Validation result: ", validationResult.Status)
	fmt.Println("Validatoin message: " + validationResult.Message)
}

func main() {
	head := &Node{Next: nil, Value: START, NextHash: NO_HASH}
	initBlockChain(head)
	blockChainValidationResult := validateBlockChain(head)
	printResults(blockChainValidationResult)
}
