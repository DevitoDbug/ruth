package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
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
	stageCmd := exec.Command("git", "add", ".")
	commitCmd := exec.Command("git", "commit", "-m", commitMessage)

	stageCmd.Dir = currentDir
	commitCmd.Dir = currentDir

	output1, err := stageCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	output2, err := commitCmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("Commit made")
	fmt.Printf("%v\n%v", string(output1), string(output2))
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
		handleCommit(strings.Join(commitMessage, " "))
		return
	} else {
		color.Red("Spelling errors in the commit message!!")
		fmt.Println("Mistakes: ")
		for index, mistake := range responseValue.Corrections {
			fmt.Printf("	%v.wrong word: ", index+1)
			color.Red("%v\n", mistake.Text)

			fmt.Print("		Suggestions: ")
			color.Green("%v\n", mistake.Candidates)

		}
	}
}
