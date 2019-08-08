package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

import (
	"github.com/JustinDFuller/github-activity/internal/types"
	"github.com/JustinDFuller/github-activity/internal/url"
)

type GithubUserResponse struct {
	ReposUrl string `json:"repos_url"`
}

func (response *GithubUserResponse) parseJSON(jsonToParse []byte) {
	if !json.Valid(jsonToParse) {
		log.Fatal("Invalid JSON received")
	}

	err := json.Unmarshal(jsonToParse, &response)

	if err != nil {
		log.Fatal("An error occured while parsing JSON.")
	}
}

func (responseTarget *GithubUserResponse) httpRequest(githubUrl string) {
	urlWithAuth := url.FormatWithAuth(githubUrl)

	fmt.Print(urlWithAuth)

	res, err := http.Get(urlWithAuth)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	responseTarget.parseJSON(body)
}

func (parsedJSON GithubUserResponse) prettyPrintJSON() {
	formattedJSON, err := json.MarshalIndent(parsedJSON, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(formattedJSON))
}

func FetchReposUrl(user types.User) string {
	var githubApi GithubUserResponse
	githubApi.httpRequest("https://api.github.com/users/" + string(user))
	return githubApi.ReposUrl
}
