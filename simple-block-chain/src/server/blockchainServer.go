package main

import (
	"crypto/sha1"
	"encoding/base64"
	"log"
	"net"
	"net/rpc"
	"sync"
	"time"

	commons "blockchain/src/commons"
)

const (
	START            = "START"
	NO_HASH          = "NO_HASH"
	VALIDATION_OK    = 1
	VALIDATION_ERROR = 2
)

type BlockChain struct {
	Head  *Node
	mutex sync.Mutex
}

type Node struct {
	Next      *Node
	Value     string
	NextHash  string
	TimeStamp string
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

func (blockChain *BlockChain) InsertBlock(request *commons.InsertNodeRequest, response *commons.InsertNodeResponse) error {
	inputValue := request.Payload
	current := blockChain.Head
	if isStringValid(inputValue) {
		next := &Node{Next: nil, Value: inputValue, NextHash: NO_HASH, TimeStamp: time.Now().Format(time.RFC850)}
		current.NextHash = getHashOfValue(inputValue)
		current.Next = next
		current = next
		log.Println("Succesfully added new value into Blockchain")
	} else {
		log.Println("Failed to add new value into Blockchain")
	}
	return nil
}

func (blockChain *BlockChain) ValidateBlockChain(_ *commons.BlockchainValidationRequest, response *commons.BlockchainValidationResponse) error {
	temp := blockChain.Head
	hash := temp.NextHash
	if temp.Next != nil {
		temp = temp.Next
	}
	for temp != nil {
		currentNodeHash := getHashOfValue(temp.Value)
		if currentNodeHash != hash {
			log.Println(currentNodeHash + "  " + hash)
			response.Valid = false
			response.Message = "Validation error! blockchain was tempered!"
			break
		}
		hash = temp.NextHash
		temp = temp.Next
	}
	response.Valid = true
	response.Message = "Blockchain is in a consistent, valid state."
	return nil
}

// func (blockChain *BlockChain) visualizeBlockChain() {

// }

func initServer(waitGroup *sync.WaitGroup) {
	blockChain := new(BlockChain)
	blockChain.Head = &Node{Next: nil, Value: START, NextHash: NO_HASH, TimeStamp: time.Now().Format(time.RFC850)}
	server := rpc.NewServer()
	server.Register(blockChain)
	l, e := net.Listen("tcp", ":8081")
	if e != nil {
		log.Fatal(e)
	}
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()
		for {
			conn, err := l.Accept()
			if err == nil {
				go server.ServeConn(conn)
			} else {
				break
			}
		}
		l.Close()
	}()

}

func main() {
	log.Println("Starting rpc server.")
	wg := new(sync.WaitGroup)
	initServer(wg)
	wg.Wait()
	log.Println("Closing RPC server program.")
}
