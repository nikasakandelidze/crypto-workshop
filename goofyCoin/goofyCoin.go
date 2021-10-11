package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	uuid "github.com/google/uuid"
)

type User struct {
	UUID       string
	name       string
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

var usersStorage = make(map[string]*User)

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
	return newUser
}

func main() {
	fmt.Println("Starting goofy coin mechanism")
	user := createUser("goofy")
	if user == nil {
		fmt.Println("Failed to create a user")
	}
	usersStorage[user.UUID] = user
	fmt.Println(usersStorage[user.UUID])
}
