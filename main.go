package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ResponseValue struct {
	Corrections []struct {
		Text          string   `json:"text"`
		BestCandidate string   `json:"best_candidate"`
		Candidates    []string `json:"candidates"`
	} `json:"corrections"`
	OriginalText string `json:"original_text"`
}

func validateWord(url string, sentence string) (*http.Response, error) {
	newURL := fmt.Sprintf("%v%v", url, sentence)
	req, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", "jfNRme3SHJ3WIGR29FpyhLJC5PT4qem0")
	client := &http.Client{}
	body, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	var responseValue ResponseValue
	var commitMessage []string
	url := "https://api.apilayer.com/spell/spellchecker?q="
	if len(os.Args) >= 2 {
		commitMessage = append(commitMessage, os.Args[1:]...)
	}

	response, err := validateWord(url, strings.Join(commitMessage, "%20"))
	if err != nil {
		panic(err)
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal([]byte(responseBody), &responseValue); err != nil {
		panic(err)
	}
	fmt.Println(responseValue)
}
