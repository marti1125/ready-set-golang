package main

import (
	"fmt"
	"os"

	"github.com/orlandovald/ready-set-golang/marvel"
)

const (
	// MarvelPrivateKey is the environment variable name for the api's private key
	MarvelPrivateKey = "MARVEL_PRIVATE_KEY"
	// MarvelPublicKey is the environment variable name for the api's public key
	MarvelPublicKey = "MARVEL_PUBLIC_KEY"
)

func main() {

	privateKey := os.Getenv(MarvelPrivateKey)
	publicKey := os.Getenv(MarvelPublicKey)

	api := marvel.NewAPI(privateKey, publicKey)
	character1 := api.GetRandomCharacter()
	fmt.Printf("%#v\n", character1)

}
