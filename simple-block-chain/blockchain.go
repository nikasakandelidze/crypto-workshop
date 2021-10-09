package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	START            = "START"
	NO_HASH          = "NO_HASH"
	VALIDATION_OK    = 1
	VALIDATION_ERROR = 2
)

type Node struct {
	Next      *Node
	Value     string
	NextHash  string
	TimeStamp string
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
	fmt.Println("> Initialization of blockchain. Please start entering values for nodes. ")
	scanner := bufio.NewScanner(os.Stdin)
	current := head
	fmt.Print("> Please enter value for next block in blockchain: ")
	for scanner.Scan() {
		inputValue := scanner.Text()
		if isStringValid(inputValue) {
			next := &Node{Next: nil, Value: inputValue, NextHash: NO_HASH, TimeStamp: time.Now().Format(time.RFC850)}
			current.NextHash = getHashOfValue(inputValue)
			current.Next = next
			current = next
			fmt.Print("> Please enter value for next block in blockchain: ")
		} else {
			fmt.Println("> Breaking Blockhain data initialization process.")
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
	fmt.Println("> Validation result: ", validationResult.Status)
	fmt.Println("> Validatoin message: " + validationResult.Message)
}

func visualizeBlockchain(withData bool, head *Node) {
	temp := head
	fmt.Println("> Visualisation of Blockchain")
	for temp != nil {
		fmt.Println("[ value: " + temp.Value + ",  timestamp: " + temp.TimeStamp + ", next Hash: " + temp.NextHash + "]")
		fmt.Println("    |")
		fmt.Println("    |")
		fmt.Println("   \\|/")
		temp = temp.Next
	}
}

func serveUser(head *Node) {
	fmt.Println("> Initialization Complete.\n Please enter:\n   visualize - to visualize blockchain.\n   validate - to validate blockchain\n")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.EqualFold(input, "visualize") {
			visualizeBlockchain(true, head)
		} else if strings.EqualFold(input, "validate") {
			validationResult := validateBlockChain(head)
			printResults(validationResult)
		}
	}
}

func main() {
	head := &Node{Next: nil, Value: START, NextHash: NO_HASH, TimeStamp: time.Now().Format(time.RFC850)}
	initBlockChain(head)
	serveUser(head)
}
