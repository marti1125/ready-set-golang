/*
Package that implements a minimal library to interact with the Marvel API.
See https://developer.marvel.com/
*/
package marvel

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	tsParam     = "ts"
	hashParam   = "hash"
	apiKeyParam = "apikey"
	offsetParam = "offset"
	baseURL     = "https://gateway.marvel.com/v1/public/characters"
	total       = 1490
)

type Api struct {
	privateKey string
	publicKey  string
}

type response struct {
	Code            int       `json:"code"`
	Status          string    `json:"status"`
	Data            container `json:"data"`
	Etag            string    `json:"etag"`
	Copyright       string    `json:"copyright"`
	AttributionText string    `json:"attributionText"`
	AttributionHTML string    `json:"attributionHTML"`
}

type container struct {
	Offset  int         `json:"offset"`
	Limit   int         `json:"limit"`
	Total   int         `json:"total"`
	Count   int         `json:"count"`
	Results []Character `json:"results"`
}

type thumbnail struct {
	Path      string `json:"path`
	Extension string `json:"extension`
}

type Character struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Thumbnail   thumbnail `json:"thumbnail"`
}

// NewAPI returns a new api given the public and private keys
func NewAPI(privateKey, publicKey string) *Api {
	rand.Seed(time.Now().UnixNano())
	return &Api{privateKey, publicKey}
}

// GetRandomeCharacter returns a random character from the Marvel API
func (api *Api) GetRandomCharacter() Character {

	ts := strconv.FormatInt(time.Now().UnixNano(), 10)
	offset := rand.Intn(total)
	hash := hash(ts, api.privateKey, api.publicKey)
	url := fmt.Sprintf("%s?%s=%s&%s=%s&%s=%s&limit=1&%s=%d", baseURL, tsParam, ts, apiKeyParam, api.publicKey, hashParam, hash, offsetParam, offset)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v", err)
		return Character{Name: "Http-Error-Man"}
	}
	fmt.Printf("Status code: %d\n", resp.StatusCode)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
		return Character{Name: "Io-Error-Man"}
	}

	var response = response{}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("%v", err)
		return Character{Name: "Json-Error-Man"}
	}

	fmt.Printf("%#v\n", response)

	return response.Data.Results[0]
}

// hash generates a hash value required by the Marvel API
// hash = md5(ts+privateKey+publicKey)
func hash(ts, privateKey, publicKey string) string {
	h := md5.New()
	h.Write([]byte(ts))
	h.Write([]byte(privateKey))
	h.Write([]byte(publicKey))
	return fmt.Sprintf("%x", h.Sum(nil))
}
