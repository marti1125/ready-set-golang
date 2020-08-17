package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/orlandovald/ready-set-golang/marvel"
)

const (
	// MarvelPrivateKey is the environment variable name for the api's private key
	MarvelPrivateKey = "MARVEL_PRIVATE_KEY"
	// MarvelPublicKey is the environment variable name for the api's public key
	MarvelPublicKey = "MARVEL_PUBLIC_KEY"
)

// Port to start listening for requests
const Port = ":8080"

var api *marvel.API

func main() {

	privateKey := os.Getenv(MarvelPrivateKey)
	publicKey := os.Getenv(MarvelPublicKey)

	api = marvel.NewAPI(privateKey, publicKey)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/character/random", characterHandler)
	http.HandleFunc("/character/random/", randomCountHandler)

	log.Printf("Starting server on port %s", Port)
	log.Fatal(http.ListenAndServe(Port, nil))

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

func randomCountHandler(w http.ResponseWriter, r *http.Request) {

	count, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/character/random/"))
	if err != nil {
		count = 1
	}
	ch := make(chan marvel.Character, count)

	data := make([]marvel.Character, count)

	for i := 0; i < count; i++ {
		go func() {
			ch <- api.GetRandomCharacter()
		}()
	}
	for i := 0; i < count; i++ {
		data[i] = <-ch
	}

	j, err := json.Marshal(data)
	if err != nil {
		j = []byte("{'error':'Error marshalling marvel character'}")
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(j)
}
