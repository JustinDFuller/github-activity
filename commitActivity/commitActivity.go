package commitActivity

import (
	"fmt"
	"log"
	"sync"
	"strings"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

import (
	"github-activity/repos"
	"github-activity/url"
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

func (responseTarget *GithubCommitResponse) httpRequest(githubUrl string, c chan GithubCommitResponse) {
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
	
	c <- *responseTarget
}

func prettyPrintJSON(parsedJSON repos.GithubReposResponse) {
	formattedJSON, err := json.MarshalIndent(parsedJSON, "", "  ")

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Print(string(formattedJSON))
}

func FetchCommitActivity(repos repos.GithubReposResponse) repos.GithubReposResponse {
  var wg sync.WaitGroup
  
	replaceShaInUrl := strings.NewReplacer("{/sha}", "")
  c := make(chan GithubCommitResponse, len(repos))

	for _, value := range repos {
	  wg.Add(1)
	  
		var commits GithubCommitResponse
		url := replaceShaInUrl.Replace(value.CommitsUrl)
		
		go func(url string) {
      defer wg.Done()
		  commits.httpRequest(url, c)
    }(url)
	}
	
	wg.Wait()
	
	for index, _ := range repos {
	  response := <-c
		repos[index].Commits = FormatCommitActivity(response)
	}
	
	return repos
}
