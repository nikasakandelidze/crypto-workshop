package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	uuid "github.com/google/uuid"
)

const (
	BOSS                 = "goofy"
	CREATE_INSTRUCTION   = "CreateCoin:"
	TRANSFER_INSTRUCTION = "TransferCoin:"
)

type User struct {
	UUID       string
	name       string
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

type Node struct {
	payload   string
	signature []*big.Int
	prev      *Node
	ownerId   string
}

//global variables

var usersStorage = make(map[string]*User) //for storing users
var ledger *Node
var GOOFYS_UUID string

func generateKeys() (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, err
}

func createUser(name string) *User {
	privateKey, publicKey, err := generateKeys()
	if err != nil {
		fmt.Println("Error while generating keys")
		return nil
	}
	newUser := &User{UUID: uuid.New().String(), name: name, privateKey: privateKey, publicKey: publicKey}
	usersStorage[newUser.UUID] = newUser
	return newUser
}

func getTokenUUIDFromPayload(payload string) string {
	splittedArray := strings.Split(payload, ":")
	if len(splittedArray) == 2 {
		return splittedArray[1]
	} else {
		return ""
	}
}

func createNewNode(payload string, ownerId string) *Node {
	user := usersStorage[ownerId]
	r, s, err := ecdsa.Sign(rand.Reader, user.privateKey, []byte(payload))
	if err != nil {
		fmt.Println("Failed to sign payload of new node")
		return nil
	}
	newNode := &Node{payload: payload, signature: []*big.Int{r, s}, prev: nil, ownerId: ownerId}
	return newNode
}

func createNewCoin(ownerId string) *Node {
	if ownerId != GOOFYS_UUID {
		fmt.Println("Oncly goofy can create coins")
		return nil
	}
	node := createNewNode(CREATE_INSTRUCTION+uuid.New().String(), ownerId)
	if node == nil {
		fmt.Println("node is nil")
		return nil
	}
	if ledger != nil {
		node.prev = ledger
	}
	ledger = node
	fmt.Println("Goofy created a new coin")
	return ledger
}

func transferCoin(fromId string, toId string, coinId string) {
	fromUser := usersStorage[fromId]
	toUser := usersStorage[toId]
	if fromUser == nil || toUser == nil {
		fmt.Println("Failed transaction! error: to user and from user ids' must be specified correctly.")
		return
	}
	newNode := createNewNode(TRANSFER_INSTRUCTION+fromId+":"+toId, fromId)
	if newNode == nil {
		fmt.Println("Failed to create new node for transaction")
	}
	if ledger == nil {
		fmt.Println("No coins created yet. Failed to make a transaction.")
	}
	newNode.prev = ledger
	ledger = newNode
	fmt.Println("User with id: " + fromId + "succesfully transfered coin to user with id: " + toId)
}

//below are verification utility functions

func verify() {
	fmt.Println("***************Starting verifying process*****************")
	currentNode := ledger
	for currentNode != nil {
		payload := currentNode.payload
		identifier := strings.Split(payload, ":")
		currentNodeOwnerId := currentNode.ownerId
		fmt.Println(currentNodeOwnerId)
		user := usersStorage[currentNodeOwnerId]
		if user == nil {
			fmt.Println("Failed to verify, couldn't find user in usersStorage")
			return
		}
		if identifier[0] == "CreateCoin" {
			goofyUser := usersStorage[GOOFYS_UUID]
			isVerified := ecdsa.Verify(goofyUser.publicKey, []byte(currentNode.payload), currentNode.signature[0], currentNode.signature[1])
			if !isVerified {
				fmt.Println("Failed to verify signature for the coin creation node. payload: " + currentNode.payload)
				return
			}
			fmt.Println("Succesfully verified coin creation by goofy for payload: " + currentNode.payload)
		}
		if identifier[0] == "TransferCoin" {
			fromId := identifier[1]
			fromUser := usersStorage[fromId]
			isVerified := ecdsa.Verify(fromUser.publicKey, []byte(currentNode.payload), currentNode.signature[0], currentNode.signature[1])
			if !isVerified {
				fmt.Println("Failed to verify signature for the coin transaction. payload: " + currentNode.payload)
				return
			}
			fmt.Println("Succesfully verified coin transaction with payload: " + currentNode.payload)
		}
		currentNode = currentNode.prev
	}
}

func main() {
	//initialize
	fmt.Println("Starting goofy coin mechanism")
	userGoofy := createUser(BOSS)
	if userGoofy == nil {
		fmt.Println("Failed to create a user")
	}
	GOOFYS_UUID = userGoofy.UUID
	goofy := usersStorage[userGoofy.UUID]
	coin := createNewCoin(goofy.UUID)
	if coin == nil {
		fmt.Println("Error while creating new coin")
	}
	userAlice := createUser("Alice")
	if userAlice == nil {
		fmt.Println("Failed to create a user")
	}
	coinId := getTokenUUIDFromPayload(coin.payload)
	transferCoin(userGoofy.UUID, userAlice.UUID, coinId)
	//verify
	verify()
}
