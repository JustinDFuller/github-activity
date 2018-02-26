package commitActivity

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

import (
	"github/repos"
	"github/url"
)

type GithubCommitResponse []struct {
	Commit struct {
		Committer struct {
			Date string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

func (response *GithubCommitResponse) parseJSON(jsonToParse []byte) {
	if !json.Valid(jsonToParse) {
		log.Fatal("Invalid JSON received")
	}

	err := json.Unmarshal(jsonToParse, &response)

	if err != nil {
		log.Fatal("An error occured while parsing JSON.", err)
	}
}

func (responseTarget *GithubCommitResponse) httpRequest(githubUrl string) {
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

func prettyPrintJSON(parsedJSON repos.GithubReposResponse) {
	formattedJSON, err := json.MarshalIndent(parsedJSON, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(formattedJSON))
}

func FetchCommitActivity(repos repos.GithubReposResponse) repos.GithubReposResponse {
	replaceShaInUrl := strings.NewReplacer("{/sha}", "")

	for index, value := range repos {
		var commits GithubCommitResponse
		url := replaceShaInUrl.Replace(value.CommitsUrl)
		commits.httpRequest(url)
		repos[index].Commits = FormatCommitActivity(commits)
	}

	return repos
}
