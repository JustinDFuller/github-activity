package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

import (
	"github-activity/url"
	"github-activity/types"
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

	fmt.Print(string(formattedJSON))
}

func FetchReposUrl(user types.User) string {
	var githubApi GithubUserResponse
	githubApi.httpRequest("https://api.github.com/users/" + string(user))
	return githubApi.ReposUrl
}
