package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

func validateWord(url string, payload []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(req.Body)

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func main() {
	var commitMessage []string
	//url := "https://words.dev-apis.com/validate-word"

	if len(os.Args) >= 2 {
		commitMessage = append(commitMessage, os.Args[1:]...)
	}
	fmt.Println(commitMessage)
}
