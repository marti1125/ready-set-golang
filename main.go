package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/orlandovald/ready-set-golang/marvel"
)

const (
	// MarvelPrivateKey is the environment variable name for the api's private key
	MarvelPrivateKey = "MARVEL_PRIVATE_KEY"
	// MarvelPublicKey is the environment variable name for the api's public key
	MarvelPublicKey = "MARVEL_PUBLIC_KEY"
)

var api *marvel.Api

func main() {

	privateKey := os.Getenv(MarvelPrivateKey)
	publicKey := os.Getenv(MarvelPublicKey)

	api = marvel.NewAPI(privateKey, publicKey)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/character/random", characterHandler)

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	character := api.GetRandomCharacter()

	j, err := json.Marshal(character)
	if err != nil {
		j = []byte("{'error':'Error marshalling marvel character'}")
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}
