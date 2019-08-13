package commitActivity

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

import (
	"github.com/JustinDFuller/github-activity/internal/repos"
	"github.com/JustinDFuller/github-activity/internal/types"
	"github.com/JustinDFuller/github-activity/internal/url"
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
		log.Fatal("An error occurred while parsing JSON.", err)
	}
}

func (responseTarget *GithubCommitResponse) httpRequest(githubUrl string, index int, repos *repos.GithubReposResponse) {
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

	(*repos)[index].Commits = FormatCommitActivity(*responseTarget)
}

func prettyPrintJSON(parsedJSON repos.GithubReposResponse) {
	formattedJSON, err := json.MarshalIndent(parsedJSON, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(formattedJSON))
}

func FetchCommitActivity(repos repos.GithubReposResponse, userName types.User) repos.GithubReposResponse {
	var waitGroup sync.WaitGroup

	urlShaReplacer := strings.NewReplacer("{/sha}", "")

	for index, value := range repos {
		waitGroup.Add(1)

		var commits GithubCommitResponse
		url := urlShaReplacer.Replace(value.CommitsUrl)

		go func(url string, index int) {
			defer waitGroup.Done()
			commits.httpRequest(url, index, &repos)
		}(url, index)
	}

	waitGroup.Wait()

	return repos
}
