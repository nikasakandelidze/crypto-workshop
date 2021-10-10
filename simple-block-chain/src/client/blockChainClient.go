package main

import (
	commons "blockchain/src/commons"
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

// func visualizeBlockchain(withData bool, head *Node) {
// 	temp := head
// 	fmt.Println("> Visualisation of Blockchain")
// 	for temp != nil {
// 		fmt.Println("[ value: " + temp.Value + ",  timestamp: " + temp.TimeStamp + ", next Hash: " + temp.NextHash + "]")
// 		fmt.Println("    |")
// 		fmt.Println("    |")
// 		fmt.Println("   \\|/")
// 		temp = temp.Next
// 	}
// }

func isStringValid(str string) bool {
	return len(str) > 0
}

func addNewValuetoBlockchain(value string) {
	if !isStringValid(value) {
		log.Fatal("Input value not valid")
		return
	}
	client, err := rpc.Dial("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	request := &commons.InsertNodeRequest{Payload: value}
	response := &commons.InsertNodeResponse{}
	log.Println("sending request")
	errorResponse := client.Call("BlockChain.InsertBlock", request, response)
	if errorResponse != nil {
		log.Fatal(errorResponse)
	}
	client.Close()
	log.Println("Blockchain node insert successful for value: " + value)
}

func serveUser() {
	fmt.Println("> Welcome to Blockchain client\nPlease enter:\n  add: {VALUE} - to add node in the blockchain.\n  validate - to validate blockchain")
	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		inputs := strings.Split(input, " ")
		if strings.EqualFold(inputs[0], "add") {
			value := inputs[1]
			addNewValuetoBlockchain(value)
		} else if strings.EqualFold(input, "validate") {
		}
		fmt.Print("> ")
	}
}

func main() {
	serveUser()
}
