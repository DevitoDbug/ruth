package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func validateWord(url string, sentence string) (*http.Response, error) {
	newURL := fmt.Sprintf("%v%v", url, sentence)
	req, err := http.NewRequest("GET", newURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", "jfNRme3SHJ3WIGR29FpyhLJC5PT4qem0")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(req.Body)
	client := &http.Client{}
	body, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

//$ go run main.go how are you doink
//{
//    "corrections": [
//        {
//            "text": "doink",
//            "best_candidate": "doing",
//            "candidates": [
//                "doing",
//                "dwink",
//                "drink"
//            ]
//        }
//    ],
//    "original_text": "how are you doink"
//}
func main() {
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

	fmt.Println(string(responseBody))
}
