package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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

func handleCommit(commitMessage string) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		os.Exit(1)
	}
	commitCmd := exec.Command("git", "commit", "-m", commitMessage)
	commitCmd.Dir = currentDir
	output, err := commitCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Commit made")
	fmt.Println("Commit message: ", string(output))
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

	if len(responseValue.Corrections) <= 0 {
		handleCommit(string(responseBody))
		return
	}

	fmt.Println("Spelling errors in commit message!!")
	fmt.Println("Mistakes: ")
	for _, mistake := range responseValue.Corrections {
		fmt.Printf("	wrong word: %v\n", mistake.Text)
		fmt.Printf("		Suggestions: %v\n", mistake.Candidates)
	}
}
