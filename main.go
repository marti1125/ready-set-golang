package main

import (
	"fmt"
	"os"

	"github.com/orlandovald/ready-set-golang/marvel"
)

const (
	MarvelPrivateKey = "MARVEL_PRIVATE_KEY"
	MarvelPublicKey  = "MARVEL_PUBLIC_KEY"
)

func main() {

	privateKey := os.Getenv(MarvelPrivateKey)
	publicKey := os.Getenv(MarvelPublicKey)

	api := marvel.NewApi(privateKey, publicKey)
	character1 := api.GetRandomCharacter()
	fmt.Printf("%#v\n", character1)

}
