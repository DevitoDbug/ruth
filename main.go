package main

import (
	"fmt"
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
	client := &http.Client{}
	body, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return body, nil
}

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

//func main() {
//	word := []string{"name", "is", "thinj"}
//	url := "https://api.apilayer.com/spell/spellchecker?q="
//	urlWords := fmt.Sprintf("%v%v", url, strings.Join(word, "%20"))
//
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", urlWords, nil)
//	req.Header.Set("apikey", "jfNRme3SHJ3WIGR29FpyhLJC5PT4qem0")
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	res, err := client.Do(req)
//	if res.Body != nil {
//		defer res.Body.Close()
//	}
//	body, err := ioutil.ReadAll(res.Body)
//
//	fmt.Println(string(body))
//}
