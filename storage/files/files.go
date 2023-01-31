package files

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Storage struct {
	basePath string
}

const defaultPerm = 0774

func New(basePath string) Storage {
	return Storage{basePath: basePath}
}

func (s Storage) RndJoke() string {
	url := "https://jokeapi.dev/joke/Any?format=txt&type=single&blacklistFlags=nsfw,racist,sexist&lang=en"
	responseBytes := getJokeData(url)
	return fmt.Sprintf(strings.TrimSpace(string(responseBytes)))
}

func getJokeData(baseAPI string) []byte {
	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)
	if err != nil {
		log.Printf("Could not request a dadjoke. %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "GenerationJoke CLI (https://github.com/NineIsCool/GenerationJoke)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}
